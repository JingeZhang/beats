// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package foo

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "foo", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJx8j0EOwiAQRfdzip/uewEW7jyFcYEyGFLoEEpje3vTooY21bf8E94LLTqeFawIAdllzwqNFWkIMDzck4vZSa9wIgDgKXJygfus/eVK62ZFEMSMngmwjr0Z1Hpo0evAH/lCniMrPJKM8b0cNLaSWnTT6bsdyX4KC9vn+0gd4kmHuP6npgQ7np+SzO72J7twLsISpVcAAAD//+HhYIk="
}
