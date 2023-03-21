package svc

import (
	"backend/internal/config"
	"backend/internal/middleware"
	"backend/internal/utils/jwt"
	"backend/model"
	"time"

	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type OTPNode struct {
	OTP  string
	Time time.Time
}

type ServiceContext struct {
	Config    config.Config
	OTPList   map[string]*OTPNode
	DB        *gorm.DB
	JWTParser rest.Middleware
	Claims    **jwt.Claims
}

func NewServiceContext(c config.Config) *ServiceContext {
	claims := &jwt.Claims{}

	return &ServiceContext{
		Config:    c,
		DB:        model.InitDB(c.DSN),
		OTPList:   map[string]*OTPNode{},
		JWTParser: middleware.NewJWTParserMiddleware(c.Auth.AccessSecret, &claims).Handle,
		Claims:    &claims,
	}
}
