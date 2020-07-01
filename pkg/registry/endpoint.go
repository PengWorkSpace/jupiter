// Copyright 2020 Douyu
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


package registry

import (
	"encoding/json"

	"github.com/douyu/jupiter/pkg/server"
)

// Endpoints ...
type Endpoints struct {
	// 服务节点列表
	Nodes map[string]server.ServiceInfo

	// 路由配置
	RouteConfigs map[string]RouteConfig
}

// RouteConfig ...
type RouteConfig struct {
	ID     string `json:"id" toml:"id"`
	Scheme string `json:"scheme" toml:"scheme"`
	Host   string `json:"host" toml:"host"`

	Deployment string   `json:"deployment"`
	URI        string   `json:"uri"`
	Upstream   Upstream `json:"upstream"`
}

// String ...
func (config RouteConfig) String() string {
	bs, _ := json.Marshal(config)
	return string(bs)
}

// Upstream ...
type Upstream struct {
	Nodes  map[string]int `json:"nodes"`
	Groups map[string]int `json:"groups"`
}
