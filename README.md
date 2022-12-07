## *covermyass* ##

[![Build status](https://github.com/sundowndev/covermyass/workflows/Go%20build/badge.svg)](https://github.com/sundowndev/covermyass/actions)
[![Tag](https://img.shields.io/github/tag/SundownDEV/covermyass.svg)](https://github.com/sundowndev/covermyass/releases)

### About ###

**Covermyass** is a post-exploitation tool to cover your tracks on various operating systems (Linux, Darwin, Windows, ...). It was designed for penetration testing "covering tracks" phase, before exiting the infected server. At any time, you can run the tool to find which log files exists on the system, then run again later to erase those files. The tool will tell you which file can be erased with the current user permissions. Files are overwritten repeatedly with random data, in order to make it harder for even very expensive hardware probing to recover the data.

### Current status ###

This tool is still in beta. Upcoming versions might bring breaking changes. For now, we're focusing Linux and Darwin support, Windows may come later.

### Installation ###

Download the latest release : 

```bash
curl -sSL https://github.com/sundowndev/covermyass/releases/latest/download/covermyass_linux_amd64 -o ./covermyass
chmod +x ./covermyass
```

### Usage ###

```
$ covermyass -h

Usage:
  covermyass [flags]

Examples:

Overwrite log files as well as those found by path /db/*.log
covermyass --write -p /db/*.log

Overwrite log files 5 times with a final overwrite with zeros to hide shredding
covermyass --write -z -n 5


Flags:
  -f, --filter strings   File paths to ignore (supports glob patterns)
  -h, --help             help for covermyass
  -n, --iterations int   Overwrite N times instead of the default (default 3)
  -l, --list             Show files in a simple list format. This will prevent any write operation
      --no-read-only     Exclude read-only files in the list. Must be used with --list
  -v, --version          version for covermyass
      --write            Erase found log files. This WILL shred the files!
  -z, --zero             Add a final overwrite with zeros to hide shredding

```

First, run an analysis. This will not erase anything.

```
covermyass
```

When you acknowledged the results, erase those files.

```
covermyass --write
```

Filter out some paths : 

```
covermyass -f '/foo/bar/*.log'
covermyass -f '/foo/bar.log'
```

### License ###

**covermyass** is licensed under the MIT license. Refer to [LICENSE](LICENSE) for more information.

## Support

<div align="center">
  <img src="https://github.com/sundowndev/static/raw/main/sponsors.svg?v=c68eba9" width="100%" heigh="auto" />
</div>
