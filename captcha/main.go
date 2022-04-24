package main

import (
	"bytes"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./*.html")
	router.Use(Session("golang-tech-stack"))
	router.GET("/captcha", func(c *gin.Context) {
		Captcha(c, 4)
	})
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/captcha/verify/", func(ctx *gin.Context) {
		value := ctx.PostForm("code")
		if CaptchaVerify(ctx, value) {
			ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": "success"})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"status": 1, "msg": "failed"})
		}
	})
	router.Run(":8099")

}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)

	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

//中间件，处理session
func Session(keyPairs string) gin.HandlerFunc {
	store := SessionConfig()
	return sessions.Sessions(keyPairs, store)
}

//配置session
func SessionConfig() sessions.Store {
	sessionMaxAge := 3600
	sessionSecret := "golang-tech-stack"
	store := cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: sessionMaxAge,
	})
	return store
}

//验证
func CaptchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	if captchaId := session.Get("captcha"); captchaId != nil {
		session.Delete("captcha")
		_ = session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

//生成验证码
func Captcha(c *gin.Context, length ...int) {
	l := captcha.DefaultLen
	w, h := 107, 36
	if len(length) == 1 {
		l = length[0]
	}
	if len(length) == 2 {
		w = length[1]
	}
	if len(length) == 3 {
		h = length[2]
	}
	captchaId := captcha.NewLen(l)
	session := sessions.Default(c)

	session.Set("captcha", captchaId)
	_ = session.Save()
	_ = Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}
