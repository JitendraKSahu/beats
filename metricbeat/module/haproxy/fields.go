// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package haproxy

import (
	"github.com/JitendraKSahu/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "haproxy", asset.ModuleFieldsPri, AssetHaproxy); err != nil {
		panic(err)
	}
}

// AssetHaproxy returns asset data.
// This is the base64 encoded gzipped contents of module/haproxy.
func AssetHaproxy() string {
	return "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"
}
