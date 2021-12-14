# roulette

Draws two names out of a list.

## Usage

**roulette** draws elements from the comma-separated environment variable `ROULETTE_LIST`.

It accepts an optional parameter `-a` to exclude an element from the list (see Example below).

### Example

```bash
$ ROULETTE_LIST="foo,bar,baz,foz" roulette -a foo
2021/12/08 17:38:39 choosing from: [foz bar baz]
selected:  [foz bar]
```

## Install

Run the following command in a terminal:

```bash
go install github.com/denischevalier/roulette@latest
```
