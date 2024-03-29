// Copyright (c) 2018 Uber Technologies, Inc.
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

package quuxhandler

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/quux/module"
	"github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/quux/workflow"
	thriftQuux "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/tchannel/quux/quux"
	"github.com/uber/zanzibar/examples/example-gateway/clients/quux"
	"github.com/uber/zanzibar/runtime"
)

type handler struct {
	quux quux.Client
}

// NewSimpleServiceEchoStringWorkflow ...
func NewSimpleServiceEchoStringWorkflow(
	deps *module.Dependencies,
) workflow.SimpleServiceEchoStringWorkflow {
	return &handler{
		quux: deps.Client.Quux,
	}
}

// Handle ...
func (h *handler) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *thriftQuux.SimpleService_EchoString_Args,
) (string, zanzibar.Header, error) {
	resp := h.quux.EchoString(r.Msg)
	return resp, nil, nil
}
