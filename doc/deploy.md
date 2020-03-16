---
title: Fab-IoT-Lab LoRa embedded stack
subtitle: How to deploy
author: François Roland
date: 11 mars 2020
papersize: A4
documentclass: scrartcl
...

# Requirements

* Linux machine with docker (must support aufs overlay)
* [balena-cli](https://github.com/balena-io/balena-cli/blob/master/INSTALL.md)
* [balenaEtcher](https://www.balena.io/etcher/)

# Steps to create fabiotlab-lora-stack gateway

1. Create a new user account on https://dashboard.balena-cloud.com/
1. Create a new application
   - choose an application name;
   - choose your gateway device type;
   - select "starter" application type (max 10 devices).
1. Add a new device
   - select development or production image;
   - download image on local computer;
   - unzip the downloaded file.
1. Fork https://github.com/fablabmons/fabiotlab-lora-stack.
1. Clone your forked GitHub repository locally.
1. Make sure [balena-cli](https://github.com/balena-io/balena-cli/blob/master/INSTALL.md) is installed locally.
1. Execute these commands:
```sh
balena login
balena push <your_application_name>
balena preload <unzipped_img> --app <your_application_name> --commit current
```
1. Write image on SD card with balenaEtcher.
