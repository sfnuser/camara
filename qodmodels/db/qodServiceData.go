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

import (
	"github.com/sfnuser/camara/qodmodels/api"
)

type FlowInfo struct {
	// Indicates the IP flow.
	FlowId uint32 `json:"flowId" bson:"flowId"`
	// Indicates the packet filters of the IP flow. Refer to subclause 5.3.8 of 3GPP TS 29.214 for encoding. It shall contain UL and/or DL IP flow description.
	FlowDescriptions *[]string `json:"flowDescriptions,omitempty" bson:"flowDescriptions,omitempty"`
}

type ServiceQoDUeSession struct {
	UeIpv4Addr              string            `json:"ueIpv4Addr" bson:"ueIpv4Addr"`
	ScsAsId                 string            `json:"scsAsId" bson:"scsAsId"`
	SessionId               string            `json:"sessionId" bson:"sessionId"`
	NefSubscriptionId       string            `json:"nefSubscriptionId" bson:"nefSubscriptionId"`
	NefSubscriptionResource string            `json:"nefSubscriptionResource" bson:"nefSubscriptionResource"`
	QosReference            string            `json:"qosReference" bson:"qosReference"`
	FlowInfo                FlowInfo          `json:"flowInfo" bson:"flowInfo"`
	SessionReq              api.CreateSession `json:"sessionReq" bson:"sessionReq"`
	SessionInfo             api.SessionInfo   `json:"sessionInfo" bson:"sessionInfo"`
}

type ServiceQoDUeFlow struct {
	UeIpv4Addr  string `json:"ueIpv4Addr" bson:"ueIpv4Addr"`
	ScsAsId     string `json:"scsAsId" bson:"scsAsId"`
	FlowCounter uint32 `json:"flowCounter" bson:"flowCounter"` // Running counter of FlowId
}
