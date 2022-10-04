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

# Installing

## macOS

Simply use [brew](https://brew.sh/):

```
brew install forquare/tools/zht
```

This works for Intel and ARM based Macs.

## Linux

[brew](https://brew.sh/) also works for Linux too!

```
brew install forquare/tools/zht
```

## Other Systems

Go to the [latest release](https://github.com/forquare/zht/releases/tag/0.1.1-test), download the relevant archive for your OS, unarchive it, and place it somewhere on your PATH.
