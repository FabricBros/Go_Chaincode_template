package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoChaincodeTemplate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "chaincode suite")
}
