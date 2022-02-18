package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	crypto_rand "crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aeternity/aepp-sdk-go/v9/account"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/kataras/iris/v12"
	iriswebsocket "github.com/kataras/iris/v12/websocket"
	"golang.org/x/net/websocket"

	aebinary "github.com/aeternity/aepp-sdk-go/v9/binary"
	"github.com/jdgcs/ed25519/extra25519"
	"golang.org/x/crypto/nacl/box"
)

type PageChat struct {
	Account     string
	PageContent template.HTML
	PageTitle   string
	ChatTopic   string
}

type Msg struct {
	Signature string
	Body      string
	Account   string
	Mtype     string
}

type ChatMsg struct {
	Mine SenderMsg
	To   ReceiverMsg
}

type SenderMsg struct {
	Username  string
	Groupname string
	Avatar    string
	Id        string
	Type      string
	Content   string
	Cid       string
	Mine      bool
	Fromid    string
	Timestamp int64
	Name      string
}
type ReceiverMsg struct {
	Username  string
	Groupname string
	Avatar    string
	Id        string
	Type      string
	Content   string
	Cid       string
	Mine      bool
	Fromid    string
	Timestamp uint64
	Sign      string
	Name      string
}

func Chaet_UI(ctx iris.Context) {
	accountname := SESS_GetAccountName(ctx)
	//MysignAccount := SESS_GetAccount(ctx)

	//sealed := MSG_SealTo(accountname, "Hello, world!")
	//fmt.Println("sealed: " + sealed)
	//opened := MSG_OpenMSG(sealed, *MysignAccount)
	//fmt.Println("opened: " + opened)
	//origData := []byte("Hello World") // 待加密的数据
	//key := []byte("0123456789abcdef") // 加密的密钥
	//encrypted := MSG_AesEncryptCBC(origData, key)
	//log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	//fmt.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	//decrypted := MSG_AesDecryptCBC(encrypted, key)
	//encrypted := MSG_SealGroupMSG("123", "Hello, world!")
	//fmt.Println("en: " + encrypted)
	//fmt.Println("de: " + MSG_OpenGroupMSG("123", encrypted))

	ctx.View("mainroad/client.php", PageChat{Account: accountname, PageTitle: "Chaet"})
}

func Chaet_SignJson(ctx iris.Context) {
	//accountname := SESS_GetAccountName(ctx)

	body := ctx.FormValue("body")

	MysignAccount := SESS_GetAccount(ctx)
	signature := base64.StdEncoding.EncodeToString(MysignAccount.Sign([]byte(body)))

	//fmt.Println("string:" + body)

	//fmt.Println("sig:" + signature)
	ctx.HTML(signature)

	//fmt.Println(c)
	//ctx.JSON(c)

}

//Start listening the pubsub channels, for the whole message system.
func PubSub_StartListen(accountname string, signAccount account.Account) {
	fmt.Println("Check IPFS status...")
	for {
		sh := shell.NewShell("127.0.0.1:5001")
		cid, err := sh.Add(strings.NewReader("Hello from AEKs!"))

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err)
		} else {
			fmt.Println("IPFS booted, " + cid)
			goto Loop
		}

		time.Sleep(time.Duration(1) * time.Second)
	}

Loop:
	fmt.Println("Start listening channels...")
	//start message listening
	go PubSub_Listening(accountname, accountname, signAccount)
	go PubSub_Listening("ak_public", accountname, signAccount)                                                 //test channel
	go PubSub_Listening("group_bKVvB7iFJKuzH6EvpzLfWKFUpG3qFxUvj8eGwdkFEb7TCTwP8_1", accountname, signAccount) //test group channel
	go PubSub_Listening("group_fCCw1JEkvXdztZxk8FRGNAkvmArhVeow89e64yX4AxbCPrVh5_2", accountname, signAccount) //test group channel

	//If I am acting as a proxy, start listening the proxy channel
	if MSG_AmIProxy(accountname) {
		go PubSub_ProxyListening(MyNodeConfig.PubsubProxy, accountname, signAccount)
	}
}

//start listening a single channel, decode&process the messages
func PubSub_ProxyListening(channel, accountname string, signAccount account.Account) {
	sh := shell.NewShell(MyNodeConfig.IPFSAPI)
	sub, err := sh.PubSubSubscribe(channel)

	if err != nil {
		fmt.Println("Sub message error", err)
	}

	for {
		r, err := sub.Next()
		if err != nil {
			fmt.Println("Message error", err)
		}

		//put messages to proxy
		if !strings.Contains(string(r.Data), "put") {

		}

		//get messages from proxy
		if !strings.Contains(string(r.Data), "get") {

		}

	}

}

//start listening a single channel, decode&process the messages
func PubSub_Listening(channel, accountname string, signAccount account.Account) {
	var origin = "http://127.0.0.1:8888/"
	var url = "ws://127.0.0.1:8888/websocket"
	ws, err := websocket.Dial(url, "", origin)

	if err != nil {
		//log.Fatal(err)
	}
	//MyNodeConfig := DB_GetConfigs()
	sh := shell.NewShell(MyNodeConfig.IPFSAPI)
	sub, err := sh.PubSubSubscribe(channel)

	if err != nil {
		fmt.Println("Sub message error", err)
	}

	for {
		r, err := sub.Next()
		//open the sealed message
		if !strings.Contains(string(r.Data), "Account") {
			if strings.Contains(channel, "group") {
				r.Data = []byte(MSG_OpenGroupMSG(channel, string(r.Data)))
			} else {
				r.Data = []byte(MSG_OpenMSG(string(r.Data), signAccount))
			}
		}
		//fmt.Println(r.From)
		//fmt.Println(string(r.Seqno))
		//fmt.Println(r.TopicIDs)
		//	fmt.Println("Pubsub " + channel + " received:" + string(r.Data))

		//if !strings.Contains(string(r.Data), "sername") {

		if err != nil {
			fmt.Println("Message error", err)
		}

		var msg Msg
		err = json.Unmarshal([]byte(r.Data), &msg)
		if err != nil {
			fmt.Println(err)
		}

		theSig, _ := base64.StdEncoding.DecodeString(msg.Signature)
		//fmt.Println("Message base64 encoded msg:" + msg.Body + "\nSig:" + msg.Signature)

		//if !strings.Contains(string(r.Data), "username") && !strings.Contains(string(r.Data), "Username") && !strings.Contains(string(r.Data), "groupname") {
		sigVerify, err := account.Verify(msg.Account, []byte(msg.Body), theSig)

		if sigVerify {
			fmt.Println("VERIFIED")
		} else {
			fmt.Println(err)
			fmt.Println("MSG UN-VERIFIED")
		}

		if !strings.Contains(string(r.Data), accountname) {
			if !strings.Contains(string(r.Data), "ping") {
				var s ChatMsg
				bosyStr, _ := base64.StdEncoding.DecodeString(msg.Body)
				err = json.Unmarshal(bosyStr, &s)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println("msgto:" + s.To.Id)

				//I am the receiver or group message and I am not the sender
				if (s.To.Id == accountname || s.To.Groupname != "") && s.Mine.Id != accountname {

					if err != nil {
						fmt.Println("error:", err)
					}

					msgstr := ""
					pubtime := strconv.FormatInt(int64(s.To.Timestamp), 10)

					if s.To.Groupname == "" {
						msgstr = "{\"username\":\"" + s.Mine.Username + "\",\"avatar\":\"" + s.Mine.Avatar + "\",\"id\":\"" + s.Mine.Id + "\",\"type\":\"friend\",\"content\":\"" + strings.Replace(html.EscapeString(s.Mine.Content), "\n", "\\n", -1) + "\",\"cid\":0,\"mine\":false,\"fromid\":\"" + s.Mine.Id + "\",\"timestamp\":" + pubtime + "}"

						DB_RecordMsgs(accountname, s.Mine.Id, s.To.Id, DB_IndexCJKText(strings.Replace(html.EscapeString(s.Mine.Content), "\n", "\\n", -1), segmenter), string(r.Data), "friend", pubtime)

					} else {
						msgstr = "{\"username\":\"" + s.Mine.Username + "\",\"groupname\":\"" + s.To.Groupname + "\",\"avatar\":\"" + s.Mine.Avatar + "\",\"id\":\"" + s.To.Id + "\",\"type\":\"group\",\"content\":\"" + strings.Replace(html.EscapeString(s.Mine.Content), "\n", "\\n", -1) + "\",\"cid\":0,\"mine\":false,\"fromid\":\"" + s.Mine.Id + "\",\"timestamp\":" + pubtime + ",\"name\":\"" + s.To.Name + "\"}"
						DB_RecordMsgs(accountname, s.Mine.Id, s.To.Id, DB_IndexCJKText(strings.Replace(html.EscapeString(s.Mine.Content), "\n", "\\n", -1), segmenter), string(r.Data), "group", pubtime)

					}

					_, err = ws.Write([]byte(msgstr))

					//fmt.Println("msgdto:" + s.To.Id + "::" + msgstr)
					if err != nil {
						fmt.Println(err)
					}
				}
			} else {
				//_, err = ws.Write(r.Data)
				if sigVerify {
					DB_RecordActiveInfo(accountname, msg.Account)
				}

			}
		} else { //Record msgs out

			fmt.Println("self ping:" + string(r.Data))

			if sigVerify {
				DB_RecordActiveInfo(accountname, msg.Account)
			}

		}
		//	}

	}
	//ws.Close()

}

//handle received websocket message,broadcast or send to pubsub
func WebSocket_handleChatMsg(message iriswebsocket.Message, nsConn *iriswebsocket.NSConn) {

	accountname := nsConn.Conn.Socket().Request().URL.Query().Get("user")
	publicChannel := "ak_public" //public topic

	sh := shell.NewShell(MyNodeConfig.IPFSAPI)

	msgBody := string(message.Body)

	fmt.Println("full body:" + msgBody)

	if strings.Contains(msgBody, "sername") {
		//return
	}

	var msg Msg
	err := json.Unmarshal([]byte(msgBody), &msg)
	if err != nil {
		fmt.Println("umarshal err: ")
		fmt.Println(err)
	}

	//not ping msg, and not plain local msg
	if msg.Mtype != "ping" && !strings.Contains(msgBody, "sername") && msg.Mtype != "online" {
		//fmt.Println("encoded body:" + string(msg.Body))
		var s ChatMsg
		bodyStr, _ := base64.StdEncoding.DecodeString(msg.Body)
		err = json.Unmarshal(bodyStr, &s)
		if err != nil {
			fmt.Println(err)
		}

		if s.Mine.Id == accountname {
			fmt.Println("Publish message to channel " + s.To.Id)

			pubtime := strconv.FormatInt(int64(s.To.Timestamp), 10)

			if msg.Mtype == "group" {
				//sealed with the target group passwords and record to the database
				rawMSG := MSG_SealGroupMSG(s.To.Id, msgBody)
				err = sh.PubSubPublish(s.To.Id, rawMSG)
				DB_RecordMsgs(accountname, s.Mine.Id, s.To.Id, DB_IndexCJKText(strings.Replace(html.EscapeString(s.Mine.Content), "\n", "\\n", -1), segmenter), msgBody, "group", pubtime)

				if MSG_CheckProxy(accountname, s.Mine.Id, s.To.Id) {
					err = sh.PubSubPublish(MyNodeConfig.PubsubProxy, s.To.Id+":"+rawMSG)
				}
			}

			if msg.Mtype == "private" {
				//sealed with the target user's channel accounts and record to the database
				rawMSG := MSG_SealTo(s.To.Id, msgBody)
				//fmt.Println(msgBody)
				err = sh.PubSubPublish(s.To.Id, rawMSG)
				DB_RecordMsgs(accountname, s.Mine.Id, s.To.Id, DB_IndexCJKText(strings.Replace(html.EscapeString(s.Mine.Content), "\n", "\\n", -1), segmenter), msgBody, "friend", pubtime)

				if MSG_CheckProxy(accountname, s.Mine.Id, s.To.Id) {
					err = sh.PubSubPublish(MyNodeConfig.PubsubProxy, s.To.Id+":"+rawMSG)
				}
			}

			//fmt.Println("Sealed: " + MSG_SealTo(s.To.Id, msgBody))

			if err != nil {
				fmt.Println("Publish message failed")

			}
		} else {
			fmt.Println("Received Msg:" + msgBody)
		}
	} else {
		if msg.Account == accountname && !strings.Contains(msgBody, "sername") {
			err := sh.PubSubPublish(publicChannel, msgBody)
			fmt.Println("broadcast ping to channel " + publicChannel)
			if err != nil {
				fmt.Println("Braoadcast ping failed.")

			}

			//fmt.Println("Broadcast ping:" + msgBody)
		} else {
			fmt.Println("Received ping:" + msgBody)
		}

	}

	if msg.Mtype == "online" && msg.Account == accountname {
		//TODO:get messages from database and broadcast
		go DB_GetUnreadMsgs(accountname, message, nsConn)
	} else {
		nsConn.Conn.Server().Broadcast(nsConn, message)
		fmt.Println("Default cast:" + string(message.Body))
	}

}

func MSG_UploadImage(ctx iris.Context) {
	//filename := ctx.FormValue("filename")
	//fmt.Println("\n\nfile:" + filename + "\n\n")
	file, info, err := ctx.FormFile("file")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}

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

	//MyNodeConfig := DB_GetConfigs()
	myfile := ""
	sh := shell.NewShell(MyNodeConfig.IPFSAPI)
	if ostype == "windows" {
		myfile = ".\\uploads\\" + fname
	} else {
		myfile = "./uploads/" + fname

	}

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
  "code": 0 
  ,"msg": "" 
  ,"data": {
    "src": "` + MyNodeConfig.IPFSNode + `/ipfs/` + cid + `" 
  }
}`

	err = os.Remove(myfile)
	if err != nil {
		fmt.Println("Delete uplaod file failed.", err)
	}
	ctx.Writef(uploadedImageValue)
}

func MSG_UploadFile(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	//filename := ctx.FormValue("filename")
	//fmt.Println("\n\nfile:" + filename + "\n\n")
	file, info, err := ctx.FormFile("file")
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

	//MyNodeConfig := DB_GetConfigs()

	sh := shell.NewShell(MyNodeConfig.IPFSAPI)
	myfile := ""
	if ostype == "windows" {
		myfile = ".\\uploads\\" + fname
	} else {
		myfile = "./uploads/" + fname

	}

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

	uploadedFileValue := `{
		"code": 0 
		,"msg": "" 
		,"data": {
		  "src": "` + MyNodeConfig.IPFSNode + `/ipfs/` + cid + `" 
		  ,"name": "` + fname + `"
		}
	  }`
	err = os.Remove(myfile)
	if err != nil {
		fmt.Println("Delete uplaod file failed.", err)
	}

	ctx.Writef(uploadedFileValue)
}

func MSG_SealGroupMSG(groupid string, message string) string {
	key := []byte(DB_GetGroupKey(groupid)) // 加密的密钥
	encrypted := MSG_AesEncryptCBC([]byte(message), key)

	return base64.StdEncoding.EncodeToString(encrypted)
}

func MSG_OpenGroupMSG(groupid string, message string) string {
	//fmt.Println(DB_GetGroupKey(groupid) + message)
	key := []byte(DB_GetGroupKey(groupid)) // 加密的密钥
	encrypted, _ := base64.StdEncoding.DecodeString(message)
	return string(MSG_AesDecryptCBC(encrypted, key))
}

//seal message with the target address
func MSG_SealTo(ToAddress, Message string) string {
	recipientPublicKey, _, _ := box.GenerateKey(crypto_rand.Reader) //assume a key
	toPublicKey, _ := aebinary.Decode(ToAddress)

	var publicKeySlice [32]byte
	copy(publicKeySlice[0:32], toPublicKey)
	myrecipientPublicKey := &publicKeySlice
	extra25519.PublicKeyToCurve25519(recipientPublicKey, myrecipientPublicKey)

	byteMSG := []byte(Message)

	sealedMsg, err := box.SealAnonymous(nil, byteMSG, recipientPublicKey, nil)
	if err != nil {
		fmt.Println("Unexpected error sealing ", err)
	}

	return base64.StdEncoding.EncodeToString(sealedMsg)
}

//open sealed message
func MSG_OpenMSG(Message string, signAccount account.Account) string {
	fmt.Println(Message)

	fmt.Println(signAccount.Address)

	recipientPublicKey, openPrivateKey, _ := box.GenerateKey(crypto_rand.Reader) //assume a key

	toPublicKey, _ := aebinary.Decode(signAccount.Address)

	var publicKeySlice [32]byte
	var privateKeySlice [64]byte

	copy(publicKeySlice[0:32], toPublicKey)
	myrecipientPublicKey := &publicKeySlice
	extra25519.PublicKeyToCurve25519(recipientPublicKey, myrecipientPublicKey)

	copy(privateKeySlice[0:64], signAccount.SigningKey)
	myrecipientPrivateKey := &privateKeySlice
	extra25519.PrivateKeyToCurve25519(openPrivateKey, myrecipientPrivateKey)

	encrypted, _ := base64.StdEncoding.DecodeString(Message)
	openedMsg, ok := box.OpenAnonymous(nil, encrypted, recipientPublicKey, openPrivateKey)

	if !ok {
		fmt.Println("failed to open private box")
	}

	return string(openedMsg)
}

//Encrypt messages with AES
func MSG_AesEncryptCBC(origData []byte, key []byte) (encrypted []byte) {
	// 分组秘钥
	// NewCipher该函数限制了输入key的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted
}

//Decrypt messages with AES
func MSG_AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)                              // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return decrypted
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//Check if we should use the proxy mode for the offline messages
func MSG_CheckProxy(accountname, fromid, toid string) bool {
	//TODO:proxy mode or offline mode
	return true
}

//Check if we should use the proxy mode for the offline messages
func MSG_AmIProxy(accountname string) bool {
	//TODO:proxy mode or offline mode
	return true
}

//Get the secret key of each group, the length MUST be 16, 24 or 32
func DB_GetGroupKey(groupid string) string {
	//TODO:get group key from list data, which can be set by the owner
	return "0123456789abcdef"
}

func DB_RecordActiveInfo(accountname, activeaccount string) {
	dbpath := "./data/accounts/" + accountname + "/chaet.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)
	sql_check := "SELECT lastactive FROM users WHERE id='" + activeaccount + "'"
	rows, err := db.Query(sql_check)
	checkError(err)

	fmt.Println("Record active of " + activeaccount)
	NeedInsert := true
	for rows.Next() {
		NeedInsert = false
	}
	lastactive := strconv.FormatInt(time.Now().Unix(), 10)
	if NeedInsert {
		sql_insert := "INSERT INTO users(id,lastactive) VALUES('" + activeaccount + "','" + lastactive + "')"
		db.Exec(sql_insert)
		fmt.Println("Insert new active user: " + activeaccount)
	} else {
		sql_update := "UPDATE users set lastactive='" + lastactive + "' WHERE id='" + activeaccount + "'"
		db.Exec(sql_update)
		fmt.Println("Update active user: " + activeaccount)
	}

	db.Close()
}

//Record private and group messages to the database
func DB_RecordMsgs(accountname, from, to, body, raw, mtype, pubtime string) {
	dbpath := "./data/accounts/" + accountname + "/chaet.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)

	sql_insert := "INSERT INTO msgs(fromid, toid, body,raw,mtype, pubtime) VALUES('" + from + "','" + to + "','" + body + "','" + raw + "','" + mtype + "','" + pubtime + "')"

	fmt.Println(sql_insert)
	_, err = db.Query(sql_insert)
	checkError(err)

	db.Close()
}

//get the messages that are not broadcasted, then rebroadcast to the UI
func DB_GetUnreadMsgs(accountname string, message iriswebsocket.Message, nsConn *iriswebsocket.NSConn) {

	time.Sleep(time.Duration(1) * time.Second)
	var origin = "http://127.0.0.1:8888/"
	var url = "ws://127.0.0.1:8888/websocket"
	ws, err := websocket.Dial(url, "", origin)

	if err != nil {
		fmt.Println("Rebroadcast websocket error")
	}

	dbpath := "./data/accounts/" + accountname + "/chaet.db"
	db, err := sql.Open("sqlite", dbpath)
	checkError(err)
	sql_check := "SELECT lastactive FROM users WHERE id='" + accountname + "'"
	rows, err := db.Query(sql_check)
	checkError(err)

	lastactive := "9999014887789"

	for rows.Next() {
		err = rows.Scan(&lastactive)
	}

	sql_getmsg := "SELECT raw,toid FROM msgs WHERE pubtime > '" + lastactive + "'"
	rows, err = db.Query(sql_getmsg)
	checkError(err)

	rawmsg := ""
	toid := ""
	for rows.Next() {
		err = rows.Scan(&rawmsg, &toid)

		var msg Msg
		err := json.Unmarshal([]byte(rawmsg), &msg)
		if err != nil {
			fmt.Println("umarshal err: ")
			fmt.Println(err)
		}

		var s ChatMsg
		bosyStr, _ := base64.StdEncoding.DecodeString(msg.Body)
		err = json.Unmarshal(bosyStr, &s)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("msgto:" + s.To.Id)

		//I am the receiver or group message and I am not the sender
		if (s.To.Id == accountname || s.To.Groupname != "") && s.Mine.Id != accountname {

			if err != nil {
				fmt.Println("error:", err)
			}

			msgstr := ""
			pubtime := strconv.FormatInt(int64(s.To.Timestamp), 10)

			if s.To.Groupname == "" {
				msgstr = "{\"username\":\"" + s.Mine.Username + "\",\"avatar\":\"" + s.Mine.Avatar + "\",\"id\":\"" + s.Mine.Id + "\",\"type\":\"friend\",\"content\":\"" + strings.Replace(html.EscapeString(s.Mine.Content), "\n", "\\n", -1) + "\",\"cid\":0,\"mine\":false,\"fromid\":\"" + s.Mine.Id + "\",\"timestamp\":" + pubtime + "}"
			} else {
				msgstr = "{\"username\":\"" + s.Mine.Username + "\",\"groupname\":\"" + s.To.Groupname + "\",\"avatar\":\"" + s.Mine.Avatar + "\",\"id\":\"" + s.To.Id + "\",\"type\":\"group\",\"content\":\"" + strings.Replace(html.EscapeString(s.Mine.Content), "\n", "\\n", -1) + "\",\"cid\":0,\"mine\":false,\"fromid\":\"" + s.Mine.Id + "\",\"timestamp\":" + pubtime + ",\"name\":\"" + s.To.Name + "\"}"
			}

			_, err = ws.Write([]byte(msgstr))
			time.Sleep(time.Duration(100) * time.Millisecond)
			//fmt.Println("msgdto:" + s.To.Id + "::" + msgstr)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("brocast from db: " + msgstr)
		}

	}
	db.Close()

	DB_RecordActiveInfo(accountname, accountname)
}
