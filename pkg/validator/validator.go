package validator

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	internalError "go-ewallet/pkg/error"
	"log"
	"net/http"
	"strings"
	"sync"
)

var (
	internalValidator *validator.Validate
	onceValidator     sync.Once
	Trans             ut.Translator
)

func init() {
	onceValidator.Do(func() {
		enLocale := en.New()
		uni := ut.New(enLocale, enLocale)

		Trans, _ = uni.GetTranslator("en")

		internalValidator = validator.New()
		err := enTranslations.RegisterDefaultTranslations(internalValidator, Trans)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func Validate(payload any) error {
	if err := internalValidator.Struct(payload); err != nil {
		return internalError.NewErr(http.StatusBadRequest, getValidatorMessage(err))
	}
	return nil
}

func getValidatorMessage(err error) string {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return ""
	}

	errs := make([]string, 0)
	for _, err := range err.(validator.ValidationErrors) {
		tmp := []string{fmt.Sprintf("validation failed on field %s with precondition '%s'", err.Field(), err.ActualTag())}

		if err.Param() != "" {
			if err.ActualTag() == "oneof" {
				tmp = append(tmp, fmt.Sprintf("want %s", strings.Replace(err.Param(), `' '`, `' or '`, -1)))
			} else {
				tmp = append(tmp, fmt.Sprintf("want %s", err.Param()))
			}
		}

		if err.Value() != nil && err.Value() != "" {
			tmp = append(tmp, fmt.Sprintf("but got %v", err.Value()))
		}

		errs = append(errs, strings.Join(tmp, " "))
	}

	b, _ := json.Marshal(errs)
	return string(b)
}
