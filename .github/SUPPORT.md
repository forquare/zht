# Support and Help

Need help getting started? Here's how!

* There is a general guide on how to use ZHT in the [README](/README.md).
* ZHT comes with documentation built in!
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
      split       Split history into multiple files
      tail        Print lines from the bottom of history
      uniq        Report or filter out repeated commands
      version     Print the version number of ZHT
    
    Flags:
      -h, --help                       help for zht
      -F, --human-date-format string   Specify the format for printing the data/time for each command, implies -t.  See https://pkg.go.dev/time#pkg-constants for available formats
      -t, --human-readable-date        Print the epoch time stamp in a human readable format
    
    Use "zht [command] --help" for more information about a command.
  ```
  You can get help on specific commands like this:
  ```shell
    %>zht help search
    Search through ZSH history using regular expressions.
    
    By default, no regex is used and the pattern is taken to be exactly as is.
    When using regex mode, by default zsh_history will use RE2 (https://github.com/google/re2/wiki/Syntax).
    POSIX compatible Extended Regular Expressions (ERE) can also be used.
    
    Usage:
      zht search [flags] <pattern> [file ...]
    
    Aliases:
      search, find
    
    Flags:
      -g, --glob string          Search using globbing
      -h, --help                 help for search
      -S, --no-ignore-self zht   When searching history, do not ignore lines that begin with zht
      -R, --posix string         Search using POSIX ERE regular expressions
      -r, --regex string         Search using regular expressions
    
    Global Flags:
      -F, --human-date-format string   Specify the format for printing the data/time for each command, implies -t.  See https://pkg.go.dev/time#pkg-constants for available formats
      -t, --human-readable-date        Print the epoch time stamp in a human readable format
  ```
* If the documentation isn't sufficient please raise an Issue
* If you know what isn't right and how to solve it please raise a Pull Request!