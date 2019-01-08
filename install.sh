#!/usr/bin/env bash

function isRoot () {
        if [ "$EUID" -ne 0 ]; then
                return 1
        fi
}

if ! isRoot; then
  echo "You need to run this as root!"
  exit 1
fi

sudo wget -O /usr/bin/covermyass https://raw.githubusercontent.com/sundowndev/covermyass/master/covermyass.sh
sudo chmod +x /usr/bin/covermyass

echo "Installation succeeded."
