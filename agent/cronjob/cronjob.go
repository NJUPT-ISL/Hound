package cronjob

import (
	"github.com/NJUPT-ISL/Hound/agent/lib"
	"github.com/NJUPT-ISL/Hound/agent/log"
	"github.com/robfig/cron"
)

func SendUpdateJob() {
	cronJob := cron.New()
	spec := "* */5 * * * ?"
	if err := cronJob.AddFunc(spec, func() {
		lib.SendUpdate()
	}); err != nil {
		log.ErrPrint(err)
	}
	cronJob.Start()
}
