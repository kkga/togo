todo.txt CLI app written in Go. Not very usable yet.

## Usage

```
Usage:
  togo [command]

Available Commands:
  add         Add todo
  clean       Move done todos to done.txt
  do          Mark todo as done
  help        Help about any command
  ls          List todos
  rm          Remove todo

Flags:
  -h, --help   help for togo

Use "togo [command] --help" for more information about a command.
```

## Examples

### Listing todos

```
> togo ls
 1 [ ] 2020-12-03 Add file configuration
 2 [ ] 2020-12-03 Add support for priorities and `pri`/`depri` command
 3 [ ] 2020-12-03 Improve output for `do` cmd, show final done state
 4 [ ] 2020-12-02 call ls in root cmd if no args
 5 [ ] 2020-12-03 implement viper config
------
5/5 todos shown
```

```
> togo ls config
 1 [ ] 2020-12-03 Add file configuration
 5 [ ] 2020-12-03 implement viper config
------
2/5 todos shown
```

### Adding todos

```
> togo add put some examples on github
Added: put some examples on github
```

```
> togo ls
 1 [ ] 2020-12-03 Add file configuration
 2 [ ] 2020-12-03 Add support for priorities and `pri`/`depri` command
 3 [ ] 2020-12-03 Improve output for `do` cmd, show final done state
 4 [ ] 2020-12-02 call ls in root cmd if no args
 5 [ ] 2020-12-03 implement viper config
 6 [ ] 2020-12-03 put some examples on github
------
6/6 todos shown
```

### Completing todos

```
> togo do 6
Marked done:
- [x] put some examples on github
```

```
> togo ls
 1 [ ] 2020-12-03 Add file configuration
 2 [ ] 2020-12-03 Add support for priorities and `pri`/`depri` command
 3 [ ] 2020-12-03 Improve output for `do` cmd, show final done state
 4 [ ] 2020-12-02 call ls in root cmd if no args
 5 [ ] 2020-12-03 implement viper config
 6 [x] 2020-12-03 2020-12-03 put some examples on github
```

### Cleaning done todos

```
> togo cl
Archived:
- [x] 2020-12-03 2020-12-03 put some examples on github
```

```
> cat done.txt
x 2020-12-03 2020-12-01 build a todo.txt cli
x 2020-12-03 2020-12-03 put some examples on github
```

## TODO

- [ ] persistent `todo.txt` file path config (currently only reads `todo.txt`/`done.txt` in cwd)
- [ ] priorities
- [ ] sorting flags for `ls` command: by date/project/context
