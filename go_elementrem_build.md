### To build from source	

***As a general rule, Do not lose your password and keystore file. In any case, recovery is impossible.***
***And it doesn't collect any personally identifiable information.***

**Prerequisite**

`sudo apt-get install -y build-essential libgmp3-dev git curl`		

* [Go lang](https://golang.org/dl/)   
Ubuntu, for instance    
`sudo apt-get install -y golang`		
```
sudo apt-get update
sudo apt-get -y upgrade
wget https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz
sudo tar -xvf go1.7.linux-amd64.tar.gz
sudo mv go /usr/local
export GOROOT=/usr/local/go
export GOPATH=$HOME/Projects/Proj1
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

* [Docker](https://www.docker.com/products/docker#/servers)   
Ubuntu, for instance    
`sudo curl -fsSL https://get.docker.com/ | sh`    
`sudo usermod -aG docker $(whoami)` (might require reboot)  
`docker pull ubuntu:16.04` (16.04 ubuntu version. if 15.04 `docker pull ubuntu:15.04`) 

**build** 

Don't expect that you can build app for all platforms on one platform.		
- If your app has native dependencies, it can be compiled only on the target platform. prebuild is a solution, but most node modules don't provide prebuilt binaries.		
- OS Code Signing works only on MacOS. Cannot be fixed.
- Elementrem Gele provides support for 64-bit only.		

```
cd go-elementrem    
make gele
    //(or, to build the other OS suite of utilities:
make gele-windows
make gele-linux
```

**Run Gele**		
- Windows		
1.Run `gele.exe`    
2.Automatically connects to the Elementrem node : Default directory will be created in the C:\Users\<User Account>\AppData\Roaming\Elementrem	
3.Run `gele attach` from command prompt window(cmd.exe) will be automatically connected to the Elementrem console.	
		<br>
- Linux		
1.Run `gele` from Terminal - If do not copy gele to /usr/bin it must be run ./gele    
2.Automatically connects to the Elementrem node : Default directory will be created in the /home/<User Account>/.elementrem		
3.Run `gele attach` from Terminal will be automatically connected to the Elementrem console.		
		<br>
- Mac OSX	
1.Run `gele` from Terminal - If do not copy gele to /usr/bin it must be run ./gele    
2.Automatically connects to the Elementrem node : Default directory will be created in the /Library/Elementrem		
3.Run `gele attach` from Terminal(launchpad -> Terminal) will be automatically connected to the Elementrem console.	    
		<br>

*Elementrem Default Network listening port = 30707		
*Elementrem Default RPC port = 7075

