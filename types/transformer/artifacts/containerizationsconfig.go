/*
 *  Copyright IBM Corporation 2021
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package artifacts

import (
	"github.com/konveyor/move2kube/common"
	transformertypes "github.com/konveyor/move2kube/types/transformer"
	"github.com/sirupsen/logrus"
)

const (
	// ContainerizationOptionsConfigType represents containerization options config type
	ContainerizationOptionsConfigType transformertypes.ConfigType = "ContainerizationOptions"
)

// ContainerizationOptionsConfig stores the containerization options config
type ContainerizationOptionsConfig []string

// Merge implements the Config interface allowing artifacts to be merged
func (co *ContainerizationOptionsConfig) Merge(newcoobj interface{}) bool {
	newcoptr, ok := newcoobj.(*ContainerizationOptionsConfig)
	if !ok {
		newco, ok := newcoobj.(ContainerizationOptionsConfig)
		if !ok {
			logrus.Error("Unable to cast to ContainerizationOptionsConfig for merge")
			return false
		}
		newcoptr = &newco
	}
	*co = common.MergeStringSlices(*co, *newcoptr...)
	return true
}
