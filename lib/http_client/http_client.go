package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"uit_payment/model"

	"github.com/sirupsen/logrus"
)

type HTTP struct{}

func HTTPInstance() HTTPCInterface {
	return &HTTP{}
}

func (h *HTTP) PostForm(endpoint string, mapParams interface{}) ([]byte, error) {
	typeOfParams := reflect.ValueOf(mapParams)

	params := make(url.Values)
	switch typeOfParams.Kind() {
	case reflect.Map:
		params = mapToParams(mapParams)
	default:
		params = parseParamToURLValue(mapParams)
	}
	res, err := http.PostForm(endpoint, params)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (h *HTTP) GetForm(endpoint string, params interface{}) ([]byte, error) {
	formParams := parseParamToURLValue(params)
	req, err := http.NewRequest("GET", endpoint, bytes.NewBufferString(formParams.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (h *HTTP) SendOrderCallback(payment *model.Payment) error {
	// req := request.CrmCallBackRequest{}
	// req = req.PopulateModel(payment)

	// // resp := map[string]string{}

	// url := payment.Partner.WebhookEndpoint
	// if url == "" {
	// 	return errors.New("Can not find url endpoint for this payment")
	// }

	// _, err := this.PostJSON(url, req)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (h *HTTP) DoRequest(method, endpoint string, params interface{}, header interface{}) ([]byte, error) {
	data := []byte{}
	if params != nil {
		data, _ = json.Marshal(params)
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewBufferString(string(data)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	addParamsToHeader(req, header)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	logrus.Info("Response\n", resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (h *HTTP) Post(endpoint string, contentType string, req interface{}, resp interface{}) error {
	byteParam, err := json.Marshal(req)
	if err != nil {
		return err
	}

	logrus.Info("Request", string(byteParam))
	httpResp, err := http.Post(endpoint, contentType, bytes.NewBuffer(byteParam))
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}
	logrus.Info("Response", string(body))

	if http.StatusBadRequest <= httpResp.StatusCode && httpResp.StatusCode <= http.StatusUnavailableForLegalReasons {
		return errors.New(string(body[:]))
	}

	if resp == nil {
		return nil
	}

	err = json.Unmarshal(body[:], resp)
	if err != nil {
		return err
	}
	return nil
}

func mapToParams(m interface{}) url.Values {
	params := make(url.Values)
	mapParams := m.(map[string]string)
	for key, value := range mapParams {
		params.Add(key, value)
	}
	return params
}

func parseParamToURLValue(req interface{}) url.Values {
	data := url.Values{}
	typeOfReq := reflect.ValueOf(req).Type()
	valueOfReq := reflect.ValueOf(req)
	for i := 0; i < typeOfReq.NumField(); i++ {
		varName := typeOfReq.Field(i).Tag.Get("form")
		varValue := valueOfReq.Field(i)

		switch varValue.Kind() {
		case reflect.String:
			data.Set(varName, varValue.String())
		case reflect.Int:
			data.Set(varName, strconv.Itoa(int(varValue.Int())))
		case reflect.Float32:
			data.Set(varName, fmt.Sprintf("%.0f", varValue.Float()))
		}
	}
	return data
}

func addParamsToHeader(req *http.Request, params interface{}) {
	typeOfParams := reflect.ValueOf(params).Type()
	valueOfReq := reflect.ValueOf(params)
	for i := 0; i < typeOfParams.NumField(); i++ {
		varName := typeOfParams.Field(i).Tag.Get("header")
		varValue := valueOfReq.Field(i)
		if varName != "" {
			switch varValue.Kind() {
			case reflect.String:
				req.Header.Add(varName, varValue.String())
			case reflect.Int:
				req.Header.Add(varName, strconv.Itoa(int(varValue.Int())))
			default:
			}
		}
	}
}
