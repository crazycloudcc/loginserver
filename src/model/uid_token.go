package model

/*
 * token data
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

import (
	"assist"
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
	tokenExpire int64 = 60 * 30
)

var (
	tokenSeed int32 = 1
)

// TokenInfo TODO
type TokenInfo struct {
	Token  string
	Expire int64
}

/************************************************************************/
// export functions.
/************************************************************************/

// TokenMake TODO
func TokenMake(uid uint64, passwd string) string {
	defer seedTick()
	baseStr := fmt.Sprintf("%d%s%d", uid, passwd, tokenSeed)
	token := sha1.New()
	// base.LogDebug("new token: ", token)
	io.WriteString(token, baseStr)
	// base.LogDebug("token info: ", baseStr, token, token.Sum(nil))
	return fmt.Sprintf("%x", token.Sum(nil))
}

// UIDTokenSet TODO
func UIDTokenSet(uid uint64, token string) error {
	err := assist.RedisHSET(dbproxy.InsRedisRemote, HASHuidtoken, uid, &TokenInfo{token, time.Now().Unix() + tokenExpire})
	return err
}

// UIDTokenGet TODO
func UIDTokenGet(uid uint64) (string, error) {
	reply, err := assist.RedisHGET(dbproxy.InsRedisRemote, HASHuidtoken, uid)
	if reply == nil || err != nil {
		return "", err
	}
	ti := new(TokenInfo)
	dbproxy.RedisUnmarshal(reply.([]byte), ti)
	if ti.Expire > time.Now().Unix() {
		return "", err
	}
	return ti.Token, nil
}

/************************************************************************/
// moudule functions.
/************************************************************************/

func seedTick() {
	tokenSeed++
}

/************************************************************************/
// unit tests.
/************************************************************************/
