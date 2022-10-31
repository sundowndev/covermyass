# covermyass

[![Build status](https://github.com/sundowndev/covermyass/workflows/Go%20build/badge.svg)](https://github.com/sundowndev/covermyass/actions)
[![Tag](https://img.shields.io/github/tag/SundownDEV/covermyass.svg)](https://github.com/sundowndev/covermyass/releases)

### About

Covermyass is a post-exploitation tool to cover your tracks on various operating systems (Linux, Darwin, Windows, ...). It was designed for penetration testing "covering tracks" phase, before exiting the infected server. At any time, you can run the tool to find which log files exists on the system, then run again later to erase those files. The tool will tell you which file can be erased with the current user permissions.

## Installation

With sudo

```bash
sudo curl -sSL https://github.com/sundowndev/covermyass/releases/latest/download/covermyass_Linux_x86_64 -o /usr/bin/covermyass
sudo chmod +x /usr/bin/covermyass
```

Without sudo :

```bash
curl -sSL https://github.com/sundowndev/covermyass/releases/latest/download/covermyass_Linux_x86_64 -o ~/.local/bin/covermyass
chmod +x ~/.local/bin/covermyass
```

Keep in mind that without sudo privileges, you *might* be unable to clear system-level log files.

## Usage

Run an analysis to find log files : 

```
covermyass
```

Clear log files instantly :

```
covermyass --write
```

Add custom file paths : 

```
covermyass -p '/db/**/*.log'
```

Filter out some paths : 

```
covermyass -f '/foo/bar/*.log'
covermyass -f '/foo/bar.log'
```
