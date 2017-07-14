// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package bar

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/thriftrw/ptr"
	"go.uber.org/zap"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
)

// ArgWithNestedQueryParamsHandler is the handler for "/bar/argWithNestedQueryParams"
type ArgWithNestedQueryParamsHandler struct {
	Clients *module.ClientDependencies
}

// NewArgWithNestedQueryParamsEndpoint creates a handler
func NewArgWithNestedQueryParamsEndpoint(
	gateway *zanzibar.Gateway,
	deps *module.Dependencies,
) *ArgWithNestedQueryParamsHandler {
	return &ArgWithNestedQueryParamsHandler{
		Clients: deps.Client,
	}
}

// Register adds the http handler to the gateway's http router
func (handler *ArgWithNestedQueryParamsHandler) Register(g *zanzibar.Gateway) error {
	return g.HTTPRouter.Register(
		"GET", "/bar/argWithNestedQueryParams",
		zanzibar.NewRouterEndpoint(
			g,
			"bar",
			"argWithNestedQueryParams",
			handler.HandleRequest,
		),
	)
}

// HandleRequest handles "/bar/argWithNestedQueryParams".
func (handler *ArgWithNestedQueryParamsHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	var requestBody endpointsBarBar.Bar_ArgWithNestedQueryParams_Args

	if requestBody.Request == nil {
		requestBody.Request = &endpointsBarBar.QueryParamsStruct{}
	}
	xUUIDValue, _ := req.Header.Get("x-uuid")
	requestBody.Request.AuthUUID = ptr.String(xUUIDValue)
	xUUID2Value, _ := req.Header.Get("x-uuid2")
	requestBody.Request.AuthUUID2 = xUUID2Value

	if requestBody.Request == nil {
		requestBody.Request = &endpointsBarBar.QueryParamsStruct{}
	}
	requestNameOk := req.CheckQueryValue("request.name")
	if !requestNameOk {
		return
	}
	requestNameQuery, ok := req.GetQueryValue("request.name")
	if !ok {
		return
	}
	requestBody.Request.Name = requestNameQuery

	requestUserUUIDOk := req.HasQueryValue("request.userUUID")
	if requestUserUUIDOk {
		requestUserUUIDQuery, ok := req.GetQueryValue("request.userUUID")
		if !ok {
			return
		}
		requestBody.Request.UserUUID = ptr.String(requestUserUUIDQuery)
	}

	workflow := ArgWithNestedQueryParamsEndpoint{
		Clients: handler.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, cliRespHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		switch errValue := err.(type) {

		default:
			req.Logger.Warn("Workflow for endpoint returned error",
				zap.String("error", errValue.Error()),
			)
			res.SendErrorString(500, "Unexpected server error")
			return
		}
	}
	// TODO(jakev): implement writing fields into response headers

	res.WriteJSON(200, cliRespHeaders, response)
}

// ArgWithNestedQueryParamsEndpoint calls thrift client Bar.ArgWithNestedQueryParams
type ArgWithNestedQueryParamsEndpoint struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w ArgWithNestedQueryParamsEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBarBar.Bar_ArgWithNestedQueryParams_Args,
) (*endpointsBarBar.BarResponse, zanzibar.Header, error) {
	clientRequest := convertToArgWithNestedQueryParamsClientRequest(r)

	clientHeaders := map[string]string{}

	clientRespBody, _, err := w.Clients.Bar.ArgWithNestedQueryParams(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		default:
			w.Logger.Warn("Could not make client request",
				zap.String("error", errValue.Error()),
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertArgWithNestedQueryParamsClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToArgWithNestedQueryParamsClientRequest(in *endpointsBarBar.Bar_ArgWithNestedQueryParams_Args) *clientsBarBar.Bar_ArgWithNestedQueryParams_Args {
	out := &clientsBarBar.Bar_ArgWithNestedQueryParams_Args{}

	if in.Request != nil {
		out.Request = &clientsBarBar.QueryParamsStruct{}
		out.Request.Name = string(in.Request.Name)
		out.Request.UserUUID = (*string)(in.Request.UserUUID)
		out.Request.AuthUUID = (*string)(in.Request.AuthUUID)
		out.Request.AuthUUID2 = string(in.Request.AuthUUID2)
	} else {
		out.Request = nil
	}

	return out
}

func convertArgWithNestedQueryParamsClientResponse(in *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	out := &endpointsBarBar.BarResponse{}

	out.StringField = string(in.StringField)
	out.IntWithRange = int32(in.IntWithRange)
	out.IntWithoutRange = int32(in.IntWithoutRange)
	out.MapIntWithRange = make(map[string]int32, len(in.MapIntWithRange))
	for key, value := range in.MapIntWithRange {
		out.MapIntWithRange[key] = int32(value)
	}
	out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
	for key, value := range in.MapIntWithoutRange {
		out.MapIntWithoutRange[key] = int32(value)
	}

	return out
}
