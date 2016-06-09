# dosql
Simple go cli to execute sql commands

# usage
```
Usage:
	dosql [ -F configFile ] [ -e environment ] <script>

Options:
	-e=env   Configuration environment [default: default]
	-F=file  Configuration file [default: /usr/local/etc/dosql/config.toml]
```

# setup
dosql reads a toml config file to connect to your database.  By default, the file is located at `/usr/local/etc/dosql/config.toml`.  You can specify an alternate file using -F.
