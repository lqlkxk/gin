package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lqlkxk/gin/dao"
	"github.com/lqlkxk/gin/model"
	"github.com/lqlkxk/gin/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

// 用户处理

type UserHandler struct{}

// 根据名字查询
func (u UserHandler) GetByName(ctx *gin.Context) { // 根据 gin上下文获取参数，设置响应参数
	// 获取参数
	username := ctx.Param("username")
	user := dao.UserDao{}.FindByname(username)
	ret := utils.SuccessWithData(user)
	// 返回信息 json
	ctx.JSON(http.StatusOK, ret)
}

func (u UserHandler) FindById(ctx *gin.Context) {
	// 获取参数
	id := ctx.Query("id")
	log.Println(id)
	// 返回信息
	ctx.String(http.StatusOK, "10086")
}

func (u UserHandler) Save(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	log.Println("[username]：" + username)
	log.Println("[password]：" + password)
	// 返回信息
	ctx.String(http.StatusOK, "success")
}
func (u UserHandler) Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	user := dao.UserDao{}.FindByname(username)
	if user.IsEmpty() {
		// 返回信息
		ctx.JSON(http.StatusOK, utils.ResponseMsg{Code: 400, Msg: "账号密码错误", Data: ""})
	}
	if user.Password != password {
		ctx.JSON(http.StatusOK, utils.ResponseMsg{Code: 400, Msg: "账号密码错误", Data: ""})
	}
	// 登录成功 设置cookie name, value string, maxAge int, path, domain string, secure, httpOnly bool
	//ctx.SetCookie("user_cookie", string(user.Id), 1000, "/", "127.0.0.1", false, true)
	// 返回信息
	oneDayOfHours, secret, _ := utils.GetJWT()
	dayOfHours, _ := strconv.Atoi(oneDayOfHours)
	expiresTime := time.Now().Unix() + int64(dayOfHours)
	claims := jwt.StandardClaims{
		Audience:  user.UserName,     // 受众
		ExpiresAt: expiresTime,       // 失效时间
		Id:        string(user.Id),   // 编号
		IssuedAt:  time.Now().Unix(), // 签发时间
		Issuer:    "gin hello",       // 签发人
		NotBefore: time.Now().Unix(), // 生效时间
		Subject:   "login",           // 主题
	}
	var jwtSecret = []byte(secret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrWithMsg("登录失败"))
	}
	token = "Bearer " + token
	ctx.JSON(http.StatusOK, utils.SuccessWithData(token))
}

func (u UserHandler) SaveJson(ctx *gin.Context) {
	var user model.User
	ctx.BindJSON(&user)
	log.Printf("[user]：%v", user)
	//data, err := ioutil.ReadAll(ctx.Request.Body)
	//if err != nil {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"error": err,
	//	})
	//} else {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"json": string(data),
	//	})
	//}
	// 返回信息
	ctx.String(http.StatusOK, "success")
}
