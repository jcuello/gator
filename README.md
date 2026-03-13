# gator
RSS Feed reader which allows you to subscribe to given RSS feeds as different users.

# Requirements
* Postgres SQL
* Go 1.25+

# Install
Install gator with the `go` command line
`go install github.com/jcuello/gator`

# Configuration
You'll need a configuration file in your home directory named `.gatorconfig.json`. It requires a connection string to postgres and a username which you can add with the `gator resgister ...` command
```
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": "joe"
}
```

# Commands
* `register <username>`
* `login <username>`
* `addfeed "<feed name>" <url>`