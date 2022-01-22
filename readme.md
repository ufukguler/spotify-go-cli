## TL;DR

```bash
go build
./spotify-go-cli set-client xxxxx
./spotify-go-cli set-secret xxxxx
./spotify-go-cli login
````

### available commands

```
$ .\spotify-go-cli --help
Usage:
  spotify-go-client [flags]
  spotify-go-client [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell      
  config      Show current configuration
  current     Get the object currently being played.                          
  devices     Get Available Devices                                           
  help        Help about any command                                          
  login       Login                                                           
  next        Skip To Next                                                    
  pause       Pause Playback                                                  
  previous    Skip To Previous                                                
  repeat      Set Repeat Mode                                                 
  set-client  Set Client ID                                                   
  set-secret  Set Secret Key                                                  
  start       Start/Resume Playback                                           
  transfer    Transfer Playback                                               
  version     Get CLI version                                                 
  volume      Set Playback Volume  

Flags:
      -h, --help            help for spotify-go-client

Use "spotify-go-client [command] --help" for more information about a command.
```

---

#### to get client id and secret key & fix callback url error

1. create an application from spotify developer console

> https://developer.spotify.com/dashboard/applications

2. add callback url (`http://localhost:8080`) to your application at spotify developer console

## TODO

- [ ] refresh token
- [ ] customize answers by status code
- [ ] add all player commands (9/14 done)


