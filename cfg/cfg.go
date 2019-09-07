package cfg

import (
	"github.com/sqsinformatique/backend/utils"
)

var DefaultCfg struct {
	VerboseLevel utils.Level `environment:"VERBOSE_LEVEL,InfoLevel"`
	Port         string      `environment:"ELEVATOR_PORT,9777"`
	AppLog       string      `environment:"ELEVATOR_APP_LOG,"`
	DSN          string      `environment:"ELEVATOR_DSN,postgres://postgres:q2w3e4r@localhost:5433/userdb?sslmode=disable""`
}

func init() {
	utils.LoadFromEnv(&DefaultCfg)
	utils.InitLogger(DefaultCfg.AppLog)

}
