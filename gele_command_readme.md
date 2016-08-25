* [Please note `web3.js` document with the `gele console` commands and instructions. Use web3.js document to see all the command displayed on the current page.](https://github.com/elementrem/web3.js/blob/master/web3_document.md)    

* [gele command line option list](command_line_option.md)
#### gele console command list 1

[admin](#admin)                         | [debug](#debug)                           |[ele](#ele)               |[ele](#ele)         |
------------                            | -------------                             |     -------------     | -------------     |
[datadir](#admindatadir)                | [backtraceAt](#debugbacktraceat)          |[accounts](#eleaccounts)  |getPendingTransactions
[nodeInfo](#adminnodeinfo)              | [blockProfile](#debugblockprofile)        |[blockNumber](#eleblocknumber)|getStorageAt
[peers](#adminpeers)                    | [chaindbProperty](#debugcpuprofile)       |[coinbase](#elecoinbase)               |getSyncing
[addPeer](#adminaddpeer)                | [cpuProfile](#debugcpuprofile)            |compile                |getTransaction
[exportChain](#adminexportchain)       | [dumpBlock](#debugdumpblock)              |defaultAccount         |[getTransactionCount](#elegettransactioncount)
[getContractInfo](#admingetcontractinfo)| [gcStats](#debuggcstats)                  |defaultBlock           |[getTransactionFromBlock](#elegettransactionfromblock)
getDatadir                              | [getBlockRlp](#debuggetblockrlp)          |[gasPrice](#elegasprice)               |[getTransactionReceipt](#elegettransactionreceipt)
getNodeInfo                             | [goTrace](#debuggotrace)                  |[hashrate](#elehashrate)               |getUncle|isSyncing
getPeers                                | [memStats](#debugmemstats)                |[mining](#elemining)                 |[getWork](#elegetwork)|iban|namereg
httpGet                                 | metrics                                   |[pendingTransactions](#elependingtransactions)    |icapNamereg
importChain                             | [printBlock](#debugprintblock)            |[syncing](#elesyncing)                |resend
[register](#adminregister)              | seedHash                                  |call                 |sendIBANTransaction
[registerUrl](#adminregisterurl)        | [setBlockProfileRate](#debugsetblockprofilerate)|contract            |sendRawTransaction
saveInfo                                | [setHead](#debugsethead)                  |estimateGas            |[sendTransaction](#elesendtransaction)
setGlobalRegistrar                      | [stacks](#debugstacks)                    |filter                 |[sign](#elesign)
setHashReg                              | [startCPUProfile](#debugstartcpuprofile)  |getAccounts            |signTransaction
[setSolc](#adminsetsolc)                | [startGoTrace](#debugstartgotrace)        |[getBalance](#elegetbalance)             |submitTransaction
setUrlHint                              | [stopCPUProfile](#debugstopcpuprofile)    |[getBlock](#elegetblock)               |[submitWork](#elesubmitwork)
[sleep](#adminsleep)                    | [stopGoTrace](#debugstopgotrace)          |[getBlockByNumber](#elegetblockbynumber) |[isSyncing](#eleissyncing)
[sleepBlocks](#adminsleepblocks)        | [traceBlock](#debugtraceblock)            |[getBlockTransactionCount](#elegetblocktransactioncount)
[startNatSpec](#adminstartnatspec)      | traceBlockByFile                              |[getBlockUncleCount](#elegetblockunclecount)
[startRPC](#adminstartrpc)              | [traceBlockByHash](#debugtraceblockbyhash)    |[getCode](#elegetcode)
[startWS](#adminstartws)                | [traceBlockByNumber](#debugtraceblockbynumber)|getCoinbase
[stopNatSpec](#adminstopnatspec)        | [traceTransaction](#debugtracetransaction)    |[getCompilers](#elegetcompilers)
[stopRPC](#adminstoprpc)                | [verbosity](#debugverbosity)                  |getGasPrice
[stopWS](#adminstopws)                  | [vmodule](#debugvmodule)                      |getHashrate
                                        | [writeBlockProfile](#debugwriteblockprofile)  |getMining
                                        | [writeMemProfile](#debugwritememprofile)      |getNatSpec

#### gele console command list 2

 [miner](miner.md)        | [net](#net)                 |[personal](#personal)                    | [txpool](#txpool)        |[web3.version](#web3version)|
------------  | ------------- | ------------| -------------                           | -------------            | ------------   |
[makeDAG](miner.md)       |[listening](#netlistening)   |[listAccounts](#personallistaccounts)    |[content](#txpoolcontent) |[api](#web3version) 
setElementbase|[peerCount](#netpeercount)   |getListAccounts                          |[inspect](#txpoolinspect) |[elementrem](#web3version)
[setExtra](miner.md)      |[version](#netversion)       |[importRawKey](#personalimportrawkey)    |[status](#txpoolstatus)   |[network](#web3version)     
[setGasPrice](miner.md)   |getListening                 |[lockAccount](#personallockaccount)      |getContent                |[node](#web3version)
[start](miner.md)         |getPeerCount                 |[newAccount](#personalnewaccount)        |getInspect                |[whisper](#web3version) 
[startAutoDAG](miner.md)  |getVersion                   |[signAndSendTransaction](#personalsignandsendtransaction)|getStatus |getElementrem 
[stop](miner.md)          |                             |[unlockAccount](#personalunlockaccount)  |                          |getNetwork
[stopAutoDAG](miner.md)   |                             |                                         |                          | getNode 
|||||getWhisper

[web3.db](#web3db)   |[web3.currentProvider](#web3currentprovider)|[web3.providers](#web3providers)|[web3.shh](#web3shh)
  ------------       |    -------------                   |                    ------------        | ------------
[getHex](#web3db)    |[newAccount](#web3currentprovider)          |[HttpProvider](#web3providers)  |[filter](#web3shh)
[getString](#web3db) |[send](#web3currentprovider)                |[IpcProvider](#web3providers)   |[hasIdentity](#web3shh)
[putHex](#web3db)    |[sendAsync](#web3currentprovider)                                           ||[newGroup](#web3shh)
[putString](#web3db) |[unlockAccount](#web3currentprovider)                                       ||[newIdentity](#web3shh)
                     |                                                                            ||[post](#web3shh)                  
                     |                                                                            ||[addToGroup](#web3shh)


                                 |              |              |              |
 ------------                    | ------------ | ------------ | ------------ |
[web3.createBatch](#batch-requests)   |web3.isAddress          |web3.toAscii            |web3.toUtf8
web3.fromAscii                        |web3.isChecksumAddress  |web3.toBigNumber        |web3.isIBAN
web3.fromDecimal                      |web3.isConnected        |web3.toChecksumAddress  
web3.fromICAP                         |web3.reset              |web3.toDecimal
[web3.fromMey](#web3frommey)                        |web3.setProvider        |web3.toHex
web3.fromUtf8                         |web3.sha3               |[web3.toMey](#web3tomey)














## admin
### admin.datadir
The datadir administrative property can be queried for the absolute path the running Gele node currently uses to store all its databases.    

 Client          | Method        |
------------  | ------------- | 
Console    |admin.datadir        |   

```bash
- Example   
>admin.datadir    
"/home/ubuntu/.elementrem"   
```
* [Back to Top](#gele-console-command-list-1)   

### admin.nodeInfo
The admin exposes the methods to manage, control or monitor your node.

 Client          | Method        |
------------  | ------------- | 
Console    |admin.nodeInfo        |   

```bash
- Example   
>admin.nodeInfo    
{
  enode: "enode://4de730c6fd4c9e64d2cc0ebb846e4feb164b145dd1c885d2150372903e2bac3cd32e6d4997d2bd76646fb78dbf976b5afc70eaebbd10acaae94eb9d932bc77e8@[::]:30707",
  id: "4de730c6fd4c9e64d2cc0ebb846e4feb164b145dd1c885d2150372903e2bac3cd32e6d4997d2bd76646fb78dbf976b5afc70eaebbd10acaae94eb9d932bc77e8",
  ip: "::",
  listenAddr: "[::]:30707",
  name: "Gele/v1.4.7-stable-64a5a363/linux/go1.6.2",
  ports: {
    discovery: 30707,
    listener: 30707
  },
  protocols: {
    ele: {
      difficulty: 20253847329,
      genesis: "0xc7130c992fb03c628c86439f8e5d1facdbf48f3f3c115ffd99a0e3abefa1016f",
      head: "0x214d56d49ca1cbb728d566aeb3a0292546c71086d7520a77bab7b4cc5a5589f2",
      network: 73733
    }
  }
}  
```
* [Back to Top](#gele-console-command-list-1)   

### admin.peers
The peers administrative property can be queried for all the information known about the connected remote nodes at the networking granularity.    

 Client          | Method        |
------------  | ------------- | 
Console    |admin.peers       |   

```bash
- Example   
>admin.peers
[{
    caps: ["ele/61", "ele/62", "ele/63"],
    id: "7e8ec02b19bec6e150645d785b57a3a60eb64aac729750c97fe9fcd15dc324fa8ffa1693f0d0ad7727a7e940bbfa2cc6e0680b5b12988f67be57fe5a146a9a8b",
    name: "Gele/v1.4.7-stable-64a5a363/linux/go1.6.2",
    network: {
      localAddress: "192.168.204.134:53342",
      remoteAddress: "000.000.000.000:30707"
    },
    protocols: {
      ele: {
        difficulty: 41160129719,
        head: "08cac5e0404d11a6e028b60e254473e1a4bddc8d38ef5259cfa2a55e904e0ae1",
        version: 63
      }
    }
}, {
    caps: ["ele/61", "ele/62", "ele/63"],
    id: "d2b05122f4ef26cce1a44450903cdab898928ebbb29b5ca85ecd2ecf7689207a7e8a75bceb8f57dde9a92315343cb357c092f2fffd56f11755818aea23394ec0",
    name: "Gele/v1.4.7-stable-64a5a363/linux/go1.6.2",
    network: {
      localAddress: "192.168.204.134:56570",
      remoteAddress: "000.000.000.000:30707"
    },
    protocols: {
      ele: {
        difficulty: 41340455879,
        head: "564d4122f054f6d0246f066d348018073bf9f079d3f1260b2a8e4fc7ccb730da",
        version: 63
      }
    }
}, 
```
* [Back to Top](#gele-console-command-list-1) 

### admin.addPeer
Pass a nodeURL to connect a to a peer on the network. The nodeURL needs to be in enode URL format. gele will maintain the connection until it shuts down and attempt to reconnect if the connection drops intermittently.   

 Client          | Method        |
------------  | ------------- | 
Console    |admin.addPeer      |   

```bash
- Example   
> admin.addPeer("enode://d2b05122f4ef26cce1a44450903cdab898928ebbb29b5ca85ecd2ecf7689207a7e8a75bceb8f57dde9a92315343cb357c092f2fffd56f11755818aea23394ec0@00.00.00.00:30707")
true
```
* [Back to Top](#gele-console-command-list-1) 

### admin.exportChain
Exports the blockchain to the given file in binary format. It is output to the gele path.

 Client          | Method        |
------------  | ------------- | 
Console    |admin.exportChain("filename")      |
RPC|curl -X POST --data '{"jsonrpc":"2.0","method":"admin_exportChain","params":["filename"],"id":67}' 127.0.0.1:7075|

```bash
> admin.exportChain("chaindata")
true
//
{"jsonrpc":"2.0","id":67,"result":true}
```

### admin.getContractInfo
this will retrieve the contract info json for a contract on the address   

 Client          | Method        |
------------  | ------------- | 
Console    |admin.getContractInfo     |   

```bash
- Example   
> info = admin.getContractInfo(contractaddress)
> source = info.source
> abi = info.abiDefinition
```
* [Back to Top](#gele-console-command-list-1) 


### admin.register
will register content hash to the codehash (hash of the code of the contract on contractaddress). The register transaction is sent from the address in the first parameter. The transaction needs to be processed and confirmed on the canonical chain for the registration to take effect.  

 Client          | Method        |
------------  | ------------- | 
Console    |admin.register    |   

```bash
- Example   
source = "contract test {\n" +
"   /// @notice will multiply `a` by 7.\n" +
"   function multiply(uint a) returns(uint d) {\n" +
"      return a * 7;\n" +
"   }\n" +
"} ";
contract = ele.compile.solidity(source).test;
txhash = ele.sendTransaction({from: primary, data: contract.code });
// after it is uncluded
contractaddress = ele.getTransactionReceipt(txhash);
filename = "/tmp/info.json";
contenthash = admin.saveInfo(contract.info, filename);
admin.register(primary, contractaddress, contenthash);
true
```
* [Back to Top](#gele-console-command-list-1) 


### admin.registerUrl
this will register a contant hash to the contract' codehash. This will be used to locate contract info json files. Address in the first parameter will be used to send the transaction.

 Client          | Method        |
------------  | ------------- | 
Console    |admin.registerUrl    |   

```bash
- Example   
source = "contract test {\n" +
"   /// @notice will multiply `a` by 7.\n" +
"   function multiply(uint a) returns(uint d) {\n" +
"      return a * 7;\n" +
"   }\n" +
"} ";
contract = ele.compile.solidity(source).test;
txhash = ele.sendTransaction({from: primary, data: contract.code });
// after it is uncluded
contractaddress = ele.getTransactionReceipt(txhash);
filename = "/tmp/info.json";
contenthash = admin.saveInfo(contract.info, filename);
admin.register(primary, contractaddress, contenthash);
admin.registerUrl(primary, contenthash, "file://"+filename);
true
```
* [Back to Top](#gele-console-command-list-1) 

### admin.setSolc
Set the solidity compiler

 Client          | Method        |
------------  | ------------- | 
Console    |admin.setSolc    |   

```bash
- Example   
admin.setSolc('/some/path/solc')
'solc v0.9.29
Solidity Compiler: /some/path/solc
'
```
* [Back to Top](#gele-console-command-list-1) 


### admin.sleep
Sleeps for s seconds.

 Client          | Method        |
------------  | ------------- | 
Console    |admin.sleep(s)    |   

* [Back to Top](#gele-console-command-list-1) 

### admin.sleepBlocks
Sleeps for n blocks.

 Client          | Method        |
------------  | ------------- | 
Console    |admin.sleepBlocks(n)    |   

* [Back to Top](#gele-console-command-list-1) 

### admin.startNatSpec
activate NatSpec: when sending a transaction to a contract, Registry lookup and url fetching is used to retrieve authentic contract Info for it. It allows for prompting a user with authentic contract-specific confirmation messages.

 Client          | Method        |
------------  | ------------- | 
Console    |admin.startNatSpec()    |   

* [Back to Top](#gele-console-command-list-1) 


### admin.startRPC
The startRPC administrative method starts an HTTP based JSON RPC API webserver to handle client requests. All the parameters are optional:   
- host: network interface to open the listener socket on (defaults to "localhost")
- port: network port to open the listener socket on (defaults to 7075)
- cors: cross-origin resource sharing header to use (defaults to "")
- apis: API modules to offer over this interface (defaults to "ele,ne0t,web3")

 Client          | Method        |
------------  | ------------- | 
Console    |admin.startRPC    |   

```bash
- Example   
> admin.startRPC("127.0.0.1", 7075)
true
```
* [Back to Top](#gele-console-command-list-1) 


### admin.startWS
The startWS administrative method starts an WebSocket based JSON RPC API webserver to handle client requests. All the parameters are optional:   
- host: network interface to open the listener socket on (defaults to "localhost")
- port: network port to open the listener socket on (defaults to 7075)
- cors: cross-origin resource sharing header to use (defaults to "")
- apis: API modules to offer over this interface (defaults to "ele,ne0t,web3")

 Client          | Method        |
------------  | ------------- | 
Console    |admin.startRPC    |   

```bash
- Example   
> admin.startWS("127.0.0.1", 7076)
true
```
* [Back to Top](#gele-console-command-list-1) 

### admin.stopNatSpec
deactivate NatSpec: when sending a transaction, the user will be prompted with a generic confirmation message, no contract info is fetched

 Client          | Method        |
------------  | ------------- | 
Console    |admin.stopNatSpec()    |   

* [Back to Top](#gele-console-command-list-1) 


### admin.stopRPC
Stops the HTTP server for the JSON-RPC.

 Client          | Method        |
------------  | ------------- | 
Console    |admin.stopRPC()    |   

* [Back to Top](#gele-console-command-list-1) 

### admin.stopWS
Stops the websocket server for the JSON-RPC.

 Client          | Method        |
------------  | ------------- | 
Console    |admin.stopWS()     |   

* [Back to Top](#gele-console-command-list-1) 








## debug
The debug API gives you access to several non-standard RPC methods, which will allow you to inspect, debug and set certain debugging flags during runtime.

### debug.backtraceAt
Sets the logging backtrace location. When a backtrace location is set and a log message is emitted at that location, the stack of the goroutine executing the log statement will be printed to stderr.

The location is specified as <filename>:<line>.  

 Client          | Method        |
------------  | ------------- | 
Console    |debug.backtraceAt(string)        |   

```bash
> debug.backtraceAt("server.go:443")
```
* [Back to Top](#gele-console-command-list-1)  



### debug.blockProfile
Turns on block profiling for the given duration and writes profile data to disk. It uses a profile rate of 1 for most accurate information. If a different rate is desired, set the rate and write the profile manually using debug_writeBlockProfile. 

 Client          | Method        |
------------  | ------------- | 
Console    |debug.blockProfile(file, seconds)       |   

* [Back to Top](#gele-console-command-list-1)  


### debug.cpuProfile
Turns on CPU profiling for the given duration and writes profile data to disk.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.cpuProfile(file, seconds)       |   

* [Back to Top](#gele-console-command-list-1)  



### debug.dumpBlock
Retrieves the state that corresponds to the block number and returns a list of accounts (including storage and code).

 Client          | Method        |
------------  | ------------- | 
Console    |debug.dumpBlock        |   

```bash
Example
> debug.dumpBlock(10)
{
  accounts: {
    79f512cd283989eca606379471d76c93f1a66f56: {
      balance: "50000000000000000000",
      code: "",
      codeHash: "c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470",
      nonce: 0,
      root: "56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
      storage: {}
    },
    a748f720f5989f2b541fa9ef3c78236808123635: {
      balance: "50000000000000000000000000",
      code: "",
      codeHash: "c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470",
      nonce: 0,
      root: "56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
      storage: {}
    }
  },
  root: "c7162345dd67efa5a3a725b5b84847d7facb8e0dfd8be516c1f257f0a06f6e67"
}
>
```
* [Back to Top](#gele-console-command-list-1)  


### debug.gcStats
Returns GC statistics.

See https://golang.org/pkg/runtime/debug/#GCStats for information about the fields of the returned object.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.gcStats()        |   

* [Back to Top](#gele-console-command-list-1)  


### debug.getBlockRlp
Returns the hexadecimal representation of the RLP encoding of the block.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.getBlockRlp(blockNumber)        |   

```bash
Example
> debug.getBlockRlp(1000)
"f90216f90211a06265ef89489451dc390ea8915edb1197d73810301747f8d6ec4bb450c8dd96faa01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d493479479f512cd283989eca606379471d76c93f1a66f56a0f2186db64c74757c93d674f7e7fa88bf63ad1fb117a5cccd0d4ddcee957d8a85a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b901000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000083033ee48203e88347e7c48084577ad1dc98d7830104078447656c6587676f312e362e32856c696e7578a0a3729eef354613c10424b80da189c895eaff704caef4137ba97918d5bc358fe6885ef9de36ce5270bdc0c0"
```
* [Back to Top](#gele-console-command-list-1)  




### debug.goTrace
Turns on Go runtime tracing for the given duration and writes trace data to disk.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.goTrace(file, seconds)        |   

* [Back to Top](#gele-console-command-list-1)  


### debug.memStats
Returns detailed runtime memory statistics.

See https://golang.org/pkg/runtime/#MemStats for information about the fields of the returned object.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.memStats()        |   

* [Back to Top](#gele-console-command-list-1)  

### debug.printBlock
Prints information about the block such as size, total difficulty, as well as header fields properly formatted.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.printBlock(blockNumber)       |   
```bash
Example
> debug.printBlock(10)
"Block(#10): Size: 535.00 B {\nMinerHash: 31b08a0672898a49ab1351aeb542f45d16da07fa879f9bd8e83309ac12cb506d\nHeader(5d1440943c2b6627285e2017265da72d455dcb542582cf02b4bae4f5736346bf):\n[\n\tParentHash:\t    bbae450122bc3242e7c7389dc81f00026358fb0fefa8659f2625e4dd14066038\n\tUncleHash:\t    1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347\n\tCoinbase:\t    79f512cd283989eca606379471d76c93f1a66f56\n\tRoot:\t\t    c7162345dd67efa5a3a725b5b84847d7facb8e0dfd8be516c1f257f0a06f6e67\n\tTxSha\t\t    56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421\n\tReceiptSha:\t    56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421\n\tBloom:\t\t    00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\n\tDifficulty:\t    131584\n\tNumber:\t\t    10\n\tGasLimit:\t    3172388\n\tGasUsed:\t    0\n\tTime:\t\t    1467665594\n\tExtra:\t\t    ׃\x01\x04\a�Gele�go1.6.2�linux\n\tMixDigest:      2658a5fc9782dba93cd42c1bca42664c388853ca06550b37d2d91e0a0811dae7\n\tNonce:\t\t    34dfde497db37171\n]\nTransactions:\n[]\nUncles:\n[]\n}\n"
```
* [Back to Top](#gele-console-command-list-1)  


### debug.setBlockProfileRate
Sets the rate (in samples/sec) of goroutine block profile data collection. A non-zero rate enables block profiling, setting it to zero stops the profile. Collected profile data can be written using debug_writeBlockProfile.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.setBlockProfileRate(rate)       |   

* [Back to Top](#gele-console-command-list-1)  

### debug.setHead
Sets the rate (in samples/sec) of goroutine block profile data collection. A non-zero rate enables block profiling, setting it to zero stops the profile. Collected profile data can be written using debug_writeBlockProfile.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.setHead(number)       |   
```bash
Example
> debug.setHead(1000)   // Your block forced to go back to 1000. And the sink begins again.
null
```
* [Back to Top](#gele-console-command-list-1)  

### debug.stacks
Returns a printed representation of the stacks of all goroutines. Note that the web3 wrapper for this method takes care of the printing and does not return the string.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.stacks()       |   

* [Back to Top](#gele-console-command-list-1)  

### debug.startCPUProfile
Turns on CPU profiling indefinitely, writing to the given file.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.startCPUProfile(file)       |   

* [Back to Top](#gele-console-command-list-1)  


### debug.startGoTrace
Starts writing a Go runtime trace to the given file.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.startGoTrace(file)      |   

* [Back to Top](#gele-console-command-list-1)  


### debug.stopCPUProfile
Stops an ongoing CPU profile.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.stopCPUProfile()      |   

* [Back to Top](#gele-console-command-list-1)  


### debug.stopGoTrace
Stops an ongoing CPU profile.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.startGoTrace(file)      |   

* [Back to Top](#gele-console-command-list-1)  


### debug.traceBlock
The traceBlock method will return a full stack trace of all invoked opcodes of all transaction that were included included in this block. Note, the parent of this block must be present or it will fail.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.traceBlock(tblockRlp, [options])       |   

* [Back to Top](#gele-console-command-list-1)  

### debug.traceBlockByHash
Similar to debug_traceBlock, traceBlockByHash accepts a block hash and will replay the block that is already present in the database.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.traceBlockByHash(hash, [options])      |   

* [Back to Top](#gele-console-command-list-1)  


### debug.traceBlockByNumber
The traceBlock method will return a full stack trace of all invoked opcodes of all transaction that were included included in this block. Note, the parent of this block must be present or it will fail.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.traceBlockByNumber(number, [options])       |   

* [Back to Top](#gele-console-command-list-1)  


### debug.traceTransaction
The traceTransaction debugging method will attempt to run the transaction in the exact same manner as it was executed on the network. It will replay any transaction that may have been executed prior to this one before it will finally attempt to execute the transaction that corresponds to the given hash.

In addition to the hash of the transaction you may give it a secondary optional argument, which specifies the options for this specific call. The possible options are:

disableStorage: BOOL. Setting this to true will disable storage capture (default = false).
disableMemory: BOOL. Setting this to true will disable memory capture (default = false).
disableStack: BOOL. Setting this to true will disable stack capture (default = false).
fullStorage: BOOL. Setting this to true will return you, for each opcode, the full storage, including everything which hasn't changed. This is a slow process and is therefor defaulted to false. By default it will only ever give you the changed storage values.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.traceTransaction(txHash, [options])       |   

```bash
> debug.traceTransaction("0xb842cfd9860c41b955cf9357d5909a2b4e7b6de200903aeae7341e37b36b0d77")
{
  gas: 21000,
  returnValue: "",
  structLogs: []
```
* [Back to Top](#gele-console-command-list-1)  


### debug.verbosity
Sets the logging verbosity ceiling. Log messages with level up to and including the given level will be printed.
The verbosity of individual packages and source files can be raised using debug_vmodule.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.verbosity(level)       |   

* [Back to Top](#gele-console-command-list-1) 


### debug.vmodule
Sets the logging verbosity pattern.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.vmodule(string)       |   

* [Back to Top](#gele-console-command-list-1) 

### debug.writeBlockProfile
Writes a goroutine blocking profile to the given file.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.writeBlockProfile(file)       |   

* [Back to Top](#gele-console-command-list-1) 

### debug.writeMemProfile
Writes an allocation profile to the given file. Note that the profiling rate cannot be set through the API, it must be set on the command line using the --memprofilerate flag.

 Client          | Method        |
------------  | ------------- | 
Console    |debug.writeMemProfile(file string)       |   

* [Back to Top](#gele-console-command-list-1) 





## ele
The ele is a shortcut for web3.ele.     
In addition to the web3 and ele interfaces exposed by web3.js a few additional calls are exposed.

### ele.accounts
Returns a list of addresses owned by client.

 Client          | Method        |
------------  | ------------- | 
Console    |ele.accounts           |   
RPC        |curl -X POST --data '{"jsonrpc":"2.0","method":"ele_accounts","params":[],"id":1}' 127.0.0.1:7075    |
```bash
> ele.accounts
["0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4"]
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"ele_accounts","params":[],"id":1}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":1,"result":["0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4"]}
```
* [Back to Top](#gele-console-command-list-1) 
 


### ele.blockNumber
Returns the number of most recent block.

 Client          | Method        |
------------  | ------------- | 
Console    |ele.blockNumber           |   
RPC        |curl -X POST --data '{"jsonrpc":"2.0","method":"ele_blockNumber","params":[],"id":83}' 127.0.0.1:7075    |
```bash
> ele.blockNumber
22505
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"ele_blockNumber","params":[],"id":83}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":83,"result":"0x57e9"}
```
* [Back to Top](#gele-console-command-list-1) 



### ele.coinbase
Returns the client coinbase address.

 Client          | Method        |
------------  | ------------- | 
Console    |ele.coinbase          |   
RPC        |curl -X POST --data '{"jsonrpc":"2.0","method":"ele_coinbase","params":[],"id":64}' 127.0.0.1:7075    |
```bash
> ele.coinbase
"0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4"
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"ele_coinbase","params":[],"id":64}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":64,"result":"0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4"}
```
* [Back to Top](#gele-console-command-list-1) 


### ele.gasPrice
Returns the current price per gas in `mey`

 Client          | Method        |
------------  | ------------- | 
Console    |ele.gasPrice          |   
RPC        |curl -X POST --data '{"jsonrpc":"2.0","method":"ele_gasPrice","params":[],"id":73}' 127.0.0.1:7075    |
```bash
> ele.gasPrice
20000000000
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"ele_gasPrice","params":[],"id":73}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":73,"result":"0x4a817c800"}
```
* [Back to Top](#gele-console-command-list-1) 

### ele.hashrate
Returns the number of hashes per second that the node is mining with.

 Client          | Method        |
------------  | ------------- | 
Console    |ele.hashrate          |   
RPC        |curl -X POST --data '{"jsonrpc":"2.0","method":"ele_hashrate","params":[],"id":71}' 127.0.0.1:7075    |
```bash
> ele.hashrate
928210
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"ele_hashrate","params":[],"id":71}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":71,"result":"0xefbe1"}
```
* [Back to Top](#gele-console-command-list-1) 

### ele.mining
Returns true if client is actively mining new blocks.

 Client          | Method        |
------------  | ------------- | 
Console    |ele.mining          |   
RPC        |curl -X POST --data '{"jsonrpc":"2.0","method":"ele_mining","params":[],"id":71}' 127.0.0.1:7075    |
```bash
> ele.mining
false
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"ele_mining","params":[],"id":71}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":71,"result":false}
```
* [Back to Top](#gele-console-command-list-1) 


### ele.pendingTransactions
Returns pending transactions that belong to one of the users ele.accounts

 Client          | Method        |
------------  | ------------- | 
Console    |ele.pendingTransactions         |   
RPC        |curl -X POST --data '{"jsonrpc":"2.0","method":"ele_pendingTransactions","params":[],"id":71}' 127.0.0.1:7075    |
```bash
> ele.pendingTransactions
[]
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"ele_pendingTransactions","params":[],"id":71}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":71,"result":[]}
```
* [Back to Top](#gele-console-command-list-1) 

### ele.syncing
Returns block syncing status. If the block returns a value of `false` information is up to date. 

 Client          | Method        |
------------  | ------------- | 
Console    |ele.syncing         |   
RPC        |curl -X POST --data '{"jsonrpc":"2.0","method":"ele_syncing","params":[],"id":71}' 127.0.0.1:7075    |
```bash
> ele.syncing
false
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"ele_syncing","params":[],"id":71}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":71,"result":false}
```
* [Back to Top](#gele-console-command-list-1) 


### ele.getBalance
Returns the balance of the account of given address.    
**Parameters**    
`DATA`, 20 Bytes - address to check for balance.    
`QUANTITY|TAG` - integer block number, or the string `"latest"`, `"earliest"` or `"pending"`   
```bash
params: [
   '0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4',
   'latest'
]
```
 Client          | Method        |
------------  | ------------- | 
Console    |ele.getBalance        |   
RPC        |curl -X POST --data '{"jsonrpc":"2.0","method":"ele_getBalance","params":["address", "latest"],"id":1}' 127.0.0.1:7075    |
```bash
> ele.getBalance("0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4")
70651268000000000000
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"ele_getBalance","params":["0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4", "latest"],"id":1}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":1,"result":"0x3d47bb9eb66ec4000"}
```
Returns `QUANTITY` - integer of the current balance in `mey` But **Simply `mey` can be converted to the `element`**
```bash
> web3.fromMey(ele.getBalance("0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4"), "element")
70.651268
```
* [Back to Top](#gele-console-command-list-1) 



### ele.getBlock
Returns the selected block information.

 Client          | Method        |
------------  | ------------- | 
Console    |ele.getBlock         |   
```bash
> ele.getBlock(10000)
{
  difficulty: 2267769,
  extraData: "0xd7830104078447656c6587676f312e362e32856c696e7578",
  gasLimit: 4712388,
  gasUsed: 0,
  hash: "0x80230edb308173f95a27a05716bcfeb992754e933793bd2bb2400d7a455a409e",
  logsBloom: "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
  miner: "0x79f512cd283989eca606379471d76c93f1a66f56",
  nonce: "0x60fa5eefb321b553",
  number: 10000,
  parentHash: "0xb7f873fb8fbfb6a7409017b8df952024ff866cb5970ed04fed785dc96d2ba421",
  receiptRoot: "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
  sha3Uncles: "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
  size: 537,
  stateRoot: "0xfa61769bb4cdaafaa477a9f8a8e195997ec2a8abb73f2bc251c5725e6abc7115",
  timestamp: 1467745620,
  totalDifficulty: 11426704498,
  transactions: [],
  transactionsRoot: "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
  uncles: []
}
```
* [Back to Top](#gele-console-command-list-1) 

### ele.getBlockByNumber   
Returns the selected block information.
Enter the block number in hexadecimal.    

 Client          | Method        |
------------  | ------------- | 
RPC    |curl -X POST --data '{"jsonrpc":"2.0","method":"ele_getBlockByNumber","params":["0x1b4", true],"id":1}' 127.0.0.1:7075 | 
```
{"jsonrpc":"2.0","id":1,"result":{"difficulty":"0x277e1","extraData":"0xd7830104078447656c6587676f312e362e32856c696e7578","gasLimit":"0x47e7c4","gasUsed":"0x0","hash":"0x48021ab27e35c43be0a3d0c94fc902b6ef1603cb857d5b5ed494099fc6d5aee9","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000 0000000000000000000000000000000000000000000000000000000000000000000000000000000 0000000000000000000000000000000000000000000000000000000000000000000000000000000 0000000000000000000000000000000000000000000000000000000000000000000000000000000 0000000000000000000000000000000000000000000000000000000000000000000000000000000 0000000000000000000000000000000000000000000000000000000000000000000000000000000 0000000000000000000000000000000000000000","miner":"0x79f512cd283989eca606379471d76c93f1a66f56","nonce":"0x77097e973d6528c3","number":"0x1b4","parentHash":"0x89f79ba99e0d38c4dea366b5784fa39dfbe9f82aefa1c42c699732c1769906b7","receiptRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0x219","stateRoot":"0x1731e1dace82c79472a45a46e2bb6f2199fd33a5f221e2328cdfd33c3312bd9f","timestamp":"0x577aceb3","totalDifficulty":"0x3cc4861","transactions":[],"transactionsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","uncles":[]}}
```
* [Back to Top](#gele-console-command-list-1) 
 
### ele.getBlockTransactionCount
Returns number of the selected block transactions.

 Client          | Method        |
------------  | ------------- | 
Console    |ele.getBlockTransactionCount         |   
```bash
> ele.getBlockTransactionCount(13568)
1
```
* [Back to Top](#gele-console-command-list-1) 

### ele.getBlockUncleCount
Returns the number of uncles in a block from a block matching the given block.

 Client          | Method        |
------------  | ------------- | 
Console    |ele.getBlockUncleCount         |   
```bash
> ele.getBlockUncleCount(block number)
```
* [Back to Top](#gele-console-command-list-1) 

### ele.getCode
Returns code at a given address.    
**Parameters**        
- DATA, 20 Bytes - address
- QUANTITY|TAG - integer block number, or the string "latest", "earliest" or "pending"
```
params: [
   '0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b',
   '0x2'  // 2
]
```
 Client          | Method        |
------------  | ------------- | 
Console    |ele.getBlockUncleCount         |   
RPC       |curl -X POST --data '{"jsonrpc":"2.0","method":"ele_getCode","params":["0x5d15a90d015ead1dfbbca1d5b11201ffa7039d77", "0x2"],"id":1}' 127.0.0.1:7075 |
```bash
> ele.getCode("0x5d15a90d015ead1dfbbca1d5b11201ffa7039d77")
"0x"
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"ele_getCode","params":["0x5d15a90d015ead1dfbbca1d5b11201ffa7039d77", "0x2"],"id":1}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":1,"result":"0x"}
```
* [Back to Top](#gele-console-command-list-1) 


### ele.getCompilers   
Returns a list of available compilers in the client.     

 Client          | Method        |
------------  | ------------- | 
Console    |ele.getCompilers()         |   
RPC       |curl -X POST --data '{"jsonrpc":"2.0","method":"ele_getCompilers","params":[],"id":1}' 127.0.0.1:7075 |

* [Back to Top](#gele-console-command-list-1) 


### ele.getTransactionCount
Returns integer of the number of transactions send from this address.   
**Parameters**        
- DATA, 20 Bytes - address
- QUANTITY|TAG - integer block number, or the string "latest", "earliest" or "pending"
```
params: [
   '0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4',
   'latest' // state at the latest block
]
```
 Client          | Method        |
------------  | ------------- | 
Console    |ele.getTransactionCount         |   
RPC       |curl -X POST --data '{"jsonrpc":"2.0","method":"ele_getTransactionCount","params":["0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4","latest"],"id":1}' 127.0.0.1:7075 |
```bash
> ele.getTransactionCount("0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4")
7
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"ele_getTransactionCount","params":["0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4","latest"],"id":1}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":1,"result":"0x7"}
```
* [Back to Top](#gele-console-command-list-1) 




### ele.getTransactionFromBlock
Returns the block transaction information.

 Client          | Method        |
------------  | ------------- | 
Console    |ele.getTransactionFromBlock         |   
```bash
> ele.getTransactionFromBlock(13568)
{
  blockHash: "0x58ce450e4ac04427a28bc02ba27967ae05c77ff0372d19172f377a320ae847ad",
  blockNumber: 13568,
  from: "0xa748f720f5989f2b541fa9ef3c78236808123635",
  gas: 121000,
  gasPrice: 20000000000,
  hash: "0xb842cfd9860c41b955cf9357d5909a2b4e7b6de200903aeae7341e37b36b0d77",
  input: "0x",
  nonce: 0,
  to: "0x5d15a90d015ead1dfbbca1d5b11201ffa7039d77",
  transactionIndex: 0,
  value: 10000000000000000000
}
```
* [Back to Top](#gele-console-command-list-1) 



### ele.getTransactionReceipt
Returns the receipt of a transaction by transaction hash. (receipt is not available for pending transactions.)    
**Parameters**          
- DATA, 32 Bytes - hash of a transaction

```
params: [
   '0xb842cfd9860c41b955cf9357d5909a2b4e7b6de200903aeae7341e37b36b0d77',
]
```

 Client          | Method        |
------------  | ------------- | 
Console    |ele.getTransactionReceipt         |   
RPC | curl -X POST --data '{"jsonrpc":"2.0","method":"ele_getTransactionReceipt","params":["0xb842cfd9860c41b955cf9357d5909a2b4e7b6de200903aeae7341e37b36b0d77"],"id":1}' 127.0.0.1:7075 |


```bash
> ele.getTransactionReceipt("0xb842cfd9860c41b955cf9357d5909a2b4e7b6de200903aeae7341e37b36b0d77")
{
  blockHash: "0x58ce450e4ac04427a28bc02ba27967ae05c77ff0372d19172f377a320ae847ad",
  blockNumber: 13568,
  contractAddress: null,
  cumulativeGasUsed: 21000,
  from: "0xa748f720f5989f2b541fa9ef3c78236808123635",
  gasUsed: 21000,
  logs: [],
  root: "2cd17a1d375006c6bc03c3a36d1db319a5f630e8bbd16cac5024a81167e5e9dd",
  to: "0x5d15a90d015ead1dfbbca1d5b11201ffa7039d77",
  transactionHash: "0xb842cfd9860c41b955cf9357d5909a2b4e7b6de200903aeae7341e37b36b0d77",
  transactionIndex: 0
}
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"ele_getTransactionReceipt","params":["0xb842cfd9860c41b955cf9357d5909a2b4e7b6de200903aeae7341e37b36b0d77"],"id":1}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":1,"result":{"blockHash":"0x58ce450e4ac04427a28bc02ba27967ae05c77ff0372d19172f377a320ae847ad","blockNumber":"0x3500","contractAddress":null,"cumulativeGasUsed":"0x5208","from":"0xa748f720f5989f2b541fa9ef3c78236808123635","gasUsed":"0x5208","logs":[],"root":"2cd17a1d375006c6bc03c3a36d1db319a5f630e8bbd16cac5024a81167e5e9dd","to":"0x5d15a90d015ead1dfbbca1d5b11201ffa7039d77","transactionHash":"0xb842cfd9860c41b955cf9357d5909a2b4e7b6de200903aeae7341e37b36b0d77","transactionIndex":"0x0"}}
```
* [Back to Top](#gele-console-command-list-1) 
 

### ele.getWork
Returns the hash of the current block, the seedHash, and the boundary condition to be met ("target").    
**Returns**          
Array - Array with the following properties:

- DATA, 32 Bytes - current block header pow-hash
- DATA, 32 Bytes - the seed hash used for the DAG.
- DATA, 32 Bytes - the boundary condition ("target"), 2^256 / difficulty.

 Client          | Method        |
------------  | ------------- | 
Console    |ele.getTransactionReceipt         |   
RPC | curl -X POST --data '{"jsonrpc":"2.0","method":"ele_getWork","params":[],"id":73}' 127.0.0.1:7075 |


```bash
ele.getWork()
["0x69f0c90dde481b5b0dd99ae0013e59ae63b6f464483beb34e01739c2423fb396", "0x0000000000000000000000000000000000000000000000000000000000000000", "0x000006957372f46c29601697dcadfbc52bfc3d8822cb0c7c1f711927ab65ca56"]
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"ele_getWork","params":[],"id":73}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":73,"result":["0x414d545e4ca36dc3ef756f909e2ba47d921043da16ed62559e9abe0d9c4556c1","0x0000000000000000000000000000000000000000000000000000000000000000","0x000006a8b4308ddada8855e5ae5751d8557126060025df80d426ccbae7688a4e"]}
```
* [Back to Top](#gele-console-command-list-1) 




### ele.sendTransaction
Creates new message call transaction or a contract creation, if the data field contains code.
**This command must be preceded by `personal.unlock Account` before run it.     
*If the RPC of gele is open*, unlock account has some security problem. You can be Element transfer without Unlock account. `personal.signAndSendTransaction`**


**Parameters**  

- from: DATA, 20 Bytes - The address the transaction is send from.
- to: DATA, 20 Bytes - (optional when creating new contract) The address the transaction is directed to.
- gas: QUANTITY - (optional, default: 90000) Integer of the gas provided for the transaction execution. It will return unused gas.
- gasPrice: QUANTITY - (optional, default: To-Be-Determined) Integer of the gasPrice used for each paid gas
- value: QUANTITY - (optional) Integer of the value send with this transaction
- data: DATA - The compiled code of a contract OR the hash of the invoked method signature and encoded parameters.
- nonce: QUANTITY - (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce.

```
params: [{
  "from": "0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4",
  "to": "0x5d15a90d015ead1dfbbca1d5b11201ffa7039d77"",
  "gas": "0x76c0", // 30400,
  "gasPrice": "0x9184e72a000", // 10000000000000
  "value": "0x9184e72a", // 2441406250
  "data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"
}]
```

 Client          | Method        |
------------  | ------------- | 
Console    |ele.sendTransaction         |   
RPC | curl -X POST --data '{"jsonrpc":"2.0","method":"ele_sendTransaction","params":[{see above}],"id":1}' 127.0.0.1:7075 |
```bash
> personal.unlockAccount(ele.coinbase)
Unlock account 0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4
Passphrase: 
true
> ele.sendTransaction({from:ele.accounts[0], to:"0x79f512cd283989eca606379471d76c93f1a66f56", value:web3.toMey(10.30667, "element"), Gas:21000}) // 10.30667 element transfer
"0xf51e58422fbd887cfdbf9157ad3dcfc6b17130ef217607d41394aaefafd54bb6"
> personal.lockAccount(ele.coinbase)
true
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"ele_sendTransaction","params":[{see above}],"id":1}'
{"id":1,"jsonrpc": "2.0", "result": "0xf51e58422fbd887cfdbf9157ad3dcfc6b17130ef217607d41394aaefafd54bb6"}
```
* [Back to Top](#gele-console-command-list-1) 



### ele.sign
Signs data with a given address. the address to sign must be unlocked.
**Parameters**  
- DATA, 20 Bytes - address
- DATA, 32 Bytes - sha3 hash of data to sign.

 Client          | Method        |
------------  | ------------- | 
Console    |ele.sign        |   
RPC | curl -X POST --data '{"jsonrpc":"2.0","method":"ele_sign","params":["<20 Bytes - address>", "<32 Bytes - sha3 hash of data to sign.>"],"id":1}' 127.0.0.1:7075 |

* [Back to Top](#gele-console-command-list-1) 

### ele.submitWork
Used for submitting a proof-of-work solution.
**Parameters**  
- DATA, 8 Bytes - The nonce found (64 bits)
- DATA, 32 Bytes - The header's pow-hash (256 bits)
- DATA, 32 Bytes - The mix digest (256 bits)

**Returns**   
Boolean - returns true if the provided solution is valid, otherwise false.

 Client          | Method        |
------------  | ------------- | 
Console    |ele.submitWork        |   
RPC | curl -X POST --data '{"jsonrpc":"2.0", "method":"ele_submitWork", "params":["8 Bytes - The nonce found (64 bits)", "DATA, 32 Bytes - The header's pow-hash (256 bits)", "DATA, 32 Bytes - The mix digest (256 bits)"],"id":73}' 127.0.0.1:7075 |

* [Back to Top](#gele-console-command-list-1) 

### ele.isSyncing
**Returns**

Object - a isSyncing object with the following methods:

- syncing.addCallback(): Adds another callback, which will be called when the node starts or stops syncing.
- syncing.stopWatching(): Stops the syncing callbacks.

**Callback return value**

- Boolean - The callback will be fired with true when the syncing starts and with false when it stopped.
- Object - While syncing it will return the syncing object:   
    startingBlock: Number - The block number where the sync started.    
    currentBlock: Number - The block number where at which block the node currently synced to already.    
    highestBlock: Number - The estimated block number to sync to.   

 Client          | Method        |
------------  | ------------- | 
Console    |ele.isSyncing        |   
```
> ele.isSyncing()
{
  callbacks: [],
  lastSyncState: false,
  pollId: "syncPoll_1",
  requestManager: {
    polls: {
      syncPoll_1: {
        data: {...},
        id: "syncPoll_1",
        callback: function(error, sync),
        uninstall: function()
      }
    },
    provider: {
      newAccount: function(),
      send: function(),
      sendAsync: function(),
      unlockAccount: function()
    },
    timeout: {},
    poll: function(),
    reset: function(keepIsSyncing),
    send: function(data),
    sendAsync: function(data, callback),
    sendBatch: function(data, callback),
    setProvider: function(p),
    startPolling: function(data, pollId, callback, uninstall),
    stopPolling: function(pollId)
  },
  addCallback: function(callback),
  stopWatching: function()
}
```

* [Back to Top](#gele-console-command-list-1) 

## net

### net.listening
Returns true if client is actively listening for network connections.

 Client          | Method        |
------------  | ------------- | 
Console    |net.listening       |   
RPC | curl -X POST --data '{"jsonrpc":"2.0","method":"net_listening","params":[],"id":67}' 127.0.0.1:7075 |
```
> net.listening
true
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"net_listening","params":[],"id":67}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":67,"result":true}
```
* [Back to Top](#gele-console-command-list-2) 

### net.peerCount
Returns number of peers currenly connected to the client.

 Client          | Method        |
------------  | ------------- | 
Console    |net.peerCount      |   
RPC | curl -X POST --data '{"jsonrpc":"2.0","method":"net_peerCount","params":[],"id":74}' 127.0.0.1:7075 |
```
> net.peerCount
5152345186
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"net_peerCount","params":[],"id":74}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":74,"result":"0x1331a8c62"}
```
* [Back to Top](#gele-console-command-list-2) 



### net.version
Returns the current networkID (Officially 73733).

 Client          | Method        |
------------  | ------------- | 
Console    |net.version     |   
RPC | curl -X POST --data '{"jsonrpc":"2.0","method":"net_version","params":[],"id":67}' 127.0.0.1:7075 |
```
> net.version
"73733"
//
$ curl -X POST --data '{"jsonrpc":"2.0","method":"net_version","params":[],"id":67}' 127.0.0.1:7075
{"jsonrpc":"2.0","id":67,"result":"73733"}
```
* [Back to Top](#gele-console-command-list-2) 



## personal

### personal.listAccounts
List all accounts(`ele.accounts` functionally identical)

 Client          | Method        |
------------  | ------------- | 
Console    |personal.listAccounts     |   

```
personal.listAccounts
["0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4"]
```
* [Back to Top](#gele-console-command-list-2) 


### personal.importRawKey
Imports the given unencrypted private key (hex string) into the key store, encrypting it with the passphrase.
Returns the address of the new account.

 Client          | Method        |
------------  | ------------- | 
Console    |personal.importRawKey(keydata, passphrase)     |   
* [Back to Top](#gele-console-command-list-2) 

### personal.lockAccount
Lock Account. The account can no longer be used to send transactions.

 Client          | Method        |
------------  | ------------- | 
Console    |personal.lockAccount("address")     | 

```
> personal.lockAccount(ele.coinbase)
true
> personal.lockAccount(ele.accounts[0])
true
> personal.lockAccount("0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4")
true
```
* [Back to Top](#gele-console-command-list-2) 


### personal.newAccount
Generates a new private key and stores it in the key store directory. The key file is encrypted with the given passphrase. Returns the address of the new account.
At the gele console, newAccount will prompt for a passphrase when it is not supplied as the argument.

 Client          | Method        |
------------  | ------------- | 
Console    |personal.newAccount()     | 
```
> personal.newAccount()
Passphrase: 
Repeat passphrase: 
"0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4"
```
The passphrase can also be supplied as a string. **However, the password will be exposed to history log. It is very dangerous. I do not recommend.**
```
> personal.newAccount("xaESCWOMAW498")
"0x1234567899876543211234569874563214569874"
```

* [Back to Top](#gele-console-command-list-2) 


### personal.signAndSendTransaction
Validate the given password and submit transaction.   
The transaction is the same argument as for `ele.sendTransaction` and contains the from address. If the passphrase can be used to decrypt the private key belogging to tx.from the transaction is verified, signed and send onto the network.     
**The account is not unlocked globally in the node and cannot be used in other RPC calls. command is not writed on history log.**

 Client          | Method        |
------------  | ------------- | 
Console    |personal.signAndSendTransaction(tx, password)     | 
```
>var tx = {from:ele.accounts[0], to: "<Element address of recipient>", value: web3.toMey(<Element amount>, "element"), Gas:21000}   
undefined       
>personal.signAndSendTransaction(tx, "<password>")      
"0x1234567890987654321234567890987654321"
//
personal.signAndSendTransaction({from:"<Element address of sender>", to: "<Element address of recipient>", value: web3.toMey(<Element amount>, "element"), Gas:21000}, "<password>")
"0x1234567890987654321234567890987654321"
```

```
> personal.signAndSendTransaction({from:ele.coinbase, to: "0x5d15a90d015ead1dfbbca1d5b11201ffa7039d77", value: web3.toMey(11.111, "element"), Gas:21000}, "<password>")
"0xa8f3dfc9e23cf22a6390dcfd22d9da83ed0ef33da54e06e5bb0dfdd728bc41b4"
```
* [Back to Top](#gele-console-command-list-2) 


### personal.unlockAccount
Decrypts the key with the given address from the key store.   
Both passphrase and unlock duration are optional when using the JavaScript console. If the passphrase is not supplied as an argument, the console will prompt for the passphrase interactively. 
The unencrypted key will be held in memory until the unlock duration expires. If the unlock duration defaults to 300 seconds. An explicit duration of zero seconds unlocks the key until gele exits.   
The account can be used with `ele_sign` and `ele_sendTransaction` while it is unlocked.   


 Client          | Method        |
------------  | ------------- | 
Console    |personal.unlockAccount(address, passphrase, duration)     | 
```
> personal.unlockAccount("0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4")
Unlock account 0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4
Passphrase: 
true

// Supplying the passphrase and unlock duration as arguments:(!!!!! The password will be exposed to history log)
> personal.unlockAccount("0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4", "foo", 30)
true

//If you want to type in the passphrase and stil override the default unlock duration, pass null as the passphrase.
> personal.unlockAccount("0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4", null, 30)
Unlock account 0x71842e5c0b8db31dcf4be4cc87b820de0c1e64e4
Passphrase: 
true
```

* [Back to Top](#gele-console-command-list-2) 


## txpool   
The txpool API gives you access to several non-standard RPC methods to inspect the contents of the transaction pool containing all the currently pending transactions as well as the ones queued for future processing.
* [Back to Top](#gele-console-command-list-2) 

### txpool.content
The content inspection property can be queried to list the exact details of all the transactions currently pending for inclusion in the next block(s), as well as the ones that are being scheduled for future execution only.

The result is an object with two fields pending and queued. Each of these fields are associative arrays, in which each entry maps an origin-address to a batch of scheduled transactions. These batches themselves are maps associating nonces with actual transactions.

Please note, there may be multiple transactions associated with the same account and nonce. This can happen if the user broadcast mutliple ones with varying gas allowances (or even complerely different transactions).

 Client          | Method        |
------------  | ------------- | 
Console    |txpool.content     | 
```
> txpool.content
{
  pending: {
    0x0216d5032f356960cd3749c31ab34eeff21b3395: {
      806: [{
        blockHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
        blockNumber: null,
        from: "0x0216d5032f356960cd3749c31ab34eeff21b3395",
        gas: "0x5208",
        gasPrice: "0xba43b7400",
        hash: "0xaf953a2d01f55cfe080c0c94150a60105e8ac3d51153058a1f03dd239dd08586",
        input: "0x",
        nonce: "0x326",
        to: "0x7f69a91a3cf4be60020fb58b893b7cbb65376db8",
        transactionIndex: null,
        value: "0x19a99f0cf456000"
      }]
    },
    0x24d407e5a0b506e1cb2fae163100b5de01f5193c: {
      34: [{
        blockHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
        blockNumber: null,
        from: "0x24d407e5a0b506e1cb2fae163100b5de01f5193c",
        gas: "0x44c72",
        gasPrice: "0x4a817c800",
        hash: "0xb5b8b853af32226755a65ba0602f7ed0e8be2211516153b75e9ed640a7d359fe",
        input: "0xb61d27f600000000000000000000000024d407e5a0b506e1cb2fae163100b5de01f5193c00000000000000000000000000000000000000000000000053444835ec580000000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
        nonce: "0x22",
        to: "0x7320785200f74861b69c49e4ab32399a71b34f1a",
        transactionIndex: null,
        value: "0x0"
      }]
    }
  },
  queued: {
    0x976a3fc5d6f7d259ebfb4cc2ae75115475e9867c: {
      3: [{
        blockHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
        blockNumber: null,
        from: "0x976a3fc5d6f7d259ebfb4cc2ae75115475e9867c",
        gas: "0x15f90",
        gasPrice: "0x4a817c800",
        hash: "0x57b30c59fc39a50e1cba90e3099286dfa5aaf60294a629240b5bbec6e2e66576",
        input: "0x",
        nonce: "0x3",
        to: "0x346fb27de7e7370008f5da379f74dd49f5f2f80f",
        transactionIndex: null,
        value: "0x1f161421c8e0000"
      }]
    },
    0x9b11bf0459b0c4b2f87f8cebca4cfc26f294b63a: {
      2: [{
        blockHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
        blockNumber: null,
        from: "0x9b11bf0459b0c4b2f87f8cebca4cfc26f294b63a",
        gas: "0x15f90",
        gasPrice: "0xba43b7400",
        hash: "0x3a3c0698552eec2455ed3190eac3996feccc806970a4a056106deaf6ceb1e5e3",
        input: "0x",
        nonce: "0x2",
        to: "0x24a461f25ee6a318bdef7f33de634a67bb67ac9d",
        transactionIndex: null,
        value: "0xebec21ee1da40000"
      }],
      6: [{
        blockHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
        blockNumber: null,
        from: "0x9b11bf0459b0c4b2f87f8cebca4cfc26f294b63a",
        gas: "0x15f90",
        gasPrice: "0x4a817c800",
        hash: "0xbbcd1e45eae3b859203a04be7d6e1d7b03b222ec1d66dfcc8011dd39794b147e",
        input: "0x",
        nonce: "0x6",
        to: "0x6368f3f8c2b42435d6c136757382e4a59436a681",
        transactionIndex: null,
        value: "0xf9a951af55470000"
      }, {
        blockHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
        blockNumber: null,
        from: "0x9b11bf0459b0c4b2f87f8cebca4cfc26f294b63a",
        gas: "0x15f90",
        gasPrice: "0x4a817c800",
        hash: "0x60803251d43f072904dc3a2d6a084701cd35b4985790baaf8a8f76696041b272",
        input: "0x",
        nonce: "0x6",
        to: "0x8db7b4e0ecb095fbd01dffa62010801296a9ac78",
        transactionIndex: null,
        value: "0xebe866f5f0a06000"
      }],
    }
  }
}
```
* [Back to Top](#gele-console-command-list-2) 

### txpool.inspect
The inspect inspection property can be queried to list a textual summary of all the transactions currently pending for inclusion in the next block(s), as well as the ones that are being scheduled for future execution only. This is a method specifically tailored to developers to quickly see the transactions in the pool and find any potential issues.

The result is an object with two fields pending and queued. Each of these fields are associative arrays, in which each entry maps an origin-address to a batch of scheduled transactions. These batches themselves are maps associating nonces with transactions summary strings.

Please note, there may be multiple transactions associated with the same account and nonce. This can happen if the user broadcast mutliple ones with varying gas allowances (or even complerely different transactions).

 Client          | Method        |
------------  | ------------- | 
Console    |txpool.inspect     | 

```
> txpool.inspect
{
  pending: {
    0x26588a9301b0428d95e6fc3a5024fce8bec12d51: {
      31813: ["0x3375ee30428b2a71c428afa5e89e427905f95f7e: 0 mey + 500000 × 20000000000 gas"]
    },
    0x2a65aca4d5fc5b5c859090a6c34d164135398226: {
      563662: ["0x958c1fa64b34db746925c6f8a3dd81128e40355e: 1051546810000000000 mey + 90000 × 20000000000 gas"],
      563663: ["0x77517b1491a0299a44d668473411676f94e97e34: 1051190740000000000 mey + 90000 × 20000000000 gas"],
      563664: ["0x3e2a7fe169c8f8eee251bb00d9fb6d304ce07d3a: 1050828950000000000 mey + 90000 × 20000000000 gas"],
      563665: ["0xaf6c4695da477f8c663ea2d8b768ad82cb6a8522: 1050544770000000000 mey + 90000 × 20000000000 gas"],
      563666: ["0x139b148094c50f4d20b01caf21b85edb711574db: 1048598530000000000 mey + 90000 × 20000000000 gas"],
      563667: ["0x48b3bd66770b0d1eecefce090dafee36257538ae: 1048367260000000000 mey + 90000 × 20000000000 gas"],
      563668: ["0x468569500925d53e06dd0993014ad166fd7dd381: 1048126690000000000 mey + 90000 × 20000000000 gas"],
      563669: ["0x3dcb4c90477a4b8ff7190b79b524773cbe3be661: 1047965690000000000 mey + 90000 × 20000000000 gas"],
      563670: ["0x6dfef5bc94b031407ffe71ae8076ca0fbf190963: 1047859050000000000 mey + 90000 × 20000000000 gas"]
    },
    0x9174e688d7de157c5c0583df424eaab2676ac162: {
      3: ["0xbb9bc244d798123fde783fcc1c72d3bb8c189413: 30000000000000000000 mey + 85000 × 21000000000 gas"]
    },
    0xb18f9d01323e150096650ab989cfecd39d757aec: {
      777: ["0xcd79c72690750f079ae6ab6ccd7e7aedc03c7720: 0 mey + 1000000 × 20000000000 gas"]
    },
    0xb2916c870cf66967b6510b76c07e9d13a5d23514: {
      2: ["0x576f25199d60982a8f31a8dff4da8acb982e6aba: 26000000000000000000 mey + 90000 × 20000000000 gas"]
    },
    0xbc0ca4f217e052753614d6b019948824d0d8688b: {
      0: ["0x2910543af39aba0cd09dbb2d50200b3e800a63d2: 1000000000000000000 mey + 50000 × 1171602790622 gas"]
    },
    0xea674fdde714fd979de3edf0f56aa9716b898ec8: {
      70148: ["0xe39c55ead9f997f7fa20ebe40fb4649943d7db66: 1000767667434026200 mey + 90000 × 20000000000 gas"]
    }
  },
  queued: {
    0x0f6000de1578619320aba5e392706b131fb1de6f: {
      6: ["0x8383534d0bcd0186d326c993031311c0ac0d9b2d: 9000000000000000000 mey + 21000 × 20000000000 gas"]
    },
    0x5b30608c678e1ac464a8994c3b33e5cdf3497112: {
      6: ["0x9773547e27f8303c87089dc42d9288aa2b9d8f06: 50000000000000000000 mey + 90000 × 50000000000 gas"]
    },
    0x976a3fc5d6f7d259ebfb4cc2ae75115475e9867c: {
      3: ["0x346fb27de7e7370008f5da379f74dd49f5f2f80f: 140000000000000000 mey + 90000 × 20000000000 gas"]
    },
    0x9b11bf0459b0c4b2f87f8cebca4cfc26f294b63a: {
      2: ["0x24a461f25ee6a318bdef7f33de634a67bb67ac9d: 17000000000000000000 mey + 90000 × 50000000000 gas"],
      6: ["0x6368f3f8c2b42435d6c136757382e4a59436a681: 17990000000000000000 mey + 90000 × 20000000000 gas", "0x8db7b4e0ecb095fbd01dffa62010801296a9ac78: 16998950000000000000 mey + 90000 × 20000000000 gas"],
      7: ["0x6368f3f8c2b42435d6c136757382e4a59436a681: 17900000000000000000 mey + 90000 × 20000000000 gas"]
    }
  }
}
```
* [Back to Top](#gele-console-command-list-2) 

### txpool.status

The status inspection property can be queried for the number of transactions currently pending for inclusion in the next block(s), as well as the ones that are being scheduled for future execution only.

The result is an object with two fields pending and queued, each of which is a counter representing the number of transactions in that particular state.


 Client          | Method        |
------------  | ------------- | 
Console    |txpool.status     | 

```
> txpool.status
{
  pending: 10,
  queued: 7
}
```
* [Back to Top](#gele-console-command-list-2) 


## web3.version

```
> web3.version
{
  api: "0.15.3",                                     //The elementrem js api version.
  elementrem: "0x3f",                                //The elementrem protocol version.
  network: "73733",                                  //The network protocol version.
  node: "Gele/v1.4.7-stable-64a5a363/linux/go1.6.2", //The client/node version.
  whisper: undefined,                                //The whisper protocol version.
  getElementrem: function(callback),
  getNetwork: function(callback),
  getNode: function(callback),
  getWhisper: function(callback)
}
```
* [Back to Top](#gele-console-command-list-2) 

## web3.db

```
> web3.db
{
  getHex: function(),    //This method should be called, when we want to get a binary data in HEX form from the local leveldb database.
  getString: function(), //This method should be called, when we want to get string from the local leveldb database.
  putHex: function(),    //This method should be called, when we want to store binary data in HEX form in the local leveldb database.
  putString: function()  //This method should be called, when we want to store a string in the local leveldb database.
}
```
* [Back to Top](#gele-console-command-list-2) 

## web3.currentProvider
Will contain the current provider, if one is set. This can be used to check if mist etc. set already a provider.

```
> web3.currentProvider
{
  newAccount: function(),
  send: function(),
  sendAsync: function(),
  unlockAccount: function()
}
```
* [Back to Top](#gele-console-command-list-2)


## web3.providers

```
> web3.providers
{
  HttpProvider: function(host),
  IpcProvider: function(path, net)
}
```
```
var Web3 = require('web3');
// create an instance of web3 using the HTTP provider.
// NOTE in mist web3 is already available, so check first if its available before instantiating
var web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:7075"));
```
* [Back to Top](#gele-console-command-list-2)


## web3.shh

```
> web3.shh
{
  addToGroup: function(),
  filter: function(fil, callback),
  hasIdentity: function(),          
  newGroup: function(),
  newIdentity: function(),          
  post: function()                  
}
```

* [Back to Top](#gele-console-command-list-2)

## Batch requests
Batch requests allow queuing up requests and processing them at once.

```js
var batch = web3.createBatch();
batch.add(web3.ele.getBalance.request('0x0000000000000000000000000000000000000000', 'latest', callback));
batch.add(web3.ele.contract(abi).at(address).balance.request(address, callback2));
batch.execute();
```
* [Back to Top](#gele-console-command-list-2)

## web3.toMey
## web3.fromMey

elementrem has the following units.
```
    'mey':      '1',
    'kmey':     '1000',
    'ada':      '1000',
    'mmey':     '1000000',
    'babbage':  '1000000',
    'gmey':     '1000000000',
    'shannon':  '1000000000',
    'szabo':    '1000000000000',
    'finney':   '1000000000000000',
    'element':    '1000000000000000000',
    'kelement':   '1000000000000000000000',
    'grand':    '1000000000000000000000',
    'einstein': '1000000000000000000000',
    'melement':   '1000000000000000000000000',
    'gelement':   '1000000000000000000000000000',
    'telement':   '1000000000000000000000000000000'
```

All units can be converted to each other.
**example**
```
> var value1 = web3.fromMey('214356897445000', 'szabo');
undefined
> var value2 = web3.toMey('8', 'element');
undefined
> console.log(value1);
214.356897445
undefined
> console.log(value2);
8000000000000000000
undefined
```
* [Back to Top](#gele-console-command-list-2)
