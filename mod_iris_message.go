package main

import (
	"encoding/json"
	"fmt"
	"html/template"

	shell "github.com/ipfs/go-ipfs-api"

	"github.com/kataras/iris/v12"
	iriswebsocket "github.com/kataras/iris/v12/websocket"
)

type PageChat struct {
	Account     string
	PageContent template.HTML
	PageTitle   string
	ChatTopic   string
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

func ListeningLocal(accountname string) {
	MyNodeConfig := DB_GetConfigs()
	sh := shell.NewShell(MyNodeConfig.IPFSAPI)
	sub, err := sh.PubSubSubscribe(accountname)

	if err != nil {
		fmt.Println("Sub message error", err)
	}

	for {
		r, err := sub.Next()
		if err != nil {
			fmt.Println("Message error", err)
		}

		fmt.Println(r.From)
		fmt.Println(string(r.Data))
	}

}

func handleChatMsg(msg iriswebsocket.Message, nsConn *iriswebsocket.NSConn) {
	//
	/*利用Message定义聊天消息的结构，需要包含
	From string
	To string
	Signature string //如果是公开消息
	Msgtype string //0-palin, 1-crypted
	Body string//json string, 包含nonce
	*/

	//SmartPrint(msg)

	//	accountname := SESS_GetAccountName(ctx)
	accountname := "ak_bKVvB7iFJKuzH6EvpzLfWKFUpG3qFxUvj8eGwdkFEb7TCTwP8"

	topic := "ak_fCCw1JEkvXdztZxk8FRGNAkvmArhVeow89e64yX4AxbCPrVh5"

	MyNodeConfig := DB_GetConfigs()
	sh := shell.NewShell(MyNodeConfig.IPFSAPI)

	msgBody := string(msg.Body)

	if msgBody != "ping" {
		fmt.Println("Ready to decode msg....")
		var s ChatMsg
		err := json.Unmarshal([]byte(msgBody), &s)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Msg From " + s.Mine.Id + " to...." + s.To.Id)
		//SmartPrint(s)

		//topic := msg.To
		//if strings.Index(msgBody, "avatar") > 0 {
		if s.Mine.Id == accountname {
			err = sh.PubSubPublish(topic, msgBody)
			if err != nil {
				fmt.Println("Publish message failed")

			}
		} else {
			fmt.Println("Received Msg:" + msgBody)
		}
	} else {
		//myapi, err := cmdenv.GetApi(myenv, myreq)
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

	nsConn.Conn.Server().Broadcast(nsConn, msg)
}
