package internal

import (
	"bytes"
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()
var apiStates = []string{"S", "E", "C"}

type Helper interface {
	UnmarshalAndValidate(data []byte, v any) error
	DecodeAndValidate(data []byte, v any) error
	ValidateStruct(s any) error
	ToggleApiState(currentState string) string
}

type helper struct {
}

func NewHelper() Helper {
	return &helper{}
}

func (h *helper) ValidateStruct(s any) error {
	return validate.Struct(s)
}

func (h *helper) UnmarshalAndValidate(data []byte, v any) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return h.ValidateStruct(v)
}

func (h *helper) DecodeAndValidate(data []byte, v any) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()

	if err := dec.Decode(v); err != nil {
		return err
	}

	return h.ValidateStruct(v)
}

func (h *helper) ToggleApiState(currentState string) string {
	for i, state := range apiStates {
		if state == currentState {
			return apiStates[(i+1)%len(apiStates)]
		}
	}
	return ""
}
