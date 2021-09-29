package controllers

import (
	"encoding/json"
	"net/http"

	"steamBackend/conf"
	"steamBackend/steam"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type data struct {
	Action string                   `json:"action"`
	Count  int                      `json:"count"`
	Data   *steam.SteamLogOnDetails `json:"data"`
}

var upGrande = websocket.Upgrader{
	//设置允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 处理WebSocket请求
func Ws(context *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrande.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		log.Error("websocket连接错误", err.Error())
	}
	startMsg := gin.H{
		"action": "connect",
		"result": "success",
		"server": conf.Conf.Server.ServerName,
	}
	err = ws.WriteJSON(startMsg)
	if err != nil {
		log.Errorf("发送WebSocket消息失败: ", err.Error())
	}
	defer ws.Close()
	for {
		// 读取ws中的数据
		_, message, err := ws.ReadMessage()
		if err != nil {
			// 客户端关闭连接时也会进入
			log.Error("请求断开，或读取消息错误", err.Error())
			break
		}
		msg := &data{}
		json.Unmarshal(message, msg)
		// fmt.Println(msg.Action)
		// fmt.Println(msg.Count)
		// fmt.Println(mt)
		switch msg.Action {
		case "ping":
			// 返回JSON字符串，借助gin的gin.H实现
			resMsg := gin.H{
				"action": "pong",
				"count":  msg.Count,
			}
			// 写入ws数据
			err = ws.WriteJSON(resMsg)
			if err != nil {
				log.Errorf("发送WebSocket消息失败: ", err.Error())
				break
			}
		case "logOn":
			log.Infof("%+v", msg.Data)
			userinfo := steam.SteamLogOn(msg.Data)
			if userinfo.Error {
				resMsg := gin.H{
					"action": "authCode",
				}
				err = ws.WriteJSON(resMsg)
				if err != nil {
					log.Errorf("发送WebSocket消息失败: ", err.Error())
					break
				}
			} else {
				resMsg := gin.H{
					"action":  "logOn",
					"result":  "success",
					"name":    userinfo.Name,
					"country": userinfo.Country,
				}
				err = ws.WriteJSON(resMsg)
				if err != nil {
					log.Errorf("发送WebSocket消息失败: ", err.Error())
					break
				}
			}
		}
	}
}
