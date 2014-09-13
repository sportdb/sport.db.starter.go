package main

import (
  "fmt"
  "net/http"
  "regexp"
  "strings"
)


func handleRoutesWorker( url_path string ) (string,bool) {

  eventsRoute, _        := regexp.Compile( "^/events$" )
  teamsByEventRoute, _  := regexp.Compile( "^/event/([a-z0-9_.]+)/teams$")

  if eventsRoute.MatchString( url_path ) {
    return GetEvents(), true
  }

  m := teamsByEventRoute.FindStringSubmatch( url_path )
  if m!= nil {
    key := m[1]
    key = strings.Replace( key, "_", "/", -1 )   // replace _ with /
    return GetTeamsByEvent( key ), true
  }

  return "", false  // not route match found
}


func handleRoutes( w http.ResponseWriter, r *http.Request ) {

  fmt.Println( "url.path: " + r.URL.Path)

  buf, handled := handleRoutesWorker( r.URL.Path)

  if handled {
    // todo: add mimetype for json too
    fmt.Fprintf( w, buf )
  } else {
    fmt.Fprintf( w, "no route match found for '" + r.URL.Path + "'" )
  }
} 



func main() {

  initDb()
  defer db.Close()

  testQueries()

  http.HandleFunc( "/", handleRoutes )
  http.ListenAndServe( ":9292", nil )

  fmt.Println( "Bye" )
}
