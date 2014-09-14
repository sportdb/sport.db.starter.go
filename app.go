package main

import (
  "log"
)


func GetEvents() interface{} {

  // step 1: fetch records

  events := FetchEvents()
  log.Println( events )

  // step 2: map to json structs for serialization/marshalling

  type JsEvent struct {
      Key     string   `json:"key"`
      Title   string   `json:"title"`
  }
  data := []*JsEvent{}

  for _,event := range events {
    data = append( data,
                   &JsEvent {
                      Key:   event.Key,
                      Title: event.Title, } )
  }

  return data
}



func GetTeamsByEvent( eventKey string ) interface{} {

  // step 1: fetch records

  event := FetchEventByKey( eventKey )
  log.Println( event )

  teams := FetchTeamsByEvent( event )
  log.Println( teams )

  // step 2: map to json structs for serialization/marshalling

  type JsTeam struct {
      Key     string  `json:"key"`
      Title   string  `json:"title"`
      Code    string  `json:"code"`
  }

  type JsEvent struct {
      Key     string  `json:"key"`
      Title   string  `json:"title"`
      Teams []*JsTeam `json:"teams"`
  }

  data := JsEvent {
    Key:   event.Key,
    Title: event.Title,
  }

  for _,team := range teams {
    data.Teams = append( data.Teams,
                         &JsTeam{ Key:   team.Key,
                                  Title: team.Title,
                                  Code:  team.Code, } )
  }


  return data
}

