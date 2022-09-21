package util

import (
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	Validate *validator.Validate
	Trans    ut.Translator
)

// InitTrans 初始化翻译器
func InitTrans() (err error) {
	en := zh.New() // 中文翻译器
	uni = ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	Trans, _ = uni.GetTranslator("zh")

	Validate = validator.New()
	zh_translations.RegisterDefaultTranslations(Validate, Trans)
	// zhT := zh.New() // 中文翻译器

	// // 第一个参数是备用（fallback）的语言环境
	// // 后面的参数是应该支持的语言环境（支持多个）
	// // uni := ut.New(zhT, zhT) 也是可以的
	// uni := ut.New(zhT, zhT)

	// // locale 通常取决于 http 请求头的 'Accept-Language'
	// var ok bool
	// // 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
	// Trans, ok = uni.GetTranslator("zh")
	// if !ok {
	// 	return fmt.Errorf("uni.GetTranslator failed")
	// }
	// validate := validator.New()
	// zhTranslations.RegisterDefaultTranslations(validate, Trans)
	//validator.Validate
	// 注册翻译器
	// switch locale {
	// case "en":
	// 	err = enTranslations.RegisterDefaultTranslations(validator.Validate, trans)
	// case "zh":
	// 	err = zhTranslations.RegisterDefaultTranslations(v, trans)
	// default:
	// 	err = enTranslations.RegisterDefaultTranslations(v, trans)
	// }
	return
}

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
