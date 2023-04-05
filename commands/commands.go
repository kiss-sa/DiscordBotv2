package commands

import "github.com/bwmarrin/discordgo"

// var (
// 	integerOptionMinValue          = 1.0
// 	dmPermission                   = false
// 	defaultMemberPermissions int64 = discordgo.PermissionManageServer
// )

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "basic-command",
		Description: "Basic command",
	},
	{
		Name:        "pun",
		Description: "Gives a random pun!",
	},
}

var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"basic-command": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hey there! Congratulations, you just executed your first slash command",
			},
		})
	},
	"pun": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		pun, err := getPun()
		if err != nil {
			panic(err)
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: pun,
			},
		})
	},
}
