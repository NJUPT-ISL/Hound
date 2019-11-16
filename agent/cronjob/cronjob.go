package cronjob

import (
	"github.com/NJUPT-ISL/Hound/agent/lib"
	"github.com/robfig/cron"
	"log"
)

func SendUpdateJob() {
	cronJob := cron.New()
	spec := "* */5 * * * ?"
	if err := cronJob.AddFunc(spec, func() {
		lib.SendUpdate()
	}); err != nil {
		log.Println(err)
	}
	cronJob.Start()
}
