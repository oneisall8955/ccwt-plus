package service

import "strings"

const (
	ProviderClaude = "claude"
	ProviderCodex  = "codex"
)

func NormalizeProvider(provider string) string {
	switch strings.ToLower(strings.TrimSpace(provider)) {
	case ProviderCodex:
		return ProviderCodex
	default:
		return ProviderClaude
	}
}
