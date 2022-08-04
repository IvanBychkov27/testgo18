/*  Будильник с множеством установок
    Пример использования функции Cron

Cron - это команда Linux для запуска скриптов по расписанию.

* * * * * *
- – – – – –
| | | | | |
| | | | | +—– day of week (0 – 7) (Sunday=0)
| | | | +——- month (1 – 12)
| | | +——— day of month (1 – 31)
| | +———– hour (0 – 23)
| +————- min (0 – 59)
+————- sec (0 – 59)
*/

package main

import (
	"context"
	cronPkg "github.com/robfig/cron/v3"
	"log"
	"os/signal"
	"sync"
	"syscall"
)

type Application struct {
	ch1 chan struct{}
	ch2 chan struct{}
}

func newApplication() *Application {
	return &Application{
		ch1: make(chan struct{}, 1),
		ch2: make(chan struct{}, 1),
	}
}

func main() {
	log.Println("Start AlarmClock")

	app := newApplication()

	err := app.run()
	if err != nil {
		log.Println("error:", err.Error())
	}
	log.Println("Done...")
}

func (app *Application) run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go app.processing(ctx, wg)
	go app.startCron(ctx, wg, "*/2 * * * * *", "*/5 * * * * *")

	<-ctx.Done()

	wg.Wait()

	return nil
}

func (app *Application) processing(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-app.ch1:
			log.Println("AlarmClock 1 ok")
		case <-app.ch2:
			log.Println("AlarmClock 2 ok")
		case <-ctx.Done():
			return
		}
	}
}

func (app *Application) startCron(ctx context.Context, wg *sync.WaitGroup, timeAlarmClock ...string) {
	defer wg.Done()

	//cron := cronPkg.New() // по минутный
	cron := cronPkg.New(cronPkg.WithSeconds()) // по секундный
	_, errCron := cron.AddFunc(timeAlarmClock[0], app.alarmClock1)
	if errCron != nil {
		log.Println("error add func to cron", errCron.Error())
		return
	}

	_, errCron = cron.AddFunc(timeAlarmClock[1], app.alarmClock2)
	if errCron != nil {
		log.Println("error add func to cron", errCron.Error())
		return
	}

	cron.Start()

	<-ctx.Done()

	cron.Stop()
}

func (app *Application) alarmClock1() {
	app.ch1 <- struct{}{}
}

func (app *Application) alarmClock2() {
	app.ch2 <- struct{}{}
}
