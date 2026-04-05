package internal

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Helper interface {
	UnmarshalAndValidate(data []byte, v interface{}) error
	ValidateStruct(s interface{}) error
}

type helper struct {
}

func NewHelper() Helper {
	return &helper{}
}

func (h *helper) ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func (h *helper) UnmarshalAndValidate(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return h.ValidateStruct(v)
}
