# PakkretQC
HPQC with no IE shit 🤷🏼‍♂️



## Development
This project is focusing on simplicity of development tools. It is included with a set of simple tool chain built-in natively without any fancy dependency.

### Front-end and Server (compile TS + Development server)
``` 
$ PAKKRETQC_ALM_ENDPOINT=https://[your-alm-hpqc-endpoint] go run devtools/cmd/appbundler/main.go -w
```
then devtool will compile typescript files and serve front-end server.


## Build
```
$ ./build.sh
$ ./bundler.sh ## optionally packing a tarball
```