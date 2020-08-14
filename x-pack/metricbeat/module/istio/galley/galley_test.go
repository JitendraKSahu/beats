// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build !integration

package galley

import (
	"testing"

	mbtest "github.com/JitendraKSahu/beats/v7/metricbeat/mb/testing"

	_ "github.com/JitendraKSahu/beats/v7/x-pack/metricbeat/module/istio"
)

func TestData(t *testing.T) {
	mbtest.TestDataFiles(t, "istio", "galley")
}
