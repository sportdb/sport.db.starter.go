package main

import (
  "log"
  "encoding/json"
)


func testQueries() {
  events := GetEvents()
  log.Println( events )

  teams := GetTeamsByEvent( "de.2014/15" )
  log.Println( teams )

  teams = GetTeamsByEvent( "de.2.2014/15" )
  log.Println( teams )
}
