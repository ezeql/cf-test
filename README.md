# cf-test

## Description ##

### [Consumer](https://github.com/ezeql/cf-test/blob/master/back/consumer.go)  ###
A Go http server listen for transactions, parses and validates them and then are pushed into a nsq messaging queue instance.

### [Processor](https://github.com/ezeql/cf-test/blob/master/back/ws.go)  ###
Listen to determinate nsq topics and procceses the messages data for delivering to the frontend. A websocket hub is used in order to mantain a single connection to each client, for all processors involved.

### [Frontend](https://github.com/ezeql/cf-test/blob/master/front/src/components/main.js)  ###
Created with React following Flux arquitecture
Shows the data coming from backend via websockets.


### [Flooder](https://github.com/ezeql/cf-test/blob/master/back/flooder.go)  ###
A small utily for feeding the consumer

## Building ##

### Backend ###

#### Required Software #####

* [golang](https://golang.org/) 
* [nsq](http://nsq.io/) 
* [redis](redis.io/) 


#### Required Go libraries ####
```go get github.com/gin-gonic/gin```

```go get github.com/gorilla/websocket```

```go get github.com/bitly/go-nsq```

```go get github.com/ezeql/go-randomdata```

```go get github.com/garyburd/redigo/redis```

```go get github.com/shopspring/decimal```

FIXME

### frontend ###

#### software ####
* [node and npm](https://nodejs.org/) 
* [grunt](http://gruntjs.com/) 


```npm install```

```grunt serve``` or ```grunt build```


##  List of used libraries ##

### Backend ###

* [gorilla/websocket](https://github.com/gorilla/websocket) 
* [bitly/go-nsq](https://github.com/bitly/go-nsq) NSQ Messaging Queue client
* [gin-gonic/gin](https://github.com/gin-gonic/gin) http framework
* [ezeql/go-randomdata](github.com/ezeql/go-randomdata) 
* [shopspring/decimal](https://github.com/shopspring/decimal) Money handling
* [garyburd/redigo](github.com/garyburd/redigo/redis) redis go client

### Frontend ###

* [react](http://facebook.github.io/react/)
* [flux](https://facebook.github.io/flux/)
* [moment.js](https://facebook.github.io/flux/)
* [fixed-data-table](https://facebook.github.io/flux/)
* [react-bootstrap](https://facebook.github.io/flux/)
* [react-d3](https://facebook.github.io/flux/)
* [react-google-charts](https://facebook.github.io/flux/)
* [react-router](https://facebook.github.io/flux/)
* [react-router-bootstrap](https://facebook.github.io/flux/)
* [react-stockcharts](https://facebook.github.io/flux/)
