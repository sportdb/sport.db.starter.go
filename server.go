package main

import (
  "fmt"
  "log"
  "net/http"
  "net/url"
  "regexp"
  "strings"
  "encoding/json"
  "errors"
)


func handleRoutes( url *url.URL ) (interface{},error) {

  eventsRoute, _        := regexp.Compile( "^/events$" )
  teamsByEventRoute, _  := regexp.Compile( "^/event/([a-z0-9_.]+)/teams$")

  if eventsRoute.MatchString( url.Path ) {
    return GetEvents(), nil
  } else if m := teamsByEventRoute.FindStringSubmatch( url.Path ); m != nil {
    key := m[1]   // capture group 1 is event key
    key = strings.Replace( key, "_", "/", -1 )   // replace _ with / e.g. 2014_15 => 2014/15    
    return GetTeamsByEvent( key ), nil
  } else {
    return nil, errors.New( "No route match found for '" + url.Path + "'" )
  }
}


func handleFunc( w http.ResponseWriter, r *http.Request ) {

  // todo: log HTTP verb/method - check query paras too? why? why not?
  log.Println( "url.path: " + r.URL.Path )

  data, err := handleRoutes( r.URL )

  if err != nil {
    fmt.Fprintf( w, err.Error() )  // error - no route match found
  } else {
    // todo: add mimetype for json too
    // todo: check write to w.write() directly? no need for fmt.Fprintf - why? why not?

    b, err := json.Marshal( data )
    checkErr( err )
    log.Println( "json: ", string(b) )
    fmt.Fprintf( w, string(b) )
  }
} 



func main() {

  // todo: use a variable for ./football.db - pass along ot initDb 

  fmt.Println( "Connecting to ./football.db ..." )
  InitDb()
  defer db.Close()

  addr := ":9292"
  fmt.Println( "Starting web server listening on " + addr + "..." )
  fmt.Println( "Use Ctrl-C to stop" )

  http.HandleFunc( "/", handleFunc )
  http.ListenAndServe( addr, nil )

  fmt.Println( "Bye" )
}
