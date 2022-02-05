package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/kataras/iris/v12"
	iriswebsocket "github.com/kataras/iris/v12/websocket"
	"golang.org/x/net/websocket"
)

type PageChat struct {
	Account     string
	PageContent template.HTML
	PageTitle   string
	ChatTopic   string
}

type Msg struct {
	Signature string
	Body      ChatMsg
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
	ctx.View("mainroad/client.php", PageChat{Account: accountname, PageTitle: "Chaet"})
}

func Chaet_SignJson(ctx iris.Context) {
	//accountname := SESS_GetAccountName(ctx)
	c := &ChatMsg{}
	fmt.Println("Ready to get signature")
	if err := ctx.ReadJSON(c); err != nil {
		fmt.Println("Json message error", err)
	} else {
		b, err := json.Marshal(c)
		if err != nil {
			fmt.Println("json err:", err)
		}

		MysignAccount := SESS_GetAccount(ctx)
		signature := base64.StdEncoding.EncodeToString(MysignAccount.Sign(b))

		fmt.Println("string:" + string(b))

		fmt.Println("sig:" + signature)
		ctx.HTML(signature)

		//fmt.Println(c)
		//ctx.JSON(c)
	}
}

func ListeningLocal(channel, accountname string) {
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

		fmt.Println("Pubsub received:" + string(r.Data))

		if err != nil {
			fmt.Println("Message error", err)
		}
		if !strings.Contains(string(r.Data), "username") && !strings.Contains(string(r.Data), "Username") && !strings.Contains(string(r.Data), "groupname") {
			msgBody := "{\"username\":\"localakak\",\"message\":\"" + string(r.Data) + "\"}"
			_, err = ws.Write([]byte(msgBody))
			fmt.Println("Msg in json:" + msgBody)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Sent: %s\n", string(r.Data))
		} else {
			var msg Msg
			err := json.Unmarshal([]byte(r.Data), &msg)
			if err != nil {
				fmt.Println(err)
			}

			s := msg.Body

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

				fmt.Println("msgdto:" + s.To.Id + "::" + msgstr)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		//fmt.Println(r.From)
		fmt.Println(string(r.Data))
	}

	ws.Close()

}

func handleChatMsg(message iriswebsocket.Message, nsConn *iriswebsocket.NSConn) {
	//
	/*利用Message定义聊天消息的结构，需要包含
	From string
	To string
	Signature string //如果是公开消息
	Msgtype string //0-palin, 1-crypted
	Body string//json string, 包含nonce
	*/

	accountname := nsConn.Conn.Socket().Request().URL.Query().Get("user")
	//accountname := "ak_bKVvB7iFJKuzH6EvpzLfWKFUpG3qFxUvj8eGwdkFEb7TCTwP8"

	topic := "ak_fCCw1JEkvXdztZxk8FRGNAkvmArhVeow89e64yX4AxbCPrVh5"

	MyNodeConfig := DB_GetConfigs()
	sh := shell.NewShell(MyNodeConfig.IPFSAPI)

	msgBody := string(message.Body)

	var msg Msg
	err := json.Unmarshal([]byte(msgBody), &msg)
	if err != nil {
		fmt.Println(err)
	}

	s := msg.Body

	if msgBody != "ping" {
		fmt.Println("Ready to decode msg....")

		fmt.Println("Msg From " + s.Mine.Id + " to...." + s.To.Id)
		//SmartPrint(s)

		if s.Mine.Id == accountname {
			err = sh.PubSubPublish(topic, msgBody)
			if err != nil {
				fmt.Println("Publish message failed")

			}
		} else {
			fmt.Println("Received Msg:" + msgBody)
		}
	} else {
		msgBody = "ping from:" + accountname

		err := sh.PubSubPublish(topic, msgBody)
		if err != nil {
			fmt.Println("Online braoadcast failed.")

		}

		fmt.Println("Broadcast ping:" + msgBody)
		//RecordPingTimestamp()
	}

	if msgBody == accountname+" Online" {
		err := sh.PubSubPublish(topic, msgBody)
		if err != nil {
			fmt.Println("Online braoadcast failed.")

		}
	}

	nsConn.Conn.Server().Broadcast(nsConn, message)
}
