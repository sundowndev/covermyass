# Covermyass

![Build status](https://img.shields.io/travis/sundowndev/covermyass/master.svg?style=flat-square)
![Tag](https://img.shields.io/github/tag/SundownDEV/covermyass.svg?style=flat-square)

CLI tool to cover your tracks on UNIX systems. Designed for pen testing "Covering Tracks" phase, before exiting the infected server. Or, even better, permanently disable bash & auth history.

This tool allows you to clear log files such as :

- `/var/log/messages` : General message and system related stuff
- `/var/log/auth.log` : Authenication logs
- `/var/log/kern.log` : Kernel logs
- `/var/log/cron.log` : Crond logs (cron job)
- `/var/log/maillog` : Mail server logs
- `/var/log/qmail/` : Qmail log directory (more files inside this directory)
- `/var/log/httpd/` : Apache access and error logs directory
- `/var/log/lighttpd/` : Lighttpd access and error logs directory
- `/var/log/boot.log` : System boot log
- `/var/log/mysqld.log` : MySQL database server log file
- `/var/log/secure` or `/var/log/auth.log` : Authentication log
- `/var/log/utmp` or `/var/log/wtmp` : Login records file
- `/var/log/yum.log` : Yum command log file.

**This tool supports zsh & bash shell.**

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
covermyass # you may need to use sudo if you want to clean auth logs
```

Follow the instructions :

```
Welcome to Cover my ass tool !

Select an option :

1) Clear auth & bash history for user root
2) Permenently disable auth & bash history
3) Restore settings to default
99) Exit tool

>
```

*NOTE: don't forget to exit the terminal session since the bash history is cached.*

Clear auth & history instantly :

```
sudo covermyass now
```

### Using cron job

Clear bash history every day at 5am :

```
0 5 * * * covermyass now >/dev/null 2>&1
```
