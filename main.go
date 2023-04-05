package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"discordv2.at/m/v2/commands"
	"discordv2.at/m/v2/config"
	"github.com/bwmarrin/discordgo"
)

func init() {
	c, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	config.Config = c
}

func main() {
	dg, err := discordgo.New("Bot " + config.Config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commands.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	if err := dg.Open(); err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands.Commands))
	for i, v := range commands.Commands {
		cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	dg.Close()
}
