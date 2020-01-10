package handlers

import (
	"cherry/base"
	"fmt"
	"model"
	"net/http"
	"strconv"
)

// LoginHandle TODO
func LoginHandle(w http.ResponseWriter, r *http.Request) {
	base.LogDebug("LoginHandle method:", r.Method)
	// if r.Method != "POST" {
	// 	return
	// }
	r.ParseForm()
	base.LogDebug("LoginHandle get request Form:", r.Form)

	base.LogDebug("LoginHandle get request body:", r.Body)

	uidStr := r.Form.Get("uid")
	passwd := r.Form.Get("passwd")
	if uidStr == "" || passwd == "" {
		base.LogError("uid or passwd is nil: ", uidStr, passwd)
		return
	}

	uid, err := strconv.ParseUint(uidStr, 10, 64)
	if err != nil {
		base.LogError("LoginHandle [uid] atoi error: ", err)
		return
	}

	retData := doLogin(uid, passwd)
	base.LogDebug("LoginHandle response info: ", retData)
	w.Write([]byte(retData))
}

func doLogin(uid uint64, passwd string) string {
	info, err := model.AccountGet(uid)
	if info == nil || err != nil {
		base.LogError("LoginHandle get account error: ", uid, passwd, err)
		return "1001|err1001"
	}

	if !info.CheckPasswordSuccess(passwd) {
		base.LogError("LoginHandle passwd error: ", uid, passwd)
		return "1002|err1002"
	}

	token := model.TokenMake(uid, passwd)

	// 设置登录令牌. 无论是否在线, 直接刷新登录状态, 原登录连接验证令牌时自动失效.(游戏服务器在新连接请求验证时踢出旧连接.)
	err = model.UIDTokenSet(uid, token)
	if err != nil {
		base.LogError("LoginHandle UIDTokenSet error: ", err, uid, token)
		return "1003|err1003"
	}

	return fmt.Sprintf("0|%s", token)
}
