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

Example:

```bash
$ check_domain -D example.com
OK: Domain will expire in 159 days|expires=2025-08-13T04:00:00Z
```

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

[Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0). Please read the [LICENSE](LICENSE) file.
