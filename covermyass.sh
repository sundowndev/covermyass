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
echo "2) Permenently disable bash history"
echo "3) Restore settings to default"
echo "99) Exit tool"
echo

printf "> "
read -r option
echo

if [[ $option == 1 ]]; then
        # Clear current history
        if [ -w /var/log/auth.log ]; then
                echo "" > /var/log/auth.log
                echo "[+] /var/log/auth.log cleaned."
        else
                echo "[!] /var/log/auth.log is not writable! Retry using sudo."
        fi

        if [ -a ~/.zsh_history ]; then
                echo "" > ~/.zsh_history
                echo "[+] ~/.zsh_history cleaned."
        fi

        echo "" > ~/.bash_history
        rm ~/.bash_history -rf
        echo "[+] ~/.bash_history cleaned."

        history -c
        echo "[+] History file deleted."

        echo
        echo "Reminder: your need to reload the session to see effects."
        echo "Type exit to do so."
elif [[ $option == 2 ]]; then
        # Permenently disable bash log
        ln /dev/null ~/.bash_history -sf
        echo "[+] Permanently sending bash_history to /dev/null"

        if [ -a ~/.zsh_history ]; then
                ln /dev/null ~/.zsh_history -sf
                echo "[+] Permanently sending zsh_history to /dev/null"
        fi

        export HISTFILESIZE=0
        export HISTSIZE=0
        unset HISTFILE
        echo "[+] Set HISTFILESIZE & HISTSIZE to 0"

        set +o history

        echo
        echo "Permenently disabled bash log."
elif [[ $option == 3 ]]; then
        # Restore default settings
        #ln /dev/null ~/.bash_history -sf
        echo "[+] Disabled sending history to /dev/null"

        export HISTFILESIZE=""
        export HISTSIZE=50000
        echo "[+] Restore HISTFILESIZE & HISTSIZE default values"

        echo
        echo "Permenently enabeld bash log."
elif [[ $option == 99 ]]; then
        exit 1
else
        echo "[!] Option not reconized. Exiting."
fi
