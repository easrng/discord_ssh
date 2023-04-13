package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/crypto/ssh"
)

func uncheckedJsonMarshal(in interface{}) []byte {
	r, _ := json.Marshal(in)
	return r
}

type Config struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("# Please provide a username as an argument.")
		os.Exit(1)
	}
	username := os.Args[1]
	u, err := user.Lookup(username)
	if err != nil {
		fmt.Printf("# Error looking up user %s\n", uncheckedJsonMarshal(username))
		os.Exit(1)
	}
	configPath := u.HomeDir + "/.ssh/config_discord"
	stat, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		fmt.Printf("# The .ssh/config_discord file for user %s does not exist.\n", uncheckedJsonMarshal(username))
		os.Exit(1)
	}
	mode := stat.Mode()
	if (mode.Perm() & 0077) != 0 {
		fmt.Printf("# The .ssh/config_discord file for user %s is accessible by others.\n", uncheckedJsonMarshal(username))
		os.Exit(1)
	}
	jsonBytes, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("# Error reading config file: %s\n", uncheckedJsonMarshal(err))
		os.Exit(1)
	}
	config := &Config{}
	err = json.Unmarshal(jsonBytes, config)
	if err != nil {
		fmt.Printf("# Error parsing JSON: %s\n", uncheckedJsonMarshal(err))
		os.Exit(1)
	}
	discord, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Printf("# Error creating client: %s\n", uncheckedJsonMarshal(err))
		os.Exit(1)
	}
	var lastMessageID string
	for {
		messages, err := discord.ChannelMessages(config.Channel, 100, lastMessageID, "", "")
		if err != nil {
			fmt.Printf("# Error getting channel messages: %s\n", uncheckedJsonMarshal(err))
			os.Exit(1)
		}
		if len(messages) == 0 {
			break
		}
		for _, message := range messages {
			publicKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(message.Content))
			if err != nil {
				continue
			}
			reserializedKey := ssh.MarshalAuthorizedKey(publicKey)
			fmt.Printf("%s", reserializedKey)
		}
		lastMessageID = messages[len(messages)-1].ID
	}
}
