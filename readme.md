Medium post: https://medium.com/@ufuk.guler/how-to-create-a-cli-in-golang-501ea927df15

## TL; DR
Spotify CLI made with Golang. (WIP)

```bash
go build
./spotify-go-cli set-client xxxxx
./spotify-go-cli set-secret xxxxx
./spotify-go-cli login
````
---

### Available Commands

```
$ .\spotify-go-cli --help

Usage:
  spotify-go-client [flags]
  spotify-go-client [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Show current configuration
  current     Get Currently Playing Track
  devices     Get Available Devices
  help        Help about any command
  login       Login
  next        Skip To Next
  pause       Pause Playback
  previous    Skip To Previous
  profile     Get Current User's Profile
  queue       Add Item to Playback Queue
  recent      Get Recently Played Tracks
  repeat      Set Repeat Mode
  search      Search for Item
  seek        Seek To Position
  set-client  Set Client ID
  set-secret  Set Secret Key
  shuffle     Toggle Playback Shuffle
  start       Start/Resume Playback
  state       Get Playback State
  transfer    Transfer Playback
  version     Get CLI version
  volume      Set Playback Volume

Flags:
      --config string   config file ($HOME/.spotify-go.yml)
  -h, --help            help for spotify-go-client

Use "spotify-go-client [command] --help" for more information about a command.
```

---

#### to get client id and secret key & fix callback url error

1. create an application from spotify developer console

> https://developer.spotify.com/dashboard/applications

2. add callback url (`http://localhost:8080`) to your application at spotify developer console

## TODO

- [X] refresh token
- [X] add all player commands (14/14 done)


