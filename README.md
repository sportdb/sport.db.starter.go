# sport.db.starter - Build Your Own HTTP JSON API (Go Edition)

The sportdb web service starter sample lets you build your own HTTP JSON API
using the
[`football.db`](https://github.com/openfootball),
[`worldcup.db`](https://github.com/openfootball/world-cup),
[`ski.db`](https://github.com/opensport/ski.db),
[`formulal1.db`](https://github.com/opensport/formula1.db), etc.


## Getting Started

Step 1: Build yourself a binary using go. Type:

    $ go build server.go app.go database.go

Step 2: Copy an SQLite database e.g. `football.db` into your folder.

Step 3: Startup the web service (HTTP JSON API). Type:

    $ ./server

That's it. Open your web browser and try some services
running on your machine on port 9292 (e.g. `localhost:9292`). Example:



List all events:

- `http://localhost:9292/events`

List all World Cup Brazil 2014 teams:

- `http://localhost:9292/event/world.2014/teams`


And so on. Now change the sources to fit your needs. Be bold. Enjoy.


## License

The `sportdb.db.starter` scripts are dedicated to the public domain.
Use it as you please with no restrictions whatsoever.


## Questions? Comments?

Send them along to the
[Open Sports & Friends Forum/Mailing List](http://groups.google.com/group/opensport).
Thanks!
