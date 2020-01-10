package model

/*
 * module model.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

// Redis Keys
const (
	// STRINGversion   string = "version"   // version
	// STRINGupdateurl string = "updateURL" // updateURL

	// remote global
	INCRseeduid    string = "incr_login_uid_seed" // global login uid seed
	HASHuidaccount string = "hash_uid_account"    // uid => AccountInfo
	HASHuidtoken   string = "hash_uid_token"      // uid => TokenInfo

	// local
	HASHsidserver string = "hash_sid_server" // sid => ServerInfo
)

/************************************************************************/
// export functions.
/************************************************************************/

/************************************************************************/
// moudule functions.
/************************************************************************/

// InitFromRedis TODO
func InitFromRedis() {
	SeedUIDInit()
}

/************************************************************************/
// unit tests.
/************************************************************************/
