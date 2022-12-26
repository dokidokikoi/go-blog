package meta

import "github.com/go-playground/validator/v10"

type ListOption struct {
	GetOption
	PageSize int `validate:"gte=1|eq=-1"`
	Page     int `validate:"gte=1|eq=-1"`
	Order    string
	Group    string
}

func (l *ListOption) Validate() []error {
	var errs = []error{}
	var validate = validator.New()
	err := validate.Struct(l)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			panic("传入值错误")
		}

		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, err)
		}
		return errs

	}

	return errs

}

const (
	defaultPageSize = 10
	defaultPage     = 1
	defaultOrder    = ""
)

type optionFunc func(option *ListOption)

func WithPageSize(size int) optionFunc {
	if size > 0 {
		return optionFunc(func(option *ListOption) {
			option.PageSize = size
		})
	}
	return optionFunc(func(option *ListOption) {
		option.PageSize = defaultPageSize
	})
}

func WithPage(page int) optionFunc {
	if page > 0 {
		return optionFunc(func(option *ListOption) {
			option.Page = page
		})
	}
	return optionFunc(func(option *ListOption) {
		option.Page = defaultPage
	})
}
func WithOrderBy(order string) optionFunc {
	if order != "" {
		return optionFunc(func(option *ListOption) {
			option.Order = order
		})
	}
	return optionFunc(func(option *ListOption) {
		option.Order = defaultOrder
	})
}

func NewListOption(include []string, options ...optionFunc) *ListOption {
	listOption := &ListOption{
		GetOption: GetOption{Include: include},
		Page:      defaultPage,
		PageSize:  defaultPageSize,
		Order:     defaultOrder,
	}
	for _, f := range options {
		f(listOption)
	}
	return listOption
}
