package main

import (
	//"bufio"
	//"bytes"
	"database/sql"
	"encoding/base64"

	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"io"

	//"math"

	//"io/ioutil"

	//"unsafe"

	"net/http"
	"os"

	//"os/exec"

	//"path"
	"strconv"
	"strings"
	"time"

	//"github.com/gomarkdown/markdown"
	//"github.com/gomarkdown/markdown/parser"

	//"github.com/shopspring/decimal"
	//_ "github.com/mattn/go-sqlite"
	_ "modernc.org/sqlite"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/kataras/iris/v12"

	"github.com/aeternity/aepp-sdk-go/v9/account"
	"github.com/aeternity/aepp-sdk-go/v9/naet"
)

type PageBlog struct {
	Account          string
	PageAid          string
	PageHash         string
	PageContent      template.HTML
	PageTitle        string
	PageAuthor       string
	PageAuthorname   string
	PageDescription  string
	PageKeywords     string
	PageTags         string
	PageCategories   string
	PageSignature    string
	Action           string
	EditPath         string
	LastHash         string
	PreLink          template.HTML
	NextLink         template.HTML
	PubTime          string
	LastModTime      string
	AuthorLink       template.HTML
	TagsLink         template.HTML
	CatgoriesLink    template.HTML
	AllTagsLink      template.HTML
	AllCatgoriesLink template.HTML
	LastTenLink      template.HTML
	SigStatus        string
	Site             SiteConfig
}

type PageList struct {
	Account         string
	PageContent     template.HTML
	PageTitle       string
	PageDescription string
	PageTags        string
	PageKeywords    string
	PageCategories  string
	EditPath        string
	PreLink         template.HTML
	NextLink        template.HTML
	TagsLink        template.HTML

	PubTime    string
	AuthorLink template.HTML
}

//	jsonstr = jsonstr + "\"title\":" + title + ",\n" + "\"body\":" + body + ",\n" + "\"pubtime\":" + pubtime + ",\n" +
// "\"lasthash\":" + lasthash + ",\n" + "\"keywords\":" + keywords + ",\n" + "\"tags\":" + tags + ",\n" +
// "\"description\":" + description + ",\n" + "\"author\":" + accountname + "}\n"

type PageInfo struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	Pubtime     string `json:"pubtime"`
	Lasthash    string `json:"lasthash"`
	Keywords    string `json:"keywords"`
	Tags        string `json:"tags"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

type Post struct {
	Aid         int
	Title       string
	Abstract    string
	LastModTime string
	Tags        template.HTML
	Filetype    string
	Hash        string
	Remark      string
	Pubkey      string
	IsOwner     string
}

type Posts struct {
	Pubkey            string
	Name              string
	LastIPFS          string
	Title             string
	Subtitle          string
	Description       string
	AuthorDescription string
	AENS              string
	Theme             string
	Categories        template.HTML
	Tags              template.HTML
	Account           string
	Posts             []Post
	Site              SiteConfig
	Paginator         Paginator
}

type IPFSConfig struct {
	Identity  IPFSIdentity
	Datastore IPFSDatastore
}

type IPFSIdentity struct {
	PeerID  string
	PrivKey string
}

type IPFSDatastore struct {
	StorageMax         string
	StorageGCWatermark string
	GCPeriod           string
	Bootstrap          []string
}

type SiteConfig struct {
	Pubkey            string
	Name              string
	LastIPFS          string
	Title             string
	Subtitle          string
	Description       string
	Author            string
	AuthorDescription string
	AENS              string
	Theme             string
}

type Paginator struct {
	HasPrev    bool
	HasNext    bool
	PageNumber int
	TotalPages int
	PrevURL    string
	NextURL    string
}

var perpage int

//var MyIPFSConfig IPFSConfig
//var MySiteConfig SiteConfig
//var lastIPFS string

//var MyUsername string
//var MyAENS string
/*
func getSiteConfig(accountname string) SiteConfig {

	configFilePath := ""
	if ostype == "windows" {
		configFilePath = "data\\site\\" + accountname + "\\site.json"
	} else {
		configFilePath = "./data/site/" + accountname + "/site.json"
	}
	_, err := os.Stat(configFilePath)

	if err != nil {
		configFilePath = "./data/config_default.json"
	}

	JsonParse := NewJsonStruct()
	readConfigfile := SiteConfig{}
	JsonParse.Load(configFilePath, &readConfigfile)

	return readConfigfile

}*/

func getIPFSConfig() IPFSConfig {
	configFilePath := ""
	if ostype == "windows" {
		configFilePath = "data\\repo\\config"
	} else {
		configFilePath = "./data/repo/config"
	}
	_, err := os.Stat(configFilePath)

	if err != nil {
		configFilePath = "./data/config_default.json"
	}

	JsonParse := NewJsonStruct()
	readConfigfile := IPFSConfig{}
	JsonParse.Load(configFilePath, &readConfigfile)

	return readConfigfile

}

//change the site settings, such as theme, name and description
func iSetSite(ctx iris.Context) {
	if !checkLogin(ctx) {
		ctx.Redirect("/")
	}
	accountname := SESS_GetAccountName(ctx)
	//ctx.ViewData("Account", accountname)
	ctx.ViewData("", DB_GetSiteConfigs(accountname))
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/haeme_settings.php")
	//ctx.View("haeme_settings.php")
}

func iSaveSetSite(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}

	accountname := SESS_GetAccountName(ctx)

	title := ctx.FormValue("title")
	aens := ctx.FormValue("aens")
	title = strings.TrimSpace(title)
	subtitle := ctx.FormValue("subtitle")
	subtitle = strings.TrimSpace(subtitle)
	sitedescription := ctx.FormValue("sitedescription")
	author := ctx.FormValue("author")
	authordescription := ctx.FormValue("authordescription")
	theme := ctx.FormValue("theme")

	dbpath := "./data/accounts/" + accountname + "/config.db"
	table := "config"

	DB_SaveItemToDB(dbpath, table, "Title", title, "")
	DB_SaveItemToDB(dbpath, table, "Subtitle", subtitle, "")
	DB_SaveItemToDB(dbpath, table, "Description", sitedescription, "")
	DB_SaveItemToDB(dbpath, table, "name", author, "")
	DB_SaveItemToDB(dbpath, table, "AuthorDescription", authordescription, "")
	DB_SaveItemToDB(dbpath, table, "AENS", aens, "")
	DB_SaveItemToDB(dbpath, table, "Theme", theme, "")

	//Save to Public DB
	dbpath = "./data/accounts/" + accountname + "/public.db"
	DB_SaveItemToDB(dbpath, table, "Title", title, "")
	DB_SaveItemToDB(dbpath, table, "Subtitle", subtitle, "")
	DB_SaveItemToDB(dbpath, table, "Description", sitedescription, "")
	DB_SaveItemToDB(dbpath, table, "name", author, "")
	DB_SaveItemToDB(dbpath, table, "AuthorDescription", authordescription, "")
	DB_SaveItemToDB(dbpath, table, "AENS", aens, "")
	DB_SaveItemToDB(dbpath, table, "Theme", theme, "")

	//Update alias in Database
	DB_UpdateAccountName(accountname, author)

	//configHugo()
	//update AENS info of the config updating
	go IPFS_UpdateIPNS(accountname, DB_GetConfigItem(accountname, "AENS"))
	ctx.HTML("<h1>Site config has been saved.</h1>")
	//ctx.ViewData("", MySiteConfig)
	//ctx.View("haeme_settings.php")
}

//show the defult homepage
func iHaeme(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	//fmt.Println("Haeme")
	myPage := PageWallet{PageId: 23, Account: accountname, PageTitle: "Haeme"}
	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/transaction.php")
	//ctx.View("transaction.php")
}

func iBlog(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}

	accountname := SESS_GetAccountName(ctx)
	fmt.Println("Haeme")
	myPage := PageBlog{Account: accountname, PageTitle: "", EditPath: ""}

	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/haeme_blog.php")
	//ctx.View("haeme_blog.php")
}

func iView(ctx iris.Context) {
	//View the page content with 2 parameters.
	if !checkLogin(ctx) {
		ctx.Redirect("/")
	}

	hash := ctx.FormValue("hash")
	aid := ctx.FormValue("aid")
	pubkey := ctx.FormValue("pubkey")
	tag := ctx.FormValue("tag")
	cat := ctx.FormValue("cat")
	viewtype := ctx.FormValue("viewtype")
	page := ctx.FormValue("page")

	dbpath := "./data/accounts/" + pubkey + "/public.db"
	//fmt.Println(dbpath)

	if hash != "" && pubkey != "" {
		ViewContent(ctx, hash, pubkey, dbpath, aid)
	}

	if tag != "" && pubkey != "" {
		ViewTag(ctx, tag, pubkey, dbpath, page)
	}

	if cat != "" && pubkey != "" {
		ViewCat(ctx, cat, pubkey, dbpath, page)
	}

	if viewtype == "author" && pubkey != "" {
		ViewHome(ctx, pubkey, dbpath, page)
	}

	if viewtype == "reader" && pubkey != "" {
		ViewHome(ctx, pubkey, dbpath, page)
	}

}
func ViewHome(ctx iris.Context, pubkey string, dbpath string, page string) {
	//View the Homepage of a user
	accountname := SESS_GetAccountName(ctx)
	if !FileExist(dbpath) {
		ctx.HTML("<h1>No such sqlite db." + dbpath + "</h1>")
		return
	}
	db, err := sql.Open("sqlite", dbpath)
	//PageContent := ""
	//PageContent = "PubKey: " + DB_GetPublicItem(pubkey, "pubkey") + ", Author: " + DB_GetPublicItem(pubkey, "name")

	pagenum, err := strconv.Atoi(page)

	querystr := "SELECT aid,title,abstract,tags,keywords,filetype,hash,lastmodtime,remark FROM aek ORDER BY lastmodtime DESC LIMIT " + strconv.Itoa(perpage) + " offset " + strconv.Itoa(pagenum*perpage)
	rows, err := db.Query(querystr)
	checkError(err)

	var aid int
	var title string
	var abstract sql.NullString
	var tags, keywords string
	var filetype string
	var hash string
	var remark sql.NullString
	var lastmodtime string

	var homePosts Posts

	categoriesList := GetAllCategoriesLink(pubkey)
	tagsList := GetAllTagsLink(pubkey)

	homePosts.Posts = []Post{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}

	postCounter := 0
	homePosts.Account = accountname

	//var author string

	for rows.Next() {

		err = rows.Scan(&aid, &title, &abstract, &tags, &keywords, &filetype, &hash, &lastmodtime, &remark)

		myTime, err := strconv.ParseInt(lastmodtime, 10, 64)
		//fmt.Println(lastmodtime)
		tm := time.Unix(myTime, 0)
		lastmodtime = tm.Format("2006-01-02 15:04:05") //2018-07-11 15:10:19

		lastmodtime = strings.Replace(lastmodtime, "T", " ", -1)
		lastmodtime = strings.Replace(lastmodtime, "Z", " ", -1)
		//fmt.Println(lastmodtime)

		//TagsLink := template.HTML(GetTagLink(tags, pubkey))

		homePosts.Posts[postCounter].Aid = aid
		homePosts.Posts[postCounter].Title = title
		homePosts.Posts[postCounter].Abstract = abstract.String
		homePosts.Posts[postCounter].Tags = template.HTML(GetCatsLink(tags, pubkey))
		homePosts.Posts[postCounter].Filetype = filetype
		homePosts.Posts[postCounter].Hash = hash
		homePosts.Posts[postCounter].Remark = remark.String
		homePosts.Posts[postCounter].LastModTime = lastmodtime
		homePosts.Posts[postCounter].Pubkey = pubkey

		if pubkey == accountname {
			homePosts.Posts[postCounter].IsOwner = "YES"
		} else {
			homePosts.Posts[postCounter].IsOwner = ""
		}

		postCounter++
		checkError(err)

	}

	homePosts.Categories = template.HTML(categoriesList)
	homePosts.Tags = template.HTML(tagsList)
	homePosts.Site = GetSiteInfo(pubkey)

	//Pagination
	totalCount := GetTotalCount(pubkey, perpage)
	homePosts.Paginator.TotalPages = totalCount
	homePosts.Paginator.PageNumber = pagenum + 1

	if pagenum+1 < totalCount {
		homePosts.Paginator.HasNext = true
		homePosts.Paginator.NextURL = "/view?pubkey=" + pubkey + "&viewtype=author" + "&page=" + strconv.Itoa(pagenum+1)
	} else {
		homePosts.Paginator.HasNext = false
	}

	if pagenum > 0 {
		homePosts.Paginator.HasPrev = true
		homePosts.Paginator.PrevURL = "/view?pubkey=" + pubkey + "&viewtype=author" + "&page=" + strconv.Itoa(pagenum-1)
	} else {
		homePosts.Paginator.HasPrev = false

	}

	ctx.ViewData("", homePosts)
	//myTheme := DB_GetConfigItem(accountname, "Theme")

	//ctx.View(myTheme + "/haeme_index.php")
	ctx.View("mainroad/haeme_index.php")
	db.Close()

}

func ViewTag(ctx iris.Context, tag string, pubkey string, dbpath string, page string) {
	//View the Homepage of a user

	accountname := SESS_GetAccountName(ctx)
	if !FileExist(dbpath) {
		ctx.HTML("<h1>No such sqlite db." + dbpath + "</h1>")
		return
	}
	db, err := sql.Open("sqlite", dbpath)
	//PageContent := ""
	//PageContent = "PubKey: " + DB_GetPublicItem(pubkey, "pubkey") + ", Author: " + DB_GetPublicItem(pubkey, "name")

	pagenum, err := strconv.Atoi(page)

	querystr := "SELECT aid,title,abstract,tags,keywords,filetype,hash,lastmodtime,remark FROM aek WHERE keywords LIKE '%" + tag + "%' ORDER BY lastmodtime DESC LIMIT " + strconv.Itoa(perpage) + " offset " + strconv.Itoa(pagenum*perpage)
	//fmt.Println(querystr)

	rows, err := db.Query(querystr)
	checkError(err)

	var aid int
	var title string
	var abstract sql.NullString
	var tags, keywords string
	var filetype string
	var hash string
	var remark sql.NullString
	var lastmodtime string

	var homePosts Posts

	categoriesList := GetAllCategoriesLink_cat(pubkey, tag)
	tagsList := GetAllTagsLink_cat(pubkey, tag)

	homePosts.Posts = []Post{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}

	postCounter := 0
	homePosts.Account = accountname

	//var author string

	for rows.Next() {

		err = rows.Scan(&aid, &title, &abstract, &tags, &keywords, &filetype, &hash, &lastmodtime, &remark)

		myTime, err := strconv.ParseInt(lastmodtime, 10, 64)
		//fmt.Println(lastmodtime)
		tm := time.Unix(myTime, 0)
		lastmodtime = tm.Format("2006-01-02 15:04:05") //2018-07-11 15:10:19

		lastmodtime = strings.Replace(lastmodtime, "T", " ", -1)
		lastmodtime = strings.Replace(lastmodtime, "Z", " ", -1)
		//fmt.Println(lastmodtime)

		//TagsLink := template.HTML(GetTagLink(tags, pubkey))

		homePosts.Posts[postCounter].Aid = aid
		homePosts.Posts[postCounter].Title = title
		homePosts.Posts[postCounter].Abstract = abstract.String
		homePosts.Posts[postCounter].Tags = template.HTML(GetCatsLink(tags, pubkey))
		homePosts.Posts[postCounter].Filetype = filetype
		homePosts.Posts[postCounter].Hash = hash
		homePosts.Posts[postCounter].Remark = remark.String
		homePosts.Posts[postCounter].LastModTime = lastmodtime
		homePosts.Posts[postCounter].Pubkey = pubkey

		if pubkey == accountname {
			homePosts.Posts[postCounter].IsOwner = "YES"
		} else {
			homePosts.Posts[postCounter].IsOwner = ""
		}

		postCounter++
		checkError(err)

	}

	homePosts.Categories = template.HTML(categoriesList)
	homePosts.Tags = template.HTML(tagsList)
	homePosts.Site = GetSiteInfo(pubkey)
	homePosts.Title = "Tag:" + tag

	//Pagination
	totalCount := GetTagCount(pubkey, tag, perpage)
	homePosts.Paginator.TotalPages = totalCount
	homePosts.Paginator.PageNumber = pagenum + 1

	if pagenum+1 < totalCount {
		homePosts.Paginator.HasNext = true
		homePosts.Paginator.NextURL = "/view?pubkey=" + pubkey + "&tag=" + tag + "&page=" + strconv.Itoa(pagenum+1)
	} else {
		homePosts.Paginator.HasNext = false
	}

	if pagenum > 0 {
		homePosts.Paginator.HasPrev = true
		homePosts.Paginator.PrevURL = "/view?pubkey=" + pubkey + "&tag=" + tag + "&page=" + strconv.Itoa(pagenum-1)
	} else {
		homePosts.Paginator.HasPrev = false

	}

	ctx.ViewData("", homePosts)
	//myTheme := DB_GetConfigItem(accountname, "Theme")

	//ctx.View(myTheme + "/haeme_index.php")
	ctx.View("mainroad/haeme_categories.php")
	db.Close()

	//TODO: 用户页面布局，按时间展示
}

func ViewCat(ctx iris.Context, tag string, pubkey string, dbpath string, page string) {
	//ctx iris.Context, tag string, pubkey string, dbpath string
	//View the Homepage of a user
	accountname := SESS_GetAccountName(ctx)
	if !FileExist(dbpath) {
		ctx.HTML("<h1>No such sqlite db." + dbpath + "</h1>")
		return
	}
	db, err := sql.Open("sqlite", dbpath)
	//PageContent := ""
	//PageContent = "PubKey: " + DB_GetPublicItem(pubkey, "pubkey") + ", Author: " + DB_GetPublicItem(pubkey, "name")

	pagenum, err := strconv.Atoi(page)
	//querystr := "SELECT aid,title,abstract,tags,keywords,filetype,hash,lastmodtime,remark FROM aek WHERE tags LIKE '%" + tag + "%' ORDER BY lastmodtime DESC LIMIT 10"
	querystr := "SELECT aid,title,abstract,tags,keywords,filetype,hash,lastmodtime,remark FROM aek WHERE tags LIKE '%" + tag + "%' ORDER BY lastmodtime DESC LIMIT " + strconv.Itoa(perpage) + " offset " + strconv.Itoa(pagenum*perpage)

	rows, err := db.Query(querystr)
	checkError(err)

	var aid int
	var title string
	var abstract sql.NullString
	var tags, keywords string
	var filetype string
	var hash string
	var remark sql.NullString
	var lastmodtime string

	var homePosts Posts

	categoriesList := GetAllCategoriesLink_cat(pubkey, tag)
	tagsList := GetAllTagsLink_cat(pubkey, tag)

	homePosts.Posts = []Post{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}

	postCounter := 0
	homePosts.Account = accountname

	//var author string

	for rows.Next() {

		err = rows.Scan(&aid, &title, &abstract, &tags, &keywords, &filetype, &hash, &lastmodtime, &remark)

		myTime, err := strconv.ParseInt(lastmodtime, 10, 64)
		//fmt.Println(lastmodtime)
		tm := time.Unix(myTime, 0)
		lastmodtime = tm.Format("2006-01-02 15:04:05") //2018-07-11 15:10:19

		lastmodtime = strings.Replace(lastmodtime, "T", " ", -1)
		lastmodtime = strings.Replace(lastmodtime, "Z", " ", -1)
		//fmt.Println(lastmodtime)

		//TagsLink := template.HTML(GetTagLink(tags, pubkey))

		homePosts.Posts[postCounter].Aid = aid
		homePosts.Posts[postCounter].Title = title
		homePosts.Posts[postCounter].Abstract = abstract.String
		homePosts.Posts[postCounter].Tags = template.HTML(GetCatsLink(tags, pubkey))
		homePosts.Posts[postCounter].Filetype = filetype
		homePosts.Posts[postCounter].Hash = hash
		homePosts.Posts[postCounter].Remark = remark.String
		homePosts.Posts[postCounter].LastModTime = lastmodtime
		homePosts.Posts[postCounter].Pubkey = pubkey

		if pubkey == accountname {
			homePosts.Posts[postCounter].IsOwner = "YES"
		} else {
			homePosts.Posts[postCounter].IsOwner = ""
		}

		postCounter++
		checkError(err)

	}

	homePosts.Categories = template.HTML(categoriesList)
	homePosts.Tags = template.HTML(tagsList)
	homePosts.Site = GetSiteInfo(pubkey)
	homePosts.Title = "Category:" + tag

	//Pagination
	totalCount := GetCatCount(pubkey, tag, perpage)
	homePosts.Paginator.TotalPages = totalCount
	homePosts.Paginator.PageNumber = pagenum + 1

	if pagenum+1 < totalCount {
		homePosts.Paginator.HasNext = true
		homePosts.Paginator.NextURL = "/view?pubkey=" + pubkey + "&cat=" + tag + "&page=" + strconv.Itoa(pagenum+1)
	} else {
		homePosts.Paginator.HasNext = false
	}

	if pagenum > 0 {
		homePosts.Paginator.HasPrev = true
		homePosts.Paginator.PrevURL = "/view?pubkey=" + pubkey + "&cat=" + tag + "&page=" + strconv.Itoa(pagenum-1)
	} else {
		homePosts.Paginator.HasPrev = false

	}

	ctx.ViewData("", homePosts)
	//myTheme := DB_GetConfigItem(accountname, "Theme")

	//ctx.View(myTheme + "/haeme_index.php")
	ctx.View("mainroad/haeme_categories.php")
	db.Close()

	//TODO: 用户页面布局，按时间展示
}

func ViewContent(ctx iris.Context, hash, pubkey, dbpath, myaid string) {
	if hash == "" || pubkey == "" {
		ctx.HTML("<h1>No such parameter.</h1> ")
		return
	}
	accountname := SESS_GetAccountName(ctx)
	if !FileExist(dbpath) {
		ctx.HTML("<h1>No such sqlite db." + dbpath + "</h1>")
		return
	}

	db, err := sql.Open("sqlite", dbpath)
	checkError(err)

	aid, err := strconv.Atoi(myaid)

	querystr := "SELECT title,author,keywords,aid,abstract,filetype,pubtime,authorname,tags,lasthash,lastmodtime,filesize,signature,body FROM aek WHERE hash='" + hash + "' and aid=" + myaid
	//fmt.Println(querystr)
	rows, err := db.Query(querystr)
	checkError(err)

	var title string
	var author string
	var keywords string
	var abstract sql.NullString
	var filetype sql.NullString
	var pubtime string
	var authorname string
	var signature string
	var body string
	var tags, lasthash, lastmodtime, filesize string
	title = ""
	for rows.Next() {
		err = rows.Scan(&title, &author, &keywords, &aid, &abstract, &filetype, &pubtime, &authorname, &tags, &lasthash, &lastmodtime, &filesize, &signature, &body)
		checkError(err)
	}
	db.Close()

	if len(title) > 0 { //View page and Save to logs
		MyNodeConfig := DB_GetConfigs()

		myTime, err := strconv.ParseInt(pubtime, 10, 64)
		logtime := strconv.FormatInt(time.Now().Unix(), 10)
		tm := time.Unix(myTime, 0)
		pubtime = tm.Format("2006-01-02 15:04:05")

		pubtime = strings.Replace(pubtime, "T", " ", -1)
		pubtime = strings.Replace(pubtime, "Z", " ", -1)

		myTime, err = strconv.ParseInt(lastmodtime, 10, 64)
		tm = time.Unix(myTime, 0)
		lastmodtime = tm.Format("2006-01-02 15:04:05")
		lastmodtime = strings.Replace(lastmodtime, "T", " ", -1)
		lastmodtime = strings.Replace(lastmodtime, "Z", " ", -1)

		PreLink := template.HTML(GetPreLink(aid, pubkey))
		NextLink := template.HTML(GetNextLink(aid, pubkey))

		TagsLink := template.HTML(GetTagsLink(keywords, pubkey))

		CatgoriesLink := template.HTML(GetCatsLink(tags, pubkey))

		theSig, _ := base64.StdEncoding.DecodeString(signature)

		sigVerify, _ := account.Verify(pubkey, []byte(hash), theSig)
		//fmt.Println(hash, signature)

		sigStatus := ""
		if sigVerify {
			sigStatus = "OK"
		} else {
			sigStatus = "FAIL"
		}

		AuthorLink := template.HTML("<a href=/view?pubkey=" + pubkey + "&viewtype=author>" + authorname + "</a>")
		sh := shell.NewShell(MyNodeConfig.IPFSAPI)
		rc, err := sh.Cat("/ipfs/" + hash)
		checkError(err)

		s, err := copyToString(rc)
		checkError(err)
		//fmt.Println("ipfs string", s)

		pageinfo := PageInfo{}
		err = json.Unmarshal([]byte(s), &pageinfo)
		if err != nil {
			checkError(err)
		}
		bodystr, _ := base64.StdEncoding.DecodeString(pageinfo.Body)

		db, err := sql.Open("sqlite", "./data/accounts/"+accountname+"/logs.db")
		sql_check := "SELECT hash FROM logs WHERE hash='" + hash + "'"
		checkError(err)

		rows, err := db.Query(sql_check)
		checkError(err)
		NeedInsert := true
		for rows.Next() {
			NeedInsert = false
		}

		if NeedInsert {
			sql_insert := "INSERT INTO logs(title, author, keywords, abstract, filetype, pubtime, authorname,hash,tags, lasthash, lastmodtime, filesize,signature,body) VALUES('" + title + "','" + author + "','" + keywords + "','" + abstract.String + "','" + filetype.String + "','" + logtime + "','" + authorname + "','" + hash + "','" + tags + "','" + lasthash + "','" + lastmodtime + "'," + filesize + ",'" + signature + "','" + body + "')" //, , ,
			//fmt.Println(sql_insert)
			_, err := db.Exec(sql_insert)
			checkError(err)
		}
		db.Close()

		//myPage := PageBlog{AuthorLink: AuthorLink, PubTime: pubtime, TagsLink: TagsLink, PreLink: PreLink, NextLink: NextLink, Account: accountname, PageTitle: title, PageContent: template.HTML(pageinfo.Body), PageTags: keywords, EditPath: author, LastHash: lasthash}

		myPage := PageBlog{AuthorLink: AuthorLink, PubTime: pubtime, TagsLink: TagsLink, CatgoriesLink: CatgoriesLink, PreLink: PreLink, NextLink: NextLink, Account: accountname, PageTitle: title, PageContent: template.HTML(string(bodystr)), EditPath: author, LastHash: lasthash, AllCatgoriesLink: template.HTML(GetAllCategoriesLink(pubkey)), AllTagsLink: template.HTML(GetAllTagsLink(pubkey)), LastTenLink: template.HTML(GetLastTenLinks(pubkey)), PageAid: myaid, PageHash: hash, PageDescription: abstract.String, SigStatus: sigStatus, PageSignature: signature}
		myPage.Site = GetSiteInfo(pubkey)
		myPage.PageAuthorname = authorname
		myPage.PageAuthor = pubkey
		myPage.LastModTime = lastmodtime

		ctx.ViewData("", myPage)

		if filetype.String == "markdown" {
			//myTheme := DB_GetConfigItem(accountname, "Theme")
			ctx.View("mainroad/haeme_page.php")
			//ctx.View("haeme_page.php")
		}

		if filetype.String == "html" {
			myTheme := DB_GetConfigItem(accountname, "Theme")
			ctx.View(myTheme + "/haeme_page_html.php")
			//ctx.View("haeme_page_html.php")
		}
	} else {
		ctx.HTML("<h1>No such hash in database</h1>")
	}

	fmt.Println("Haeme")
}

func copyToString(r io.Reader) (res string, err error) {
	var sb strings.Builder
	if _, err = io.Copy(&sb, r); err == nil {
		res = sb.String()
	}
	return
}

func GetCatsLink(tags string, pubkey string) string {
	//Generate tag links
	TmpStr := strings.Split(tags, ",")
	taglink := ""
	for i := 0; i < len(TmpStr); i++ {
		taglink = taglink + "<a href=/view?pubkey=" + pubkey + "&cat=" + TmpStr[i] + ">" + TmpStr[i] + "</a>,"
	}
	taglink = strings.TrimRight(taglink, ",")

	return taglink
}

func GetAllCategoriesLink(pubkey string) string {
	dbpath := "./data/accounts/" + pubkey + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)
	querystr := "SELECT tags FROM aek;"

	rows, err := db.Query(querystr)
	checkError(err)
	var tags string
	categoriesList := ""

	for rows.Next() {
		err = rows.Scan(&tags)
		TmpStr := strings.Split(tags, ",")
		for i := 0; i < len(TmpStr); i++ {
			catLink := "<li class=\"widget__item\"><a lass=\"widget__link\" href=/view?pubkey=" + pubkey + "&cat=" + TmpStr[i] + ">" + TmpStr[i] + "</a></li>"

			if !strings.Contains(categoriesList, catLink) && len(tags) > 2 {
				categoriesList = categoriesList + catLink
			}
		}

		checkError(err)
	}

	db.Close()

	return categoriesList
}

func GetAllCategoriesLink_cat(pubkey string, cat string) string {
	dbpath := "./data/accounts/" + pubkey + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)

	querystr := ""

	if len(cat) > 0 {
		querystr = "SELECT tags FROM aek WHERE tags LIKE '%" + cat + "%';"
	} else {
		querystr = "SELECT tags FROM aek;"
	}

	rows, err := db.Query(querystr)
	checkError(err)
	var tags string
	categoriesList := ""

	for rows.Next() {
		err = rows.Scan(&tags)
		TmpStr := strings.Split(tags, ",")
		for i := 0; i < len(TmpStr); i++ {
			catLink := "<li class=\"widget__item\"><a lass=\"widget__link\" href=/view?pubkey=" + pubkey + "&cat=" + TmpStr[i] + ">" + TmpStr[i] + "</a></li>"

			if !strings.Contains(categoriesList, catLink) && len(tags) > 2 {
				categoriesList = categoriesList + catLink
			}
		}

		checkError(err)
	}

	db.Close()

	return categoriesList
}

func GetAllTagsLink(pubkey string) string {
	dbpath := "./data/accounts/" + pubkey + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)
	querystr := "SELECT keywords FROM aek;"

	rows, err := db.Query(querystr)
	checkError(err)
	var keywords string
	tagsList := ""

	for rows.Next() {
		err = rows.Scan(&keywords)
		TmpStr := strings.Split(keywords, ",")
		for i := 0; i < len(TmpStr); i++ {
			catLink := "<a class=\"widget-taglist__link widget__link btn\" href=\"/view?pubkey=" + pubkey + "&tag=" + TmpStr[i] + "\" title=\"" + TmpStr[i] + "\">" + TmpStr[i] + "</a>\n"

			if !strings.Contains(tagsList, catLink) && len(keywords) > 2 {
				tagsList = tagsList + catLink
			}
		}

		checkError(err)
	}

	db.Close()

	return tagsList
}

func GetTagCount(pubkey string, tag string, perpage int) int {
	dbpath := "./data/accounts/" + pubkey + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)
	querystr := "SELECT count(*) FROM aek WHERE keywords LIKE '%" + tag + "%';"

	var count int
	count = 0
	err = db.QueryRow(querystr).Scan(&count)

	return (count / perpage) + 1

}

func GetCatCount(pubkey string, cat string, perpage int) int {
	dbpath := "./data/accounts/" + pubkey + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)
	querystr := "SELECT count(*) FROM aek WHERE tags LIKE '%" + cat + "%';"

	var count int
	count = 0
	err = db.QueryRow(querystr).Scan(&count)

	return (count / perpage) + 1

}

func GetTotalCount(pubkey string, perpage int) int {
	dbpath := "./data/accounts/" + pubkey + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)
	querystr := "SELECT count(*) FROM aek;"

	var count int
	count = 0
	err = db.QueryRow(querystr).Scan(&count)

	return (count / perpage) + 1

}

func GetAllTagsLink_cat(pubkey string, key string) string {
	dbpath := "./data/accounts/" + pubkey + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)

	querystr := ""

	if len(key) > 0 {
		querystr = "SELECT tags FROM aek WHERE keywords LIKE '%" + key + "%';"
	} else {
		querystr = "SELECT keywords FROM aek;"
	}

	rows, err := db.Query(querystr)
	checkError(err)
	var keywords string
	tagsList := ""

	for rows.Next() {
		err = rows.Scan(&keywords)
		TmpStr := strings.Split(keywords, ",")
		for i := 0; i < len(TmpStr); i++ {
			catLink := "<a class=\"widget-taglist__link widget__link btn\" href=\"/view?pubkey=" + pubkey + "&tag=" + TmpStr[i] + "\" title=\"" + TmpStr[i] + "\">" + TmpStr[i] + "</a>\n"

			if !strings.Contains(tagsList, catLink) && len(keywords) > 2 {
				tagsList = tagsList + catLink
			}
		}

		checkError(err)
	}

	db.Close()

	return tagsList
}

func GetSiteInfo(pubkey string) SiteConfig {

	dbpath := "./data/accounts/" + pubkey + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)

	querystr := "SELECT item,value FROM config"
	rows, err := db.Query(querystr)
	checkError(err)
	var SiteConfigs SiteConfig
	var item, value string
	for rows.Next() {

		err = rows.Scan(&item, &value)
		if item == "pubkey" {
			SiteConfigs.Pubkey = value
		}

		if item == "name" {
			SiteConfigs.Name = value
		}

		if item == "LastIPFS" {
			SiteConfigs.LastIPFS = value
		}

		if item == "Title" {
			SiteConfigs.Title = value
		}

		if item == "Subtitle" {
			SiteConfigs.Subtitle = value
		}

		if item == "Description" {
			SiteConfigs.Description = value
		}

		if item == "AuthorDescription" {
			SiteConfigs.AuthorDescription = value
		}

		if item == "AENS" {
			SiteConfigs.AENS = value
		}

		if item == "Theme" {
			SiteConfigs.Theme = value
		}

	}

	return SiteConfigs

}

func GetLastTenLinks(pubkey string) string {
	dbpath := "./data/accounts/" + pubkey + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)
	querystr := "SELECT aid,hash,title FROM aek ORDER BY lastmodtime DESC LIMIT 10;"

	rows, err := db.Query(querystr)
	checkError(err)
	var hash, title string
	var aid string

	lastTenLinks := ""

	for rows.Next() {
		err = rows.Scan(&aid, &hash, &title)
		//<li class="widget__item"><a class="widget__link" href="/view?pubkey=ak_fCCw1JEkvXdztZxk8FRGNAkvmArhVeow89e64yX4AxbCPrVh5&hash={{.Hash}}&aid={{.Aid}}">{{.Title}}</a></li>

		catLink := "<li class=\"widget__item\"><a class=\"widget__link\" href=\"/view?pubkey=" + pubkey + "&hash=" + hash + "&aid=" + aid + "\" title=\"" + title + "\">" + title + "</a></li>\n"

		if !strings.Contains(lastTenLinks, catLink) && len(hash) > 2 {
			lastTenLinks = lastTenLinks + catLink
		}

		checkError(err)
	}

	db.Close()

	return lastTenLinks

}

func GetTagsLink(keywords string, pubkey string) string {
	//Generate tag links
	TmpStr := strings.Split(keywords, ",")
	taglink := ""
	for i := 0; i < len(TmpStr); i++ {
		//<li class="tags__item"><a class="tags__link btn" href="/tags/css/" rel="tag">CSS</a></li>
		taglink = taglink + "<li class=\"tags__item\"><a class=\"tags__link btn\"  href=/view?pubkey=" + pubkey + "&tag=" + TmpStr[i] + " rel=\"tag\">" + TmpStr[i] + "</a></li>"
	}
	taglink = strings.TrimRight(taglink, ",")

	return taglink
}

func GetPreLink(aid int, pubkey string) string {
	//Get the previous page and link
	dbpath := "./data/accounts/" + pubkey + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)
	querystr := "SELECT title,hash,aid FROM aek WHERE aid<" + strconv.Itoa(aid) + " ORDER BY aid desc LIMIT 1;"

	rows, err := db.Query(querystr)
	checkError(err)
	title := ""
	hash := ""
	myaid := ""

	for rows.Next() {
		err = rows.Scan(&title, &hash, &myaid)
		checkError(err)
	}

	db.Close()
	return "<a href=/view?pubkey=" + pubkey + "&hash=" + hash + "&aid=" + myaid + ">" + title + "</a>"
}

func GetNextLink(aid int, pubkey string) string {
	//Get the next page and link
	dbpath := "./data/accounts/" + pubkey + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)
	querystr := "SELECT title,hash,aid FROM aek WHERE aid>" + strconv.Itoa(aid) + " ORDER BY aid asc LIMIT 1;"

	rows, err := db.Query(querystr)
	checkError(err)
	title := ""
	hash := ""
	myaid := ""

	for rows.Next() {
		err = rows.Scan(&title, &hash, &myaid)
		checkError(err)
	}

	db.Close()
	return "<a href=/view?pubkey=" + pubkey + "&hash=" + hash + "&aid=" + myaid + ">" + title + "</a>"
}

func iSaveBlog(ctx iris.Context) {

	if !checkLogin(ctx) {
		return
	}
	jsonstr := "{"
	accountname := SESS_GetAccountName(ctx)

	//url := ctx.FullRequestURI()

	title := ctx.FormValue("title")
	title = strings.TrimSpace(title)

	//keywords := ctx.FormValue("keywords")
	//keywords = strings.Replace(keywords, "，", ",", -1)
	tags := ctx.FormValue("tags")
	tags = strings.Replace(tags, "，", ",", -1)

	keywords := ctx.FormValue("keywords")
	keywords = strings.Replace(keywords, "，", ",", -1)

	content := ctx.FormValue("content")
	description := ctx.FormValue("description")
	aid := ctx.FormValue("aid")
	action := ctx.FormValue("action")
	lasthash := ctx.FormValue("lasthash")
	//draft := ctx.FormValue("draft")

	body := html.EscapeString(content)
	description = html.EscapeString(description)

	MyNodeConfig := DB_GetConfigs()
	sh := shell.NewShell(MyNodeConfig.IPFSAPI)

	dbpath := "./data/accounts/" + accountname + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)

	filesize := strconv.Itoa(strings.Count(body, ""))
	pubtime := strconv.FormatInt(time.Now().Unix(), 10)

	jsonstr = jsonstr + "\"title\":\"" + base64.StdEncoding.EncodeToString([]byte(title)) + "\",\n" + "\"body\":\"" + base64.StdEncoding.EncodeToString([]byte(body)) + "\",\n" + "\"pubtime\":\"" + pubtime + "\",\n" + "\"lasthash\":\"" + lasthash + "\",\n" + "\"keywords\":\"" + base64.StdEncoding.EncodeToString([]byte(keywords)) + "\",\n" + "\"tags\":\"" + base64.StdEncoding.EncodeToString([]byte(tags)) + "\",\n" + "\"description\":\"" + base64.StdEncoding.EncodeToString([]byte(description)) + "\",\n" + "\"author\":\"" + accountname + "\"}\n"

	hash, err := sh.Add(strings.NewReader(jsonstr))
	fmt.Println("posted hash: " + hash)
	checkError(err)

	MyUsername := DB_GetAccountName(accountname)
	MysignAccount := SESS_GetAccount(ctx)
	signature := base64.StdEncoding.EncodeToString(MysignAccount.Sign([]byte(hash)))

	//fmt.Println(url, "==>", action)
	if aid == "" || action == "fork" {

		sql_insert := "INSERT INTO aek(title,abstract,hash,tags,keywords,author,pubtime,filetype,authorname,filesize,lasthash,body,signature,lastmodtime) VALUES('" + title + "','" + description + "','" + hash + "','" + tags + "','" + keywords + "','" + accountname + "'," + pubtime + ",'markdown','" + MyUsername + "'," + filesize + ",'" + lasthash + "','" + body + "','" + signature + "'," + pubtime + ")"
		//sql_insert := "INSERT INTO aek(title,abstract,hash,tags,keywords,author,pubtime,filetype,authorname,filesize,lasthash) VALUES('" + title + "','" + description + "','" + hash + "','" + tags + "','" + keywords + "','" + accountname + "'," + pubtime + ",'markdown','" + MyUsername + "'," + filesize + ",'" + lasthash + "')"

		fmt.Println(sql_insert)
		_, err = db.Exec(sql_insert)

		checkError(err)
	} else {
		aid := ctx.FormValue("aid")
		lasthash := ctx.FormValue("lasthash")
		sql_update := "UPDATE aek SET title='" + title + "',abstract='" + description + "',hash='" + hash + "',keywords='" + keywords + "',tags='" + tags + "',author='" + accountname + "',filesize='" + filesize + "',authorname='" + MyUsername + "',lasthash='" + lasthash + "',lastmodtime=" + pubtime + ",body='" + body + "',signature='" + signature + "' WHERE aid=" + aid
		//sql_update := "UPDATE aek SET title='" + title + "',abstract='" + description + "',hash='" + hash + "',keywords='" + keywords + "',tags='" + tags + "',author='" + accountname + "',pubtime='" + pubtime + "',filesize='" + filesize + "',authorname='" + MyUsername + "',lasthash='" + lasthash + "',lastmodtime=" + pubtime + " WHERE aid=" + aid
		//fmt.Println(sql_update)
		_, err = db.Exec(sql_update)

		checkError(err)
	}

	db.Close()

	//add new database to ipfs
	pubfile, err := os.Open(dbpath)
	if err != nil {
		fmt.Print("Failed openfile", err)

	}
	cid, err := sh.Add(pubfile)

	//lastIPFS = cid
	DB_UpdatePublicConfigItem(accountname, "LastIPFS", cid, "")
	DB_UpdateConfigs(accountname, "LastIPFS", cid) //update lastipfs
	fmt.Println("Pub: " + cid)
	if err != nil {
		panic(err)
	} else {
		//Update AENS name.
		//AENS_UpdateAENS(MyAENS, "ak_ipfsD1iUfRLdnJjQMEczjSzzphPbNnSQudnqUAe1vPJetmMK9", cid, accountname)
		//fmt.Println("Posted and updated." + IPFS_Update(accountname))

		go IPFS_UpdateIPNS(accountname, DB_GetConfigItem(accountname, "AENS"))
		ctx.Redirect("/view?pubkey=" + accountname + "&viewtype=author")
	}

	//TODO:detail the section of title and other field in UI, and polish UI
}

func BindMyDomain(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	MyDomain := DB_GetConfigItem(accountname, "AENS")

	myPage := PageBlog{Account: accountname, PageTitle: "", PageContent: "", PageTags: "", PageKeywords: "", EditPath: ""}
	//myPage := PageWallet{PageId: 23, Account: globalAccount.Address, PageTitle: "Haeme"}
	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/haeme_settings.php")
	//ctx.View("haeme_settings.php")
	fmt.Println(MyDomain)
}
func iNewBlog(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	myPage := PageBlog{Account: accountname, PageTitle: "", PageContent: "", PageTags: "", PageKeywords: "", EditPath: ""}
	//myPage := PageWallet{PageId: 23, Account: globalAccount.Address, PageTitle: "Haeme"}
	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/haeme_newblog.php")
	//ctx.View("haeme_newblog.php")
}

func iDelBlog(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	aid := ctx.URLParam("aid")

	dbpath := "./data/accounts/" + accountname + "/public.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)

	CleanIPFSChain(accountname, aid)

	sql_del := "DELETE FROM aek WHERE aid=" + aid
	fmt.Println(sql_del)
	_, err = db.Exec(sql_del)
	checkError(err)
	if err == nil {
		go IPFS_UpdateIPNS(accountname, DB_GetConfigItem(accountname, "AENS"))
		ctx.Redirect("/view?pubkey=" + accountname + "&viewtype=author")
	} else {

		fmt.Println("Delete Failed", err)
	}

}

func CleanIPFSChain(accountname, aid string) {
	//TODO:Clean all the ipfs onchain

}

func iEditBlog(ctx iris.Context) {
	//if !checkLogin(ctx) {
	//	return
	//}
	accountname := SESS_GetAccountName(ctx)
	pubkey := ctx.URLParam("pubkey")
	aid := ctx.URLParam("aid")
	action := ctx.URLParam("action")
	dbpath := "./data/accounts/" + pubkey + "/public.db"
	db, err := sql.Open("sqlite", dbpath)

	querystr := "SELECT title,hash,abstract,keywords,tags,author,filetype FROM aek WHERE aid=" + aid
	rows, err := db.Query(querystr)
	checkError(err)

	var title string
	var hash string
	var abstract string
	var keywords string
	var tags string
	var author string
	var filetype string
	//PageContent := ""
	for rows.Next() {
		err = rows.Scan(&title, &hash, &abstract, &keywords, &tags, &author, &filetype)
		checkError(err)
	}

	db.Close()
	fmt.Println(dbpath, hash)
	MyNodeConfig := DB_GetConfigs()
	sh := shell.NewShell(MyNodeConfig.IPFSAPI)
	rc, err := sh.Cat("/ipfs/" + hash)
	ipfsstr, err := copyToString(rc)
	checkError(err)

	pageinfo := PageInfo{}
	err = json.Unmarshal([]byte(ipfsstr), &pageinfo)
	if err != nil {
		checkError(err)
	}
	bodystr, _ := base64.StdEncoding.DecodeString(pageinfo.Body)
	pageinfo.Body = string(bodystr)

	myPage := PageBlog{Account: accountname, PageTitle: title, PageDescription: abstract, PageContent: template.HTML(pageinfo.Body), PageTags: tags, PageKeywords: keywords, EditPath: aid, LastHash: hash, Action: action}
	ctx.ViewData("", myPage)
	myTheme := DB_GetConfigItem(accountname, "Theme")
	ctx.View(myTheme + "/haeme_editblog.php")
	//ctx.View("haeme_editblog.php")

}

func iBlogUploadFile(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	//filename := ctx.FormValue("filename")
	//fmt.Println("\n\nfile:" + filename + "\n\n")
	file, info, err := ctx.FormFile("editormd-image-file")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}

	defer file.Close()
	fname := info.Filename
	fmt.Println(fname)
	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile("./uploads/"+fname,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	defer out.Close()

	io.Copy(out, file)

	MyNodeConfig := DB_GetConfigs()

	sh := shell.NewShell(MyNodeConfig.IPFSAPI)
	myfile := ".\\uploads\\" + fname
	fmt.Println("Open file ", myfile)
	pubfile, err := os.Open(myfile)
	if err != nil {
		fmt.Println("Open file failed.", err)
	}
	cid, err := sh.Add(pubfile)
	if err != nil {
		fmt.Println("Add file failed.", err)
	}
	pubfile.Close()

	uploadedImageValue := `{
  "success": 1,
"message" : "` + fname + `", 
  "url": "/ipfs/` + cid + "" + `"  
}`

	fmt.Println(uploadedImageValue)
	ctx.Writef(uploadedImageValue)
	//remove the temp uploaded file
	//time.Sleep(1 * time.Second)
	if ostype == "windows" {
		err = os.Remove(".\\uploads\\" + fname)
		if err != nil {
			fmt.Println("Delete uplaod file failed.", err)
		}
	} else {
		err = os.Remove("./uploads/" + fname)
		if err != nil {
			fmt.Println("Delete uplaod file failed.")
		}
	}
}

func iGoAENS(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	accountname := SESS_GetAccountName(ctx)
	aensname := ctx.URLParam("aensname")
	refresh := ctx.URLParam("refresh")
	gohome := ctx.URLParam("gohome")
	MyNodeConfig := DB_GetConfigs()
	//Go home firstly
	if gohome == "gohome" {
		//ctx.Redirect(MyNodeConfig.IPFSNode + "/ipfs/" + DB_GetConfigItem(accountname, "LastIPFS"))
		ctx.Redirect("/view?pubkey=" + accountname + "&viewtype=author")
		return
	}

	//Do normal AENS resolve
	myurl := MyNodeConfig.PublicNode + "/v2/names/" + aensname

	str := httpGet(myurl)
	fmt.Println(myurl)

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
	redirecturl := MyNodeConfig.IPFSNode
	IsRedirect := false

	if myPagedata.IPNSAddress != "" {
		redirecturl = MyNodeConfig.IPFSNode + "/" + "ipns/" + myPagedata.IPNSAddress
		if refresh == "refresh" {
			t := time.Now()
			timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
			redirecturl = MyNodeConfig.IPFSNode + "/" + "ipns/" + myPagedata.IPNSAddress + "/?" + timestamp
		}
		IsRedirect = true
	}

	if myPagedata.IPFSAddress != "" {
		redirecturl = MyNodeConfig.IPFSNode + "/" + "ipfs/" + myPagedata.IPFSAddress
		if refresh == "refresh" {
			t := time.Now()
			timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
			redirecturl = MyNodeConfig.IPFSNode + "/" + "ipfs/" + myPagedata.IPFSAddress + "/?" + timestamp
		}
		IsRedirect = true
	}

	if IsRedirect {
		fmt.Println(redirecturl)
		ctx.Redirect(redirecturl)
	}

	ctx.HTML("No IPFS or IPNS pointer,AENS info:<br/>" + str)
}

func AENS_GetData(ctx iris.Context) {

	accountname := SESS_GetAccountName(ctx)
	aensname := ctx.URLParam("aensname")

	MyNodeConfig := DB_GetConfigs()

	//Do normal AENS resolve
	fmt.Println("Start Resolve")
	myurl := MyNodeConfig.PublicNode + "/v2/names/" + aensname
	str := httpGet(myurl)
	fmt.Println(myurl)

	var s AENSInfo
	var reposites Reposite
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

	theAccount := s.OWNER

	myPointers := s.Pointers

	var i int

	for i = 0; i < len(myPointers); i++ {
		if myPointers[i].Key == "account_pubkey" {
			myPagedata.AEAddress = myPointers[i].ID
			theAccount = myPointers[i].ID
			fmt.Println(theAccount)
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

	theIPFS := ""
	theIPNS := ""
	theBlock_height := ""

	if myPagedata.IPNSAddress != "" {
		ipnsurl := MyNodeConfig.IPFSNode + "/ipns/" + myPagedata.IPNSAddress
		theIPNS = myPagedata.IPNSAddress
		str = httpGet(ipnsurl)
		fmt.Println(str)
		err = json.Unmarshal([]byte(str), &reposites)
		if err != nil {
			fmt.Println(err)
		}

		for i = 0; i < len(reposites.Reposites); i++ {
			if reposites.Reposites[i].Name == aensname {
				theIPFS = reposites.Reposites[i].Hash
				theBlock_height = reposites.Reposites[i].Metainfo
			}
		}

	} else {
		if myPagedata.IPFSAddress != "" {
			theIPFS = myPagedata.IPFSAddress
			theBlock_height = "0"
		}
	}

	if AENS_needUpdate(aensname, theIPFS, theBlock_height) {
		fileUrl := MyNodeConfig.IPFSNode + "/" + "ipfs/" + theIPFS
		//fmt.Println(theAccount + "=>" + theIPFS)
		// Get the data
		resp, err := http.Get(fileUrl)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// 创建一个文件用于保存
		filename := "./data/accounts/" + theAccount + "/public.db"
		filedir := "./data/accounts/" + theAccount
		if !FileExist(filedir) {
			os.Mkdir(filedir, 0755)
		}
		//filename := "./data/accounts/" + theAccount + ".db"
		out, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		// 然后将响应流和文件流对接起来
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("Update data from " + aensname)

		AENS_UpdateOnce(aensname, theIPNS, theIPFS, accountname, theBlock_height)
	}

	redirecturl := "/view?pubkey=" + theAccount + "&viewtype=author"
	ctx.Redirect(redirecturl)

}

func AENS_needUpdate(aensname, ipfs, block_height string) bool {

	if ipfs == "" {
		return false
	}

	dbpath := "./data/aens.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)
	sql_check := "SELECT ipfs,block_height FROM aens WHERE aensname='" + aensname + "' ORDER BY aid DESC LIMIT 1"
	fmt.Println(sql_check)
	rows, err := db.Query(sql_check)
	checkError(err)
	NeedUpdate := true
	old_ipfs := ""
	old_block_height := ""

	for rows.Next() {
		err = rows.Scan(&old_ipfs, &old_block_height)
		fmt.Println(old_ipfs, old_block_height)
		if old_ipfs == ipfs {
			NeedUpdate = false
		} else {
			if block_height > old_block_height {
				NeedUpdate = true
			}
		}
	}

	db.Close()

	return NeedUpdate

}

func AENS_UpdateOnce(aensname, ipns, ipfs, accountname, block_height string) {
	if ipfs == "" {

		fmt.Println("err: NULL IPFS")
		return
	}
	dbpath := "./data/aens.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)

	MyNodeConfig := DB_GetConfigs()
	node := naet.NewNode(MyNodeConfig.PublicNode, false)
	new_block_height, err := node.GetHeight()

	sql_insert := "INSERT INTO aens(aensname,ipns,ipfs,block_height) VALUES('" + aensname + "','" + ipns + "','" + ipfs + "','" + strconv.FormatUint(new_block_height, 10) + "')"
	fmt.Println(sql_insert)
	_, err = db.Exec(sql_insert)
	checkError(err)

	db.Close()
}

func AENS_UpdateALLOnce(ctx iris.Context) {
	dbpath := "./data/aens.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)

	sql_check := "SELECT DISTINCT(aensname) FROM aens;"
	rows, err := db.Query(sql_check)
	checkError(err)

	aensname := ""
	sql_insert := ""
	for rows.Next() {
		err = rows.Scan(&aensname)
		MyNodeConfig := DB_GetConfigs()

		//Do normal AENS resolve
		fmt.Println("Start Resolve")
		myurl := MyNodeConfig.PublicNode + "/v2/names/" + aensname
		str := httpGet(myurl)
		fmt.Println(myurl)

		var s AENSInfo
		var reposites Reposite
		err := json.Unmarshal([]byte(str), &s)
		if err != nil {
			fmt.Println(err)
		}

		var myPagedata PageUpdateAENS

		myPagedata.NameID = s.ID
		myPagedata.NameTTL = s.TTL
		myPagedata.NameJson = template.HTML(str)
		myPagedata.AENSName = aensname
		//myPagedata.Account = accountname

		theAccount := s.OWNER

		myPointers := s.Pointers

		var i int
		theIPFS := ""
		theIPNS := ""
		theAENS := ""
		theBlock_height := ""

		for i = 0; i < len(myPointers); i++ {
			if myPointers[i].Key == "account_pubkey" {
				myPagedata.AEAddress = myPointers[i].ID
				theAccount = myPointers[i].ID
				fmt.Println(theAccount)
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

		if myPagedata.IPNSAddress != "" {
			ipnsurl := MyNodeConfig.IPFSNode + "/ipns/" + myPagedata.IPNSAddress
			theIPNS = myPagedata.IPNSAddress
			str = httpGet(ipnsurl)
			fmt.Println(str)
			err = json.Unmarshal([]byte(str), &reposites)
			if err != nil {
				fmt.Println(err)
			}

			//Update N reposites?
			for i = 0; i < len(reposites.Reposites); i++ {
				//if reposites.Reposites[i].Name == aensname {
				theIPFS = reposites.Reposites[i].Hash
				theBlock_height = reposites.Reposites[i].Metainfo
				theAENS = reposites.Reposites[i].Name

				sql_check := "SELECT ipfs,block_height FROM aens WHERE aensname='" + theAENS + "' ORDER BY aid DESC LIMIT 1"
				fmt.Println(sql_check)
				rows1, err := db.Query(sql_check)
				checkError(err)
				NeedUpdate := true
				old_ipfs := ""
				old_block_height := ""

				for rows1.Next() {
					err = rows1.Scan(&old_ipfs, &old_block_height)
					fmt.Println(old_ipfs, old_block_height)
					if old_ipfs == theIPFS {
						NeedUpdate = false
					} else {
						if theBlock_height > old_block_height {
							NeedUpdate = true
						}
					}
				}

				if theIPFS == "" {
					NeedUpdate = false
				}

				if NeedUpdate {
					sql_insert = sql_insert + "INSERT INTO aens(aensname,ipns,ipfs,block_height) VALUES('" + theAENS + "','" + theIPNS + "','" + theIPFS + "','" + theBlock_height + "');\n"
					//fmt.Println(sql_insert)
					//_, err = db.Exec(sql_insert)
					//checkError(err)
				}
				//}
			}

		}

	}

	db.Close()

	//insert databases
	db, err = sql.Open("sqlite", dbpath)
	checkError(err)
	fmt.Println(sql_insert)
	_, err = db.Exec(sql_insert)
	checkError(err)
	db.Close()
}
