package main

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "log"
)


func checkErr(err error) {
  if err != nil {
    log.Fatalln(err)
  }
}

// note: "package global" db object
var db *sql.DB


func initDb() {
  var err error
  db, err = sql.Open( "sqlite3", "./football.db" )
  checkErr( err )
}


func FetchTeamsByEvent( event EventRow ) []TeamRow {

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

  columns, err := rows.Columns()
  checkErr( err )
  log.Println( columns )

  teams := []TeamRow{}

  for rows.Next() {
    var r TeamRow
    err = rows.Scan( &r.Key, &r.Title, &r.Code )
    checkErr( err )

    teams = append( teams, r ) // add new row
  }
  rows.Close()
  return teams
}



func FetchEventByKey( key string ) EventRow {
  query :=
  `SELECT
      e.[key]                    AS event_key,
      l.title || ' ' || s.title  AS event_name
   FROM events e
        INNER JOIN seasons s ON s.id = e.season_id
        INNER JOIN leagues l ON l.id = e.league_id
   WHERE e.[key] = ?`

  var r EventRow
  
  err := db.QueryRow( query, key ).Scan( &r.Key, &r.Title )
  checkErr( err )

  return r
}


func FetchEvents() []EventRow {
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

  events := []EventRow{}

  for rows.Next() {
    var r EventRow
    err = rows.Scan( &r.Key, &r.Title )
    checkErr( err )

    events = append( events, r ) // add new row
  }
  rows.Close()
  return events
}

