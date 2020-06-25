package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"uit_payment/entities/payment/delivery/response"

	"github.com/go-playground/form"
	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type HandlerInterface interface {
	SetContext(w http.ResponseWriter, req *http.Request)
	Handle()
}

type Handler struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

func New(handler HandlerInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		handler.SetContext(w, req)
		handler.Handle()
	}
}

func (handler *Handler) SetContext(w http.ResponseWriter, req *http.Request) {
	handler.Writer = w
	handler.Request = req
}

func (handler *Handler) RenderSuccess(data interface{}) {
	handler.Writer.Header().Set("Content-Type", "application/json")
	handler.Writer.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(data)
	fmt.Printf("Response:\n%s\n", string(resp))
	handler.Writer.Write(resp)
}

func (handler *Handler) RenderError(message string) {
	handler.Writer.Header().Set("Content-Type", "application/json")
	handler.Writer.WriteHeader(http.StatusBadRequest)
	data := response.Error{Message: message}
	resp, _ := json.Marshal(data)
	fmt.Printf("Response:\n%s\n", string(resp))
	handler.Writer.Write(resp)
}

func (handler *Handler) RenderNoContent() {
	handler.Writer.WriteHeader(http.StatusNoContent)
}

func (handler *Handler) RenderErrorWithJSON(data interface{}) {
	handler.Writer.Header().Set("Content-Type", "application/json")
	handler.Writer.WriteHeader(http.StatusBadRequest)
	resp, _ := json.Marshal(data)
	fmt.Printf("Response:\n%s\n", string(resp))
	handler.Writer.Write(resp)
}

func (handler *Handler) ParseParam(param interface{}) error {
	contentType := handler.Request.Header.Get("Content-Type")
	switch contentType {
	case "application/json", "application/json; charset=utf-8":
		decoder := json.NewDecoder(handler.Request.Body)
		err := decoder.Decode(&param)
		if err != nil {
			log.Printf("Decode error: #%v", err)
			return err
		}
	default:
		err := handler.Request.ParseForm()
		if err != nil {
			log.Printf("Parse params: #%v", err)
			return err
		}
		decoder := form.NewDecoder()
		err = decoder.Decode(param, handler.Request.Form)
		if err != nil {
			log.Printf("Decode: #%v", err)
			return err
		}
	}

	req, err := json.Marshal(param)
	if err != nil {
		log.Printf("Hanlder.MarshalParamError: #%v", req)
	}
	fmt.Printf("Request params:\n%s\n", string(req))

	return nil
}

func (handler *Handler) Vars() map[string]string {
	return mux.Vars(handler.Request)
}
