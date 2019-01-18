#!/usr/bin/env bash

function isRoot () {
        if [ "$EUID" -ne 0 ]; then
                return 1
        fi
}

function menu () {
        echo
        echo "Welcome to Cover my ass tool !"

        echo
        echo "Select an option :"
        echo
        echo "1) Clear auth & bash history for user $USER"
        echo "2) Permenently disable auth & bash history"
        echo "3) Restore settings to default"
        echo "99) Exit tool"
        echo

        printf "> "
        read -r option
        echo
}

function disableAuth () {
        if [ -w /var/log/auth.log ]; then
                ln /dev/null /var/log/auth.log -sf
                echo "[+] Permanently sending /var/log/auth.log to /dev/null"
        else
                echo "[!] /var/log/auth.log is not writable! Retry using sudo."
        fi
}

function disableHistory () {
        ln /dev/null ~/.bash_history -sf
        echo "[+] Permanently sending bash_history to /dev/null"

        if [ -a ~/.zsh_history ]; then
                ln /dev/null ~/.zsh_history -sf
                echo "[+] Permanently sending zsh_history to /dev/null"
        fi

        export HISTFILESIZE=0
        export HISTSIZE=0
        echo "[+] Set HISTFILESIZE & HISTSIZE to 0"

        set +o history
        echo "[+] Disabled history library"

        echo
        echo "Permenently disabled bash log."
}

function enableAuth () {
        if [ -w /var/log/auth.log ] && [ -L /var/log/auth.log ]; then
                rm -rf /var/log/auth.log
                echo "" > /var/log/auth.log
                echo "[+] Disabled sending auth logs to /dev/null"
        else
                echo "[!] /var/log/auth.log is not writable! Retry using sudo."
        fi
}

function enableHistory () {
        if [[ -L ~/.bash_history ]]; then
                rm -rf ~/.bash_history
                echo "" > ~/.bash_history
                echo "[+] Disabled sending history to /dev/null"
        fi

        if [[ -L ~/.zsh_history ]]; then
                rm -rf ~/.zsh_history
                echo "" > ~/.zsh_history
                echo "[+] Disabled sending zsh history to /dev/null"
        fi

        export HISTFILESIZE=""
        export HISTSIZE=50000
        echo "[+] Restore HISTFILESIZE & HISTSIZE default values."

        set -o history
        echo "[+] Enabled history library"

        echo
        echo "Permenently enabled bash log."
}

function clearAuth () {
        if [ -w /var/log/auth.log ]; then
                echo "" > /var/log/auth.log
                echo "[+] /var/log/auth.log cleaned."
        else
                echo "[!] /var/log/auth.log is not writable! Retry using sudo."
        fi
}

function clearHistory () {
        if [ -a ~/.zsh_history ]; then
                echo "" > ~/.zsh_history
                echo "[+] ~/.zsh_history cleaned."
        fi

        echo "" > ~/.bash_history
        echo "[+] ~/.bash_history cleaned."

        history -c
        echo "[+] History file deleted."

        echo
        echo "Reminder: your need to reload the session to see effects."
        echo "Type exit to do so."
}

function exitTool () {
        exit 1
}

clear # Clear output

# "now" option
if [ -n "$1" ] && [ "$1" == 'now' ]; then
        clearAuth
        clearHistory
        exitTool
fi

menu

if [[ $option == 1 ]]; then
        # Clear current history
        clearAuth
        clearHistory
elif [[ $option == 2 ]]; then
        # Permenently disable auth & bash log
        disableAuth
        disableHistory
elif [[ $option == 3 ]]; then
        # Restore default settings
        enableAuth
        enableHistory
elif [[ $option == 99 ]]; then
        # Exit tool
        exitTool
else
        echo "[!] Option not reconized. Exiting."
fi
