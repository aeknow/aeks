package main

import (
	crypto_rand "crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
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

func PubSub_Listening(channel, accountname string, signAccount account.Account) {
	var origin = "http://127.0.0.1:8888/"
	//var url = "ws://127.0.0.1:8888/websocket?user=" + accountname
	var url = "ws://127.0.0.1:8888/websocket"
	ws, err := websocket.Dial(url, "", origin)

	if err != nil {
		//log.Fatal(err)
	}
	MyNodeConfig := DB_GetConfigs()
	sh := shell.NewShell(MyNodeConfig.IPFSAPI)
	sub, err := sh.PubSubSubscribe(channel)

	if err != nil {
		fmt.Println("Sub message error", err)
	}

	for {
		r, err := sub.Next()
		//open the sealed message
		if !strings.Contains(string(r.Data), "Account") {
			r.Data = []byte(MSG_OpenMSG(string(r.Data), signAccount))
		}

		fmt.Println("Pubsub " + channel + " received:" + string(r.Data))

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
					pubtime := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)

					if s.To.Groupname == "" {
						msgstr = "{\"username\":\"" + s.Mine.Username + "\",\"avatar\":\"" + s.Mine.Avatar + "\",\"id\":\"" + s.Mine.Id + "\",\"type\":\"friend\",\"content\":\"" + s.Mine.Content + "\",\"cid\":0,\"mine\":false,\"fromid\":\"" + s.Mine.Id + "\",\"timestamp\":" + pubtime + "}"
					} else {
						msgstr = "{\"username\":\"" + s.Mine.Username + "\",\"groupname\":\"" + s.To.Groupname + "\",\"avatar\":\"" + s.Mine.Avatar + "\",\"id\":\"" + s.To.Id + "\",\"type\":\"group\",\"content\":\"" + s.Mine.Content + "\",\"cid\":0,\"mine\":false,\"fromid\":\"" + s.Mine.Id + "\",\"timestamp\":" + pubtime + ",\"name\":\"" + s.To.Name + "\"}"
					}

					_, err = ws.Write([]byte(msgstr))

					//fmt.Println("msgdto:" + s.To.Id + "::" + msgstr)
					if err != nil {
						fmt.Println(err)
					}
				}
			} else {
				//_, err = ws.Write(r.Data)
				DB_RecordActiveInfo(string(r.Data))
			}
		} else {
			//fmt.Println("self msg:" + string(r.Data))
			fmt.Println("self msg")
		}
		//	}

	}

	//ws.Close()

}

func DB_RecordActiveInfo(pingInfo string) {
	fmt.Println("process income ping :" + pingInfo)
}

//handle received websocket message,broadcast or send to pubsub
func WebSocket_handleChatMsg(message iriswebsocket.Message, nsConn *iriswebsocket.NSConn) {

	accountname := nsConn.Conn.Socket().Request().URL.Query().Get("user")
	//accountname := "ak_bKVvB7iFJKuzH6EvpzLfWKFUpG3qFxUvj8eGwdkFEb7TCTwP8"

	topic := "ak_public" //public topic
	MyNodeConfig := DB_GetConfigs()
	sh := shell.NewShell(MyNodeConfig.IPFSAPI)

	msgBody := string(message.Body)

	fmt.Println("full body:" + msgBody)

	if strings.Contains(msgBody, "sername") {
		//return
	}

	//not ping msg, and not plain local msg
	if !strings.Contains(string(msgBody), "ping") && !strings.Contains(msgBody, "sername") {

		var msg Msg
		err := json.Unmarshal([]byte(msgBody), &msg)
		if err != nil {
			fmt.Println(err)
		}

		//fmt.Println("encoded body:" + string(msg.Body))
		var s ChatMsg
		bodyStr, _ := base64.StdEncoding.DecodeString(msg.Body)
		err = json.Unmarshal(bodyStr, &s)
		if err != nil {
			fmt.Println(err)
		}

		if s.Mine.Id == accountname && !strings.Contains(msgBody, "sername") {
			fmt.Println("Publish message to channel " + s.To.Id)
			//sealed with the target channel accounts
			err = sh.PubSubPublish(s.To.Id, MSG_SealTo(s.To.Id, msgBody))

			//fmt.Println("Sealed: " + MSG_SealTo(s.To.Id, msgBody))

			if err != nil {
				fmt.Println("Publish message failed")

			}
		} else {
			fmt.Println("Received Msg:" + msgBody)
		}
	} else {
		//msgBody = "ping from:" + accountname
		var msg Msg
		err := json.Unmarshal([]byte(msgBody), &msg)
		if err != nil {
			fmt.Println(err)
		}

		if msg.Account == accountname && !strings.Contains(msgBody, "sername") {
			err := sh.PubSubPublish(topic, msgBody)
			fmt.Println("braoadcast to channel " + topic)
			if err != nil {
				fmt.Println("Online braoadcast failed.")

			}

			fmt.Println("Broadcast ping:" + msgBody)
		} else {
			fmt.Println("Received ping:" + msgBody)
		}

	}

	nsConn.Conn.Server().Broadcast(nsConn, message)
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

	MyNodeConfig := DB_GetConfigs()
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

	MyNodeConfig := DB_GetConfigs()

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
		fmt.Println("failed to open box")
	}

	return string(openedMsg)
}
