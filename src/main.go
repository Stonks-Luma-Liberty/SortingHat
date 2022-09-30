package main

import (
	"github.com/Stonks-Luma-Liberty/SortingHat.git/src/config"
	"github.com/Stonks-Luma-Liberty/SortingHat.git/src/discord"
	"github.com/bwmarrin/discordgo"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

var session *discordgo.Session
var cfg config.Config
var err error

func userHasRole(memberRoles []string, roleName string) bool {
	for _, role := range memberRoles {
		if role == roleName {
			return true
		}
	}
	return false
}

// Loads the environment variables from the .env file, parses them into the config struct, and then starts a new Discord
// session
func init() {
	log.Println("Loading environment variables.")
	dotenvErr := godotenv.Load()
	if dotenvErr != nil {
		log.Fatal("Error loading .env file")
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	log.Println("Starting SortingHat bot session.")
	session, err = discordgo.New("Bot " + cfg.DiscordBotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", dotenvErr)
	}
}

// Defining handlers for components and commands.
var (
	componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"role_crypto_trader": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			roleName := "Crypto Trader"
			content := ""
			role := discord.GetRole(s, i.GuildID, roleName)
			user := i.Member.User
			log.Printf("Assigning role [%s] to %s", role.Name, user.Username)

			member, err := s.GuildMember(i.GuildID, user.ID)
			if err != nil {
				return
			}

			if userHasRole(member.Roles, roleName) {
				err = s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, role.ID)
				if err != nil {
					panic(err)
				}
				content = "You have been granted the " + role.Name + " role"
			} else {
				err = s.GuildMemberRoleRemove(i.GuildID, i.Member.User.ID, role.ID)
				if err != nil {
					panic(err)
				}
				content = "You have been removed from the " + role.Name + " role"
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
					Flags:   uint64(discordgo.MessageFlagsEphemeral),
				},
			})
			if err != nil {
				panic(err)
			}
		},
		"role_nft_trader": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			roleName := "NFT Trader"
			content := ""
			role := discord.GetRole(s, i.GuildID, roleName)
			user := i.Member.User
			log.Printf("Assigning role [%s] to %s", role.Name, user.Username)

			member, err := s.GuildMember(i.GuildID, user.ID)
			if err != nil {
				return
			}

			if userHasRole(member.Roles, roleName) {
				err = s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, role.ID)
				if err != nil {
					panic(err)
				}
				content = "You have been granted the " + role.Name + " role"
			} else {
				err = s.GuildMemberRoleRemove(i.GuildID, i.Member.User.ID, role.ID)
				if err != nil {
					panic(err)
				}
				content = "You have been removed from the " + role.Name + " role"
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
					Flags:   uint64(discordgo.MessageFlagsEphemeral),
				},
			})
			if err != nil {
				panic(err)
			}
		},
		"role_tweeter": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			roleName := "Tweeter"
			content := ""
			role := discord.GetRole(s, i.GuildID, roleName)
			user := i.Member.User
			log.Printf("Assigning role [%s] to %s", role.Name, user.Username)

			member, err := s.GuildMember(i.GuildID, user.ID)
			if err != nil {
				return
			}

			if userHasRole(member.Roles, roleName) {
				err = s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, role.ID)
				if err != nil {
					panic(err)
				}
				content = "You have been granted the " + role.Name + " role"
			} else {
				err = s.GuildMemberRoleRemove(i.GuildID, i.Member.User.ID, role.ID)
				if err != nil {
					panic(err)
				}
				content = "You have been removed from the " + role.Name + " role"
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
					Flags:   uint64(discordgo.MessageFlagsEphemeral),
				},
			})
			if err != nil {
				panic(err)
			}
		},
	}
	commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"roles": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Click buttons below to get your desired roles",
					//Flags:   uint64(discordgo.MessageFlagsEphemeral),
					// Buttons and other components are specified in Components field.
					Components: []discordgo.MessageComponent{
						// ActionRow is a container of all buttons within the same row.
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Label:    "Crypto Trader",
									Style:    discordgo.PrimaryButton,
									Disabled: false,
									CustomID: "role_crypto_trader",
									Emoji: discordgo.ComponentEmoji{
										Name: "📊",
									},
								},
								discordgo.Button{
									Label:    "NFT Trader",
									Style:    discordgo.PrimaryButton,
									Disabled: false,
									CustomID: "role_nft_trader",
									Emoji: discordgo.ComponentEmoji{
										Name: "🎨",
									},
								},
								discordgo.Button{
									Label:    "Tweeter",
									Style:    discordgo.PrimaryButton,
									Disabled: false,
									CustomID: "role_tweeter",
									Emoji: discordgo.ComponentEmoji{
										Name: "🐦",
									},
								},
							},
						},
					},
				},
			})
			if err != nil {
				panic(err)
			}
		},
	}
)

func main() {
	AppId := cfg.ApplicationId
	GuildId := cfg.GuildId

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is ready!")
	})
	// Components are part of interactions, so we register InteractionCreate handler
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commandsHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:

			if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})
	_, err := session.ApplicationCommandCreate(AppId, GuildId, &discordgo.ApplicationCommand{
		Name:        "roles",
		Description: "Start reaction roles",
	})

	if err != nil {
		log.Fatalf("Cannot create slash command: %v", err)
	}

	err = session.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer func(session *discordgo.Session) {
		err := session.Close()
		if err != nil {
			log.Fatalf("Encountered problems while closing session: %v", err)
		}
	}(session)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}
