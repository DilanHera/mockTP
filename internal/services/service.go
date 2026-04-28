package services

import (
	"encoding/json"
	"fmt"

	"github.com/DilanHera/mockTP/internal/store"
)

type Service interface {
	GetApiInfo(apiName string, v any) (store.ApiInfo, error)
	InitServiceStore(apiNames []string)
}

type service struct {
	apiInfoStore *store.ApiInfoStore
}

func NewService(store *store.ApiInfoStore) *service {
	return &service{
		apiInfoStore: store,
	}
}

func (s *service) InitServiceStore(apiNames []string) {
	for _, apiName := range apiNames {
		s.apiInfoStore.Create(apiName, "", "S", 200)
	}
}

func (s *service) GetApiInfo(apiName string, v any) (store.ApiInfo, error) {
	res, err := s.apiInfoStore.Get(apiName)
	if err != nil {
		return store.ApiInfo{}, err
	}
	if res == nil || res.Resp == "" {
		if res == nil {
			return store.ApiInfo{}, fmt.Errorf("custom response not set")
		}
		return *res, fmt.Errorf("custom response not set")
	}

	if res.Resp != "" {
		if err := CreateResponse([]byte(res.Resp), v); err != nil {
			return *res, err
		}
	}

	return *res, nil
}

func CreateResponse(resp []byte, v any) error {
	if err := json.Unmarshal(resp, v); err != nil {
		return fmt.Errorf("unmarshal failed: %w", err)
	}
	return nil
}
