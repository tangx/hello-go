package main

import (
	"log"
	"testing"

	"github.com/robfig/cron/v3"
)

func main() {
	i := 0
	// cr := cron.New(cron.WithSeconds())
	cr := cron.New()
	spec := "* * * * * ?"
	id, err := cr.AddFunc(spec, func() {
		i++
		log.Println("logger ", i)
	})

	if err != nil {
		log.Panic(err)
	}
	log.Println("EntryID", id)
	cr.Run()
}

func TestMain(t *testing.T) {
	main()
}

func main2() {
	i := 0
	p := cron.WithParser(cron.NewParser(
		cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow,
	))

	c := cron.New(p)
	spec := `* * * * *`

	_, _ = c.AddFunc(spec, func() {
		i++
		log.Printf("main2 %d\n", i)
	})

	log.Println("start")
	c.Run()
}

func TestMain2(t *testing.T) {
	main2()
}
