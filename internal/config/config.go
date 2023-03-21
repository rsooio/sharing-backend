package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DSN  string
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
