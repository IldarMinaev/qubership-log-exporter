// Copyright 2024 Qubership
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
    "fmt"
    "runtime"
)

var (
    Version   string
    BuildDate string
    Branch    string
    Revision  string
    GoVersion = runtime.Version()
    Platform  = runtime.GOOS + "-" + runtime.GOARCH
)

func versionString() string {
    return fmt.Sprintf("log-exporter version: %v (build date: %v, branch: %v, revision: %v, go version: %v, platform: %v)", Version, BuildDate, Branch, Revision, GoVersion, Platform)
}
