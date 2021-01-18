package models

import (
	"github.com/bradfitz/gomemcache/memcache"
	gsm "github.com/bradleypeabody/gorilla-sessions-memcache"
)

var mc *memcache.Client

func init() {

	conn := memcache.New("127.0.0.1:11211")
	mc = conn

}

func GetMC() *memcache.Client {
	return mc
}

func GetCookiesStore() *gsm.MemcacheStore {
	memClient := gsm.NewGoMemcacher(GetMC())
	///gsm.NewMemcacherStoreWithValueStorer(memClient, &gsm.HeaderStorer{HeaderFieldName:"X-CUSTOM-HEADER"}, "session_prefix_", []byte("secret-key-goes-here"))
	return gsm.NewMemcacherStore(memClient, "session_prefix_", []byte("secret-key-goes-here"))
}
