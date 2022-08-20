package config

type Config struct {
	DiscordBotToken string `env:"DISCORD_BOT_TOKEN"`
	ApplicationId   string `env:"APPLICATION_ID"`
	GuildId         string `env:"GUILD_ID"`
}
