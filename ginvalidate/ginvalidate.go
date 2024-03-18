package ginvalidate

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"sync"
)

var (
	uni        *ut.UniversalTranslator
	validate   *validator.Validate
	utLangList []locales.Translator
	_once      sync.Once
)

const (
	english = "en"
	chinese = "zh"
)

var languagesList = [...]string{english, chinese}

func init() {
	lazyInit()
}

// 注册语言包
func lazyInit() {
	_once.Do(func() {
		utLangList = []locales.Translator{
			zh2.New(),
			en2.New(),
		}
		uni = ut.New(utLangList[0], utLangList...)
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			validate = v
		} else {
			validate = validator.New()
		}
		registerLang := func() error {
			var translations func(*validator.Validate, ut.Translator) (err error)
			for _, val := range utLangList {
				tans, ok := uni.GetTranslator(val.Locale())
				if !ok {
					return errors.New("获取语言包实例错误")
				}
				switch val.Locale() {
				case english:
					translations = enTranslations.RegisterDefaultTranslations
				case chinese:
					translations = zhTranslations.RegisterDefaultTranslations
				default:
					translations = enTranslations.RegisterDefaultTranslations
				}
				err := translations(validate, tans)
				if err != nil {
					return err
				}
			}
			return nil
		}
		err := registerLang()
		if err != nil {
			panic(err)
		}
	})
}

func inLanguages(language string) bool {
	for _, val := range languagesList {
		if val == language {
			return true
		}
	}
	return false
}

func SimpleValidate(ctx *gin.Context, param any, localParam ...string) (bool, []error) {
	local := ""
	if len(localParam) == 0 {
		local = ctx.DefaultQuery("local", "zh")
	} else {
		local = localParam[0]
	}
	if !inLanguages(local) {
		return false, []error{errors.New("目前只支持中文和英文，无法识别：" + local)}
	}
	if err := ctx.ShouldBind(param); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok || len(errs) == 0 {
			panic(err)
		}
		tans, ok := uni.GetTranslator(local) // 获取转换的实例
		if !ok {
			return false, []error{errors.New("验证器获取语言包实例失败，请检查是否配置正确")}
		}
		// errMsg := errs.Translate(tans) // 全部错误(map)
		// errMsg := errs[0].Translate(tans) // 第一个错误
		errList := make([]error, 0, len(errs)) // 全部错误
		for _, val := range errs {
			errList = append(errList, errors.New(val.Translate(tans)))
		}
		return false, errList
	}
	return true, nil
}
