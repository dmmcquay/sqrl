# sqrl

# Overview
A system inventory program created by HPC academy students designed to gather extensive system information.

This program monitors RAM, OS, CPU, HDD, swap, and network information across systems, logging the information periodically.

* Logs amount of free, total, and used RAM.
* Logs amount of free, total, and used Swap.
* Logs hostname, version, and kernel information.
* Logs CPU, number of cores, speed of CPU, and Model of CPU.
* Logs names, addresses, flags, and speeds of networks. 


# Basic Contents
* 'main.go'
* 'README.md'

# Getting Started
If you have a working Go environment on your computer, you can pull down this repo using:

	> go get github.com/dmmcquay/sqrl

# Usage

## sqrl 

Running 'sqrl [mode][optional flag]' will log various system information.

The mode options are:
* sqrl cpu
* sqrl hdd
* sqrl net
* sqrl ros
* sqrl all

Here is an example:

```
sqrl all 
```
This will log all system information. Using specifc modes will only log that particular information.

You can also output the logged information to the screen using the -v flag.

```
sqrl all -v
```

#Authors:
* Delaney Gill-Sommerhauser
* Danielle Larson
* Christopher Moussa

# Acknowledgements
* Derek McQuay
* github.com/shirou/gopsutil
* github.com/dustin/go-humanize
* github.com/likexian/host-stat-go
* github.com/spf13/cobra
