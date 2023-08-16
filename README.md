# cycletime

Keep track of how long it takes you to complete an issue.

Works locally on any git repo where you tag your commit messages with `#NUMBER`.

- no API tokens or interaction with e.g. Github
- doesn't care about your issue lifecycle, only your commits

## Example output

Fields are: date, issue, hours between first and last commit, one dot per 8h

```
2023-07-24 #2439    336.5 ················································>>
2023-07-24 #2480      0.9 ·
2023-07-25 #2481     23.7 ···
2023-07-25 #2462    191.6 ························
2023-07-27 #2485     19.4 ···
2023-07-31 #2491     99.7 ·············
2023-08-01 #2508      3.0 ·
```

## Usage
```
Usage: cycletime [--exclude=AUTHOR_REGEX] [PATH]

Hours between first and last commit tagged with an issue number

PATH defaults to the current working directory
```

## Install

Two options:
  - clone the repo and `go install`
  - download prebuilt binaries from [releases page](https://github.com/erkannt/cycletime/releases)
