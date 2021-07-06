package main

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_Main(t *testing.T) {
	// something that uses gomock
	gomock.NewController(t)
}
