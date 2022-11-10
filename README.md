# `zht` - The ZSH History Tool

ZHT aims to make working with your ZSH History easier:

* Search with ease using fixed strings, globbing, or regex
* Sort commands in one or more history files by date
* Deduplicate, or search for duplicates in one or more history files

## Installing

### macOS

Simply use [brew](https://brew.sh/):

```
brew install forquare/tools/zht
```

This works for Intel and ARM based Macs.

### Linux

#### Arch

ZHT is available in the aur!

If you have [yay](https://github.com/Jguer/yay) installed simply:

```shell
yay -S zht-bin
```

Otherwise you can install it manually:

```shell
git clone https://aur.archlinux.org/zht-bin.git
cd zht-bin
makepkg -is
```

#### Homebrew for Linux

[brew](https://brew.sh/) also works for Linux too!

```
brew install forquare/tools/zht
```

### Other Systems

Go to the [latest release](https://github.com/forquare/zht/releases/latest), download the relevant archive for your OS, unarchive it, and place it somewhere on your `PATH`.

## Prerequisites

ZSH is designed to work with ZSH history files that use the `EXTENDED_HISTORY`
history format which, according to the man page `zshoptions(1)`, uses the
following format:

```
: <beginning time>:<elapsed seconds>;<command>
```

## How to use ZHT

### Simple searching

Search for a simple string, like "`foo`":

```shell
%>zht search 'foo'
: 1665325783:0;cat /tmp/foo
: 1667142921:0;cat << EOF\
foo\
EOF
```

Two results are returned here, the first is a simple `cat /tmp/foo` command.  The second result is more complex and ZHT has searched across multiple lines of the history entry.

### Searching with regex

For more complex searching you can use [regex](https://github.com/google/re2/wiki/Syntax):

```shell
%>zht search -r '^curl.*github' 
: 1667143560:1;curl -s 'https://api.github.com/repos/forquare/zht/releases/latest' | jq -r '.assets[]'
: 1667143578:0;curl -s 'https://api.github.com/repos/forquare/zht/releases/latest' | jq -r '.assets[0].browser_download_url'
```

Regex searching will also search inside multiline commands.

### Using a different history file

By default ZHT will use `.zsh_history` in your home directory, however you can specify other files:

```shell
%>zht search 'bash' /tmp/tmp-hist
: 1588091653:0;rm .bash*
```

### Input from a pipe

If you have multiple history files, need to do some preprocessing in a different tool, **or want to chain together multiple ZHT subcommands** ZHT will allow you to pipe in one or more history files:

```shell
%>zht search -r '^curl.*github' | zht tail -n 3
: 1667143557:0;curl -s 'https://api.github.com/repos/forquare/zht/releases/latest' | jq -r '.assets[0].browser_download_url'
: 1667143560:1;curl -s 'https://api.github.com/repos/forquare/zht/releases/latest' | jq -r '.assets[]'
: 1667143578:0;curl -s 'https://api.github.com/repos/forquare/zht/releases/latest' | jq -r '.assets[0].browser_download_url'

%>zht search -r '^curl.*github' | zht tail -n 3 | zht uniq -d
: 1667143560:1;curl -s 'https://api.github.com/repos/forquare/zht/releases/latest' | jq -r '.assets[]'
: 1667143578:0;curl -s 'https://api.github.com/repos/forquare/zht/releases/latest' | jq -r '.assets[0].browser_download_url'
```

Above there is a command showing the output of a search, then limiting the output to the last three lines.  In the second command `zht uniq` is used to filter out duplicates.

### Switching date formats

The first field of each command is a [UNIX timestamp](https://en.wikipedia.org/wiki/Unix_time), and not very human-readable!  ZHT has a built-in universal flag to make it a little easier to read:

```shell
%>zht --human-readable-date search -r '^curl.*github.*assets\[\]'
: Sun, 30 Oct 2022 15:26:00 GMT:1;curl -s 'https://api.github.com/repos/forquare/zht/releases/latest' | jq -r '.assets[]'
```

Furthermore, you can specify your own format using the [Go time package constants](https://pkg.go.dev/time#pkg-constants):

```shell
%>zht --human-date-format 'Jan _2 15:04:05' search -r '^curl.*github.*assets\[\]'
: Oct 30 15:26:00:1;curl -s 'https://api.github.com/repos/forquare/zht/releases/latest' | jq -r '.assets[]'
```

### Getting help

ZHT has built-in help!

```shell
%>zht help
zht is a tool to parse and manipulate Z Shell History.

Primarily it is can be used to sort a history if two history files have been merged.

Usage:
  zht [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  head        Print lines from the top of history
  help        Help about any command
  parse       Run history through parser and print
  search      Search history using regular expressions
  sort        Sort one or more history files
  tail        Print lines from the bottom of history
  uniq        Report or filter out repeated commands
  version     Print the version number of ZHT

Flags:
  -h, --help                       help for zht
  -F, --human-date-format string   Specify the format for printing the data/time for each command, implies -t.  See https://pkg.go.dev/time#pkg-constants for available formats
  -t, --human-readable-date        Print the epoch time stamp in a human readable format

Use "zht [command] --help" for more information about a command.
```

You can also get help with a specific sub command:

```shell
%>zht help uniq
Parse all provided history files (or ~/.zsh_history by default) and prints out a uniq history.

By default will only match adjacent duplicates, if a command occurs more than once in another part of the history file it will not get filtered out (see --unique to stop this).

Usage:
  zht uniq [flags] [file ...]

Aliases:
  uniq, unique

Flags:
  -D, --all-repeated   Print all entries that are repeated.
  -c, --count          Display the number of times each entry occurs - omits date and duration.
  -h, --help           help for uniq
  -d, --repeated       Print a single copy of each entry that is repeated - the last occurrence will be printed.
  -s, --sort           Sort input by date before making output uniq. When used with -c it will sort commands (ascending) by the number of times the occur.
  -u, --unique         Only print entries that are not repeated.

Global Flags:
  -F, --human-date-format string   Specify the format for printing the data/time for each command, implies -t.  See https://pkg.go.dev/time#pkg-constants for available formats
  -t, --human-readable-date        Print the epoch time stamp in a human readable format
```

## Contributing

Please read [CONTRIBUTING.md](.github/CONTRIBUTING.md) for details on the code of conduct, and the process for submitting Pull Requests.

## Versioning

[Semantic Versioning](http://semver.org/) is used for versioning.

## Authors

  - **Ben Lavery-Griffiths** - [Forquare](https://github.com/forquare) - [Website](https://hashbang0.com)

See also the list of [contributors](https://github.com/forquare/zht/contributors) who have participated in this project.

## License

This project is licensed under the [MIT Licence](LICENSE) - see the [LICENSE](LICENSE) file for details

## Releasing

Releasing is done via a Github action that uses [GoReleaser](https://goreleaser.com/).

1. Apply code changes via Pull Requests
2. Update local main branch to latest
3. Create an annotated tag with the new version  
  ```git tag -a x.x.z -m "Escriptive message"```
4. Push the tag  
  ```git push --tags```
5. The release action will run using GoReleaser to create the next release

All tags are [protected](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/managing-repository-settings/configuring-tag-protection-rules) so that only certain users can create them.
