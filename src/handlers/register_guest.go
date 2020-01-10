package handlers

import (
	"cherry/base"
	"fmt"
	"model"
	"net/http"
)

// RegisterGuestHandle TODO
func RegisterGuestHandle(w http.ResponseWriter, r *http.Request) {
	base.LogDebug("RegisterGuestHandle method:", r.Method)
	// if r.Method != "POST" {
	// 	return
	// }
	r.ParseForm()
	base.LogDebug("RegisterGuestHandle get request info:", r.Form)

	retData := doRegister()
	base.LogDebug("RegisterGuestHandle response info: ", retData)
	w.Write([]byte(retData))
}

func doRegister() string {
	info, err := model.AccountAddGuest(0, 0)
	if info == nil || err != nil {
		base.LogError("RegisterGuestHandle add guest account error:", info, err)
		return "1101|err1101"
	}
	return fmt.Sprintf("0|%d|%s", info.UID, info.Passwd)
}
