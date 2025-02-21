/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package httpsvr

import (
	"errors"

	"github.com/emicklei/go-restful"
)

func (h *HttpServerClient) listNode(request *restful.Request, response *restful.Response) {
	nodeName := request.QueryParameter("node_name")
	nodeInternalIP := request.QueryParameter("node_ip")
	if nodeName == "" && nodeInternalIP == "" {
		_, _ = response.Write(CreateResponseData(errors.New("empty parameter: node_name and node_ip"), "", nil))
		return
	}

	var nodeIPs []string
	var err error
	if nodeName != "" {
		nodeIPs, err = h.NodeCache.GetNodeExternalIPsByName(nodeName)
		if err != nil {
			_, _ = response.Write(CreateResponseData(err, "", nil))
			return
		}
	}
	if nodeInternalIP != "" {
		nodeIPs, err = h.NodeCache.GetNodeExternalIPsByIP(nodeInternalIP)
		if err != nil {
			_, _ = response.Write(CreateResponseData(err, "", nil))
			return
		}
	}

	_, _ = response.Write(CreateResponseData(nil, "", nodeIPs))
}
