package main

import (
	"testing"

	"github.com/KevinGong2013/apkgo/cmd/shared"
)

func TestDo(t *testing.T) {

	cp := &CustomPublisher{}

	cp.Do(shared.PublishRequest{})
}
