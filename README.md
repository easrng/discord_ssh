# discord_ssh
manage your authorized_keys with discord.

## should i use this
please don't. i made it as a joke. if you use it i will cry.

## setup
### system
  - install discord_ssh with proper permissions
    ```sh
    go install github.com/easrng/discord_ssh@latest && sudo sh -c 'cp '"${GOPATH:-$HOME/go}"'/bin/discord_ssh /sbin/discord_ssh && chmod 700 /sbin/discord_ssh'
    ```
  - edit your `/etc/ssh/sshd_config` file to include
    ```
    AuthorizedKeysCommand /sbin/discord_ssh
    AuthorizedKeysCommandUser root
    ```
  - restart sshd to apply the changes
    ```sh
    sudo systemctl restart sshd
    ```
### user
  - make a file named `~/.ssh/config_discord` that only your account can access
    ```sh
    touch ~/.ssh/config_discord
    chmod 600 ~/.ssh/config_discord
    ```
  - edit it to include your discord bot token and the id of the channel that will contain your authorized keys
    ```json
    {
        "token": "MDAwMDAwMDAwMDAwMDAwMDAwMA.GACckG.eqfP9yG2Irjn6tdJG7Y5LU5OWFSjzdHwTEPomQ",
        "channel": "1000000000000000000"
    }
    ```
that's it, you're good to go