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

package linux

import (
	"github.com/JitendraKSahu/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "linux", asset.ModuleFieldsPri, AssetLinux); err != nil {
		panic(err)
	}
}

// AssetLinux returns asset data.
// This is the base64 encoded gzipped contents of module/linux.
func AssetLinux() string {
	return "eJy8l0tv4zYQx+/+FIMcF91Hsm8fCgTdRVEUKYIGe+mhwogcWawpUuWQybqfviBl2YoiOUpTiZcgIfWfn/6ch/IStrRbg1YmfF8BeOU1reEs/X62AnCkCZnWkJPHFYAkFk7VXlmzhh9XANA8C5WVQdMKoFCkJa/T1kswWNFRPi6/q2kNG2dDvf/LgOZ9ma6UsMZ4h2J72BmSjKuP3q7BcM0aEu+DdGE4VBW63b29MZxHQse1lwNbgCmyAwywR6/YK8E/pDMkAYWzzPDT9TcQ1hH3tIagu+DS2T7bkVxbsxnYfAQ+rhrFljwn+ZokyEDg7dFWKFBpFRyNghE6vctmwjtykPFO0RHUW6hwS+CsraCwDgzdgTUPfO2ANgozULZsyoAvqQPtMdfjzhU2GDkDDgchiLkIWu+ACZ0oSY6+fkujNsYOXPP/l2JMZAC1I5S75BEJ31wk9u65X54dSMPkfBaTkuaw7rdQ5eRiObd3eleSI9CK/T54+6NhgBOot6jVMyEfyj9w1JfoQaAx1kNOkFwc8ObQAFM+ZI7Yo/MzWJhyHrS121BH+5QoocR0zznBPu6x0zTHHbH6p5OcLe2Wqzmmxn3Zk/PCo+/XzH+eFr/eXHUmwxMHQI0b4oxLdLNk/k0SbqLERhaYXk1gUYPxng2jPDHs9Ztoj8EEM5s134z6O9AJjAoPGLdWo1cDTf/5GNfpakSJZhNd8dZCgez3hdW8/bhLcRpkLNDMMQCPbdOrihgqcptU1006lXhLkMdmFQHMKUxO/SAzVlImSlSz4DZOpsaZ0BxhGpEVfs8icZvZ0zBlqOf1VIZaK4GxacYO0svDbhkoU9g5+uWA9qmmmQcpd1nvgXGgCYZ8QY9QOFvB69pZ8TpFiAEasVgPgQ/tK9+BdZLcEzvsl6vL0XscYp7AndivLhNXeomBE2NYXbSzN2eD+4+k2UTCROGIQJTBbDnm3MWfb15cX/78Nbv55Y+vp9HOF0c7n4p2sTjaxVS0t4ujvZ2K9m5xtHdT0d4vjvZ+KtqHxdE+TEX7uDjax6lonxZH+zQV7fPiaJ8nt9zlx8H52DxooeK3EL96MTjxbf4Xif4/mxNIfsc7QK2twHgK0sDvfAXEqRoD3PvSWP0bAAD//4Pi/m8="
}
