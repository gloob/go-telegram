package tg

import (
	"github.com/codegangsta/cli"
)

// Application constants.
const (
	AppName        = "go-tg"
	AppVersion     = "0.1"
	AppDescription = "Command-line interface for Telegram in Go."
	AppTgName      = "Federator"
	AppId          = 10604
	AppHash        = "34be6d99874fb9607fe932dbb86fe4a3"
)

// Arrays aren't immutable so we are defining this outside of the const
// declaration.
var AppAuthors = []cli.Author{
	{Name: "Alejandro Leiva", Email: "gloob@litio.org"},
}

// Directories and file locations.
const (
	ConfigDirectory    = "." + AppName
	DownloadsDirectory = "downloads"
	ConfigFile         = "config"
	AuthKeyFile        = "auth"
	StateFile          = "state"
	SecretChatFile     = "secret"
	BinlogFile         = "binlog"
)
