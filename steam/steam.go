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

type ErrMsg struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type UserInfo struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

func SteamLogOn(sld *SteamLogOnDetails) (userinfo UserInfo, err ErrMsg) {
	myLoginInfo := new(steam.LogOnDetails)
	if sld.Username != "" {
		myLoginInfo.Username = sld.Username
	} else {
		err.Status = true
		err.Message = "用户名不能为空"
		return
	}
	if sld.Password != "" {
		myLoginInfo.Password = sld.Password
	} else {
		err.Status = true
		err.Message = "密码不能为空"
		return
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
			err.Status = false
			client.Disconnect()
			goto END

		case *steam.LogOnFailedEvent:
			// TODO 判断失败类型
			log.Infof(e.Result.String())
			err.Status = true
			err.Message = e.Result.String()
			goto END
		}
	}
END:
	return
}
