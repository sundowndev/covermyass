#!/usr/bin/env bash

function isRoot () {
        if [ "$EUID" -ne 0 ]; then
                return 1
        fi
}

clear # Clear output

echo
echo "Welcome to Cover my ass tool !"

echo
echo "Select an option :"
echo
echo "1) Clear auth & bash history for user $USER"
echo "2) Permenently disable bash log"
echo "3) Kill current session"
echo "4) Restore settings to default"
echo "99) Exit tool"
echo

printf "Choice: "
read -r option
echo

if [[ $option == 1 ]]; then
        # Clear current history
        if [ -w /var/log/auth.log ]; then
                echo "" > /var/log/auth.log
        else
                echo "[!] /var/log/auth.log is not writable! Retry using sudo."
        fi
        
        echo "" > ~/.bash_history
        rm ~/.bash_history -rf
        history -c
        
        echo "Bash history cleaned."
        echo "Reminder: your need to kill current terminal session to see changes."
elif [[ $option == 2 ]]; then
        # Permenently disable bash log
        ln /dev/null ~/.bash_history -sf # Permanently send history to /dev/null
        echo "Permenently disabled bash log."
elif [[ $option == 3 ]]; then
        kill -9 $$ # Kill current session
elif [[ $option == 4 ]]; then
        # Restore settings to default
        echo "4"
elif [[ $option == 99 ]]; then
        exit 1
else
        echo "Option not reconized. Exiting."
fi
