package interfaces

import (
	"bytes"
	"encoding/json"
	"go-course/demo/app/shared/log"
	"go-course/demo/app/shared/utils"
	"net/http"
)

// private repository: IRepository
type HttpApiRepository struct {
	client *http.Client
}

type Response struct {
	StatusCode int
	Body       interface{}
}

func NewHttpApiRepository(client *http.Client) *HttpApiRepository {
	if client != nil {
		return &HttpApiRepository{
			client: client,
		}
	}
	panic("You must create client")
}

func (h *HttpApiRepository) DoGet(
	url string,
	header map[string]string,
	query map[string]string) (*Response, error) {
	log.Info("transform is OK")
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Error("cannot create a new request", err)
		return nil, err
	}

	if request != nil {
		for key, value := range header {
			request.Header.Set(key, value)
		}

		queryParams := request.URL.Query()
		for key, value := range query {
			queryParams.Add(key, value)
		}
		request.URL.RawQuery = queryParams.Encode()
	}

	log.Info("building request")
	httpResponse, err := h.client.Do(request)
	if err != nil {
		log.Error("an error has occurred while trying to do get request", err)
		return nil, err
	}
	response, err := h.response(httpResponse)
	return response, nil
}

func (h *HttpApiRepository) DoPost(
	url string,
	body interface{},
	header map[string]string,
	query map[string]string) (*Response, error) {
	log.Info("transforming json request")
	requestBody, err := json.Marshal(body)
	if err != nil {
		log.Error("cannot transform request body", err)
		return nil, err
	}
	log.Info("transform is OK")
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Error("cannot create a new request", err)
		return nil, err
	}
	if request != nil {
		for key, value := range header {
			request.Header.Set(key, value)
		}
		queryParams := request.URL.Query()
		for key, value := range query {
			queryParams.Add(key, value)
		}
		request.URL.RawQuery = queryParams.Encode()
	}
	log.Info("building request")
	httpResponse, err := h.client.Do(request)
	if err != nil {
		log.Error("an error has occurred while trying to do post request", err)
		return nil, err
	}
	response, err := h.response(httpResponse)
	return response, nil
}

func (h *HttpApiRepository) DoPut(
	url string,
	body interface{},
	header map[string]string,
	query map[string]string) (*Response, error) {
	log.Info("transforming json request")
	requestBody, err := json.Marshal(body)
	if err != nil {
		log.Error("cannot transform request body", err)
		return nil, err
	}
	log.Info("transform is OK")
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Error("cannot update request", err)
		return nil, err
	}
	if request != nil {
		for key, value := range header {
			request.Header.Set(key, value)
		}
		queryParams := request.URL.Query()
		for key, value := range query {
			queryParams.Add(key, value)
		}
		request.URL.RawQuery = queryParams.Encode()
	}
	log.Info("building request")
	httpResponse, err := h.client.Do(request)
	if err != nil {
		log.Error("an error has occurred while trying to do put request", err)
		return nil, err
	}
	response, err := h.response(httpResponse)
	return response, nil
}

func (h *HttpApiRepository) response(httpResponse *http.Response) (*Response, error) {
	stringResponse := utils.ConvertHttpResponseBodyToString(httpResponse)
	var entity interface{}
	err := json.Unmarshal([]byte(stringResponse), &entity)
	response := &Response{
		StatusCode: httpResponse.StatusCode,
		Body:       entity,
	}
	if response.StatusCode != http.StatusCreated &&
		response.StatusCode != http.StatusOK &&
		response.StatusCode != http.StatusAccepted {
		return nil, err
	}
	log.Info("request is resolved OK")
	return response, nil
}
