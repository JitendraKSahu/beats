// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package ibmmq

import (
	"github.com/JitendraKSahu/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "ibmmq", asset.ModuleFieldsPri, AssetIbmmq); err != nil {
		panic(err)
	}
}

// AssetIbmmq returns asset data.
// This is the base64 encoded gzipped contents of module/ibmmq.
func AssetIbmmq() string {
	return "eJy8VM1u2zwQvPspBjknfgAfvsPX5hCgCRogRXulyZW4CLUrk1TkvH1ByUkkR3brS3WyxeH8cKi9wTO9bsDbptmtgMw50AZXw/+rFeAo2chtZpUN/lsBGLG4V9cFWgGRAplEG9RmBVRMwaXNALyBmIY+yMuTX9sCjdq1hzcLCnOaKRXFqDFo/b5wtP/u/3vcP44wBK3TBPhZeklpYltSNiGYwjxZPuH47XnynMAJ2dOMYSBF79l6WCPYEmp+IYHJc1zmhtbvpEfst8b6OVyrt9Aq+PFw9+sa31i6/TWMOPxkcdqna3iTYNAJ7zoCO5LMFVPEs2gvKIvy2e4aT4spOKFL5JAVJiW1bDIhe5Y6IXXWF75dRx2hMWJqimkwY1Uqrrs4MlUcKKHn7I+1T6Qf+3um116jW+hr19Tx73t6KEG0GnqaeV3jcW69jfrCbkSx1EgUX9hSGuK3bWA7uE7jiY/b0gdv+WkythRU6rIne2rWFyUzkcvpJor5dMAv3khd/FmVTJKxNaUjlfFrWLO7SNNq05Rb8m9Vh00TgdPC34epA+McTuFpb5q2jLMJ9vBuwVSmfV5ytG+DkT8MgdsCYjnUPk4fFjQaCY6y4XBZ3/a83FeqBrG+XKyscIrek0zE1doupgsLd3Qm4UBbMOeu7u8AAAD//yFO4G0="
}
