package producers

import (
	"fmt"
	"github.com/rhiaqey/sdk-go/pkg/sdk"
	"log"
	"time"
)

func defaultInterval() uint64 {
	return 1000 // 1 second
}

type TickerSettings struct {
	Interval *uint64 `json:"Interval,omitempty"`
}

type Ticker struct {
	ticker   *time.Ticker
	settings TickerSettings
	channel  chan sdk.ProducerMessage
}

type TickerBody struct {
	Timestamp uint64 `json:"timestamp"`
}

func (t *Ticker) Setup(config interface{}) (chan sdk.ProducerMessage, error) {
	log.Println("setup in progress...")

	settings := config.(TickerSettings)

	channel := make(chan sdk.ProducerMessage)
	t.channel = channel
	t.SetSettings(settings)

	return channel, nil
}

func (t *Ticker) Start() {
	log.Println("starting ticker...")

	for {
		now := uint64(time.Now().UnixNano() / 1e6)
		tag := fmt.Sprintf("%d", now)

		t.channel <- sdk.ProducerMessage{
			Tag:       &tag,
			Key:       "timestamp",
			Value:     TickerBody{Timestamp: now},
			Category:  nil,
			Size:      nil,
			Timestamp: &now,
			UserIds:   nil,
			ClientIds: nil,
		}

		time.Sleep(time.Duration(*t.settings.Interval) * time.Millisecond)
	}
}

func (t *Ticker) SetSettings(config interface{}) {
	log.Println("setting settings")

	settings := config.(TickerSettings)

	if settings.Interval == nil {
		newInterval := defaultInterval()
		settings.Interval = &newInterval
	}

	t.settings = settings
}

func (t *Ticker) Schema() []byte {
	return []byte(`{
		
	}`)
}

func (t *Ticker) Metrics() []byte {
	return []byte(`{
		
	}`)
}

func (t *Ticker) Kind() string {
	return "ticker"
}
