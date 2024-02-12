// Copyright 2023 The Wait4X Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package waiter

import (
	"math"
	"time"
)

func exponentialBackoff(retries int, backoffCoefficient float64, initialInterval, maxInterval time.Duration) time.Duration {
	interval := initialInterval * time.Duration(math.Pow(backoffCoefficient, float64(retries)))
	if interval > maxInterval {
		return maxInterval
	}
	return interval
}
