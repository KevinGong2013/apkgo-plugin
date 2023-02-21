package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/KevinGong2013/apkgo/cmd/shared"
)

type CustomPublisher struct{}

func (cp *CustomPublisher) Name() string {
	return "custom_publisher"
}

func (p *CustomPublisher) Do(req shared.PublishRequest) error {

	r := rand.Intn(10)
	time.Sleep(time.Second * time.Duration(10+r))

	if r%2 == 0 {
		return fmt.Errorf("upload %s failed", req.AppName)
	}

	return nil
}
