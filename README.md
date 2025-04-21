# GO assertion lib

A library to add __Tiger Style__ assertions to your go code.

> Assertions detect programmer errors. Unlike operating errors, which are expected and which must be handled, assertion failures are unexpected. The only correct way to handle corrupt code is to crash. Assertions downgrade catastrophic correctness bugs into liveness bugs. Assertions are a force multiplier for discovering bugs by fuzzing. (src: [https://github.com/tigerbeetle/tigerbeetle/blob/main/docs/TIGER_STYLE.md](https://github.com/tigerbeetle/tigerbeetle/blob/main/docs/TIGER_STYLE.md))

## Features

### Mode / DefaultMode

* `Panic`  (default) => print the stack trace then `panic()` 
* `Exit` =>  print the stack trace then `os.Exit(1)` 
* `SIGTERM` =>  print the stack trace then find the current process and send `syscall.SIGTERM` to itself (and panic the current goroutine)  
* `SKIP` => just print the stack trace (but no panic, exit or SIGTERM)

You can set the mode for each assertion:
```go
// sigterm assertion 
assert.NoErr(err, assert.SIGTERM)

// use default mode
assert.NoErr(err, assert.DefaultMode)
```

Or you can override the DefaultMode globally:

```go
// override DefaultMode
assert.DefaultMode = assert.Exit
```


### Color

```go
// set isTTY to true to
// add colors to output
assert.IsTTY = true
  
```


