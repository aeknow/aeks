package main

import (
	"encoding/json"

	//"bufio"
	//crypto_rand "crypto/rand"
	"database/sql"
	"encoding/base64"

	//"encoding/hex"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	_ "modernc.org/sqlite"

	//"time"

	"github.com/aeternity/aepp-sdk-go/v9/account"
	aeconfig "github.com/aeternity/aepp-sdk-go/v9/config"
	"github.com/aeternity/aepp-sdk-go/v9/naet"
	"github.com/aeternity/aepp-sdk-go/v9/transactions"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	qrcode "github.com/skip2/go-qrcode"

	//"github.com/jdgcs/ed25519/extra25519"
	utils "github.com/aeternity/aepp-sdk-go/v9/utils"
	"github.com/tyler-smith/go-bip39"

	//"golang.org/x/crypto/nacl/box"

	shell "github.com/ipfs/go-ipfs-api"
)

/////////////////////////////////////////header///////////////////////////////////
type AccountInfo struct {
	// Name of authorization function for generalized account
	AuthFun string `json:"auth_fun,omitempty"`

	// Balance
	// Required: true
	Balance utils.BigInt `json:"balance"`

	// Id of authorization contract for generalized account
	ContractID string `json:"contract_id,omitempty"`

	// Public key
	// Required: true
	ID string `json:"id"`

	// kind
	// Enum: [basic generalized]
	Kind string `json:"kind,omitempty"`

	// Nonce
	// Required: true
	Nonce uint64 `json:"nonce"`

	// Payable
	Payable bool `json:"payable,omitempty"`
}

type HandleFnc func(http.ResponseWriter, *http.Request)

//HOMEPAGE
var aecommands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

type PageReg struct {
	PageTitle string
	SubTitle  string
	Register  string
	Lang      langFile
}

type PageLogin struct {
	Options template.HTML
	Lang    langFile
}

type AeknowConfig struct {
	PublicNode  string
	APINode     string
	IPFSNode    string
	IPFSAPI     string
	LocalWeb    string
	PubsubProxy string
}

//var myAccount account.Account

//var globalAccount account.Account
//var NodeConfig AeknowConfig

//AENS
type TTLer func(offset uint64) (ttl uint64, err error)

type AENSNames struct {
	Aensname      string
	Expire_height int64
}

type NameSlice struct {
	Names []AENSNames
}

type AENSBidinfo struct {
	Aensname   string
	Lastbidder string
	Lastprice  string
}

type NameBidSlice struct {
	Names []AENSBidinfo
}

type AENSInfo struct {
	ID       string
	TTL      uint64
	OWNER    string
	Pointers []transactions.NamePointer
}

type PageAENS struct {
	PageId      int
	PageContent template.HTML
	PageTitle   string
	Account     string
	Balance     string
	Nonce       uint64
}

///////////////////////////////////////////////////////////////////////////////////////
var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

type PageWallet struct {
	PageId       int
	PageContent  template.HTML
	PageTitle    string
	Account      string
	Balance      string
	Nonce        uint64
	Recipient_id string
	Payload      string
	Amount       string
}

var NodeOnline bool

///////////////////IPFS repo//////////////////////
type Reposite struct {
	Reposites []DomainSite
}

type DomainSite struct {
	Name     string
	Hash     string
	Metainfo string
}

func AE_WEB_RegisterNew(ctx iris.Context) {
	var myPage PageReg
	myPage.PageTitle = "Registering Page"
	myPage.SubTitle = "Decentralized knowledge system without barrier."
	myPage.Register = "Register"

	//myPage.Lang = getPageString(getPageLang(r))

	ctx.ViewData("", myPage)
	myTheme := DB_GetGlobalConfigItem("Theme")
	ctx.View(myTheme + "/register.php")
}

func AE_WEB_ImportUI(ctx iris.Context) {
	myTheme := DB_GetGlobalConfigItem("Theme")
	ctx.View(myTheme + "/import.php")
}

func AE_WEB_ExportFromMnemonic(ctx iris.Context) {
	accountname := SESS_GetAccountName(ctx)
	db, err := sql.Open("sqlite", "./data/accounts/accounts.db")
	checkError(err)

	sql_account := "SELECT mnemonic FROM accounts WHERE account='" + accountname + "'"
	rows, err := db.Query(sql_account)
	checkError(err)

	for rows.Next() {
		//	needStore = false
	}

}

func AE_WEB_ImportFromMnemonic(ctx iris.Context) {
	password := ctx.FormValue("password")
	password_repeat := ctx.FormValue("password_repeat")
	mnemonic := ctx.FormValue("mnemonic")
	account_index, _ := strconv.ParseInt(ctx.FormValue("account_index"), 10, 32)
	address_index, _ := strconv.ParseInt(ctx.FormValue("address_index"), 10, 32)

	if (password == password_repeat) && len(password) > 1 {
		seed, err := account.ParseMnemonic(mnemonic)
		if err != nil {
			fmt.Println(err)
		}

		// Derive the subaccount m/44'/457'/3'/0'/1'
		key, err := account.DerivePathFromSeed(seed, uint32(account_index), uint32(address_index))
		if err != nil {
			fmt.Println(err)
		}

		// Deriving the aeternity Account from a BIP32 Key is a destructive process
		mykey, err := account.BIP32KeyToAeKey(key)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(mykey.Address)

		jks, err := account.KeystoreSeal(mykey, password)
		alias := "Import"
		//check the database
		db, err := sql.Open("sqlite", "./data/accounts/accounts.db")
		checkError(err)

		sql_account := "SELECT account,alias FROM accounts WHERE account='" + mykey.Address + "'"
		rows, err := db.Query(sql_account)
		checkError(err)

		needStore := true
		for rows.Next() {
			needStore = false
		}

		mnemonic = "" //Do not save the mnemonic words when import

		//Store the account and initial the enviroment:config,pubdata,privatedata and logs
		if needStore {
			sql_insert := "INSERT INTO accounts(account,alias,keystore,mnemonic) VALUES ('" + mykey.Address + "','" + alias + "','" + string(jks) + "','" + mnemonic + "')"
			db.Exec(sql_insert)
			db.Close()
			//Create database for each new account
			DB_InitDatabase(mykey.Address, alias)
		} else {
			ctx.HTML("<h1>Account Exist</h1>")
		}

		db.Close()

		ctx.Redirect("/")
	} else {
		ctx.HTML("<h1>Passwords must be the same.</h1>")
	}
}

func AE_WEB_DoRegister(ctx iris.Context) {
	password := ctx.FormValue("password")
	password_repeat := ctx.FormValue("password_repeat")
	alias := ctx.FormValue("alias")
	if (password == password_repeat) && len(password) > 1 {

		//Gnerate new account's mnemonic
		entropy, _ := bip39.NewEntropy(256)
		mnemonic, _ := bip39.NewMnemonic(entropy)
		seed, err := account.ParseMnemonic(mnemonic)

		// Derive the subaccount m/44'/457'/3'/0'/1'
		key, err := account.DerivePathFromSeed(seed, 0, 0)
		if err != nil {
			fmt.Println(err)
		}

		// Deriving the aeternity Account from a BIP32 Key is a destructive process
		mykey, err := account.BIP32KeyToAeKey(key)
		if err != nil {
			fmt.Println(err)
		}

		jks, err := account.KeystoreSeal(mykey, password)
		//fmt.Println(string(jks), alias)

		//check the database
		db, err := sql.Open("sqlite", "./data/accounts/accounts.db")
		checkError(err)

		sql_account := "SELECT account,alias FROM accounts WHERE account='" + mykey.Address + "'"
		rows, err := db.Query(sql_account)
		checkError(err)

		needStore := true
		for rows.Next() {
			needStore = false
		}

		mnemonic = SealMSGTo(mykey.Address, mnemonic, mykey) //Crypt mnemonic
		//Store the account and initial the enviroment:config,pubdata,privatedata and logs
		if needStore {
			sql_insert := "INSERT INTO accounts(account,alias,keystore,mnemonic) VALUES ('" + mykey.Address + "','" + alias + "','" + string(jks) + "','" + mnemonic + "')"
			db.Exec(sql_insert)
			db.Close()
			//Create database for each new account
			DB_InitDatabase(mykey.Address, alias)
		}

		ctx.Redirect("/")
	} else {
		ctx.HTML("<h1>Passwords must be the same.</h1>")
	}
}

func AE_WEB_LogOut(ctx iris.Context) {
	//globalAccount.Address = ""
	session := sess.Start(ctx)
	NodeOnline = false
	AE_WEB_loginoutFile()

	session.Destroy()
	ctx.Redirect("/")
}

func IPFS_killIPFS() {

	if ostype == "windows" {
		c := "TASKKILL /IM ipfs.exe /F"
		fmt.Println(c)
		cmd := exec.Command("cmd", "/c", c)
		output, err := cmd.Output()

		if err != nil {
			fmt.Printf("Execute Shell:%s failed with error:%s", c, err.Error())
			return
		}
		fmt.Printf("Execute Shell:%s finished with output:\n%s", c, string(output))
	} else {
		//kill ipfs firstly
		c := `killall ipfs`
		fmt.Println(c)
		cmd := exec.Command("sh", "-c", c)
		output, err := cmd.Output()

		if err != nil {
			fmt.Printf("Execute Shell:%s failed with error:%s", c, err.Error())
			return
		}
		fmt.Printf("Execute Shell:%s finished with output:\n%s", c, string(output))
	}
}

func AE_WEB_Index(ctx iris.Context) {
	//MyAENS = ""
	needReg := true
	NeedLogin := true
	AccountsLists := ""

	accountname := SESS_GetAccountName(ctx)
	MyIPFSConfig := IPFS_GetConfig("default")
	//Check if there is a logined account
	if len(accountname) > 6 {
		if !checkLogin(ctx) {
			return
		}

		needReg = false
		NeedLogin = false
		//dbpath := "./data/accounts/" + globalAccount.Address + "/public.db"
		//db, err := sql.Open("sqlite", dbpath)
		//checkError(err)
		//sql_index := "SELECT title FROM aek WHERE author='" + globalAccount.Address + "'"

		myPage := PageWallet{PageId: 23, Account: accountname, PageTitle: MyIPFSConfig.Identity.PeerID}
		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/index.php")

		err := qrcode.WriteFile(accountname, qrcode.Medium, 256, "./views/qr_ak.png")
		checkError(err)

	} else {
		//list accounts
		dbpath := "./data/accounts/accounts.db"
		if !FileExist(dbpath) {
			db, _ := sql.Open("sqlite", dbpath)
			sql_account := `
CREATE TABLE if not exists "accounts"(
"aid" INTEGER PRIMARY KEY AUTOINCREMENT,
"account" TEXT NULL,
"keystore" TEXT NULL,
"alias" TEXT NULL,
"mnemonic" TEXT NULL,
"lastlogin" INTEGER NULL,
"remark" TEXT NULL
);
`
			db.Exec(sql_account)
			db.Close()
		}
		db, err := sql.Open("sqlite", dbpath)
		checkError(err)

		sql_account := "SELECT account,alias FROM accounts ORDER by lastlogin desc"
		rows, err := db.Query(sql_account)
		checkError(err)

		for rows.Next() {
			var account string
			var alias string
			err = rows.Scan(&account, &alias)
			AccountsLists = AccountsLists + "<option value=" + account + ">" + alias + "(" + account + ")</option>\n"
			needReg = false
		}

		db.Close()

	}

	if needReg {
		var myPage PageReg
		myPage.PageTitle = "Registering Page"
		myPage.SubTitle = "Decentralized knowledge system without barrier."
		myPage.Register = "Register"

		//myPage.Lang = getPageString(getPageLang(r))

		//myPage = getPageString(getPageLang(r), "register")
		ctx.ViewData("", myPage)
		globalTheme := DB_GetGlobalConfigItem("Theme")
		ctx.View(globalTheme + "/register.php")

		//ctx.View("register.php")
	} else {
		if NeedLogin {
			var myoption template.HTML
			myoption = template.HTML(AccountsLists)
			myPage := PageLogin{Options: myoption}
			ctx.ViewData("", myPage)
			globalTheme := DB_GetGlobalConfigItem("Theme")
			ctx.View(globalTheme + "/login.php")
			//ctx.View("aeknow/login.php")
		}
	}

}

func AE_WEB_HomePage(ctx iris.Context) {
	//MyAENS = ""
	needReg := true

	AccountsLists := ""

	accountname := SESS_GetAccountName(ctx)
	MyIPFSConfig := IPFS_GetConfig("default")
	//Check if there is a logined account
	if len(accountname) > 6 {
		if !checkLogin(ctx) {
			return
		}

		needReg = false

		myPage := PageWallet{PageId: 23, Account: accountname, PageTitle: MyIPFSConfig.Identity.PeerID}
		ctx.ViewData("", myPage)

		myTheme := DB_GetConfigItem(accountname, "Theme")

		ctx.View(myTheme + "/dashboard.php")

		err := qrcode.WriteFile(accountname, qrcode.Medium, 256, "./views/qr_ak.png")
		checkError(err)

	} else {
		//list accounts
		dbpath := "./data/accounts/accounts.db"
		if !FileExist(dbpath) {
			db, _ := sql.Open("sqlite", dbpath)
			sql_account := `
CREATE TABLE if not exists "accounts"(
"aid" INTEGER PRIMARY KEY AUTOINCREMENT,
"account" TEXT NULL,
"keystore" TEXT NULL,
"alias" TEXT NULL,
"mnemonic" TEXT NULL,
"lastlogin" INTEGER NULL,
"remark" TEXT NULL
);
`
			db.Exec(sql_account)
			db.Close()
		}
		db, err := sql.Open("sqlite", dbpath)
		checkError(err)

		sql_account := "SELECT account,alias FROM accounts ORDER by lastlogin desc"
		rows, err := db.Query(sql_account)
		checkError(err)

		for rows.Next() {
			var account string
			var alias string
			err = rows.Scan(&account, &alias)
			AccountsLists = AccountsLists + "<option value=" + account + ">" + alias + "(" + account + ")</option>\n"
			needReg = false
		}

		db.Close()

	}

	if needReg {
		var myPage PageReg
		myPage.PageTitle = "Registering Page"
		myPage.SubTitle = "Decentralized knowledge system without barrier."
		myPage.Register = "Register"

		ctx.ViewData("", myPage)
		globalTheme := DB_GetGlobalConfigItem("Theme")
		ctx.View(globalTheme + "/register.php")
	}

}

func AE_WEB_CheckLogin(ctx iris.Context) {
	accountname := ctx.FormValue("accountname")
	password := ctx.FormValue("password")
	db, err := sql.Open("sqlite", "./data/accounts/accounts.db")
	checkError(err)

	sql_account := "SELECT keystore,alias FROM accounts WHERE account='" + accountname + "'"
	rows, err := db.Query(sql_account)
	checkError(err)

	var keystore string
	var alias string
	for rows.Next() {
		err = rows.Scan(&keystore, &alias)
		checkError(err)
	}

	//myAccount, err := account.LoadFromKeyStoreFile("data/accounts/"+accountname, password)
	myAccount, err := account.KeystoreOpen([]byte(keystore), password)
	//MyUsername := DB_GetAccountName(accountname)

	if err != nil {
		fmt.Println("Could not create myAccount's Account:", err)
		myPage := PageWallet{PageTitle: "Password error:Could not Read Account"}
		ctx.ViewData("", myPage)
		globalTheme := DB_GetGlobalConfigItem("Theme")
		ctx.View(globalTheme + "/error.php")

		//ctx.View("error.php")

	} else { //init the settings
		//globalAccount = *myAccount //作为呈现账号
		// Set user as authenticated
		session := sess.Start(ctx)
		session.Set("authenticated", true)
		session.Set("account", myAccount.Address)
		session.Set("password", password)
		//DB_GetConfigs(myAccount.Address)
		//NodeConfig = getConfigString() //读取节点设置
		//MyIPFSConfig = getIPFSConfig() //读取IPFS节点配置
		//MySiteConfig = getSiteConfig() //读取网站设置
		//lastIPFS = ""
		//signGlobalAccount = *myAccount
		//go bootIPFS()
		NodeOnline = true
		AE_WEB_loginedFile()
		//go IPFS_bootIPFS()
		//go ConnetDefaultNodes()
		//lastIPFS = DB_getLastIPFS(accountname)
		lastlogin := strconv.FormatInt(time.Now().Unix(), 10)
		sql_update := "UPDATE accounts SET lastlogin=" + lastlogin + " WHERE account='" + myAccount.Address + "'"
		//fmt.Println(sql_update)
		db.Exec(sql_update)

		go PubSub_StartListen(myAccount.Address, *myAccount)
	}
	db.Close()

	ctx.Redirect("/")

}

func SESS_GetAccount(ctx iris.Context) *account.Account {
	session := sess.Start(ctx)
	password := session.GetString("password")
	accountname := session.GetString("account")

	db, err := sql.Open("sqlite", "./data/accounts/accounts.db")
	checkError(err)

	sql_account := "SELECT keystore FROM accounts WHERE account='" + accountname + "'"
	rows, err := db.Query(sql_account)
	checkError(err)

	var keystore string
	for rows.Next() {
		err = rows.Scan(&keystore)
		checkError(err)
	}

	myAccount, err := account.KeystoreOpen([]byte(keystore), password)

	if err != nil {
		fmt.Println("Could not create myAccount's Account:", err)
		myPage := PageWallet{PageTitle: "Password error:Could not Read Account"}
		ctx.ViewData("", myPage)
		globalTheme := DB_GetGlobalConfigItem("Theme")
		ctx.View(globalTheme + "/error.php")
		//ctx.View("error.php")

	}

	db.Close()

	return myAccount

}

func SESS_GetAccountName(ctx iris.Context) string {
	session := sess.Start(ctx)
	return session.GetString("account")
}

func AE_WEB_loginedFile() {
	loginedFile := ""
	if ostype == "windows" {
		loginedFile = ".\\data\\online.lock"
	} else {
		loginedFile = "./data/online.lock"
	}

	if FileExist(loginedFile) {
	} else {
		err := ioutil.WriteFile(loginedFile, []byte("ONLINE"), 0644)
		if err != nil {
			panic(err)
		}
	}
}

func AE_WEB_loginoutFile() {
	loginedFile := ""
	if ostype == "windows" {
		loginedFile = ".\\data\\online.lock"
	} else {
		loginedFile = "./data/online.lock"
	}

	if FileExist(loginedFile) {
		err := os.Remove(loginedFile)

		if err != nil {
			// 删除失败
			fmt.Println("logout failed")

		} else {
			// 删除成功
			fmt.Println("logout")
		}
	}
}

func AE_CheckActiveNode() {
	//Check active node from pubsub and other ways, recorded into the database

}
func IPFS_bootIPFS(repo string) { //boot IPFS independently
	NeedBoot := true

	IPFS_checkRepo(repo)

	sh := shell.NewShell("127.0.0.1:5001")
	cid, err := sh.Add(strings.NewReader("Hello from AEKs!"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		//os.Exit(1)
	} else {
		fmt.Println("IPFS has booted, abort=>" + cid)
		NeedBoot = false
	}

	if NeedBoot {
		if ostype == "windows" {
			fileExec := ".\\bin\\ipfs.exe"
			//c := "set IPFS_PATH=data\\site\\" + globalAccount.Address + "\\repo\\&& " + fileExec + " daemon --enable-pubsub-experiment"
			c := "set IPFS_PATH=data\\repo\\" + repo + "&& " + fileExec + " daemon --enable-pubsub-experiment --enable-namesys-pubsub"
			fmt.Println(c)
			cmd := exec.Command("cmd", "/c", c)
			out, _ := cmd.Output()
			fmt.Println(string(out))

		} else {
			fileExec := "./bin/ipfs"

			//c := "export IPFS_PATH=./data/site/" + globalAccount.Address + "/repo/&& " + fileExec + " daemon --enable-pubsub-experiment"
			c := "export IPFS_PATH=./data/repo/" + repo + "&& " + fileExec + " daemon --enable-pubsub-experiment --enable-namesys-pubsub"
			cmd := exec.Command("sh", "-c", c)
			fmt.Println(c)
			out, _ := cmd.Output()
			fmt.Println(string(out))

		}
	}
}

func IPFS_GetConfig(repo string) IPFSConfig {
	configFilePath := "./data/repo/" + repo + "/config"
	_, err := os.Stat(configFilePath)
	checkError(err)
	JsonParse := NewJsonStruct()
	readConfigfile := IPFSConfig{}
	JsonParse.Load(configFilePath, &readConfigfile)

	return readConfigfile

}

func IPFS_UpdateIPNS(accountname, domain string) string {
	dbfile := "./data/accounts/" + accountname + "/public.db"
	sh := shell.NewShell("127.0.0.1:5001")
	pubfile, err := os.Open(dbfile)
	cid, err := sh.Add(pubfile)

	defer pubfile.Close()
	//fmt.Println(dbfile)

	if err != nil {
		fmt.Println("Failed IPFS ADD DB")
		return "DB failed"
	}

	//MySiteConfig := DB_GetSiteConfigs(accountname)

	//lastTenArticle := base64.StdEncoding.EncodeToString([]byte(DB_GetLastTenArticle(accountname)))

	//siteConfigStr := "{\"Title\":\"" + MySiteConfig.Title + "\",\"Subtitle\":\"" + MySiteConfig.Subtitle + "\",\"Description\":\"" + MySiteConfig.Description + "\",\"Author\":\"" + MySiteConfig.Author + "\",\"AuthorDescription\":\"" + MySiteConfig.AuthorDescription + "\",\"Theme\":\"" + MySiteConfig.Theme + "\",\"AENS\":\"" + MySiteConfig.AENS + "\"," + "\"lastten\":\"" + lastTenArticle + "\"}"
	//metehash, err := sh.Add(bytes.NewBufferString(siteConfigStr))

	//MyNodeConfig := DB_GetConfigs()
	node := naet.NewNode(MyNodeConfig.PublicNode, false)
	block_height, err := node.GetHeight()

	//if err != nil {
	//	fmt.Println("Failed IPFS ADD Meta")
	//	return "Meta Failed"
	//}

	//fmt.Println(siteConfigStr + "=>" + metehash)

	sitefile := "./data/sites.json"
	if FileExist(sitefile) {
		JsonParse := NewJsonStruct()
		reposites := Reposite{}
		JsonParse.Load(sitefile, &reposites)
		if domain == "" {
			domain = accountname
		}
		//fmt.Println(reposites)
		NeedAdd := true
		jsonstr := "{\"reposites\":["

		for i := 0; i < len(reposites.Reposites); i++ {
			if reposites.Reposites[i].Name == domain {
				jsonstr = jsonstr + "{\"name\":\"" + domain + "\",\"hash\":\"" + cid + "\",\"metainfo\":\"" + strconv.FormatUint(block_height, 10) + "\"},"
				//reposites.Reposites[i].Hash = cid
				NeedAdd = false
			} else {
				jsonstr = jsonstr + "{\"name\":\"" + reposites.Reposites[i].Name + "\",\"hash\":\"" + reposites.Reposites[i].Hash + "\",\"metainfo\":\"" + reposites.Reposites[i].Metainfo + "\"},"
			}
		}

		fmt.Println("ORG=>" + jsonstr)

		if NeedAdd {
			jsonstr = jsonstr + "{\"name\":\"" + domain + "\",\"hash\":\"" + cid + "\",\"metainfo\":\"" + strconv.FormatUint(block_height, 10) + "\"},"
		}

		jsonstr = jsonstr + "]}END"
		jsonstr = strings.Replace(jsonstr, ",]}END", "]}", -1)

		err = ioutil.WriteFile(sitefile, []byte(jsonstr), 0666)

		checkError(err)
		fmt.Println(jsonstr)
		//fmt.Println(reposites.Sites[1].Name)

	} else {
		fmt.Println("No config file")
	}
	pubfile, err = os.Open(sitefile)
	cid, err = sh.Add(pubfile)

	HashForIPNS := "/ipfs/" + cid
	fmt.Println("Ready to pulish IPNS..", HashForIPNS)

	resp, err := sh.PublishWithDetails(HashForIPNS, "", 0, 0, false)
	if err != nil {
		fmt.Print("Failed IPNS Pub")
		return "BAD IPNS"
	}

	fmt.Println(resp)

	return cid
}

func AE_WEB_MakeTranscaction(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)

	sender_id := ctx.FormValue("sender_id")
	recipient_id := ctx.FormValue("recipient_id")
	amount := ctx.FormValue("amount")
	payload := ctx.FormValue("payload")
	password := ctx.FormValue("password")

	famount, err := strconv.ParseFloat(amount, 64)
	bigfloatAmount := big.NewFloat(famount)
	imultiple := big.NewFloat(1000000000000000000) //18 dec
	fmyamount := big.NewFloat(1)
	fmyamount.Mul(bigfloatAmount, imultiple)

	myamount := new(big.Int)
	fmyamount.Int(myamount)
	//transfer tokens to .chain name
	if strings.Index(recipient_id, ".chain") > -1 {
		recipient_id = AENS_getAccountFromAENS(recipient_id, accountname)
	}

	db, err := sql.Open("sqlite", "./data/accounts/accounts.db")
	checkError(err)

	sql_account := "SELECT keystore FROM accounts WHERE account='" + sender_id + "'"
	rows, err := db.Query(sql_account)
	checkError(err)

	var keystore string
	for rows.Next() {
		err = rows.Scan(&keystore)
		checkError(err)
	}

	MyAccount, err := account.KeystoreOpen([]byte(keystore), password)

	if err != nil {

		myPage := PageWallet{PageId: 23, Account: accountname, PageTitle: "Password error"}
		ctx.ViewData("", myPage)
		globalTheme := DB_GetGlobalConfigItem("Theme")
		ctx.View(globalTheme + "/error.php")
		//ctx.View("error.php")
		return
	}

	// create a connection to a node, represented by *Node
	//MyNodeConfig := DB_GetConfigs()
	node := naet.NewNode(MyNodeConfig.PublicNode, false)

	// create the closures that autofill the correct account nonce and transaction TTL
	ttlnoncer := transactions.NewTTLNoncer(node)

	// create the SpendTransaction

	tx, err := transactions.NewSpendTx(MyAccount.Address, recipient_id, myamount, []byte(payload), ttlnoncer)
	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	} else {
		fmt.Println(tx)
	}

	//_, myTxhash, _, _, _, err := SignBroadcastWaitTransaction(tx, alice, node, aeconfig.Node.NetworkID, 10)
	_, myTxhash, _, err := SignBroadcastTransaction(tx, MyAccount, node, aeconfig.Node.NetworkID)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)

		myPage := PageWallet{PageId: 23, Account: accountname, PageTitle: "Failed"}
		ctx.ViewData("", myPage)
		ctx.View(myTheme + "/transaction.php")
		//ctx.View("transaction.php")
	} else {
		myPage := PageWallet{PageId: 23, Account: accountname, PageTitle: myTxhash}
		ctx.ViewData("", myPage)

		ctx.View(myTheme + "/transaction.php")
		//ctx.View("transaction.php")
	}
}

func AE_WEB_Wallet(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	needReg := true
	ak := ""
	AccountsLists := ""
	recipient_id := ""
	payload := ""
	amountstr := ""

	recipient_id = ctx.URLParam("recipient_id")
	//payloadByte = ctx.URLParam("payload")
	payloadByte, _ := base64.StdEncoding.DecodeString(ctx.URLParam("payload"))
	payload = string(payloadByte)

	amountstr = ctx.URLParam("amount")
	//MyNodeConfig := DB_GetConfigs()
	node := naet.NewNode(MyNodeConfig.PublicNode, false)

	akBalance, err := node.GetAccount(accountname)
	var thisamount string
	var myNonce uint64
	if err != nil {
		fmt.Println("Account not exist.")
		thisamount = "0"
		myNonce = 0
	} else {
		bigstr := akBalance.Balance.String()
		myBalance := ToBigFloat(bigstr)
		imultiple := big.NewFloat(0.000000000000000001) //18 dec
		thisamount = new(big.Float).Mul(myBalance, imultiple).String()
		myNonce = *akBalance.Nonce

	}

	merr := filepath.Walk("data/accounts/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "ak_") {

			ak = filepath.Base(path)
			if len(ak) > 0 {
				AccountsLists = AccountsLists + "<option>" + ak + "</option>\n"
			}

			needReg = false
		}
		//fmt.Println(path)
		return nil
	})
	//fmt.Println("address:" + globalAccount.Address)
	if len(accountname) > 1 {
		needReg = false
		ak := accountname

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Wallet", Balance: thisamount, Nonce: myNonce, Recipient_id: recipient_id, Amount: amountstr, Payload: payload}
		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/wallet.php")
		//ctx.View("wallet.php")

		err := qrcode.WriteFile(ak, qrcode.Medium, 256, "./views/qr_ak.png")
		err = qrcode.WriteFile("https://www.aeknow.org/v2/accounts/"+ak, qrcode.Medium, 256, "./views/qr_account.png")
		if err != nil {
			fmt.Println("write error")
		}
	} else {

		var myoption template.HTML
		myoption = template.HTML(AccountsLists)
		myPage := PageLogin{Options: myoption}
		ctx.ViewData("", myPage)
		myTheme := DB_GetGlobalConfigItem("Theme")
		ctx.View(myTheme + "/login.php")
		//ctx.View("login.php")
	}

	if merr != nil {
		fmt.Println("error")
	}

	if needReg {

		var myPage PageReg
		myPage.PageTitle = "Registering Page"
		myPage.SubTitle = "Decentralized knowledge system without barrier."
		myPage.Register = "Register"

		myPage.Lang = getPageString(getPageLang(ctx.Request()))

		ctx.ViewData("", myPage)
		myTheme := DB_GetGlobalConfigItem("Theme")
		ctx.View(myTheme + "/register.php")
		//ctx.View("register.php")
	}
}

//Simple version login check for local user
func checkLogin(ctx iris.Context) bool {
	if len(SESS_GetAccountName(ctx)) > 6 {
		return true
	}

	return false
}

func IPFS_checkRepo(RepoName string) {
	IPFSCheck := "./data/repo/" + RepoName + "/version"

	if !FileExist(IPFSCheck) {
		if ostype == "windows" {
			IPFS_PATH := "data\\repo\\" + RepoName
			c := "mkdir " + IPFS_PATH + "&& set IPFS_PATH=" + IPFS_PATH + "\\&& bin\\ipfs.exe init"
			fmt.Println(c)
			cmd := exec.Command("cmd", "/c", c)
			out, _ := cmd.Output()

			fmt.Println(string(out))
		} else {
			IPFS_PATH := "./data/repo/" + RepoName
			c := "mkdir " + IPFS_PATH + "&& export IPFS_PATH=" + IPFS_PATH + "/&& ./bin/ipfs init"
			fmt.Println(c)
			cmd := exec.Command("sh", "-c", c)
			out, _ := cmd.Output()

			fmt.Println(string(out))

		}
	}

}

func DB_InitDatabase(pubkey, name string) {
	dbpathDir := "./data/accounts/" + pubkey

	if !FileExist(dbpathDir) {
		err := os.Mkdir(dbpathDir, os.ModePerm)
		checkError(err)
	}

	dbpath := "file:./data/accounts/" + pubkey + "/public.db?auto_vacuum=1"
	db, _ := sql.Open("sqlite", dbpath)
	//Create main data table for articles
	sql_main := `
CREATE TABLE if not exists "aek"(
"aid" INTEGER PRIMARY KEY AUTOINCREMENT,
"title" TEXT NULL,
"author" TEXT NULL,
"authorname" TEXT NULL,
"hash" TEXT NULL,
"abstract" TEXT NULL,
"keywords" TEXT NULL,
"tags" TEXT NULL,
"body" TEXT NULL,
"jsonstr" TEXT NULL,
"signature" TEXT NULL,
"filetype" TEXT NULL,
"filesize" INTEGER NULL,
"pubtime" INTEGER NULL,
"lastmodtime" INTEGER NULL,
"lasthash" TEXT NULL,
"remark" TEXT NULL
);
`
	db.Exec(sql_main)
	//Create author's info table
	sql_site := `
CREATE TABLE if not exists "config"(
"item" text NULL,
"value" TEXT NULL,
"remark" TEXT NULL
);
`

	db.Exec(sql_site)
	sql_init := `INSERT INTO config(item,value) VALUES('pubkey','` + pubkey + `');`
	db.Exec(sql_init)
	sql_init = `INSERT INTO config(item,value) VALUES('name','` + name + `');`
	db.Exec(sql_init)
	db.Close()

	//create private database
	dbpath = "file:./data/accounts/" + pubkey + "/private.db?auto_vacuum=1"
	db, _ = sql.Open("sqlite", dbpath)
	sql_main = `
CREATE TABLE if not exists "aek"(
"aid" INTEGER PRIMARY KEY AUTOINCREMENT,
"title" TEXT NULL,
"author" TEXT NULL,
"authorname" TEXT NULL,
"hash" TEXT NULL,
"abstract" TEXT NULL,
"keywords" TEXT NULL,
"tags" TEXT NULL,
"body" TEXT NULL,
"jsonstr" TEXT NULL,
"signature" TEXT NULL,
"filetype" TEXT NULL,
"filesize" INTEGER NULL,
"pubtime" INTEGER NULL,
"lasthash" TEXT NULL,
"lastmodtime" INTEGER NULL,
"remark" TEXT NULL
);
`
	db.Exec(sql_main)
	//Create author's info table
	sql_site = `
CREATE TABLE if not exists "config"(
"item" text NULL,
"value" TEXT NULL,
"remark" TEXT NULL
);
`

	db.Exec(sql_site)
	sql_init = `INSERT INTO config(item,value) VALUES('pubkey','` + pubkey + `');`
	db.Exec(sql_init)
	sql_init = `INSERT INTO config(item,value) VALUES('name','` + name + `');`
	db.Exec(sql_init)
	db.Close()

	//create config database
	dbpath = "file:./data/accounts/" + pubkey + "/config.db?auto_vacuum=1"
	db, _ = sql.Open("sqlite", dbpath)

	sql_table := `
CREATE TABLE if not exists "config"(
"item" text NULL,
"value" TEXT NULL,
"remark" TEXT NULL
);
`
	db.Exec(sql_table)
	//Default settings
	sql_init = `INSERT INTO config(item,value) VALUES('PublicNode','http://52.220.198.72:3013');`
	db.Exec(sql_init)

	sql_init = `INSERT INTO config(item,value) VALUES('APINode','https://www.aeknow.org');`
	db.Exec(sql_init)

	sql_init = `INSERT INTO config(item,value) VALUES('IPFSNode','http://127.0.0.1:8080');`
	db.Exec(sql_init)

	sql_init = `INSERT INTO config(item,value) VALUES('IPFSAPI','http://127.0.0.1:5001');`
	db.Exec(sql_init)

	sql_init = `INSERT INTO config(item,value) VALUES('LocalWeb','http://127.0.0.1:8888');`
	db.Exec(sql_init)

	sql_init = `INSERT INTO config(item,value) VALUES('Theme','aeknow');`
	db.Exec(sql_init)

	db.Close()

	//create logs database
	dbpath = "file:./data/accounts/" + pubkey + "/logs.db?auto_vacuum=1"
	db, _ = sql.Open("sqlite", dbpath)
	//Create log data table for articles
	sql_main = `
CREATE TABLE if not exists "logs"(
"aid" INTEGER PRIMARY KEY AUTOINCREMENT,
"title" TEXT NULL,
"author" TEXT NULL,
"authorname" TEXT NULL,
"hash" TEXT NULL,
"abstract" TEXT NULL,
"keywords" TEXT NULL,
"tags" TEXT NULL,
"body" TEXT NULL,
"jsonstr" TEXT NULL,
"signature" TEXT NULL,
"filetype" TEXT NULL,
"filesize" INTEGER NULL,
"pubtime" INTEGER NULL,
"lasthash" TEXT NULL,
"lastmodtime" INTEGER NULL,
"remark" TEXT NULL
);
`
	db.Exec(sql_main)

	//create indexs for search
	sql_index := `create index bodyindex on aek(body);`
	db.Exec(sql_index)

	sql_index = `create index titleindex on aek(title);`
	db.Exec(sql_index)

	sql_index = `create index keywordsindex on aek(keywords);`
	db.Exec(sql_index)

	sql_index = `create index authorindex on aek(author);`
	db.Exec(sql_index)

	db.Close()

	//Create FTS5 index data for all datas
	dbpath = "file:./data/accounts/" + pubkey + "/index.db?auto_vacuum=1"
	db, _ = sql.Open("sqlite", dbpath)
	sql_index = `CREATE VIRTUAL TABLE pages USING fts5(title, author, authorname,keywords,abstract, body,source,id,hash);`
	db.Exec(sql_index)

	//create chaet database
	dbpath = "file:./data/accounts/" + pubkey + "/chaet.db?auto_vacuum=1"
	db, _ = sql.Open("sqlite", dbpath)
	//Create main data table for messages,body for full text search
	sql_msg := `CREATE VIRTUAL TABLE msgs USING fts5(fromid, toid, body,raw,mtype, pubtime);`
	db.Exec(sql_msg)

	sql_user := `
CREATE TABLE if not exists "users"(
"uid" INTEGER PRIMARY KEY AUTOINCREMENT,
"username" TEXT NULL,
"id" TEXT NULL,
"avatar" TEXT NULL,
"sign" TEXT NULL,
"lastactive" TEXT NULL,
"isfriend" bool NULL,
"status" TEXT NULL,
"groupname" TEXT NULL,
"aens" TEXT NULL,
"description" TEXT NULL,
"remark" TEXT NULL
);
`
	db.Exec(sql_user)

	sql_group := `
CREATE TABLE if not exists "groups"(
"gid" INTEGER PRIMARY KEY AUTOINCREMENT,
"groupname" TEXT NULL,
"id" TEXT NULL,
"groupkey" TEXT NULL,
"avatar" TEXT NULL,
"sign" TEXT NULL,
"status" TEXT NULL,
"members" TEXT NULL,
"aens" TEXT NULL,
"description" TEXT NULL,
"remark" TEXT NULL
);
`
	db.Exec(sql_group)
	db.Close()

}

func DB_UpdateConfigs(pubkey, item, value string) {
	//update or insert configs
	dbpath := "./data/accounts/" + pubkey + "/config.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)
	sql_check := "SELECT value FROM config WHERE item='" + item + "'"
	rows, err := db.Query(sql_check)
	checkError(err)

	NeedInsert := true
	for rows.Next() {
		NeedInsert = false
	}

	if NeedInsert {
		sql_insert := "INSERT INTO config(item,value) VALUES('" + item + "','" + value + "')"
		db.Exec(sql_insert)
	} else {
		sql_update := "UPDATE config set value='" + value + "' WHERE item='" + item + "'"
		db.Exec(sql_update)
	}

	db.Close()
}

func DB_SaveItemToDB(dbpath, table, item, value, remark string) {
	db, _ := sql.Open("sqlite", dbpath)
	//fmt.Println(dbpath)
	sql_query := "SELECT item FROM " + table + " WHERE item='" + item + "';"
	//fmt.Println(sql_query)
	rows, err := db.Query(sql_query)
	checkError(err)

	ItemCout := 0
	for rows.Next() {
		ItemCout = ItemCout + 1
	}

	if ItemCout > 0 {
		sql_query = "UPDATE " + table + " SET value='" + value + "',remark='" + remark + "' WHERE item='" + item + "';"
		//fmt.Println(sql_query)
	} else {
		sql_query = "INSERT INTO " + table + "(item,value,remark) VALUES('" + item + "','" + value + "','" + remark + "');"
		//fmt.Println(sql_query)
	}

	_, err = db.Exec(sql_query)
	checkError(err)
	db.Close()

}

func DB_GetSiteConfigs(pubkey string) SiteConfig {
	dbpath := "./data/accounts/" + pubkey + "/config.db"
	db, _ := sql.Open("sqlite", dbpath)
	sql_query := "SELECT item, value FROM config"
	rows, err := db.Query(sql_query)
	checkError(err)

	var GetConfig SiteConfig

	GetConfig.Author = pubkey
	GetConfig.AuthorDescription = pubkey
	GetConfig.Description = "This is the default new site, ready to build my knowledge base!"
	GetConfig.Title = "New Start"
	GetConfig.Subtitle = "A new step to the knowledge blockchain."
	GetConfig.Pubkey = pubkey

	for rows.Next() {
		var item string
		var value string
		err = rows.Scan(&item, &value)
		switch item {
		//site config
		case "name":
			GetConfig.Author = value
		case "Title":
			GetConfig.Title = value
		case "AuthorDescription":
			GetConfig.AuthorDescription = value
		case "Description":
			GetConfig.Description = value
		case "Subtitle":
			GetConfig.Subtitle = value
		case "AENS":
			GetConfig.AENS = value
		case "Pubkey":
			GetConfig.Pubkey = value
		default:
			//fmt.Println("No Such item: " + item)
		}

	}

	db.Close()

	return GetConfig
	//NodeConfig = getConfigString() //读取节点设置
	//MyIPFSConfig = getIPFSConfig() //读取IPFS节点配置
	//MySiteConfig = getSiteConfig() //读取网站设置
}

func DB_GetConfigs() AeknowConfig {
	//dbpath := "./data/accounts/" + pubkey + "/config.db"
	//if pubkey == "ak_fCCw1JEkvXdztZxk8FRGNAkvmArhVeow89e64yX4AxbCPrVh5" { //default
	dbpath := "./data/config.db"
	//}
	db, _ := sql.Open("sqlite", dbpath)
	sql_query := "SELECT item, value FROM config"
	rows, err := db.Query(sql_query)
	checkError(err)
	var NodeConfig AeknowConfig
	for rows.Next() {
		var item string
		var value string
		err = rows.Scan(&item, &value)
		switch item {
		//network config
		case "PublicNode":
			NodeConfig.PublicNode = value
		case "APINode":
			NodeConfig.APINode = value
		case "IPFSNode":
			NodeConfig.IPFSNode = value
		case "IPFSAPI":
			NodeConfig.IPFSAPI = value
		case "LocalWeb":
			NodeConfig.LocalWeb = value
		case "PubsubProxy":
			NodeConfig.PubsubProxy = value
		//case "LastIPFS":
		//lastIPFS = value
		default:
			//fmt.Println("No Such item: " + item)
		}

	}

	db.Close()
	//NodeConfig = getConfigString() //读取节点设置
	//MyIPFSConfig = getIPFSConfig() //读取IPFS节点配置
	//MySiteConfig = getSiteConfig() //读取网站设置
	return NodeConfig
}

func DB_GetAccountName(pubkey string) string {
	db, err := sql.Open("sqlite", "./data/accounts/accounts.db")
	checkError(err)

	sql_account := "SELECT alias FROM accounts WHERE account='" + pubkey + "'"
	rows, err := db.Query(sql_account)
	checkError(err)

	alias := ""
	for rows.Next() {
		err = rows.Scan(&alias)
		checkError(err)
	}
	db.Close()

	return alias
}

func DB_UpdateAccountName(pubkey, alias string) {
	db, err := sql.Open("sqlite", "./data/accounts/accounts.db")
	checkError(err)

	sql_account := "UPDATE accounts SET alias='" + alias + "' WHERE account='" + pubkey + "'"
	db.Exec(sql_account)
	db.Close()
}

func DB_GetConfigItem(accountname, item string) string {
	dbpath := "./data/accounts/" + accountname + "/config.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)

	value := "NULL"
	if FileExist(dbpath) {
		sql_query := "SELECT value FROM config WHERE item='" + item + "'"

		rows, err := db.Query(sql_query)
		checkError(err)

		for rows.Next() {
			err = rows.Scan(&value)
		}
	}
	db.Close()
	return value
}

func DB_GetPublicItem(accountname, item string) string {
	dbpath := "./data/accounts/" + accountname + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)

	value := "NULL"
	if FileExist(dbpath) {
		sql_query := "SELECT value FROM config WHERE item='" + item + "'"

		rows, err := db.Query(sql_query)
		checkError(err)

		for rows.Next() {
			err = rows.Scan(&value)
		}
	}
	db.Close()
	return value
}

func DB_UpdatePublicConfigItem(accountname, item, value, remark string) {
	dbpath := "./data/accounts/" + accountname + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	sql_query := "SELECT item FROM config WHERE item='" + item + "';"
	//fmt.Println(sql_query)
	rows, err := db.Query(sql_query)
	checkError(err)

	ItemCout := 0
	for rows.Next() {
		ItemCout = ItemCout + 1
	}

	if ItemCout > 0 {
		sql_query = "UPDATE config SET value='" + value + "',remark='" + remark + "' WHERE item='" + item + "';"
		//fmt.Println(sql_query)
	} else {
		sql_query = "INSERT INTO config(item,value,remark) VALUES('" + item + "','" + value + "','" + remark + "');"
		//fmt.Println(sql_query)
	}

	_, err = db.Exec(sql_query)
	checkError(err)
	db.Close()
}

func DB_GetLastTenArticle(accountname string) string {
	dbpath := "./data/accounts/" + accountname + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)

	value := ""
	title := ""
	hash := ""
	if FileExist(dbpath) {
		sql_query := "SELECT title,hash FROM aek ORDER BY pubtime DESC LIMIT 10"

		rows, err := db.Query(sql_query)
		checkError(err)

		for rows.Next() {
			err = rows.Scan(&title, &hash)
			value = value + title + "::$::" + hash + "\n"
		}
	}
	db.Close()
	return value
}

func DB_GetGlobalConfigItem(item string) string {
	dbpath := "./data/config.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)

	value := "NULL"
	if FileExist(dbpath) {
		sql_query := "SELECT value FROM config WHERE item='" + item + "'"

		rows, err := db.Query(sql_query)
		checkError(err)

		for rows.Next() {
			err = rows.Scan(&value)
		}
	}
	db.Close()
	return value
}

///////////////////////////////////////////AENS //////////////////////////////////////////////////
type PageUpdateAENS struct {
	AENSName        string
	NameID          string
	PointerJson     template.HTML
	NameJson        template.HTML
	NameTTL         uint64
	Account         string
	Balance         string
	Nonce           uint64
	AEAddress       string
	IPFSAddress     string
	IPNSAddress     string
	ContractAddress string
	OracleAddress   string
	BTCAddress      string
	ETHAddress      string
	EmailAddress    string
	WebAddress      string
}

func AENS_WEB_ExpertDoUpdateAENS(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	myPointerJson := ctx.FormValue("pointerjson")
	aensname := ctx.FormValue("aensname")
	ak := accountname

	//fmt.Println(myPointerJson)

	var s []*transactions.NamePointer

	err := json.Unmarshal([]byte(myPointerJson), &s)
	if err != nil {
		fmt.Println(err)
	}
	//MyNodeConfig := DB_GetConfigs()
	node := naet.NewNode(MyNodeConfig.PublicNode, false)
	ttlnoncer := transactions.NewTTLNoncer(node)

	tx, err := transactions.NewNameUpdateTx(accountname, aensname, s, 50000, ttlnoncer)

	//fmt.Println(tx)

	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	}
	MysignAccount := SESS_GetAccount(ctx)
	_, myTxhash, _, err := SignBroadcastTransaction(tx, MysignAccount, node, aeconfig.Node.NetworkID)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Failed"}
		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/transaction.php")
		//ctx.View("transaction.php")
	} else {
		//fmt.Println("TxHash:" + myTxhash)

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: myTxhash}
		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/transaction.php")
		//ctx.View("transaction.php")
	}
}

func AENS_WEB_DoUpdateAENS(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)

	MyIPFSConfig := IPFS_GetConfig("default")

	aensname := ctx.FormValue("aensname")
	aeaddress := ctx.FormValue("aeaddress")
	ipfsaddress := ctx.FormValue("ipfsaddress")
	//TODO:Update IPFS every time??
	ipnsaddress := ctx.FormValue("ipnsaddress")

	contractaddress := ctx.FormValue("contractaddress")
	oracleaddress := ctx.FormValue("oracleaddress")
	btcaddress := ctx.FormValue("btcaddress")
	ethaddress := ctx.FormValue("ethaddress")
	emailaddress := ctx.FormValue("emailaddress")
	webaddress := ctx.FormValue("webaddress")

	//fmt.Println(aensname)

	myPointerJson := "["
	if strings.TrimSpace(aeaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"` + aeaddress + `","key":"account_pubkey"},`
	} else {
		myPointerJson = myPointerJson + `{"id":"` + accountname + `","key":"account_pubkey"},`
	}

	if strings.TrimSpace(ipfsaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"ch_ipfsD1iUfRLdnJjQMEczjSzzphPbNnSQudnqUAe1vPJetmMK9","key":"` + ipfsaddress + `"},`
	} else {
		//NOT Always Update IPFS
		//if DB_GetConfigItem(accountname, "LastIPFS") != "NULL" {
		//	myPointerJson = myPointerJson + `{"id":"ch_ipfsD1iUfRLdnJjQMEczjSzzphPbNnSQudnqUAe1vPJetmMK9","key":"` + DB_GetConfigItem(accountname, "LastIPFS") + `"},`
		//}
	}

	if strings.TrimSpace(ipnsaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"ch_ipnsoMiJmYq1joKGXFtLRDrSJ3mUjapNB7gcPud7mmpVUXssM","key":"` + ipnsaddress + `"},`
	} else {
		if MyIPFSConfig.Identity.PeerID != "" {
			myPointerJson = myPointerJson + `{"id":"ch_ipnsoMiJmYq1joKGXFtLRDrSJ3mUjapNB7gcPud7mmpVUXssM","key":"` + MyIPFSConfig.Identity.PeerID + `"},`
		}
	}

	if strings.TrimSpace(contractaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"` + contractaddress + `","key":"contract_pubkey"},`
	}
	if strings.TrimSpace(oracleaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"` + oracleaddress + `","key":"oracle_pubkey"},`
	}
	if strings.TrimSpace(btcaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"ch_btcqM2NycfJaeLYhYY9uPGKj98iVkwL9VLw7ZP5WzzWHHj2sP","key":"` + btcaddress + `"},`
	}
	if strings.TrimSpace(ethaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"ch_ethe795mCkWMAkguuc3ay9k2JSMikZ61L6VfEMDrujEwCiaiB","key":"` + ethaddress + `"},`
	}

	if strings.TrimSpace(emailaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"ch_em3io3Ntov4qJ1y9mDoyQgHTaWBnBZd1CBu7wnH6iyuF5jf5m","key":"` + emailaddress + `"},`
	}
	if strings.TrimSpace(webaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"ch_webcVNwKZujeYcxDMjAH5ZUPNwCdcFL4QgYD34pFHZi6KEnzS","key":"` + webaddress + `"},`
	}

	myPointerJson = myPointerJson + "]"
	myPointerJson = strings.Replace(myPointerJson, ",]", "]", -1)
	fmt.Println(myPointerJson)
	var s []*transactions.NamePointer

	err := json.Unmarshal([]byte(myPointerJson), &s)
	if err != nil {
		fmt.Println(err)
	}
	//MyNodeConfig := DB_GetConfigs()
	node := naet.NewNode(MyNodeConfig.PublicNode, false)
	ttlnoncer := transactions.NewTTLNoncer(node)
	//p := s
	tx, err := transactions.NewNameUpdateTx(accountname, aensname, s, 180000, ttlnoncer)

	fmt.Println(tx)

	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	}

	MysignAccount := SESS_GetAccount(ctx)
	_, myTxhash, _, err := SignBroadcastTransaction(tx, MysignAccount, node, aeconfig.Node.NetworkID)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)
		myPage := PageWallet{PageId: 23, Account: accountname, PageTitle: "Failed"}
		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/transaction.php")
		//ctx.View("transaction.php")
	} else {

		myPage := PageWallet{PageId: 23, Account: accountname, PageTitle: myTxhash}

		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/transaction.php")
		//ctx.View("transaction.php")
	}
}

func AENS_UpdateAENS(aensname, aensitem, aenspointer, accountname string, ctx iris.Context) {
	//MyNodeConfig := DB_GetConfigs()
	myurl := MyNodeConfig.PublicNode + "/v2/names/" + aensname
	str := httpGet(myurl)
	//fmt.Println(myurl)
	var s AENSInfo
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(s.Pointers); i++ {
		if s.Pointers[i].ID == aensitem {
			s.Pointers[i].Key = aenspointer
		}
	}

	myPointerJson := "["
	for i := 0; i < len(s.Pointers); i++ {
		myPointerJson = myPointerJson + `{"id":"` + s.Pointers[i].ID + `","key":"` + s.Pointers[i].Key + `"},`
	}

	myPointerJson = myPointerJson + "]"
	myPointerJson = strings.Replace(myPointerJson, ",]", "]", -1)

	if len(s.Pointers) == 0 {
		myPointerJson = `[{"id":"` + aensitem + `","key":"` + aenspointer + `"}]`
	}

	//fmt.Println(myPointerJson)
	var s_new []*transactions.NamePointer

	err = json.Unmarshal([]byte(myPointerJson), &s_new)
	if err != nil {
		fmt.Println(err)
	}

	node := naet.NewNode(MyNodeConfig.PublicNode, false)
	ttlnoncer := transactions.NewTTLNoncer(node)
	//p := s
	tx, err := transactions.NewNameUpdateTx(accountname, aensname, s_new, 50000, ttlnoncer)

	//fmt.Println(tx)

	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	}

	MysignAccount := SESS_GetAccount(ctx)
	_, myTxhash, _, err := SignBroadcastTransaction(tx, MysignAccount, node, aeconfig.Node.NetworkID)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)

	} else {

		fmt.Println(aensname + "Updated: " + myTxhash)
	}
}

func AENS_WEB_UpdateAENS(ctx iris.Context) {
	aensname := ctx.URLParam("aensname")
	accountname := SESS_GetAccountName(ctx)

	//MyNodeConfig := DB_GetConfigs()
	myurl := MyNodeConfig.PublicNode + "/v2/names/" + aensname
	str := httpGet(myurl)
	//fmt.Println(myurl)

	var s AENSInfo
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(err)
	}

	var myPagedata PageUpdateAENS

	myPagedata.NameID = s.ID
	myPagedata.NameTTL = s.TTL
	myPagedata.NameJson = template.HTML(str)
	myPagedata.AENSName = aensname
	myPagedata.Account = accountname

	myPointers := s.Pointers

	var i int

	for i = 0; i < len(myPointers); i++ {
		if myPointers[i].Key == "account_pubkey" {
			myPagedata.AEAddress = myPointers[i].ID
		}
		if myPointers[i].ID == "ch_ipfsD1iUfRLdnJjQMEczjSzzphPbNnSQudnqUAe1vPJetmMK9" {
			myPagedata.IPFSAddress = myPointers[i].Key
		}
		if myPointers[i].ID == "ch_ipnsoMiJmYq1joKGXFtLRDrSJ3mUjapNB7gcPud7mmpVUXssM" {
			myPagedata.IPNSAddress = myPointers[i].Key
		}
		if myPointers[i].Key == "contract_pubkey" {
			myPagedata.ContractAddress = myPointers[i].ID
		}
		if myPointers[i].Key == "oracle_pubkey" {
			myPagedata.OracleAddress = myPointers[i].ID
		}
		if myPointers[i].ID == "ch_btcqM2NycfJaeLYhYY9uPGKj98iVkwL9VLw7ZP5WzzWHHj2sP" {
			myPagedata.BTCAddress = myPointers[i].Key
		}
		if myPointers[i].ID == "ch_ethe795mCkWMAkguuc3ay9k2JSMikZ61L6VfEMDrujEwCiaiB" {
			myPagedata.ETHAddress = myPointers[i].Key
		}

		if myPointers[i].ID == "ch_em3io3Ntov4qJ1y9mDoyQgHTaWBnBZd1CBu7wnH6iyuF5jf5m" {
			myPagedata.EmailAddress = myPointers[i].Key
		}

		if myPointers[i].ID == "ch_webcVNwKZujeYcxDMjAH5ZUPNwCdcFL4QgYD34pFHZi6KEnzS" {
			myPagedata.WebAddress = myPointers[i].Key
		}

	}

	ctx.ViewData("", myPagedata)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/aens_update.php")
	//ctx.View("aens_update.php")
}

func AENS_WEB_DoTransferAENS(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	aensname := ctx.FormValue("aensname")
	toaddress := ctx.FormValue("toaddress")
	ak := accountname

	//MyNodeConfig := DB_GetConfigs()
	node := naet.NewNode(MyNodeConfig.PublicNode, false)
	ttlnoncer := transactions.NewTTLNoncer(node)

	tx, err := transactions.NewNameTransferTx(accountname, aensname, toaddress, ttlnoncer)

	MysignAccount := SESS_GetAccount(ctx)
	_, myTxhash, _, err := SignBroadcastTransaction(tx, MysignAccount, node, aeconfig.Node.NetworkID)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Failed"}

		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/transaction.php")
		//ctx.View("transaction.php")
	} else {

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: myTxhash}

		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/transaction.php")
		//ctx.View("transaction.php")
	}
}

func AENS_WEB_TransferAENS(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	aensname := ctx.URLParam("aensname")
	myPage := PageAENS{PageContent: template.HTML(aensname)}
	ctx.ViewData("", myPage)
	myTheme := DB_GetGlobalConfigItem("Theme")
	ctx.View(myTheme + "/aens_transfer.php")
	//ctx.View("aens_transfer.php")
}

func AENS_WEB_DoBidAENS(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	aensname := ctx.FormValue("aensname")
	aensprice := ctx.FormValue("aensprice")
	recommendprice := ctx.FormValue("recommendprice")

	var myprice float64

	if strings.TrimSpace(aensprice) == "" {
		myprice, _ = strconv.ParseFloat(recommendprice, 64)
	} else {
		myprice, _ = strconv.ParseFloat(aensprice, 64)
	}

	bigfloatAmount := big.NewFloat(myprice)
	imultiple := big.NewFloat(1000000000000000000) //18 dec
	fmyamount := big.NewFloat(1)
	fmyamount.Mul(bigfloatAmount, imultiple)

	myamount := new(big.Int)
	fmyamount.Int(myamount)

	//MyNodeConfig := DB_GetConfigs()
	node := naet.NewNode(MyNodeConfig.PublicNode, false)
	ttlnoncer := transactions.NewTTLNoncer(node)
	nameSalt := big.NewInt(0)
	tx, _ := transactions.NewNameClaimTx(accountname, aensname, nameSalt, myamount, ttlnoncer)

	MysignAccount := SESS_GetAccount(ctx)
	_, claimTxhash, _, _ := SignBroadcastTransaction(tx, MysignAccount, node, aeconfig.Node.NetworkID)

	hashInfo := "Bidding Tx hash is <a href=" + MyNodeConfig.APINode + "/block/transaction/" + claimTxhash + ">" + claimTxhash + "</a><br /><br /><br />"
	ak := accountname
	myPage := PageWallet{PageId: 23, Account: ak, PageTitle: claimTxhash, PageContent: template.HTML(hashInfo)}

	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/transaction.php")
	//ctx.View("transaction.php")
}
func AENS_WEB_DoRegAENS(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	aensname := ctx.FormValue("aensname")
	aensprice := ctx.FormValue("aensprice")

	var myprice float64

	if strings.TrimSpace(aensprice) == "" {
		myprice = calcAENSFee(aensname)
	} else {
		myprice, _ = strconv.ParseFloat(aensprice, 64)
	}

	bigfloatAmount := big.NewFloat(myprice)
	imultiple := big.NewFloat(1000000000000000000) //18 dec
	fmyamount := big.NewFloat(1)
	fmyamount.Mul(bigfloatAmount, imultiple)

	myamount := new(big.Int) //regfee
	fmyamount.Int(myamount)

	//MyNodeConfig := DB_GetConfigs()
	//MyNodeConfig.PublicNode = "http://192.168.0.105:6013"
	//aeconfig.Node.NetworkID = "aec"
	node := naet.NewNode(MyNodeConfig.PublicNode, false)

	ttlnoncer := transactions.NewTTLNoncer(node)
	tx_pre, nameSalt, err := transactions.NewNamePreclaimTx(accountname, aensname, ttlnoncer)
	fmt.Println("Ready to Precliam the AENS Name " + aensname)
	MysignAccount := SESS_GetAccount(ctx)

	_, preClaimTxhash, _, err := SignBroadcastTransaction(tx_pre, MysignAccount, node, aeconfig.Node.NetworkID)
	hashInfo := "Preclaim Tx hash is <a href=" + MyNodeConfig.APINode + "/block/transaction/" + preClaimTxhash + ">" + preClaimTxhash + "</a><br /><br /><br />"
	fmt.Println(aensname + " was preclaimed.\nReady to claim the AENS Name " + aensname + "\n\nPlease waiting for several minutes for 1 block...")
	time.Sleep(3 * time.Second)
	//sleep 3 secs before ClaimTx

	tx, err := transactions.NewNameClaimTx(accountname, aensname, nameSalt, myamount, ttlnoncer)

	fmt.Println(aensname + " was registered successfully.")

	_, claimTxhash, _, err := SignBroadcastTransaction(tx, MysignAccount, node, aeconfig.Node.NetworkID)
	hashInfo = hashInfo + "Claim Tx hash is <a href=" + MyNodeConfig.APINode + "/block/transaction/" + claimTxhash + ">" + claimTxhash + "</a><br />Please <b>WAIT ~1 block </b> time and check this transaction."

	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)
		ak := accountname
		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Failed"}

		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/transaction.php")
	} else {
		ak := accountname
		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: claimTxhash, PageContent: template.HTML(hashInfo)}

		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/transaction.php")
	}
}

func AENS_WEB_QueryAENS(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}

	aensname := ctx.FormValue("aensname")
	if strings.Index(aensname, ".chain") == -1 {
		aensname = aensname + ".chain"
	}
	accountname := SESS_GetAccountName(ctx)

	//MyNodeConfig := DB_GetConfigs()
	myurl := MyNodeConfig.APINode + "/api/aensquery/" + aensname
	//fmt.Println(myurl)
	str := httpGet(myurl)
	status := ""
	if strings.Index(str, "NONE") > -1 {
		regFee := calcAENSFeeStr(aensname)
		status = `  <div class="box"><div class="col-md-9"> <div class="box-footer">
		<form action="/regaens" method="post">
<input type="hidden" name="aensname" value="` + aensname + `">
                    <div class="input-group">
                      ` + aensname + `:<input type="text" name="aensprice" placeholder="Default price: ` + regFee + ` AE" class="form-control">
                      <br/><br/><br/>                    
                            <button type="submit" class="btn btn-warning btn-flat">Register & Wait several minutes</button>
                         
                    </div>
                  </form></div></div></div>`
		//status = aensname + "=>" + str
	}
	if strings.Index(str, "DONE") > -1 {
		s := strings.Split(str, ":")
		myBalance := ToBigFloat(s[1])
		imultiple := big.NewFloat(0.000000000000000001) //18 dec
		thisamount := new(big.Float).Mul(myBalance, imultiple).String()
		status = "<a href=" + MyNodeConfig.APINode + "/" + aensname + ">" + aensname + "</a>" + "=>Registed.\n<br /><br />Price=>" + thisamount + " AE\n<br /><br /><a href=" + MyNodeConfig.APINode + "/aens/viewbids/" + aensname + ">Check bidding details</a>"

		//status = aensname + "=>" + thisamount
	}

	if strings.Index(str, "BIDDING") > -1 {
		s := strings.Split(str, ":")
		myBalance := ToBigFloat(s[1])
		imultiple := big.NewFloat(0.000000000000000001) //18 dec
		thisamount := new(big.Float).Mul(myBalance, imultiple).String()

		imultiple = big.NewFloat(0.00000000000000000105) //18 dec
		recommendamount := new(big.Float).Mul(myBalance, imultiple).String()

		status = `  <div class="box"><div class="col-md-9"> <div class="box-footer">
		<form action="/bidaens" method="post">
<input type="hidden" name="aensname" value="` + aensname + `">
<input type="hidden" name="recommendprice" value="` + recommendamount + `">
                    <div class="input-group">
                      ` + aensname + "=>Last bidding price:" + thisamount + ` AE(<a href=` + MyNodeConfig.APINode + `/aens/viewbids/` + aensname + ` target=_blank>View bidding details</a>) <input type="text" name="aensprice" placeholder="Recommend bidding price: ` + recommendamount + ` AE" class="form-control">
                      <br/><br/><br/>                    
                            <button type="submit" class="btn btn-warning btn-flat">Bidding with my price</button>
                         
                    </div>
                  </form></div></div></div>`

	}

	queryResults := template.HTML(status)

	myPage := PageAENS{PageContent: queryResults}

	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/aens_query.php")
	//ctx.View("aens_query.php")
}

func AENS_getAENSBidding(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)

	//MyNodeConfig := DB_GetConfigs()

	node := naet.NewNode(MyNodeConfig.PublicNode, false)

	akBalance, err := node.GetAccount(accountname)
	var thisamount string
	var myNonce uint64
	if err != nil {
		fmt.Println("Account not exist.")
		thisamount = "0"
		myNonce = 0
	} else {
		bigstr := akBalance.Balance.String()
		myBalance := ToBigFloat(bigstr)
		imultiple := big.NewFloat(0.000000000000000001) //18 dec
		thisamount = new(big.Float).Mul(myBalance, imultiple).String()
		myNonce = *akBalance.Nonce

	}

	//fmt.Println(myNonce)
	//fmt.Println(thisamount)

	myPage := PageAENS{PageId: 1, Account: accountname, PageTitle: "Wallet", Balance: thisamount, Nonce: myNonce, PageContent: ""}
	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/aens_bidding.php")
	//ctx.View("aens_bidding.php")
}

func AENS_getAENS(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}

	accountname := SESS_GetAccountName(ctx)
	//MyNodeConfig := DB_GetConfigs()

	node := naet.NewNode(MyNodeConfig.PublicNode, false)

	akBalance, err := node.GetAccount(accountname)
	//	topHeight, _ := node.GetHeight()
	var thisamount string
	var myNonce uint64
	if err != nil {
		fmt.Println("Account not exist.")
		thisamount = "0"
		myNonce = 0
	} else {
		bigstr := akBalance.Balance.String()
		myBalance := ToBigFloat(bigstr)
		imultiple := big.NewFloat(0.000000000000000001) //18 dec
		thisamount = new(big.Float).Mul(myBalance, imultiple).String()
		myNonce = *akBalance.Nonce

	}

	//fmt.Println(myNonce)
	//fmt.Println(thisamount)

	myPage := PageAENS{PageId: 1, Account: accountname, PageTitle: "Wallet", Balance: thisamount, Nonce: myNonce, PageContent: ""}

	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/aens.php")
	//ctx.View("aens.php")
}

//get account address from aens name
func AENS_getAccountFromAENS(aensname, accountname string) string {
	//MyNodeConfig := DB_GetConfigs()
	myurl := MyNodeConfig.PublicNode + "/v2/names/" + aensname
	str := httpGet(myurl)
	//fmt.Println(myurl)

	var s AENSInfo
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(err)
	}

	var myPagedata PageUpdateAENS

	myPagedata.NameID = s.ID
	myPagedata.NameTTL = s.TTL
	myPagedata.NameJson = template.HTML(str)
	myPagedata.AENSName = aensname
	myPagedata.Account = s.OWNER

	myPointers := s.Pointers

	var i int

	for i = 0; i < len(myPointers); i++ {
		if myPointers[i].Key == "account_pubkey" {
			return myPointers[i].ID
		}
	}
	return "NULL"
}

///////////////////////////////////Contract////////////////////////////////////////////////

type TokenInfo struct {
	Tokenname string
	Decimal   int64
	Contract  string
	Balance   string
}

type TokenSlice struct {
	Tokens []TokenInfo
}

type CallResutInfo struct {
	Caller_id    string
	Caller_nonce uint64
	Contract_id  string
	Gas_price    uint64
	Gas_used     uint64
	Height       uint64
	Log          []string
	Return_type  string
	Return_value string
}

type CallResutSlice struct {
	Call_info CallResutInfo
}

type PageDeployContract struct {
	Options      template.HTML
	Account      string
	PageTitle    string
	CallContract string
	Callfunc     string
}

var ostype = runtime.GOOS

//decode the cb_strings
func Contract_WEB_DoDecodeContractCall(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	contract_name := ctx.FormValue("contract_name")
	callfunc := ctx.FormValue("callfunc")
	call_result := strings.Trim(ctx.FormValue("call_result"), "\n")

	if callfunc == "" {
		callfunc = "transfer"
	}

	callStr := strings.Replace(call_result, `"`, `\"`, -1)
	callData := Contract_getCallResult(callStr, contract_name, callfunc)

	var myoption template.HTML
	myoption = template.HTML(callData)
	myPage := PageDeployContract{Options: myoption, Account: accountname}
	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/contract_decoded.php")
	//ctx.View("contract_decoded.php")
}

func Contract_WEB_DecodeContractCall(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	contract_name := ""
	ContractsLists := ""
	filepath.Walk("./contracts/decode/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".aes") {

			contract_name = filepath.Base(path)
			if len(contract_name) > 0 {
				ContractsLists = ContractsLists + "<option>" + contract_name + "</option>\n"
			}

		}

		return nil
	})

	var myoption template.HTML
	myoption = template.HTML(ContractsLists)
	myPage := PageDeployContract{Options: myoption, Account: accountname}
	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/contract_decode.php")
	//ctx.View("contract_decode.php")
}

//deploy AEX-9 token UI
func Contract_WEB_DeployTokenUI(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	myPage := PageWallet{PageId: 23, Account: accountname, PageTitle: "AEX-9 Token"}
	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/token_create.php")
	//ctx.View("token_create.php")

}

//do deploy AEX-9 token
func Contrat_WEB_DoDeployToken(ctx iris.Context) {
	// NewContractCreateTx(ownerID string, bytecode string, vmVersion, abiVersion uint16, deposit, amount, gasLimit, gasPrice *big.Int, callData string, ttlnoncer TTLNoncer) (tx *ContractCreateTx, err error)
	accountname := SESS_GetAccountName(ctx)
	name := ctx.FormValue("name")
	symbol := ctx.FormValue("symbol")
	decimals := ctx.FormValue("decimals")
	total_supply := ctx.FormValue("total_supply")
	contract_name := ctx.FormValue("contract_name")

	PayloadStr := "AEX9#" + name + "#" + symbol + "#" + decimals

	decimals_int, _ := strconv.Atoi(decimals)
	decimals_long := "000000000000000000000000000000"
	total_supply = total_supply + decimals_long[0:decimals_int]

	//callData := getCallData("init(\""+name+"\","+decimals+",\""+symbol+"\","+total_supply+")", contract_name)
	//callStr := "init(\\\"" + name + "\\\"," + decimals + ",\\\"" + symbol + "\\\",Some(" + total_supply + "))"
	callStr := ""
	if ostype == "windows" {
		callStr = `init("` + name + `",` + decimals + `,"` + symbol + `",Some(` + total_supply + `))`
		//callStr = `init(\"` + name + `\",` + decimals + `,\"` + symbol + `\",Some(` + total_supply + `))`
	} else {
		callStr = "init(\\\"" + name + "\\\"," + decimals + ",\\\"" + symbol + "\\\",Some(" + total_supply + "))"
	}

	fmt.Println(callStr, contract_name)

	callData := Contract_getCallData(callStr, contract_name)
	fmt.Println(callData)

	vmVersion := uint16(5)
	abiVersion := uint16(3)
	deposit := big.NewInt(0)
	amount := big.NewInt(0)
	gasLimit := big.NewInt(10000)
	gasPrice := big.NewInt(1000000000)

	byteCode := Contract_getByteCode(contract_name)
	//MyNodeConfig := DB_GetConfigs()

	node := naet.NewNode(MyNodeConfig.PublicNode, false)
	ttlnoncer := transactions.NewTTLNoncer(node)
	ownerID := accountname

	tx, err := transactions.NewContractCreateTx(ownerID, byteCode, vmVersion, abiVersion, deposit, amount, gasLimit, gasPrice, callData, ttlnoncer)
	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	} else {
		//fmt.Println(tx)
	}

	MysignAccount := SESS_GetAccount(ctx)
	_, myTxhash, _, err := SignBroadcastTransaction(tx, MysignAccount, node, aeconfig.Node.NetworkID)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)
		ak := accountname

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Failed"}
		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/transaction.php")
		//ctx.View("transaction.php")
	} else {
		ak := accountname
		//TODO:get contract id and redirect to another submit token to browser page
		//contract_id := getContractIDFromHash(myTxhash)

		call_result_url := MyNodeConfig.PublicNode + "/v2/transactions/" + myTxhash + "/info"
		fmt.Println(call_result_url)
		txinfo := httpGet(call_result_url)
		//fmt.Println(txinfo)
		var s CallResutSlice
		err = json.Unmarshal([]byte(txinfo), &s)
		if err != nil {
			fmt.Println(err)
		}

		PayloadStr = PayloadStr + "#" + s.Call_info.Contract_id
		input := []byte(PayloadStr)

		encodeString := base64.StdEncoding.EncodeToString(input)
		PublishLink := "wallet?amount=1&recipient_id=ak_fCCw1JEkvXdztZxk8FRGNAkvmArhVeow89e64yX4AxbCPrVh5&payload=" + encodeString
		fmt.Println(PublishLink)

		myPage := PageWallet{PageId: 23, Account: ak, Recipient_id: symbol, PageTitle: myTxhash, PageContent: template.HTML(PublishLink)}
		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/token_deploy.php")
		//ctx.View("token_deploy.php")
	}
}

func Contract_WEB_CallContractUI(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	callcontract := ""
	callfunc := ctx.URLParam("func")
	callcontract = ctx.URLParam("contract_id")
	callstr := ""
	if callfunc == "mint" {
		callstr = "mint(" + accountname + ",123)"
	}

	if callfunc == "allow" {
		callstr = "create_allowance(for_account,123)"
	}

	if callfunc == "burn" {
		callstr = "burn(123)"
	}

	if callfunc == "transfer_allowance" {
		callstr = "transfer_allowance(ak_from,ak_to,123)"
	}

	if callfunc == "change_allowance" {
		callstr = "transfer_allowance(for_account,123)"
	}

	ContractsLists := ""
	contract_name := ""
	//myLang := getPageString(getPageLang(ctx.Request()))
	//language := ctx.GetLocale().Language()
	//fmt.Println(myLang.Register)

	filepath.Walk("./contracts/call/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".aes") {

			contract_name = filepath.Base(path)
			if len(contract_name) > 0 {
				ContractsLists = ContractsLists + "<option>" + contract_name + "</option>\n"
			}

		}

		return nil
	})

	var myoption template.HTML
	myoption = template.HTML(ContractsLists)

	myPage := PageDeployContract{Options: myoption, Account: accountname, CallContract: callcontract, Callfunc: callstr}
	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/contract_call.php")
	//ctx.View("contract_call.php")

}

func Contract_WEB_DoCallContract(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	callfunc := ctx.FormValue("callfunc")
	contract_id := ctx.FormValue("contract_id")
	contract_name := ctx.FormValue("contract_name")
	//MyNodeConfig := DB_GetConfigs()

	node := naet.NewNode(MyNodeConfig.PublicNode, false)
	ttlnoncer := transactions.NewTTLNoncer(node)
	callerID := accountname

	abiVersion := uint16(3)
	amount := big.NewInt(0)
	gasLimit := big.NewInt(10000)
	gasPrice := big.NewInt(1000000000)
	callStr := strings.Replace(callfunc, `"`, `\"`, -1)
	callData := Contract_getCallData(callStr, contract_name)

	tx, err := transactions.NewContractCallTx(callerID, contract_id, amount, gasLimit, gasPrice, abiVersion, callData, ttlnoncer)
	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	} else {
		//fmt.Println(tx)
	}

	MysignAccount := SESS_GetAccount(ctx)
	_, myTxhash, _, err := SignBroadcastTransaction(tx, MysignAccount, node, aeconfig.Node.NetworkID)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)
		ak := accountname

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Failed"}
		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/transaction.php")
		//ctx.View("transaction.php")
	} else {

		//TODO:return call result
		call_result_url := MyNodeConfig.PublicNode + "/v2/transactions/" + myTxhash + "/info"
		fmt.Println(call_result_url)
		txinfo := httpGet(call_result_url)
		//fmt.Println(txinfo)
		var s CallResutSlice
		err = json.Unmarshal([]byte(txinfo), &s)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(s.Call_info.Return_value)
		braIndex := strings.Index(callfunc, "(")

		CallResutStr := Contract_getCallResult(s.Call_info.Return_value, contract_name, Substr(callfunc, 0, braIndex))
		fmt.Println(CallResutStr)

		var myoption template.HTML
		myoption = template.HTML(CallResutStr)
		myPage := PageWallet{PageId: 23, Account: accountname, PageTitle: myTxhash, PageContent: myoption}
		ctx.ViewData("", myPage)

		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/transaction.php")
	}
}

//deploy any contracts UI
func Contract_WEB_DeployContractUI(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	ContractsLists := ""
	contract_name := ""
	//myLang := getPageString(getPageLang(ctx.Request()))
	//language := ctx.GetLocale().Language()
	//fmt.Println(myLang.Register)

	filepath.Walk("./contracts/deploy/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".aes") {

			contract_name = filepath.Base(path)
			if len(contract_name) > 0 {
				ContractsLists = ContractsLists + "<option>" + contract_name + "</option>\n"
			}

		}

		return nil
	})
	var myoption template.HTML
	myoption = template.HTML(ContractsLists)
	myPage := PageDeployContract{Options: myoption, Account: accountname}
	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/contract_deploy.php")
	//ctx.View("contract_deploy.php")
}

//deploy any contracts
func Contract_WEB_DoDeployContract(ctx iris.Context) {
	// NewContractCreateTx(ownerID string, bytecode string, vmVersion, abiVersion uint16, deposit, amount, gasLimit, gasPrice *big.Int, callData string, ttlnoncer TTLNoncer) (tx *ContractCreateTx, err error)
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	contract_name := ctx.FormValue("contract_name")
	init := ctx.FormValue("init")

	callStr := strings.Replace(init, `"`, `\"`, -1)

	fmt.Println(callStr)

	callData := Contract_getCallData(callStr, contract_name)
	vmVersion := uint16(5)
	abiVersion := uint16(3)
	deposit := big.NewInt(0)
	amount := big.NewInt(0)
	gasLimit := big.NewInt(10000)
	gasPrice := big.NewInt(1000000000)

	byteCode := Contract_getByteCode(contract_name)
	//MyNodeConfig := DB_GetConfigs()

	node := naet.NewNode(MyNodeConfig.PublicNode, false)
	ttlnoncer := transactions.NewTTLNoncer(node)
	ownerID := accountname

	tx, err := transactions.NewContractCreateTx(ownerID, byteCode, vmVersion, abiVersion, deposit, amount, gasLimit, gasPrice, callData, ttlnoncer)
	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	} else {
		//fmt.Println(tx)
	}
	MysignAccount := SESS_GetAccount(ctx)
	_, myTxhash, _, err := SignBroadcastTransaction(tx, MysignAccount, node, aeconfig.Node.NetworkID)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)
		ak := accountname

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Failed"}
		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/transaction.php")

		//ctx.View("transaction.php")
	} else {
		ak := accountname
		//TODO:get contract id and redirect to another submit token to browser page
		//contract_id := getContractIDFromHash(myTxhash)
		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: myTxhash}
		ctx.ViewData("", myPage)

		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/transaction.php")
	}
}

//build transfering token transaction and post it
func Contratc_WEB_TokenTransfer(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	//sender_id := ctx.FormValue("sender_id")
	recipient_id := ctx.FormValue("recipient_id")
	//transferamount := ctx.FormValue("amount")
	amountstr := ctx.FormValue("amount")
	contractID := ctx.FormValue("contractID")
	password := ctx.FormValue("password")

	//convert transfer amout to bigint string
	famount, err := strconv.ParseFloat(amountstr, 64)
	bigfloatAmount := big.NewFloat(famount)
	imultiple := big.NewFloat(1000000000000000000) //18 dec
	fmyamount := big.NewFloat(1)
	fmyamount.Mul(bigfloatAmount, imultiple)
	myamount := new(big.Int)
	fmyamount.Int(myamount)

	transferamount := myamount.String()
	//MyNodeConfig := DB_GetConfigs()

	node := naet.NewNode(MyNodeConfig.PublicNode, false)
	ttlnoncer := transactions.NewTTLNoncer(node)
	ownerID := accountname
	//contractID = "ct_M9yohHgcLjhpp1Z8SaA1UTmRMQzR4FWjJHajGga8KBoZTEPwC"
	//vmVersion := uint16(5)
	abiVersion := uint16(3)
	//deposit := big.NewInt(0)
	amount := big.NewInt(0)
	gasLimit := big.NewInt(10000)
	gasPrice := big.NewInt(1000000000)
	//callData := "cb_KxGEoV2hK58AoMLlAP6SFrYeiuRrxi5A5rNjruumGuhbIsuZStUbvgZYbyS7wfSV"
	callData := Contract_getCallData("transfer("+recipient_id+","+transferamount+")", "aex9.aes")
	//callData := Contract_getCallData("meta_info()")

	//NewContractCallTx(callerID string, contractID string, amount, gasLimit, gasPrice *big.Int, abiVersion uint16, callData string, ttlnoncer TTLNoncer) (tx *ContractCallTx, err error) {
	if strings.Index(recipient_id, ".chain") > -1 {
		recipient_id = AENS_getAccountFromAENS(recipient_id, accountname)
	}

	//Re-check passwd
	session := sess.Start(ctx)
	if password != session.GetString("password") {
		ak := accountname
		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Password error"}
		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/error.php")
		return
	}

	TokenSignAccount := SESS_GetAccount(ctx)

	tx, err := transactions.NewContractCallTx(ownerID, contractID, amount, gasLimit, gasPrice, abiVersion, callData, ttlnoncer)
	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	} else {
		//fmt.Println(tx)
	}

	//_, myTxhash, _, _, _, err := SignBroadcastWaitTransaction(tx, TokenSignAccount, node, aeconfig.Node.NetworkID, 10)
	_, myTxhash, _, err := SignBroadcastTransaction(tx, TokenSignAccount, node, aeconfig.Node.NetworkID)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)
		ak := accountname

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Failed"}
		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/transaction.php")
	} else {
		ak := accountname
		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: myTxhash}
		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/transaction.php")
	}
}

//get the deployed contract's hash by txhash
//func getContractIDFromHash(myTxhash string) string {
//myurl := NodeConfig.PublicNode + "/v2/transactions/" + myTxhash + "/info"
//str := httpGet(myurl)
//}

//get decoded call result
func Contract_getCallResult(callStr, callContract, callfunc string) string {
	if ostype == "windows" {
		c := "bin\\sophia\\erts\\bin\\escript.exe bin\\sophia\\aesophia_cli  contracts\\decode\\" + callContract + " -b fate --call_result " + callStr + " --call_result_fun " + callfunc
		cmd := exec.Command("cmd", "/c", c)
		fmt.Println(c)
		out, _ := cmd.Output()
		callData := strings.Trim(strings.Replace(string(out), "Calldata:", "", 1), "\n")
		fmt.Println(callData)
		return callData
	} else {
		c := "./bin/sophia/erts/bin/escript ./bin/sophia/aesophia_cli  ./contracts/decode/" + callContract + " -b fate --call_result " + callStr + " --call_result_fun " + callfunc
		cmd := exec.Command("sh", "-c", c)
		fmt.Println(c)
		out, _ := cmd.Output()
		callData := strings.Trim(strings.Replace(string(out), "Calldata:", "", 1), "\n")
		fmt.Println(callData)
		return callData
	}

	//cmd := exec.Command("sh", "-c", c)
	//return ""
}

//get call data
func Contract_getCallData(callStr, callContract string) string {

	if ostype == "windows" {
		c := "bin\\sophia\\erts\\bin\\escript.exe bin\\sophia\\aesophia_cli --create_calldata contracts\\deploy\\" + callContract + " --call " + callStr
		cmd := exec.Command("cmd", "/c", c)
		fmt.Println(c)
		out, _ := cmd.Output()
		callData := strings.Trim(strings.Replace(string(out), "Calldata:", "", 1), "\n")
		fmt.Println("Exec result:" + string(out))
		fmt.Println(callData)
		return callData
	} else {
		c := "./bin/sophia/erts/bin/escript ./bin/sophia/aesophia_cli --create_calldata ./contracts/deploy/" + callContract + " --call \"" + callStr + "\""
		fmt.Println(c)
		cmd := exec.Command("sh", "-c", c)
		out, _ := cmd.Output()
		callData := strings.Trim(strings.Replace(string(out), "Calldata:", "", 1), "\n")
		fmt.Println(callData)
		return callData
	}

	//cmd := exec.Command("sh", "-c", c)
	//return ""

}

//compie bytecode of the contract
func Contract_getByteCode(callContract string) string {
	if ostype == "windows" {
		c := "bin\\sophia\\erts\\bin\\escript.exe bin\\sophia\\aesophia_cli contracts\\deploy\\" + callContract
		cmd := exec.Command("cmd", "/c", c)
		out, _ := cmd.Output()
		outStr := strings.Trim(strings.Replace(string(out), "Bytecode:", "", 1), "\n")
		fmt.Println(outStr)

		return outStr
	} else {
		c := "./bin/sophia/erts/bin/escript ./bin/sophia/aesophia_cli ./contracts/deploy/" + callContract
		cmd := exec.Command("sh", "-c", c)
		out, _ := cmd.Output()
		outStr := strings.Trim(strings.Replace(string(out), "Bytecode:", "", 1), "\n")
		fmt.Println(outStr)

		return outStr
	}

	//return ""

}

//Token main page
func Contract_WEB_getToken(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	//MyNodeConfig := DB_GetConfigs()

	node := naet.NewNode(MyNodeConfig.PublicNode, false)

	akBalance, err := node.GetAccount(accountname)
	var thisamount string
	var myNonce uint64
	if err != nil {
		fmt.Println("Account not exist.")
		thisamount = "0"
		myNonce = 0
	} else {
		bigstr := akBalance.Balance.String()
		myBalance := ToBigFloat(bigstr)
		imultiple := big.NewFloat(0.000000000000000001) //18 dec
		thisamount = new(big.Float).Mul(myBalance, imultiple).String()
		myNonce = *akBalance.Nonce

	}

	//fmt.Println(myNonce)
	//fmt.Println(thisamount)
	/*
	   	myurl := NodeConfig.APINode + "/api/token/" + globalAccount.Address
	   	str := httpGet(myurl)
	   	fmt.Println(str)
	   	var s TokenSlice
	   	err = json.Unmarshal([]byte(str), &s)
	   	if err != nil {
	   		fmt.Println(err)
	   	}

	   	var i int
	   	myNames := ""
	   	for i = 0; i < len(s.Tokens); i++ {
	   		//fmt.Println(s.Names[i].Aensname)
	   		bigstr := s.Tokens[i].Balance
	   		myBalance := ToBigFloat(bigstr)
	   		imultiple := big.NewFloat(0.000000000000000001) //18 dec
	   		//thisamount = new(big.Float).Mul(myBalance, imultiple).String()
	   		thistokenamount := fmt.Sprintf("%.2f", new(big.Float).Mul(myBalance, imultiple))

	   		myNames = myNames + `<tr>
	                       <td><a href=` + NodeConfig.APINode + `/token/view/` + s.Tokens[i].Tokenname + ` target=_blank>` + s.Tokens[i].Tokenname + `</a></td>
	                       <td><a href="">` + strconv.FormatInt(s.Tokens[i].Decimal, 10) + `</a></td>
	                       <td><a href="">` + thistokenamount + `</a></td>
	                       <td align="center">
	                         <div class="btn-group">
	   						  <a href=/viewtoken?contractid=` + s.Tokens[i].Contract + `><button type="button" class="btn btn-success">Transfer</button></a> &nbsp;

	   						</div>
	                       </td>
	                     </tr>`
	   	}

	   	myAENSLists := template.HTML(myNames)
	   	aensCount := len(s.Tokens)*/
	myPage := PageAENS{PageId: 99, Account: accountname, PageTitle: "Wallet", Balance: thisamount, Nonce: myNonce, PageContent: ""}

	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/tokenhome.php")
	//ctx.View("tokenhome.php")
	//TODO:get the balance of each account quickly.
}

//Token management page
func Contract_WEB_Token(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	contractid := ctx.URLParam("contractid")
	recipient_id := ctx.URLParam("recipient_id")
	sendamount := ctx.URLParam("amount")

	needReg := true
	ak := ""
	AccountsLists := ""
	//MyNodeConfig := DB_GetConfigs()

	node := naet.NewNode(MyNodeConfig.PublicNode, false)

	akBalance, err := node.GetAccount(accountname)
	var thisamount string
	var myNonce uint64
	if err != nil {
		fmt.Println("Account not exist.")
		thisamount = "0"
		myNonce = 0
	} else {
		bigstr := akBalance.Balance.String()
		myBalance := ToBigFloat(bigstr)
		imultiple := big.NewFloat(0.000000000000000001) //18 dec
		thisamount = new(big.Float).Mul(myBalance, imultiple).String()
		myNonce = *akBalance.Nonce

	}

	merr := filepath.Walk("data/accounts/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "ak_") {

			ak = filepath.Base(path)
			if len(ak) > 0 {
				AccountsLists = AccountsLists + "<option>" + ak + "</option>\n"
			}

			needReg = false
		}
		//fmt.Println(path)
		return nil
	})
	//fmt.Println("address:" + globalAccount.Address)
	if len(accountname) > 1 {
		needReg = false
		ak := accountname

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: contractid, Balance: thisamount, Nonce: myNonce, Recipient_id: recipient_id, Amount: sendamount}
		ctx.ViewData("", myPage)
		myTheme := DB_GetConfigItem(accountname, "Theme")
		ctx.View(myTheme + "/token.php")
		//ctx.View("token.php")

		err := qrcode.WriteFile(ak, qrcode.Medium, 256, "./views/qr_ak.png")
		err = qrcode.WriteFile("https://www.aeknow.org/v2/accounts/"+ak, qrcode.Medium, 256, "./views/qr_account.png")
		if err != nil {
			fmt.Println("write error")
		}
	} else {

		var myoption template.HTML
		myoption = template.HTML(AccountsLists)
		myPage := PageLogin{Options: myoption}
		ctx.ViewData("", myPage)
		myTheme := DB_GetGlobalConfigItem("Theme")
		ctx.View(myTheme + "/login.php")
		//ctx.View("login.php")
	}

	if merr != nil {
		fmt.Println("error")
	}

	if needReg {

		var myPage PageReg
		myPage.PageTitle = "Registering Page"
		myPage.SubTitle = "Decentralized knowledge system without barrier."
		myPage.Register = "Register"

		myPage.Lang = getPageString(getPageLang(ctx.Request()))

		ctx.ViewData("", myPage)
		myTheme := DB_GetGlobalConfigItem("Theme")
		ctx.View(myTheme + "/register.php")
		//ctx.View("register.php")
	}
}

func Contract_WEB_ContractsHome(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	myPage := PageWallet{Account: accountname}
	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/contract_home.php")
	//ctx.View("contract_home.php")
}
