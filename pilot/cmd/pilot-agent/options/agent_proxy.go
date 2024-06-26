// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package options

import (
	istioagent "istio.io/istio/pkg/istio-agent"
)

// ProxyArgs provides all of the configuration parameters for the Pilot proxy.
type ProxyArgs struct {
	istioagent.Proxy

	StsPort            int
	TokenManagerPlugin string

	MeshConfigFile string

	// proxy config flags (named identically)
	ServiceCluster         string
	ProxyLogLevel          string
	ProxyComponentLogLevel string
	Concurrency            int
	TemplateFile           string
	OutlierLogPath         string

	PodName      string
	PodNamespace string

	// enableProfiling enables profiling via web interface host:port/debug/pprof/
	EnableProfiling bool
}

// NewProxyArgs constructs proxyArgs with default values.
func NewProxyArgs() ProxyArgs {
	p := ProxyArgs{}

	// Apply Default Values.
	p.applyDefaults()

	return p
}

// applyDefaults apply default value to ProxyArgs
func (node *ProxyArgs) applyDefaults() {
	node.PodName = PodNameVar.Get()
	node.PodNamespace = PodNamespaceVar.Get()
}
