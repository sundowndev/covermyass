# Cover my ass

CLI tool to cover your tracks on UNIX systems. Designed for pen testing "Covering Tracks" phase, before exiting the infected server. Or, even better, permanently disable bash & auth history.

**This tool supports zsh shell.**

## Installation

Read the install script before running it. You'll need sudo privileges.

```
curl -sSL https://raw.githubusercontent.com/sundowndev/covermyass/master/install.sh | bash
```

Without sudo :

```
cd $HOME
curl -sSL https://raw.githubusercontent.com/sundowndev/covermyass/master/covermyass.sh -o ./covermyass
chmod +x ./covermyass
```

Keep in mind that without sudo privileges, you'll be unable to clean auth logs.

## Usage

Simply type :

```
covermyass # use sudo if you want to clean auth logs
```

Follow the instructions :

```
Welcome to Cover my ass tool !

Select an option :

1) Clear auth & bash history for user sundowndev
2) Permenently disable bash log
3) Restore settings to default
99) Exit tool

>
```
