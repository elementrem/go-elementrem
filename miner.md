## miner

### CPU Mining with Gele

using the console.    
`> miner.start()`   
We'll have to wait a little bit while your node generates its Directed Acyclic Graph (DAG). This process is what helps the Elementrem network be resistant to ASIC mining
The DAG should be stored in a 1GB dump (for the initial epoch, anyway), in a directory:   
Mac/Linux: `$(HOME)/.elhash/ `  
Windows: `c:/Users/(User Account)/AppData/Elhash/`    
`> miner.stop()`    
    
### [GPU mining click Here](https://github.com/elementrem/webthree-umbrella/blob/master/GPU_mining_command.md)   
    
### miner.start   
`miner.start(threadCount)`        
Starts mining on with the given threadNumber of parallel threads. This is an optional argument.       
**Returns**   
`true` on success, otherwise `false.`   
**Example**   
```
>miner.start()
true
```

### miner.stop
`miner.stop(threadCount)`   
Stops threadCount miners. This is an optional argument.
**Returns**   
`true` on success, otherwise `false.`   
**Example**   
```
>miner.stop()
true
```

### miner.startAutoDAG
`miner.startAutoDAG()`   
Starts automatic pregeneration of the `elhash DAG.` This process make sure that the DAG for the subsequent epoch is available allowing mining right after the new epoch starts. 
If this is used by most network nodes, then blocktimes are expected to be normal at epoch transition. 
Auto DAG is switched on automatically when mining is started and switched off when the miner stops.   
**Returns**   
`true` on success, otherwise `false.`   

### miner.stopAutoDAG
`miner.stopAutoDAG()`   
Stops automatic pregeneration of the elhash DAG. Auto DAG is switched off automatically when mining is stops.   
**Returns**       
`true` on success, otherwise `false.`  

### miner.makeDAG   
`miner.makeDAG(blockNumber, dir)`       
Generates the DAG for epoch blockNumber/epochLength. dir specifies a target directory, If dir is the empty string, then ethash will use the default directories ~/.ethash on Linux and MacOS, and ~\AppData\Ethash on Windows. The DAG file's name is full-<revision-number>R-<seedhash>    
**Returns**       
`true` on success, otherwise `false.` 

### miner.makeDAG
`miner.setExtra("extra data")`    
Sets the extra data for the block when finding a block. Limited to 32 bytes.

### miner.setGasPrice
`miner.setGasPrice(gasPrice)`   
Sets the the gasprice for the miner

### miner.setEtherbase
`miner.setEtherbase(account)`   
Sets the the gasprice for the miner

* [Back to Page](gele_command_readme.md)
