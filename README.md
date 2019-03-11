Test code for Readline() cleanup issue.
=======

## Summary

Readline() looks like cleanup up the last line
when it does not include the tailing \n.

## Build

'''
$ go build
'''

## Run

'''
$ ./rltest
'''

"This line will be cleaned up after 10 seconds..." will be printed
but cleaned up after 10 seconds.
