package controller

import (
	"fmt"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"

	"blog-backend/config"
	"blog-backend/model"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/gin-contrib/sessions"

	"github.com/afocus/captcha"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gin-gonic/gin"
)

func GetLatestArticles(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func GetPersonalSummary(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func GetMessageBoard(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

var cap *captcha.Captcha

func GetCaptchaURL(c *gin.Context) {
	uuid := c.Param("uuid")
	cap = captcha.New()
	if err := cap.SetFont("./config/comic.ttf"); err != nil { //路径是看main.go文件，而不是调用函数位置
		panic(err.Error())
	}
	cap.SetSize(100, 38)
	cap.SetDisturbance(captcha.NORMAL)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{124, 58, 237, 255})

	img, str := cap.Create(4, captcha.NUM)

	//存储图片str到redis数据库
	err := model.SaveImgCode(str, uuid)
	if err != nil {
		return
	}

	png.Encode(c.Writer, img)
}

var phone string
var Code string

func GetPhoneCode(c *gin.Context) {
	uuid := c.Param("uuid")
	phone = c.Param("phone")
	captcha := c.Param("captcha")

	data := make(map[string]string)

	boolS := model.CheckImg(uuid, captcha)
	if !boolS {
		data["message"] = "captchaCode false; 图片验证码错误,请重新输入"
		data["ok"] = "0"
	} else {
		data["message"] = "captchaCode success; 获取成功，请注意短信 "
		data["ok"] = "1"
		_main(tea.StringSlice(os.Args[1:]))
	}

	c.JSON(200, data)

}

func CreateClient() (_result *openapi.Client, _err error) {
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考。
	// 建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html。
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")),
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &openapi.Client{}
	_result, _err = openapi.NewClient(config)
	return _result, _err
}

func CreateApiInfo() (_result *openapi.Params) {
	params := &openapi.Params{
		// 接口名称
		Action: tea.String("SendSms"),
		// 接口版本
		Version: tea.String("2017-05-25"),
		// 接口协议
		Protocol: tea.String("HTTPS"),
		// 接口 HTTP 方法
		Method:   tea.String("POST"),
		AuthType: tea.String("AK"),
		Style:    tea.String("RPC"),
		// 接口 PATH
		Pathname: tea.String("/"),
		// 接口请求体内容格式
		ReqBodyType: tea.String("json"),
		// 接口响应体内容格式
		BodyType: tea.String("json"),
	}
	_result = params
	return _result
}

func _main(args []*string) (_err error) {

	// 设置随机数种子，确保每次运行结果不同
	rand.Seed(time.Now().UnixNano())
	// 生成 0 到 9999 之间的随机数
	randomNumber := rand.Intn(10000)
	// 将随机数格式化为四位数，不足四位数前面补零，并赋值给 code 变量
	Code = fmt.Sprintf("%04d", randomNumber)
	err := model.SavePhoneCode(Code, phone)
	if err != nil {
		fmt.Println("redis save false", err)
	}

	client, _err := CreateClient()
	if _err != nil {
		return _err
	}

	params := CreateApiInfo()
	// query params
	queries := map[string]interface{}{}
	queries["PhoneNumbers"] = tea.String(phone)
	queries["SignName"] = tea.String("zlican博客")
	queries["TemplateCode"] = tea.String("SMS_474916099")
	queries["TemplateParam"] = tea.String("{\"code\": " + Code + "}")
	queries["SmsUpExtendCode"] = nil
	queries["OutId"] = nil
	// runtime options
	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{
		Query: openapiutil.Query(queries),
	}
	// 复制代码运行请自行打印 API 的返回值
	// 返回值实际为 Map 类型，可从 Map 中获得三类数据：响应体 body、响应头 headers、HTTP 返回的状态码 statusCode。
	_, _err = client.CallApi(params, request, runtime)
	if _err != nil {
		return _err
	}
	return _err
}

type RegisterData struct {
	Phone     string `json:"phone"`
	PhoneCode string `json:"phoneCode"`
	Password  string `json:"password"`
}

func PostRegister(ctx *gin.Context) {
	var registerData RegisterData
	if err := ctx.ShouldBindJSON(&registerData); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if model.CheckCode(registerData.Phone, registerData.PhoneCode) {
		// 保存用户信息
		if err := model.SaveUserInfo(registerData.Phone, registerData.Password); err != nil {
			ctx.JSON(500, gin.H{
				"message": "注册失败",
				"error":   err.Error(),
			})
			return
		}

		// 设置session
		s := sessions.Default(ctx)
		s.Set("phoneNumber", phone)
		s.Save()

		ctx.JSON(200, gin.H{
			"message": "注册成功",
		})
	} else {
		ctx.JSON(400, gin.H{
			"message": "验证码错误",
		})
	}
}

type LoginData struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func PostLogin(ctx *gin.Context) {
	var logindata LoginData
	var userinfo model.UserInfo
	ctx.ShouldBindJSON(&logindata)

	config.DB.Where("phone_number= ? ", logindata.Phone).Find(&userinfo)
	if logindata.Password == userinfo.PassWord {
		s := sessions.Default(ctx)
		s.Set("phoneNumber", logindata.Phone)
		s.Save()
		ctx.JSON(200, gin.H{
			"message": "登入成功",
		})
	} else {
		ctx.JSON(400, gin.H{
			"message": "密码错误",
		})
	}
}

func CheckLogin(ctx *gin.Context) {
	s := sessions.Default(ctx)
	result := s.Get("phoneNumber")
	if result == "" {
		ctx.JSON(200, gin.H{
			"message": "未登入",
		})
		return
	}
	var userinfo model.UserInfo

	config.DB.Where("phone_number = ?", result).Find(&userinfo)

	ctx.JSON(200, gin.H{
		"username": userinfo.UserName,
	})
}
