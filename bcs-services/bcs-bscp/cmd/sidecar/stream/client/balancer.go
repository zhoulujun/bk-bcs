/*
Tencent is pleased to support the open source community by making Basic Service Configuration Platform available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "as IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

func newBalancer(endpoints []string) (*balancer, error) {
	if len(endpoints) == 0 {
		return nil, errors.New("no endpoints is set, can initial the client round-robin balancer")
	}

	rand.Seed(time.Now().UnixNano())

	return &balancer{
		lo:        sync.Mutex{},
		index:     uint(rand.Intn(len(endpoints))),
		max:       uint(len(endpoints) - 1),
		endpoints: endpoints,
	}, nil

}

type balancer struct {
	lo        sync.Mutex
	index     uint
	max       uint
	endpoints []string
}

// PickOne pick one endpoint.
func (r *balancer) PickOne() string {
	r.lo.Lock()
	defer r.lo.Unlock()

	if r.index == r.max {
		r.index = 0
		return r.endpoints[r.max]
	}

	r.index++

	return r.endpoints[r.index-1]
}
