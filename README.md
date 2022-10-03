# `zht` - The ZSH History Tool

ZHT aims to make working with your ZSH History easier:

* Search with ease using fixed strings, globbing, or regex
* Sort commands in one or more history files by date
* Deduplicate, or search for duplicates in one or more history files

ZSH is designed to work with ZSH history files using the `EXTENDED_HISTORY`
history format which, according to the man page `zshoptions(1)`, uses the
following format:

```
: <beginning time>:<elapsed seconds>;<command>
```
