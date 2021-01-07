package models

import (
	"github.com/bradfitz/gomemcache/memcache"
)
var mc *memcache.Client

func init() {

	conn := memcache.New("127.0.0.1:11211")
	mc = conn


}

func GetMC() *memcache.Client {
	return mc
}