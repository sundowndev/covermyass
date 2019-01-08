#!/usr/bin/env bash

function isRoot () {
        if [ "$EUID" -ne 0 ]; then
                return 1
        fi
}

clear

echo
echo "Welcome to Cover my ass tool !"

echo
echo "Select an option :"
echo
echo "1) Clear current bash history"
echo "2) Permenently disable bash log"
echo "3) Kill current session"
echo "4) Restore settings to default"
echo "99) Exit"
echo

printf "Choice: "
read -r option
echo

if [[ $option == 1 ]]; then
        # Clear current history
        if [ -w /var/log/auth.log ]; then
                echo "" > /var/log/auth.log
        else
                echo "[!] /var/log/auth.log is not writable! Skipping."
        fi
        echo "" > ~/.bash_history
        #rm ~/.bash_history -rf
        history -c
        echo "Bash history cleaned."
        echo "Reminder: your need to restart current terminal session to see changes."
elif [[ $option == 2 ]]; then
        echo "2"
elif [[ $option == 3 ]]; then
        echo "3"
elif [[ $option == 4 ]]; then
        echo "4"
elif [[ $option == 99 ]]; then
        exit 1
else
        echo "Option not reconized. Exiting."
fi

#export HISTFILESIZE=0
#export HISTSIZE=0
#unset HISTFILE

# Kill current session
#kill -9 $$

# Permanently send history to /dev/null
#ln /dev/null ~/.bash_history -sf
