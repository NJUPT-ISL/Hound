package cronjob

import (
	"github.com/NJUPT-ISL/Hound/agent/lib"
	"github.com/NJUPT-ISL/Hound/agent/log"
	"github.com/robfig/cron"
)

func AddSendUpdateJob(c *cron.Cron) {
	spec := "0 */5 * * * ?"
	if err := c.AddFunc(spec, func() {
		lib.SendUpdate()
	}); err != nil {
		log.ErrPrint(err)
	}
}

func AddDockerPruneJob(c *cron.Cron) {
	if err := c.AddFunc("0 */10 * * * ?", func() {
		if _, err := lib.ImagesPrune(); err != nil {
			log.ErrPrint(err)
		}
	}); err != nil {
		log.ErrPrint(err)
	}
}
