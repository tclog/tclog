# tclog - Twitch Chat Logger

`tclog` is a console application to log [Twitch](https://twitch.tv) chats.

This project is still in its infancy. I'm hoping to get feedback from people using it.

If you spot any bugs or have features that you'd really like to see in `tclog`, please submit [an issue](https://github.com/tclog/tclog/issues).

This project is maintained by [tastyfriedtofu](https://twitch.tv/tastyfriedtofu).

## Usage

Join a channel as an anonymous user and log every messages to `./tclog_[channel]_[YYYYMMDD].log`.

```bash
tclog [channel]
```

The log use JSON for structured logging. 

That's it, for now. I will add more functionality and configurations little by little.

## Documentations

Read our [wiki page](https://github.com/tclog/tclog/wiki) to learn how to install and use `tclog`.

## Installation

### Windows

`tclog` is available via [scoop](https://scoop.sh).

Install:

```powershell
scoop bucket add tclog https://github.com/tclog/scoop-tclog.git
scoop install tclog
```

Upgrade:

```powershell
scoop update tclog
```

You can also download the prebuilt binary as a [zip file][releases page].

### macOS

You can download the prebuilt binary from the [releases page][].

### Linux

#### DEB package

For Debian or Ubuntu distribution.

Install and upgrade:

1. Download the `.deb` file from the [releases page][]
2. Run `sudo apt install ./tclog_*_linux_*.deb`

#### RPM package

Install and upgrade:

1. Download the `.rpm` file from the [releases page][]
2. Run:
    - Fedora:
        ```bash
        sudo dnf install tclog_*_linux_*.rpm
        ```
    - Centos:
        ```bash
        sudo yum localinstall tclog_*_linux_*.rpm
        ```
    - openSUSE/SUSE:
        ```bash
        sudo zypper in tclog_*_linux_*.rpm
        ```
    
### Other platform

Install a prebuilt binary from the [releases page][].

### Build from source

I use [goreleaser](https://goreleaser.com/) to build this project on my local workstation. I am too lazy to write a Makefile. If you want to, please submit a PR. 

1. [Install goreleaser](https://goreleaser.com/install/)
2. Clone this repository: `git clone https://github.com/tclog/tclog && cd tclog`
2. Build the project: `goreleaser --snapshot --skip-publish --rm-dist`
3. Look for `./dist` directory for the built binaries
4. Move the resulting binary executable to somewhere in your PATH
5. Run `tclog version` to check if it worked. 

[releases page]: https://github.com/tclog/tclog/releases/latest

## Security

If you discover a security issue in this repository, please email directly to tclog.issues@gmail.com.
