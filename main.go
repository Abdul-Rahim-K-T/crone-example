package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// CustomJob is a custom job struct
type CustomJob struct {
	Name string
}

// Run implements the cron.Job Interface
func (cj CustomJob) Run() {
	fmt.Printf("Running job: %s at %s\n", cj.Name, time.Now())
}

func main() {
	c := cron.New()

	// Add a job with a different time zone
	c.AddFunc("* * * * *", func() {
		fmt.Println("Running job every minute in UTC+8:", time.Now())
	})

	// Add a custome job
	entryID, _ := c.AddJob("* * * * *", CustomJob{Name: "My Custom Job"})

	c.Start()

	// List scheduled jobs
	for _, entry := range c.Entries() {
		fmt.Printf("ID: %d, Next: %v, Prev: %v\n", entry.ID, entry.Next, entry.Prev)
	}

	// Simulate some work
	time.Sleep(2 * time.Minute)

	// Remove the custom job
	c.Remove(entryID)
	fmt.Println("Custome job removed")

	select {}
}
