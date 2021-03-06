// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build !windows

package reexec

import (
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"

	"github.com/jksroot/beats/v7/x-pack/elastic-agent/pkg/core/logger"
)

func reexec(log *logger.Logger, executable string) error {
	// force log sync, before re-exec
	_ = log.Sync()

	args := []string{filepath.Base(executable)}
	args = append(args, os.Args[1:]...)
	return unix.Exec(executable, args, os.Environ())
}
