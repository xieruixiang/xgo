package server

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
}

func NewContext(response http.ResponseWriter, request *http.Request) Context {
	return Context{
		request,
		response,
	}
}

func (receiver *Context) ReadJson(any interface{}) error {
	all, err := io.ReadAll(receiver.Request.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(all, any)
	if err != nil {
		return err
	}

	return nil
}

func (receiver *Context) WriteJson(code int, any interface{}) error {
	receiver.Response.WriteHeader(code)
	marshal, err := json.Marshal(any)
	if err != nil {
		return nil
	}
	_, err = receiver.Response.Write(marshal)
	if err != nil {
		return err
	}
	return nil
}

func (receiver *Context) Ok(any interface{}) error {
	return receiver.WriteJson(http.StatusOK, any)
}

func (receiver *Context) Error(any interface{}) error {
	return receiver.WriteJson(http.StatusInternalServerError, any)
}
