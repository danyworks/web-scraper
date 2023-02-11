package cron

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Job interface {
	Execute()
}

type jobTicker struct {
	timer *time.Timer
}

type cronSchedule struct {
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
}

func parseCronSchedule(schedule string) *cronSchedule {
	if ok, err := regexp.MatchString(`^(((\d+)|(\d+-\d+)|(\*\/([1-9]\d*))|(\*)) ?){5}$`, schedule); !ok || err != nil {
		log.Fatalf("Could not parse schedule '%s'", schedule)
	}
	sched := strings.Split(schedule, " ")

	return &cronSchedule{sched[0], sched[1], sched[2], sched[3], sched[4]}
}

func evaluateSlot(slot string, current int) bool {
	if ok, err := regexp.MatchString(`([0-9]+)|(\*/[1-9][0-9]*)|(\*)`, slot); !ok || err != nil {
		log.Printf("Value '%s' does not match cron definition. Error: %v", slot, err)
		return false
	}

	if slot == "*" {
		return true
	}

	if strings.Contains(slot, "/") {
		split := strings.Split(slot, "/")
		interval, err := strconv.Atoi(split[1])
		if err != nil {
			log.Printf("Couldn't parse interval: %v", err)
			return false
		}
		if current%interval == 0 {
			return true
		} else {
			return false
		}
	}

	if strings.Contains(slot, "-") {
		split := strings.Split(slot, "-")
		left, err := strconv.Atoi(split[0])
		if err != nil {
			log.Printf("Couldn't parse slot with value '%s': %v", slot, err)
			return false
		}

		right, err := strconv.Atoi(split[1])
		if err != nil {
			log.Printf("Couldn't parse slot with value '%s': %v", slot, err)
			return false
		}

		for i := left; i <= right; i++ {
			if i == current {
				return true
			}
		}
	} else {
		val, err := strconv.Atoi(slot)
		if err != nil {
			log.Printf("Couldn't parse slot with value '%s': %v", slot, err)
			return false
		}

		if val == current {
			return true
		}
	}

	return false
}

func evaluateSchedule(schedule *cronSchedule) bool {
	return evaluateSlot(schedule.Minute, time.Now().Minute()) &&
		evaluateSlot(schedule.Hour, time.Now().Hour()) &&
		evaluateSlot(schedule.Month, int(time.Now().Month())) &&
		evaluateSlot(schedule.DayOfMonth, time.Now().Day()) &&
		evaluateSlot(schedule.DayOfWeek, int(time.Now().Weekday()))
}

func (t *jobTicker) updateTimer(schedule *cronSchedule) {
	nextTick := time.Now().Add(time.Minute)

	diff := time.Until(nextTick)
	if t.timer == nil {
		t.timer = time.NewTimer(diff)
	} else {
		t.timer.Reset(diff)
	}
}

// minute hour day_of_month month day_of_week
func CronJob(schedule string, job func()) {
	jobTicker := &jobTicker{}
	cronSchedule := parseCronSchedule(schedule)
	jobTicker.updateTimer(cronSchedule)
	for {
		<-jobTicker.timer.C
		jobTicker.updateTimer(cronSchedule)
		if evaluateSchedule(cronSchedule) {
			job()
		}
	}
}
