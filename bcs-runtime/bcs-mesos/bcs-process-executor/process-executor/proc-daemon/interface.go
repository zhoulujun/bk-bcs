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

package procdaemon

import (
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-process-executor/process-executor/types"
)

// ProcDaemon xxx
type ProcDaemon interface {
	// CreateProcess xxx
	// create process object
	CreateProcess(*types.ProcessInfo) error

	// InspectProcessStatus xxx
	// inspect process status
	InspectProcessStatus(procId string) (*types.ProcessStatusInfo, error)

	// StopProcess xxx
	// stop process
	StopProcess(procId string, timeout int) error

	// DeleteProcess xxx
	// Delete process
	DeleteProcess(procId string) error

	// set process envs
	// types.BcsKV: key = env.key, value = env.value
	// SetProcessEnvs([]types.BcsKV)error

	// ReloadProcess xxx
	// reload process, exec reloadCmd
	ReloadProcess(procId string) error

	// RestartProcess xxx
	// restart process, exec restartCmd
	RestartProcess(procId string) error
}
