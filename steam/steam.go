package steam

import (
	"github.com/Philipp15b/go-steam/v2"
	log "github.com/sirupsen/logrus"
)

type SteamLogOnDetails struct {
	Username      string
	Password      string
	AuthCode      string
	TwoFactorCode string
}

type resInfo struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Error   bool   `json:"error"`
}

func SteamLogOn(sld *SteamLogOnDetails) (userinfo resInfo) {
	myLoginInfo := new(steam.LogOnDetails)
	if sld.Username != "" {
		myLoginInfo.Username = sld.Username
	}
	if sld.Password != "" {
		myLoginInfo.Password = sld.Password
	}
	if sld.AuthCode != "" {
		myLoginInfo.AuthCode = sld.AuthCode
	}
	if sld.TwoFactorCode != "" {
		myLoginInfo.TwoFactorCode = sld.TwoFactorCode
	}

	log.Infof("%s 用户正在登录...", myLoginInfo.Username)

	client := steam.NewClient()

	client.Connect()
	for event := range client.Events() {
		switch e := event.(type) {
		case *steam.ConnectedEvent:
			client.Auth.LogOn(myLoginInfo)

		case *steam.AccountInfoEvent:
			userinfo.Name = e.PersonaName
			userinfo.Country = e.Country
			userinfo.Error = false
			client.Disconnect()
			goto END

		case *steam.LogOnFailedEvent:
			// TODO 判断失败类型
			userinfo.Name = "error"
			userinfo.Country = "error"
			userinfo.Error = true
			goto END
		}
	}
END:
	return
}
