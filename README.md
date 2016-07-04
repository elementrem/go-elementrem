# Elementrem Smart Contract BlockChain
----------------------
**Without source build, All versions of Elementrem(Gele) is available for download. **
###https://github.com/elementrem/go-elementrem/releases###

- Build from Source		
- Console line command		
- CPU Mining with Gele
- Elementrem monetary unit

<br>
## Build from Source
**You must have Go-language installed in your system to Build from source code.**
- Installing go language(Ubuntu)   
`curl -O https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz`  
`sudo tar -C /usr/local -xzf go1.4.2.linux-amd64.tar.gz`  
`mkdir -p ~/go; echo "export GOPATH=$HOME/go" >> ~/.bashrc`   
`echo "export PATH=$PATH:$HOME/go/bin:/usr/local/go/bin" >> ~/.bashrc`  
`source ~/.bashrc`  

- Installing go language(Windows)   
https://golang.org/doc/install    

- Installing go language(Mac OSX)   
`/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`  (Install Homebrew)    
`brew install gmp go`

**You must have Docker installed in your system to makes gele of another os.   
(For instance, make the windows Gele in ubuntu.)**   
`sudo curl -fsSL https://get.docker.com/ | sh`    
`sudo usermod -aG docker $(whoami)` (might require reboot)  
`docker pull ubuntu:16.04` (16.04 ubuntu version. if 15.04 `docker pull ubuntu:15.04`)    


**Source bulid**    
`git clone https://github.com/elementrem/go-elementrem`   
`cd go-elementrem`    
`make gele` or, to build the full OS suite of utilities: `make all`
		
		
**Run Gele**		
- Windows		<br>
1.Run `gele.exe`
	(If you see a message "Synchronisation failed", exit and re-Run)		<br>
2.Automatically connects to the Elementrem node : Default directory will be created in the C:\Users\<User Account>\AppData\Roaming\Elementrem		<br>
3.Run `gele attach` from command prompt window(cmd.exe) will be automatically connected to the Elementrem console.		<br>
		<br>
- Linux		<br>
1.Run `gele` from Terminal - If do not copy gele to /usr/bin it must be run ./gele
	(If you see a message "Synchronisation failed", exit and re-Run)		<br>
2.Automatically connects to the Elementrem node : Default directory will be created in the /home/<User Account>/.elementrem		<br>
3.Run `gele attach` from Terminal will be automatically connected to the Elementrem console.		<br>
		<br>
- Mac OSX		<br>
1.Run `gele` from Terminal - If do not copy gele to /usr/bin it must be run ./gele
	(If you see a message "Synchronisation failed", exit and re-Run)				<br>
2.Automatically connects to the Elementrem node : Default directory will be created in the /Library/Elementrem		<br>
3.Run `gele attach` from Terminal(launchpad -> Terminal) will be automatically connected to the Elementrem console.		<br>
		<br>
*Elementrem Network listening port = 30707		<br>

------------------------
##Console line command##
- When you run the following command will output information about that command:    
`web3`		
`personal`    
`ele`   
`admin`   

- some examples of commands : 		  

Command | Explanation
------------ | -------------
`personal.newAccount()` | Create a new address. Do not ever forget your password.
`ele.accounts`| Accounts list
`web3.fromMey(ele.getBalance(ele.accounts[0]), "element")` | First account balance check. `accounts[1]` Secend.
`personal.unlockAccount(ele.accounts[0])` | Unlock Account    
`ele.sendTransaction({from:ele.accounts[0], to:"<Element address of recipient>", value:web3.toMey(<Element amount>, "element"), Gas:21000})` | Element transfer   
`personal.lockAccount(ele.accounts[0])` | Lock Account    
view your transactions:`ele.getTransaction("<Copy and paste transaction>")` or visit www.elementrem.net   
####***However, unlock account has some security problem. You can be Element transfer without Unlock account.***

Command | Explanation
------------ | -------------
`personal.signAndSendTransaction({from:ele.accounts[0], to: "<Element address of recipient>", value: web3.toMey(<Element amount>, "element"), Gas:21000}, "passphrase")` | Element transfer. *don't worrt. command is not writed on history log.*

or also specify a variable.

	>var tx = {from:ele.accounts[0], to: "<Element address of recipient>", value: web3.toMey(<Element amount>, "element"), Gas:21000}			
	undefined		
	>personal.signAndSendTransaction(tx, "passphrase")		
	"0x1234567890987654321234567890987654321"		
---------------------------------
##CPU Mining with Gele##
using the console.    
`> miner.start()`     
We'll have to wait a little bit while your node generates its Directed Acyclic Graph (DAG). This process is what helps the Elementrem network be resistant to ASIC mining   
The DAG should be stored in a 1GB dump (for the initial epoch, anyway), in a directory:		
**Mac/Linux: $(HOME)/`.elhash`/**		
**Windows: c:/Users/(User Account)/AppData/`Elhash`/**			
`> miner.stop()`		

GPU mining is not currently available. When enough peer to participate, it will update the GPU mining.

--------------------------------
## Elementrem monetary unit

    'noelement':      '0',
    'mey':          '1', (smallest unit)
    'kmey':         '1000',
    'Kmey':         '1000',
    'babbage':      '1000',
    'femtoelement':   '1000',
    'mmey':         '1000000',
    'Mmey':         '1000000',
    'lovelace':     '1000000',
    'picoelement':    '1000000',
    'gmey':         '1000000000',
    'Gmey':         '1000000000',
    'shannon':      '1000000000',
    'nanoelement':    '1000000000',
    'nano':         '1000000000',
    'szabo':        '1000000000000',
    'microelement':   '1000000000000',
    'micro':        '1000000000000',
    'finney':       '1000000000000000',
    'millielement':    '1000000000000000',
    'milli':         '1000000000000000',
    'element':        '1000000000000000000', (general unit)
    'kelement':       '1000000000000000000000',
    'grand':        '1000000000000000000000',
    'melement':       '1000000000000000000000000',
    'gelement':       '1000000000000000000000000000',
    'telement':       '1000000000000000000000000000000'
<br>
--------------------------------
#More objects and wiki are coming in the next update pack!


##License##
---------------------------------
The go-elementrem library (i.e. all code outside of the cmd directory) is licensed under the GNU Lesser General Public License v3.0, also included in our repository in the COPYING.LESSER file. http://www.gnu.org/licenses/lgpl-3.0.en.html

The go-elementrem binaries (i.e. all code inside of the cmd directory) is licensed under the GNU General Public License v3.0, also included in our repository in the COPYING file. http://www.gnu.org/licenses/gpl-3.0.en.html
