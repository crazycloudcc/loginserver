package model

/*
 * account data
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

import (
	"assist"
	"cherry/base"
	"cherry/dbproxy"
	"crypto/sha1"
	"fmt"
	"io"
	"time"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

const (
	// VALUESTRING TODO
	VALUESTRING = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// VALUEMAX TODO
	VALUEMAX = 0x3FFFFFFF

	// VALUELEN TODO
	VALUELEN = 0x0000003D
)

// AccountInfo TODO
type AccountInfo struct {
	UID        uint64 // user id.
	Passwd     string // user password. (unused)
	RegTime    int64  // frist login time.
	PlatformID int32  // platform exam: iOS/Android
	ChannelID  int32  // channel exam: weibo/facebook
	Email      string // thrid account email.
	Phone      string // thrid account phone.
	IsDeleted  bool   // delete flag.
}

/************************************************************************/
// export functions.
/************************************************************************/

// AccountAddGuest TODO
func AccountAddGuest(platformID, channelID int32) (*AccountInfo, error) {
	uid, err := SeedUIDGet()
	if uid == 0 || err != nil {
		return nil, err
	}

	regTime := time.Now().Unix()
	passwd := passwdMake(uid, platformID, channelID, regTime)

	a := &AccountInfo{uid, passwd, regTime, platformID, channelID, "", "", false}
	err = AccountAdd(a)
	return a, err
}

// AccountAdd TODO
func AccountAdd(info *AccountInfo) error {
	status, err := assist.RedisHEXISTS(dbproxy.InsRedisRemote, HASHuidaccount, info.UID)
	if status == 0 {
		err = AccountSet(info)
	}
	return err
}

// AccountDel TODO
func AccountDel(uid uint64) (*AccountInfo, error) {
	info, err := AccountGet(uid)
	if info == nil {
		return nil, err
	}
	info.IsDeleted = true
	return info, nil
}

// AccountSet TODO
func AccountSet(info *AccountInfo) error {
	err := assist.RedisHSET(dbproxy.InsRedisRemote, HASHuidaccount, info.UID, info)
	return err
}

// AccountGet TODO
func AccountGet(uid uint64) (*AccountInfo, error) {
	reply, err := assist.RedisHGET(dbproxy.InsRedisRemote, HASHuidaccount, uid)
	if reply == nil || err != nil {
		return nil, err
	}
	a := new(AccountInfo)
	dbproxy.RedisUnmarshal(reply.([]byte), a)
	return a, nil
}

// CheckPasswordSuccess TODO
func (owner *AccountInfo) CheckPasswordSuccess(passwd string) bool {
	return (owner.Passwd == passwd)
}

/************************************************************************/
// moudule functions.
/************************************************************************/

func passwdMake(uid uint64, platformID, channelID int32, regTime int64) string {
	passwd := sha1.New()
	baseStr := fmt.Sprintf("%d,%d,%d,%d", uid, platformID, channelID, regTime)
	io.WriteString(passwd, baseStr)
	ret := fmt.Sprintf("%x", passwd.Sum(nil))[:8]
	base.LogDebug("passwd make: ", uid, platformID, channelID, regTime, ret)
	return ret
}

/************************************************************************/
// unit tests.
/************************************************************************/
