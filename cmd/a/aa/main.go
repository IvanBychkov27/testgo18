package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Application struct {
	metricsDataMx sync.RWMutex
	metricsData   map[string]*Item
	pool          *PoolData
	poolData      map[string]*Item
}

type Item struct {
	Value     int64
	Timestamp int64
}

func NewApp() *Application {
	app := &Application{
		metricsData: map[string]*Item{},
		pool:        NewPoolData(),
	}
	app.poolData = app.pool.Data[0]
	return app
}

type PoolData struct {
	DataMx sync.RWMutex
	Data   []map[string]*Item
	count  int
}

func NewPoolData() *PoolData {
	//p := &PoolData{
	//	Data: make([]map[string]*Item, 0, 2),
	//}
	//
	//p.Data = append(p.Data, map[string]*Item{"1": {1, 1}})
	//p.Data = append(p.Data, map[string]*Item{"2": {2, 2}})

	//return p

	return &PoolData{
		Data: []map[string]*Item{{"1": {1, 1}}, {"2": {2, 2}}},
	}
}

func (p *PoolData) Next() map[string]*Item {
	p.count++
	return p.Data[p.count%2]
}

// задать новую базу метрик
func (p *PoolData) Set() map[string]*Item {
	p.count++
	return p.Data[p.count%2]
}

// получить текущую базу метрик
func (p *PoolData) Get() map[string]*Item {
	return p.Data[p.count%2]
}

// очистить текущую базу метрик
func (p *PoolData) Close() {
	p.Data[p.count%2] = map[string]*Item{}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	app := NewApp()

	interval := 2

	wg := &sync.WaitGroup{}
	wg.Add(1)
	app.saveMetricsData(ctx, wg, interval)

	<-ctx.Done()
	wg.Wait()

	fmt.Println("Done")
}

func (app *Application) saveMetricsData(ctx context.Context, wg *sync.WaitGroup, interval int) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second * time.Duration(interval))
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			app.inc()

			//for k, v := range app.poolData {
			//	fmt.Printf("pool data: [%s] = %d \n", k, v.Value)
			//}
			data := app.pool.Get()
			for k, v := range data {
				fmt.Printf("pool data: [%s] = %d \n", k, v.Value)
			}
			fmt.Println(data)
			app.poolData = app.pool.Set()
		}
	}
}

func (app *Application) inc() {
	for k, v := range app.poolData {
		if k == "1" {
			v.Value++
			app.poolData[k] = v
		}
		if k == "2" {
			v.Value += 10
			app.poolData[k] = v
		}
	}
}
