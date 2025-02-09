//
// Copyright (C) 2019-2022 vdaas.org vald team <vald@vdaas.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Package metric provides metrics functions for grpc
package metric

import (
	"go.opencensus.io/plugin/ocgrpc"
)

// ServerHandler is a type alias of ocgrpc.ServerHandler to record OpenCensus stats and traces.
type ServerHandler = ocgrpc.ServerHandler

// NewServerHandler returns the server handler of OpenCensus stats and traces.
func NewServerHandler(opts ...ServerOption) *ServerHandler {
	handler := new(ServerHandler)

	for _, opt := range append(serverDefaultOpts, opts...) {
		opt(handler)
	}

	return handler
}
