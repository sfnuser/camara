// Copyright 2023 Spry Fox Networks
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

package db

type ProvQoDAppServerData struct {
	AsIpv4Addr string            `json:"asIpv4Addr" bson:"asIpv4Addr"`
	ScsAsId    string            `json:"scsAsId" bson:"scsAsId"`
	QoSMap     map[string]string `json:"qosMap" bson:"qosMap"` // Mapping of qosProfile to qosReference map
}
