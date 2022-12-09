# TMHI - T-Mobile Home Internet CLI [![CI Status](https://circleci.com/gh/cloud-unpacked/tmhi.svg?style=shield)](https://app.circleci.com/pipelines/github/cloud-unpacked/tmhi) [![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/cloud-unpacked/tmhi/trunk/LICENSE)

TMHI is a CLI to manage your local T-Mobile Home Internet 5G Gateway/Router.
This tool can pull some data from your T-Mobile gateway such as signal metrics and information on the gateway itself.


## Table of Contents

- [Compatibility](#compatibility)
- [Installation](#installation)
  - [Linux](#linux)
  - [macOS](#macos)
  - [Windows](#windows)
- [Configuring](#configuring)
- [Features](#features)


## Compatibility

### Operating Systems

Designed to work on Linux, macOS, and Windows computers.

### Gateway Models

This has been tested with the Arcadyan KVD21 gateway only.
The Sagemcom Fast 5688W should also work in theory, but hasn't been tested yet.
The Nokia gateway hasn't been tested and I don't believe it will work at all.


## Installation

### Linux

#### Debian Package (.deb)
You can install TMHI into an Debian/Apt based computer by download the `.deb` file to the desired system.

For graphical systems, you can download it from the [GitHub Releases page][gh-releases].
Many distros allow you to double-click the file to install.
Via terminal, you can do the following:

```bash
wget https://github.com/cloud-unpacked/tmhi/releases/download/v0.1.0/tmhi_0.1.0_amd64.deb
sudo dpkg -i tmhi_0.1.0_amd64.deb
```

`0.1.0` and `amd64` may need to be replaced with your desired version and CPU architecture respectively.

#### Binary Install
You can download and run the raw TMHI binary from the [GitHub Releases page][gh-releases] if you don't want to use any package manager.
Simply download the tarball for your OS and architecture and extract the binary to somewhere in your `PATH`.
Here's one way to do this with `curl` and `tar`:

```bash
dlURL="https://github.com/cloud-unpacked/tmhi/releases/download/v0.1.0/tmhi-v0.1.0-linux-amd64.tar.gz"
curl -sSL $dlURL | sudo tar -xz -C /usr/local/bin tmhi
```

`0.1.0` and `amd64` may need to be replaced with your desired version and CPU architecture respectively.

### macOS

There are two ways you can install TMHI on a macOS system.

#### Brew (recommended)

Installing TMHI via brew is a simple one-liner:

```bash
brew install cloud-unpacked/tap/tmhi
```

#### Binary Install
You can download and run the raw TMHI binary from the [GitHub Releases page][gh-releases] if you don't want to use any package manager.
Simply download the tarball for your OS and architecture and extract the binary to somewhere in your `PATH`.
Here's one way to do this with `curl` and `tar`:

```bash
dlURL="https://github.com/cloud-unpacked/tmhi/releases/download/v0.1.0/tmhi-v0.1.0-macos-amd64.tar.gz"
curl -sSL $dlURL | sudo tar -xz -C /usr/local/bin tmhi
```

`0.1.0` and `amd64` may need to be replaced with your desired version and CPU architecture respectively.

### Windows

TMHI supports Windows 10 by downloading and installing the binary.
Chocolately support is likely coming in the future.
If there's a Windows package manager you'd like support for (including Chocolately), please open and Issue and ask for it.

#### Binary Install (exe)
You can download and run the TMHI executable from the [GitHub Releases page][gh-releases].
Simply download the zip for architecture and extract the exe.


## Configuring

`tmhi login` needs to be run so that many other commands can work.
The username and password will be the same ones used to login via the T-Mobile Home Internet app.
Typically the username is `admin` unless you changed it.


## Features

With `tmhi` you can:


- view data on the gateway
- view signal data
- reboot the gateway
- more coming soon

Run `tmhi help` to see all commands available.


## License

This repository is licensed under the MIT license.
The license can be found [here](./LICENSE).



[gh-releases]: https://github.com/cloud-unpacked/tmhi/releases
