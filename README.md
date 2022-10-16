# Media server

## Configuration
**Golang Version:** go1.19.2 darwin/amd64

**Docker support:** none


## Description
This server can send any mediafiles with using http requests.

## Setup
### Storage root folder
All data are stored in `./data/`. If you want to change it - use config file (`config.yml` as default).
    
*`./data/` file is in `.gitignore`.*

### Media request 
This server version support images and videos loading

### Images
All images are stored in `./data/images/`. If you want to change it - use config file (`config.yaml` as default). 

### Videos
All images are stored in `./data/videos/`. If you want to change it - use config file (`config.yaml` as default). 

## Flags
*-config* - path to config. Default value is `config.yaml`.

*-h* - get all aviable flags