package handlers

import (
	"os"

	logger "github.com/anikhasibul/log"
)

// custom logger
// nolint: gochecknoglobals
var log = logger.New(os.Stdout)
