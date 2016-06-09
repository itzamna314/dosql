package main

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/docopt/docopt-go"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"os"
)

const (
	usage = `
Usage:
	dosql [ -F configFile ] [ -e environment ] [<script>]

Arguments:
	<script>  The script to run.  If it is not provided, reads from stdin

Options:
	-e=env   Configuration environment [default: default]
	-F=file  Configuration file [default: /usr/local/etc/dosql/config.toml]
`
	version = `dosql 0.0.1`
)

func main() {
	dict, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		fmt.Println("Failed to parse command")
		os.Exit(1)
	}

	cfgFile := dict["-F"].(string)
	cfgEnv := dict["-e"].(string)

	connStr, driver, err := LoadConnectionString(cfgFile, cfgEnv)
	if err != nil {
		fmt.Printf("Failed to load connection string: %s\n", err)
		os.Exit(2)
	}

	db, err := sqlx.Connect(driver, connStr)
	if err != nil {
		fmt.Printf("Failed to connect to database: %s\n", err)
		os.Exit(3)
	}

	script, ok := dict["<script>"].(string)
	if !ok {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Printf("Failed to read from stdin: %s\n", err)
			os.Exit(1)
		}

		script = string(b[:])
	}

	rows, err := db.Queryx(script)
	if err != nil {
		fmt.Printf("Failed to execute query: %s\n", err)
		os.Exit(4)
	}
	defer rows.Close()
	if columns, err := rows.Columns(); err != nil {
		fmt.Printf("Failed to list columns: %s\n", err)
		os.Exit(5)
	} else {
		for _, c := range columns {
			fmt.Printf("%s\t", c)
		}
		fmt.Printf("\n")
	}

	for rows.Next() {
		r, err := rows.SliceScan()
		if err != nil {
			fmt.Printf("Failed to scan result: %s\n", err)
			os.Exit(6)
		}

		for _, c := range r {
			fmt.Printf("%v\t", c)
		}
		fmt.Printf("\n")
	}
}
