# anvil-tests

Quick tests to validate anvil works with `debug_traceTransaction` - I may add more tests to make sure ethereum-go remains compatible with anvil

I followed instructions here to get a tx and passed this as argument
https://book.getfoundry.sh/tutorials/forking-mainnet-with-cast-anvil

```
go run main.go 0x472c87f813e096d9cd501754a9d0f95c3660e255b1a17702fac5af542fe70f69
(main.CallTracerResult) {
 Type: (string) (len=4) "CALL",
 From: (common.Address) (len=20 cap=20) 0xfc2eE3bD619B7cfb2dE2C797b96DeeCbD7F68e46,
 To: (common.Address) (len=20 cap=20) 0x6B175474E89094C44Da98b954EedeAC495271d0F,
 Value: (*hexutil.Big)(0x14000144520)(0x0),
 Gas: (*hexutil.Big)(0x140001444e0)(0xca76),
 GasUsed: (*hexutil.Big)(0x14000144500)(0xca76),
 Input: (hexutil.Bytes) (len=68 cap=68) 0xa9059cbb000000000000000000000000f39fd6e51aad88f6f4ce6ab8827279cfffb92266000000000000000000000000000000000000000000003f870857a3e0e3800000,
 Output: (hexutil.Bytes) (len=32 cap=32) 0x0000000000000000000000000000000000000000000000000000000000000001,
 Error: (string) "",
 Time: (string) "",
 Calls: ([]main.CallTracerResultCall) <nil>
}
```
