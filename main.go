package main

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

// CallTracerResult is the result of a call tracer
type CallTracerResult struct {
	Type    string         `json:"type"`
	From    common.Address `json:"from"`
	To      common.Address `json:"to"`
	Value   *hexutil.Big   `json:"value,omitempty"`
	Gas     *hexutil.Big   `json:"gas,omitempty"`
	GasUsed *hexutil.Big   `json:"gasUsed,omitempty"`
	Input   hexutil.Bytes  `json:"input,omitempty"`
	Output  hexutil.Bytes  `json:"output,omitempty"`
	Error   string         `json:"error,omitempty"`
	Time    string
	Calls   []CallTracerResultCall `json:"calls,omitempty"`
}

type CallTracerResultCall struct {
	Type  string
	From  common.Address
	To    common.Address
	Value *hexutil.Big
	Error string
	Calls []CallTracerResultCall
}

func main() {
	txHash := os.Args[1]
	rpcClient, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		panic(err)
	}

	var result CallTracerResult
	err = rpcClient.Call(&result, "debug_traceTransaction", txHash, map[string]interface{}{"tracer": "callTracer"})
	if err != nil {
		panic(err)
	}

	spew.Dump(result)

}
