package main

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "log"
)


func checkErr(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

// note: "package global" db object
var db *sql.DB


func InitDb() {
  var err error
  db, err = sql.Open( "sqlite3", "./football.db" )
  checkErr( err )
}


func FetchTeamsByEvent( event Event ) []Team {

  query :=
  `SELECT
      t.[key],
      t.title,
      t.code
   FROM teams t
        INNER JOIN events_teams et ON et.team_id = t.id
        INNER JOIN events e ON e.id = et.event_id
   WHERE e.[key] = ?`


  rows, err := db.Query( query, event.Key )
  checkErr( err )
  defer rows.Close()

  teams := []Team{}

  for rows.Next() {
    var r Team
    err = rows.Scan( &r.Key, &r.Title, &r.Code )
    checkErr( err )

    teams = append( teams, r ) // add new row
  }
  return teams
}



func FetchEventByKey( key string ) Event {

  query :=
  `SELECT
      e.[key]                    AS event_key,
      l.title || ' ' || s.title  AS event_name
   FROM events e
        INNER JOIN seasons s ON s.id = e.season_id
        INNER JOIN leagues l ON l.id = e.league_id
   WHERE e.[key] = ?`

  var r Event

  err := db.QueryRow( query, key ).Scan( &r.Key, &r.Title )
  checkErr( err )

  return r
}


func FetchEvents() []Event {

  query :=
  `SELECT
      e.[key]                    AS event_key,
      l.title || ' ' || s.title  AS event_name
   FROM events e
        INNER JOIN seasons s ON s.id = e.season_id
        INNER JOIN leagues l ON l.id = e.league_id`

  rows, err := db.Query( query )
  checkErr( err )
  defer rows.Close()

  events := []Event{}

  for rows.Next() {
    var r Event
    err = rows.Scan( &r.Key, &r.Title )
    checkErr( err )

    events = append( events, r ) // add new row
  }
  return events
}

