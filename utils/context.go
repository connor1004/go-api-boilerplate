package utils

import (
	"encoding/json"
	"net/http"
)

// Context a struct for writing responses
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Params         []string
}

// Respond : send a cors json response
func (ctx *Context) Respond(code int, data map[string]interface{}) {
	ctx.ResponseWriter.Header().Add("Content-Type", "application/json")
	ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	ctx.ResponseWriter.WriteHeader(code)

	json.NewEncoder(ctx.ResponseWriter).Encode(data)
}
