package main

import (
  "fmt"
  "log"
  "net/http"
  "net/url"
  "regexp"
  "strings"
)


func handleRoutesWorker( url *url.URL ) (string,bool) {

  eventsRoute, _        := regexp.Compile( "^/events$" )
  teamsByEventRoute, _  := regexp.Compile( "^/event/([a-z0-9_.]+)/teams$")

  if eventsRoute.MatchString( url.Path ) {
    return GetEvents(), true
  } else if m := teamsByEventRoute.FindStringSubmatch( url.Path ); m != nil {
    key := m[1]
    key = strings.Replace( key, "_", "/", -1 )   // replace _ with / e.g. 2014_15 => 2014/15
    return GetTeamsByEvent( key ), true
  } else {
    return "", false  // not route match found
  }
}


func handleRoutes( w http.ResponseWriter, r *http.Request ) {

  // todo: log HTTP verb/method - check query paras too? why? why not?
  log.Println( "url.path: " + r.URL.Path )

  buf, success := handleRoutesWorker( r.URL )

  if success {
    // todo: add mimetype for json too
    // todo: check write to w.write() directly? no need for fmt.Fprintf - why? why not?
    fmt.Fprintf( w, buf )
  } else {
    fmt.Fprintf( w, "no route match found for '" + r.URL.Path + "'" )
  }
} 



func main() {

  // todo: use a variable for ./football.db - pass along ot initDb 

  fmt.Println( "Connecting to ./football.db ..." )
  initDb()
  defer db.Close()

  // testQueries()

  addr := ":9292"
  fmt.Println( "Starting web server listening on " + addr + "..." )
  http.HandleFunc( "/", handleRoutes )
  http.ListenAndServe( addr, nil )

  fmt.Println( "Bye" )
}
