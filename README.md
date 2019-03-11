Test code for Readline() cleanup issue.
=======

## Summary

Readline() looks like cleaning up the last line
when it does not include a \n at the end.

## Build

```
$ go build
```

## Run

```
$ ./rltest
```

"This line will be cleaned up after 10 seconds..." will be printed
but cleaned up after 10 seconds.
