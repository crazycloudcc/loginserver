package model

/*
 * seed uid model.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

import (
	"assist"
	"cherry/base"
	"cherry/dbproxy"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

/************************************************************************/
// export functions.
/************************************************************************/

const (
	// seedStart = uint64(1<<32) + 1
	seedStart uint64 = 1 // DEBUG
)

// SeedUIDInit TODO
func SeedUIDInit() {
	status, err := assist.RedisExists(dbproxy.InsRedisRemote, INCRseeduid)
	if status == 0 {
		if err := assist.RedisSET(dbproxy.InsRedisRemote, INCRseeduid, seedStart); err != nil {
			base.LogError("create uid seed error: ", err)
		}
	} else if status == -1 {
		base.LogError("SeedUIDInit error: ", err)
	}
}

// SeedUIDGet TODO
func SeedUIDGet() (uint64, error) {
	reply, err := assist.RedisINCR(dbproxy.InsRedisRemote, INCRseeduid)
	if reply == 0 || err != nil {
		base.LogError("SeedUIDGet error: ", err)
		return 0, err
	}
	return reply, nil
}

/************************************************************************/
// moudule functions.
/************************************************************************/

/************************************************************************/
// unit tests.
/************************************************************************/
