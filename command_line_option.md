```bash
$ ./gele --help
NAME:
   gele-linux-amd64 - the go-elementrem command line interface

USAGE:
   gele-linux-amd64 [options] command [command options] [arguments...]
   
VERSION:
   1.4.12-stable-927f38df
   
COMMANDS:
   import	import a blockchain file
   export	export blockchain into file
   upgradedb	upgrade chainblock database
   removedb	Remove blockchain and state databases
   dump		dump a specific block from storage
   monitor	Gele Monitor: node metrics monitoring and visualization
   account	manage accounts
   wallet	elementrem presale wallet
   console	Gele Console: interactive JavaScript environment
   attach	Gele Console: interactive JavaScript environment (connect to node)
   js		executes the given JavaScript files in the Gele JavaScript VM
   makedag	generate elhash dag (for testing)
   gpuinfo	gpuinfo
   gpubench	benchmark GPU
   version	print elementrem version numbers
   init		bootstraps and initialises a new genesis block (JSON)
   help, h	Shows a list of commands or help for one command
   
ELEMENTREM OPTIONS:
  --datadir "/home/chain/.elementrem"	Data directory for the databases and keystore
  --keystore 				Directory for the keystore (default = inside the datadir)
  --networkid value			Network identifier (default: 73733)
  --testnet1				Test1 network is deprecated
  --testnet2				Test2 network is deprecated
  --dev					Developer mode: pre-configured private network with several debugging flags
  --identity value			Custom node name
  --fast				Enable fast syncing through state downloads
  --lightkdf				Reduce key-derivation RAM & CPU usage at some expense of KDF strength
  --cache value				Megabytes of memory allocated to internal caching (min 16MB / database forced) (default: 128)
  --blockchainversion value		Blockchain version (integer) (default: 3)
  
ACCOUNT OPTIONS:
  --unlock value	Comma separated list of accounts to unlock
  --password value	Password file to use for non-inteactive password input
  
API AND CONSOLE OPTIONS:
  --rpc			Enable the HTTP-RPC server
  --rpcaddr value	HTTP-RPC server listening interface (default: "localhost")
  --rpcport value	HTTP-RPC server listening port (default: 7075)
  --rpcapi value	API's offered over the HTTP-RPC interface (default: "ele,net,web3")
  --ws			Enable the WS-RPC server
  --wsaddr value	WS-RPC server listening interface (default: "localhost")
  --wsport value	WS-RPC server listening port (default: 7076)
  --wsapi value		API's offered over the WS-RPC interface (default: "ele,net,web3")
  --wsorigins value	Origins from which to accept websockets requests
  --ipcdisable		Disable the IPC-RPC server
  --ipcapi value	API's offered over the IPC-RPC interface (default: "admin,debug,ele,miner,net,personal,shh,txpool,web3")
  --ipcpath "gele.ipc"	Filename for IPC socket/pipe within the datadir (explicit paths escape it)
  --rpccorsdomain value	Comma separated list of domains from which to accept cross origin requests (browser enforced)
  --jspath loadScript	JavaScript root path for loadScript and document root for `admin.httpGet` (default: ".")
  --exec value		Execute JavaScript statement (only in combination with console/attach)
  --preload value	Comma separated list of JavaScript files to preload into the console
  
NETWORKING OPTIONS:
  --bootnodes value	Comma separated enode URLs for P2P discovery bootstrap
  --port value		Network listening port (default: 30707)
  --maxpeers value	Maximum number of network peers (network disabled if set to 0) (default: 25)
  --maxpendpeers value	Maximum number of pending connection attempts (defaults used if set to 0) (default: 0)
  --nat value		NAT port mapping mechanism (any|none|upnp|pmp|extip:<IP>) (default: "any")
  --nodiscover		Disables the peer discovery mechanism (manual peer addition)
  --nodekey value	P2P node key file
  --nodekeyhex value	P2P node key as hex (for testing)
  
MINER OPTIONS:
  --mine			Enable mining
  --minerthreads value		Number of CPU threads to use for mining (default: 2)
  --minergpus value		List of GPUs to use for mining (e.g. '0,1' will use the first two GPUs found)
  --autodag			Enable automatic DAG pregeneration
  --elementbase value		Public address for block mining rewards (default = first account created) (default: "0")
  --targetgaslimit value	Target gas limit sets the artificial target gas floor for the blocks to mine (default: "4712388")
  --gasprice value		Minimal gas price to accept for mining a transactions (default: "20000000000")
  --extradata value		Block extra data set by the miner (default = client version)
  
GAS PRICE ORACLE OPTIONS:
  --gpomin value	Minimum suggested gas price (default: "20000000000")
  --gpomax value	Maximum suggested gas price (default: "500000000000")
  --gpofull value	Full block threshold for gas price calculation (%) (default: 80)
  --gpobasedown value	Suggested gas price base step down ratio (1/1000) (default: 10)
  --gpobaseup value	Suggested gas price base step up ratio (1/1000) (default: 100)
  --gpobasecf value	Suggested gas price base correction factor (%) (default: 110)
  
VIRTUAL MACHINE OPTIONS:
  --jitvm		Enable the JIT VM
  --forcejit		Force the JIT VM to take precedence
  --jitcache value	Amount of cached JIT VM programs (default: 64)
  
LOGGING AND DEBUGGING OPTIONS:
  --metrics			Enable metrics collection and reporting
  --fakepow			Disables proof-of-work verification
  --verbosity value		Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=core, 5=debug, 6=detail (default: 3)
  --vmodule value		Per-module verbosity: comma-separated list of <pattern>=<level> (e.g. ele/*=6,p2p=5)
  --backtrace value		Request a stack trace at a specific logging statement (e.g. "block.go:271") (default: :0)
  --pprof			Enable the pprof HTTP server
  --pprofport value		pprof HTTP server listening port (default: 6060)
  --memprofilerate value	Turn on memory profiling with the given rate (default: 524288)
  --blockprofilerate value	Turn on block profiling with the given rate (default: 0)
  --cpuprofile value		Write CPU profile to the given file
  --trace value			Write execution trace to the given file
  
EXPERIMENTAL OPTIONS:
  --shh		Enable Whisper
  --natspec	Enable NatSpec confirmation notice
  
MISCELLANEOUS OPTIONS:
  --solc value	Solidity compiler command to be used (default: "solc")
  --help, -h	show help

```