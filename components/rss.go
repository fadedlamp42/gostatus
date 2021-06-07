package components

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type RSS struct {
	width int
	position int
	line string
}

func NewRSS(width int, seperator string) *RSS {
	home, found := os.LookupEnv("HOME")

	var path string
	if found {
		path = home + "/.config/RSS Guard/database/local/database.db"
	} else {
		fmt.Println("WARNING: $HOME not set, using ./database.db")
		path = "./database.db"
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		fmt.Printf("ERROR: could not open rss guard db (%v)\n", err);
		os.Exit(1)
	}

	rows, err := db.Query("SELECT title FROM Messages ORDER BY random()")
	if err != nil {
		fmt.Printf("ERROR: could not execute query (%v)\n", err);
		os.Exit(1)
	}

	titles := []string{}
	var title string

	for rows.Next() {
		rows.Scan(&title)
		titles = append(titles, title)
	}

	line := strings.Join(titles, seperator) + " | "

	return &RSS{
		width: width,
		position: 0,
		line: line,
	}
}

func (r *RSS) Render() string {
	var window string
	if r.position + r.width < len(r.line) {
		window = r.line[r.position:r.position+r.width]
	} else {
		window = r.line[r.position:]
		underflow := r.width - len(window)
		window += r.line[0:underflow]
	}

	r.position++
	if r.position == len(r.line) {
		r.position = 0
	}

	return fmt.Sprintf("RSS: %s", window)
}
