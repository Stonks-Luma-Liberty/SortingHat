package discord

import (
	"github.com/bwmarrin/discordgo"
	"golang.org/x/exp/slices"
)

// GetRole returns a role from a guild by name
func GetRole(s *discordgo.Session, guildID string, roleName string) *discordgo.Role {
	guildRoles, err := s.GuildRoles(guildID)
	if err != nil {
		panic(err)
	}

	role := guildRoles[slices.IndexFunc(guildRoles, func(c *discordgo.Role) bool { return c.Name == roleName })]
	return role
}
