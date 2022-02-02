package main

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

type clientPage struct {
	Title string
	Host  string
}

func main() {

	perpage = 10
	go IPFS_bootIPFS("default")
	go AE_CheckActiveNode()
	app := iris.New()
	fmt.Println("Web UI is booting...")
	app.HandleDir("/views", "./views")
	app.HandleDir("/themes", "./themes")
	app.HandleDir("/uploads", "./uploads")
	//app.RegisterView(iris.HTML("./views", ".php"))
	//accountname := SESS_GetAccountName(ctx)
	//myTheme := DB_GetConfigItem(accountname, "Theme")

	//template
	tmp := iris.HTML("./themes", ".php")
	tmp.Reload(true)
	app.RegisterView(tmp)
	//System
	app.Get("/", AE_WEB_Index)
	app.Get("/dashboard", AE_WEB_HomePage)
	app.Post("/register", AE_WEB_DoRegister)
	app.Get("/registernew", AE_WEB_RegisterNew)
	app.Post("/login", AE_WEB_CheckLogin)
	app.Get("/logout", AE_WEB_LogOut)
	//import
	app.Get("/import", AE_WEB_ImportUI)
	app.Post("/doimport", AE_WEB_ImportFromMnemonic)
	//export
	app.Get("/export", AE_WEB_ExportFromMnemonic)

	//Wallet
	app.Get("/wallet", AE_WEB_Wallet)
	app.Post("/transaction", AE_WEB_MakeTranscaction)

	//Haeme
	app.Get("/haeme", iHaeme)
	//app.Get("/updatestatic", iUpdateStatic)
	app.Get("/blog", iBlog)
	app.Get("/newblog", iNewBlog)
	app.Post("/uploadblogimage", iBlogUploadFile)
	app.Post("/saveblog", iSaveBlog)
	//app.Get("/buildblog", iBuildSite)
	app.Get("/editpage", iEditBlog)
	app.Get("/delpage", iDelBlog)
	app.Get("/setsite", iSetSite)
	app.Post("/savesitesetting", iSaveSetSite)

	app.Get("/goaens", iGoAENS)

	//New data management
	app.Get("/view", iView)

	//AENS
	app.Get("/aens", AENS_getAENS)
	app.Get("/updateallaens", AENS_UpdateALLOnce)

	app.Get("/aensbidding", AENS_getAENSBidding)
	app.Post("/queryaens", AENS_WEB_QueryAENS)
	app.Post("/regaens", AENS_WEB_DoRegAENS)
	app.Post("/bidaens", AENS_WEB_DoBidAENS)
	app.Get("/transfername", AENS_WEB_TransferAENS)
	app.Post("/dotransferaens", AENS_WEB_DoTransferAENS)
	app.Get("/updatename", AENS_WEB_UpdateAENS)
	app.Post("/updatenamepointer", AENS_WEB_DoUpdateAENS)
	app.Post("/expertupdatenamepointer", AENS_WEB_ExpertDoUpdateAENS)

	app.Get("/getaens", AENS_GetData)

	//Contracts
	app.Get("/contracts", Contract_WEB_ContractsHome)
	app.Get("/deploy", Contract_WEB_DeployContractUI)
	app.Post("/dodeploy", Contract_WEB_DoDeployContract)
	app.Get("/call", Contract_WEB_CallContractUI)
	app.Post("/docall", Contract_WEB_DoCallContract)
	app.Get("/decodecall", Contract_WEB_DecodeContractCall)
	app.Post("/dodecode", Contract_WEB_DoDecodeContractCall)
	//aex-9 tokens
	app.Get("/deploytoken", Contract_WEB_DeployTokenUI)
	app.Post("/dodeploytoken", Contrat_WEB_DoDeployToken)
	//aex-9 token
	app.Get("/viewtoken", Contract_WEB_Token)
	app.Get("/token", Contract_WEB_getToken)
	app.Post("/transfertoken", Contratc_WEB_TokenTransfer)

	ws := websocket.New(websocket.DefaultGorillaUpgrader, websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			fmt.Printf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())
			//msg.To = globalAccount.Address
			//handleChatMsg(msg, nsConn)

			return nil
		},
	})

	ws.OnConnect = func(c *websocket.Conn) error {
		fmt.Printf("[%s] Connected to server!", c.ID())
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		fmt.Printf("[%s] Disconnected from server", c.ID())
	}

	app.Get("/websocket", websocket.Handler(ws))

	app.Get("/chattest", func(ctx iris.Context) {
		ctx.View("client.php", clientPage{"Client Page", "localhost:8888"})
	})

	//handle proxy ipfs content for editor.md
	app.Get("/ipfs/{anythingparameter:path}", func(ctx iris.Context) {
		MyNodeConfig := DB_GetConfigs()
		paramValue := ctx.Params().Get("anythingparameter")
		ipfsUrl := MyNodeConfig.IPFSNode + "/ipfs/" + paramValue
		resp, err := http.Get(ipfsUrl)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		//分片逐步写入
		buf := make([]byte, 4096)
		for {
			n, err := resp.Body.Read(buf)
			ctx.Write(buf[:n])
			if err != nil {
				break
			}

		}
	})

	//handle proxy ipns content for editor.md
	app.Get("/ipns/{anythingparameter:path}", func(ctx iris.Context) {
		MyNodeConfig := DB_GetConfigs()

		paramValue := ctx.Params().Get("anythingparameter")
		ipnsUrl := MyNodeConfig.IPFSNode + "/ipns/" + paramValue
		//fmt.Println("ipnsurl:" + ipnsUrl)
		resp, err := http.Get(ipnsUrl)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		//分片逐步写入
		buf := make([]byte, 4096)
		for {
			n, err := resp.Body.Read(buf)
			ctx.Write(buf[:n])
			if err != nil {
				break
			}

		}
	})

	//test functions for ipfs
	app.Get("/ipks/{anythingparameter:path}", func(ctx iris.Context) {

		MyNodeConfig := DB_GetConfigs()
		paramValue := ctx.Params().Get("anythingparameter")
		ipfsUrl := MyNodeConfig.IPFSNode + "/ipfs/" + paramValue
		fmt.Println("http.Get =>", ipfsUrl)

		resp, err := http.Get(ipfsUrl)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		//分片逐步写入
		buf := make([]byte, 4096)
		for {
			n, err := resp.Body.Read(buf)
			if err != nil {
				break
			}
			ctx.Write(buf[:n])
		}

		//ctx.Exec("GET", ipfsUrl)
		//fmt.Println("Got?")

	})

	app.Run(iris.Addr(":8888"))

}
