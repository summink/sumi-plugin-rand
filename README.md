# Sumi Plugin Rand

## Description

Generate random values, including integers, floats, and strings.

## Install

```bash
sumi plugin add rand@latest
```

## Usage

```bash
sumi rand [options]
```

### Options

- `--int` - generate integer.

- `--float` - generate float.

- `--string` - generate string.

- `--range <chars>` - set random charsets.

  - default: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789".

- `--min <n>` - Random minimum value.

  - default: 1

- `--max <n>` - Random maximum value.

  - default: 10000

- `--copy` - If set, copy result to clipboard.

## Examples

Generate a random integer:

```bash
sumi rand --int
```

Generate a random number within a range:

```bash
sumi rand --int --min 10 --max 100
```

Generate a random float:

```bash
sumi rand --float
```

Generate a random string:

```bash
sumi rand --string
```

Generate a random string with a custom charset:

```bash
sumi rand --string --range abc123
```

Copy result to clipboard:

```bash
sumi rand --string --copy
```
