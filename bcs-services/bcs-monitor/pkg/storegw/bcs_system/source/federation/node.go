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
 */

package federation

import (
	"context"
	"time"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-monitor/pkg/storegw/bcs_system/source/base"
	"github.com/prometheus/prometheus/prompb"
)

// GetNodeInfo 节点信息
func (m *Federation) GetNodeInfo(ctx context.Context, projectID, clusterID, node string, t time.Time) (*base.NodeInfo,
	error) {
	return nil, nil
}

// GetNodeCPUUsage 节点CPU使用率
func (m *Federation) GetNodeCPUUsage(ctx context.Context, projectID, clusterID, node string, start, end time.Time,
	step time.Duration) ([]*prompb.TimeSeries, error) {
	return nil, nil
}

// GetNodeCPURequestUsage 节点CPU装箱率
func (m *Federation) GetNodeCPURequestUsage(ctx context.Context, projectID, clusterID, node string,
	start, end time.Time, step time.Duration) ([]*prompb.TimeSeries, error) {
	return nil, nil
}

// GetNodeMemoryUsage 节点内存使用率
func (m *Federation) GetNodeMemoryUsage(ctx context.Context, projectID, clusterID, node string, start, end time.Time,
	step time.Duration) ([]*prompb.TimeSeries, error) {
	return nil, nil
}

// GetNodeMemoryRequestUsage 节点内存装箱率
func (m *Federation) GetNodeMemoryRequestUsage(ctx context.Context, projectID, clusterID, node string,
	start, end time.Time, step time.Duration) ([]*prompb.TimeSeries, error) {
	return nil, nil
}

// GetNodeDiskUsage 节点磁盘使用率
func (m *Federation) GetNodeDiskUsage(ctx context.Context, projectID, clusterID, node string, start, end time.Time,
	step time.Duration) ([]*prompb.TimeSeries, error) {
	return nil, nil
}

// GetNodeDiskioUsage 接触磁盘IO使用率
func (m *Federation) GetNodeDiskioUsage(ctx context.Context, projectID, clusterID, node string, start, end time.Time,
	step time.Duration) ([]*prompb.TimeSeries, error) {
	return nil, nil
}

// GetNodePodCount PodCount
func (m *Federation) GetNodePodCount(ctx context.Context, projectID, clusterID, node string, start, end time.Time,
	step time.Duration) ([]*prompb.TimeSeries, error) {
	return nil, nil
}

// GetNodePodTotal PodTotal
func (m *Federation) GetNodePodTotal(ctx context.Context, projectID, clusterID, node string, start, end time.Time,
	step time.Duration) ([]*prompb.TimeSeries, error) {
	return nil, nil
}

// GetNodeContainerCount 容器数量
func (m *Federation) GetNodeContainerCount(ctx context.Context, projectID, clusterID, node string, start, end time.Time,
	step time.Duration) ([]*prompb.TimeSeries, error) {
	return nil, nil
}

// GetNodeNetworkTransmit 网络发送量
func (m *Federation) GetNodeNetworkTransmit(ctx context.Context, projectID, clusterID, node string,
	start, end time.Time, step time.Duration) ([]*prompb.TimeSeries, error) {
	return nil, nil
}

// GetNodeNetworkReceive 节点网络接收量
func (m *Federation) GetNodeNetworkReceive(ctx context.Context, projectID, clusterID, node string, start, end time.Time,
	step time.Duration) ([]*prompb.TimeSeries, error) {
	return nil, nil
}
