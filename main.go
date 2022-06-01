package main

import (
	"fmt"
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/surajresearch/fabric-samples/chaincode/student/studentcc"
)

func main() {
	err := shim.Start(&studentcc.StudentChaincode{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exiting student.StudentChaincode: %s", err)
		os.Exit(2)
	}
}
