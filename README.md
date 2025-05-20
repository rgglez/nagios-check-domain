[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![GitHub all releases](https://img.shields.io/github/downloads/rgglez/nagios-check-domain/total)
![GitHub issues](https://img.shields.io/github/issues/rgglez/nagios-check-domain)
![GitHub commit activity](https://img.shields.io/github/commit-activity/y/rgglez/nagios-check-domain)
[![Go Report Card](https://goreportcard.com/badge/github.com/rgglez/nagios-check-domain)](https://goreportcard.com/report/github.com/rgglez/nagios-check-domain)
[![GitHub release](https://img.shields.io/github/release/rgglez/nagios-check-domain.svg)](https://github.com/rgglez/nagios-check-domain/releases/)

# nagios-check-domain

**check_domain** is a nagios plugin written in [Go](https://go.dev/) to check the
expiration date of a given domain and notify if it is about to expire.

This plugin queries public [whois](https://www.rfc-editor.org/rfc/rfc3912.txt) servers. It
tries its best to query the right server for the publix suffix or TLD.

## Command line options

* `--domain` `-D` string, the domain name to check.
* `--warn` `-w` integer, the number of days after which a warning will be considered a warning condition. Default: 30.
* `--crit` `-c` integer, the number of days after which a warning will be considered a critical condition. Default: 15.
* `--servers` `-s` string, the path to the file containing the list of WHOIS servers.

## Build and installation

### Build

* Get the source code:

```bash
$ git clone https://github.com/rgglez/nagios-check-domain.git
```

* Compile the code:

```bash
$ cd nagios-check-domain
$ make build
```

### Installation

To install the binary to the default path (```/usr/local/nagios/libexec```), execute:

```bash
# make install
```

Or just copy the executable to your regular Nagios plugins directory.

## Execution

Basic example using default server (whois.iana.org):

```bash
check_domain -D example.com
OK: Domain will expire in 159 days|expires=2025-08-13T04:00:00Z
```

Using the `servers.json` file:

```bash
check_domain -D example.com --servers=/path/to/servers.json
```

## Server list

A list of WHOIS servers is included in the [data/servers.json](data/servers.json) file.
This is a JSON file which has the [TLD](https://en.wikipedia.org/wiki/Top-level_domain) 
as the key and the corresponding WHOIS server as the value.

You can provide your own file. See the command line options above.

## Dependencies

This program has the following external dependencies:

* [github.com/likexian/whois](https://github.com/likexian/whois)
* [github.com/likexian/whois-parser](https://github.com/likexian/whois-parser)
* [github.com/spf13/pflag](https://github.com/spf13/pflag)
* [github.com/xorpaul/go-nagios](https://github.com/xorpaul/go-nagios)
* [github.com/ztrue/tracerr](https://github.com/ztrue/tracerr)
* [golang.org/x/net/publicsuffix](golang.org/x/net/publicsuffix)

## License

Copyright 2025 Rodolfo González González.

[GPL v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html). Please read the [LICENSE](LICENSE.md) file.
