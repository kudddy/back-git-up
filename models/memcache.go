package models

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	gsm "github.com/bradleypeabody/gorilla-sessions-memcache"
	"github.com/golang/glog"
	"os"
)

var mc *memcache.Client

func init() {

	host, check := os.LookupEnv("memcache_addres")
	if check {
		glog.Info(fmt.Sprintf("----INIT MEMCACHE on host %s ----", host))
	} else {
		glog.Error(fmt.Sprintf("----ERROR TO CONNECT MEMCACHED, ENV NOT FOUNT ----"))
		panic("----ERROR TO CONNECT MEMCACHED, ENV NOT FOUNT ----")
	}

	conn := memcache.New(host)
	mc = conn

}

func GetMC() *memcache.Client {
	return mc
}

func GetCookiesStore() *gsm.MemcacheStore {
	memClient := gsm.NewGoMemcacher(GetMC())
	///gsm.NewMemcacherStoreWithValueStorer(memClient, &gsm.HeaderStorer{HeaderFieldName:"X-CUSTOM-HEADER"}, "session_prefix_", []byte("secret-key-goes-here"))
	return gsm.NewMemcacherStore(memClient, "", []byte("secret-key-goes-here"))
}
