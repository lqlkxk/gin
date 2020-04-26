package test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/lqlkxk/gin/initRouter"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// 单元测试 **_test  单元测试 以_test结尾, cmd go test
var router *gin.Engine

// 初始化路由
func init() {
	router = initRouter.SteupRouter()
}

func TestRouter(t *testing.T) {
	//构造响应体
	w := httptest.NewRecorder()
	// 构造请求体
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// 创建请求
	router.ServeHTTP(w, req)
	// 断言 响应码为 200
	assert.Equal(t, http.StatusOK, w.Code)
	// 断言 返回 welcome
	assert.Equal(t, "welcome", w.Body.String())
}

func TestFindUserByName(t *testing.T) {
	// 构造讲求
	req := httptest.NewRequest(http.MethodGet, "/user/findByName/test", nil)
	//构造响应体
	rsp := httptest.NewRecorder()
	// 创建请求
	router.ServeHTTP(rsp, req)
	assert.Equal(t, http.StatusOK, rsp.Code)
	assert.Equal(t, "test", rsp.Body.String())
}

func TestFindById(t *testing.T) {
	// 构造讲求
	req := httptest.NewRequest(http.MethodGet, "/user/findById?id=10086", nil)
	//构造响应体
	rsp := httptest.NewRecorder()
	// 创建请求
	router.ServeHTTP(rsp, req)
	assert.Equal(t, http.StatusOK, rsp.Code)
	assert.Equal(t, "10086", rsp.Body.String())

}

func TestSaveUser(t *testing.T) {
	// 构造json
	value := url.Values{}
	value.Add("username", "lqlkxk")
	value.Add("password", "123321")
	// 构造讲求
	req := httptest.NewRequest(http.MethodPost, "/user/save", bytes.NewBufferString(value.Encode()))
	//构造响应体
	rsp := httptest.NewRecorder()
	// 设置content-type 为form提交
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	// 创建请求
	router.ServeHTTP(rsp, req)
	assert.Equal(t, http.StatusOK, rsp.Code)
	assert.Equal(t, "success", rsp.Body.String())

}
