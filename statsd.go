package statsd

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/peterbourgon/g2s"
	"log"
	"strconv"
	"time"
)

// StatsdMiddleware is a Go-Json-Rest Middleware. It sends statistic about the current
// request/response to a statsd server. It depends on rest.TimerMiddleware and
// rest.RecorderMiddleware that should be in the wrapped middlewares.
// The two metrics are in the form:
// "[<Prefix>.]response.status_code.<StatusCode>": Counter.
// "[<Prefix>.]response.elasped_time": Timer.
type StatsdMiddleware struct {

	// IP and port of the statsd server. Optional. Default to "127.0.0.1:8125".
	IpPort string

	// Prefix added to the metric keys. Optional.
	Prefix string
}

// MiddlewareFunc makes StatsdMiddleware implement the rest.Middleware interface.
func (mw *StatsdMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {

	ipPort := mw.IpPort
	if ipPort == "" {
		ipPort = "127.0.0.1:8125"
	}

	statsd, err := g2s.Dial("udp", ipPort)
	if err != nil {
		log.Fatal(err)
	}

	keyBase := ""
	if mw.Prefix != "" {
		keyBase += mw.Prefix + "."
	}
	keyBase += "response."

	return func(writer rest.ResponseWriter, request *rest.Request) {

		handler(writer, request)

		if request.Env["STATUS_CODE"] != nil {
			statusCode := request.Env["STATUS_CODE"].(int)
			statsd.Counter(1.0, keyBase+"status_code."+strconv.Itoa(statusCode), 1)
		}

		if request.Env["ELAPSED_TIME"] != nil {
			elapsedTime := request.Env["ELAPSED_TIME"].(*time.Duration)
			statsd.Timing(1.0, keyBase+"elapsed_time", *elapsedTime)
		}
	}
}
