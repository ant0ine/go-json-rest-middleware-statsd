Statsd Middleware for Go-Json-Rest
==================================

*Currently not used in production, API subject to change*

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ant0ine/go-json-rest-middleware-statsd) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/ant0ine/go-json-rest-middleware-statsd/master/LICENSE)

This is a middleware for [Go-Json-Rest](https://github.com/ant0ine/go-json-rest).
It uses [g2s](https://github.com/peterbourgon/g2s) to send statistics about the current request/response to a statsd server.
It depends on `rest.TimerMiddleware` and `rest.RecorderMiddleware` that should be in the wrapped middlewares.

The two metrics are in the form:
* `[<Prefix>.]response.status_code.<StatusCode>`: Counter.
* `[<Prefix>.]response.elasped_time`: Timer.

Copyright (c) 2015 Antoine Imbert

[MIT License](https://github.com/ant0ine/go-json-rest-examples/blob/master/LICENSE)

[![Analytics](https://ga-beacon.appspot.com/UA-309210-4/go-json-rest-middleware-statsd/master/readme)](https://github.com/igrigorik/ga-beacon)

