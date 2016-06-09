# dosql
Simple cli to execute sql commands

# usage
```
Usage:
	dosql [ -F configFile ] [ -e environment ] <script>

Options:
	-e=env   Configuration environment [default: default]
	-F=file  Configuration file [default: /usr/local/etc/dosql/config.toml]
```

# setup
dosql reads a toml config file to connect to your database.  By default, the file is located at `/usr/local/etc/dosql/config.toml`.  You can specify an alternate file using -F.  Each toml file may have multiple environments.  You can specify which environment to use with -e.

# sample config file
```
[default]
server="foo.bar.com"
port="1433"
database="MyDatabase"
user_id="iamauser"
password="password"
encrypt="true"
trust_server_certificate="true"
connection_timeout="30"
driver="mssql"
[another_db]
server="bat.baz.com"
port="1433"
database="MyOtherDatabase"
user_id="iamanotheruser"
password="thesamepasswordonoes!"
encrypt="true"
trust_server_certificate="true"
connection_timeout="30"
driver="mssql"
```

# drivers
Currently, only the mssql driver is supported.  Adding new driver support should be a simple matter of importing the desired driver and building a correctly-formatted connection string in config.go.  If you would like to add support for a new driver, please submit a pull request.
