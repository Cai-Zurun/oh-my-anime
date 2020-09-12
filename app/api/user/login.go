package user

import (
	"errors"
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"net/http"
	userMod "oh-my-anime_gf/app/model/user"
	userSer "oh-my-anime_gf/app/service/user"
	"oh-my-anime_gf/library/response"
	"time"
)

var (
	GfJWTMiddleware *jwt.GfJWTMiddleware
)

type SignInRequest struct {
	Passport string `v:"required#账号不能为空"`
	Password string `v:"required#密码不能为空"`
}

func init() {
	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour * 168,
		MaxRefresh:      time.Minute * 168,
		IdentityKey:     "id",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		LoginResponse:   LoginResponse,
		RefreshResponse: RefreshResponse,
		Unauthorized:    Unauthorized,
		IdentityHandler: IdentityHandler,
		PayloadFunc:     PayloadFunc,
	})
	if err != nil {
		glog.Fatal("JWT Error:" + err.Error())
	}
	GfJWTMiddleware = authMiddleware
}

func Authenticator(r *ghttp.Request) (interface{}, error) {
	var data *SignInRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	}
	res, err := userMod.FindOne("passport", data.Passport)
	if res == nil {
		return nil, errors.New("用户名或密码错误")
	}
	reqPwd, err := gmd5.Encrypt(data.Password + userSer.SALT)
	if err != nil {
		glog.Error("md5加密异常", err)
	}
	if reqPwd != res.Password {
		return nil, errors.New("用户名或密码错误")
	}
	return g.Map{
		"username" : res.Passport,
		"id" : res.Id,
	}, nil
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler sets the identity for JWT.
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	return claims["id"]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(r *ghttp.Request, code int, message string) {
	r.Response.WriteJson(g.Map{
		"code": http.StatusBadRequest,
		"msg":  message,
	})
	r.ExitAll()
}

// LoginResponse is used to define customized login-successful callback function.
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	r.Response.WriteJson(g.Map{
		"code":   http.StatusOK,
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
	r.ExitAll()
}

// RefreshResponse is used to get a new token no matter current token is expired or not.
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	r.Response.WriteJson(g.Map{
		"code":   http.StatusOK,
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
	r.ExitAll()
}

