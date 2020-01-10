package assist

/*
 * [file desc]
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

import "sync"

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

// IDMngr TODO
type IDMngr struct {
	sync.Mutex
	sid  int32 // always 1.
	seed uint64
}

/************************************************************************/
// export functions.
/************************************************************************/

// GetSeed TODO
func (owner *IDMngr) GetSeed() uint64 {
	defer owner.seedTick()
	return owner.seed
}

/************************************************************************/
// moudule functions.
/************************************************************************/

func (owner *IDMngr) seedTick() {
	owner.seed++
}

func (owner *IDMngr) init() {
}

func (owner *IDMngr) start() {
}

/************************************************************************/
// unit tests.
/************************************************************************/
