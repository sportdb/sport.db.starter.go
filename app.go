package main

import (
  "log"
  "encoding/json"
)


func GetEvents() string {

  type Event struct {
    Key     string   `json:"key"`
    Title   string   `json:"title"`
  }
  
  eventRows := FetchEvents()
  log.Println( eventRows )
  
  events := []Event{}

  for _,row := range eventRows {
    event := &Event {
      Key:   row.Key,
      Title: row.Title,
    }
    events = append( events, *event )
  }

  buf, err := json.Marshal( events )
  checkErr( err )
  return string(buf)
}



func GetTeamsByEvent( eventKey string ) string {

  type Team struct {
     Key     string  `json:"key"`
     Title   string  `json:"title"`
     Code    string  `json:"code"`
  }

  type Event struct {
     Key     string `json:"key"`
     Title   string `json:"title"`
     Teams[] Team   `json:"teams"`
  }

  eventRow := FetchEventByKey( eventKey )
  log.Println( eventRow )

  teamRows := FetchTeamsByEvent( eventRow )
  log.Println( teamRows )

  teams := []Team{}

  for _,row := range teamRows {
    team := &Team {
      Key:   row.Key,
      Title: row.Title,
      Code:  row.Code,
    }
    teams = append( teams, *team )
  }

  data := &Event {
    Key:   eventRow.Key,
    Title: eventRow.Title,
    Teams: teams,
  }  

  buf, err := json.Marshal( data )
  checkErr( err )
  return string(buf)
}


func testQueries() {
  events := GetEvents()
  log.Println( events )

  teams := GetTeamsByEvent( "de.2014/15" )
  log.Println( teams )

  teams = GetTeamsByEvent( "de.2.2014/15" )
  log.Println( teams )
}

