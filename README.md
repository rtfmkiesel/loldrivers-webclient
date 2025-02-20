# LOLDrivers-webclient

A PoC web client for [LOLDrivers](https://github.com/magicsword-io/LOLDrivers) (Living Off The Land Drivers) by [MagicSword](https://www.magicsword.io/). Scan your computer for known vulnerable and known malicious Windows drivers.

![A demo picture of the web application](assets/demo.png)

*This was hacked together in like 2,5 hours, please have mercy*

## Usage

Due to changes in the File System Access API, scanning the default Windows folders containing drivers is no longer possible. See [this](https://developer.chrome.com/docs/capabilities/web-apis/file-system-access#restricted_folders).

<del>Click [here](https://rtfmkiesel.github.io/loldrivers-webclient), wait a few seconds, and then select the directory to scan. The default Windows driver directories are:</del>  
<del>+ C:\Windows\System32\drivers</del>  
<del>+ C:\Windows\System32\DriverStore\FileRepository</del>  
<del>+ C:\WINDOWS\inf</del>  

## Build
To build `loldrivers-webclient.wasm` yourself:
```sh
# linux
GOOS=js GOARCH=wasm go build -o .\assets\loldrivers-webclient.wasm

# windows
$Env:GOOS = "js"
$Env:GOARCH ="wasm"
go build -o .\assets\loldrivers-webclient.wasm
```
