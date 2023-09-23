# cycletime

Keep track of how long it takes you to complete an issue.

Works locally on any git repo where you tag your commit messages with `#NUMBER`.

- doesn't care about your issue lifecycle, only your commits
- can call [github cli](https://cli.github.com/) to fetch issue titles
- no API tokens needed
- can exclude commit authors with regex

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

My go-to invocation to look at last 14 days and exclude Dependabot:

```
$ cycletime --exclude=bot --days 14 -gh
2023-09-11 Display curation statements with a serif font                #2608        4.3 ·
2023-09-12 Update images and text of home page value statements         #2610        2.1 ·
2023-09-12 Change layout of home page value statements to altern...     #2617        0.7 ·
2023-09-13 Spike extraction of constructReviewingGroups                 #2618        2.0 ·
2023-09-14 Make article cards look more like a journal                  #2609       48.5 ·······
2023-09-14 Remove the whitespace at the bottom of the home page ...     #2620        0.0
2023-09-14 Improve contrast on home page curation teasers               #2628        0.1 ·
2023-09-14 Improve alignment on the search results page                 #2629        0.0
2023-09-14 Convey the opportunity and solution space to readers ...     #2619        3.7 ·
2023-09-14 Make triangles appear on laptops                             #2630        1.0 ·
2023-09-15 Investigate our db instances hitting end of life and ...     #2631        3.9 ·
2023-09-19 Make rendering of annotation author dynamic                  #2639       25.5 ····
2023-09-19 Ensure that users can delete annotations by removing ...     #2644        0.7 ·
2023-09-19 Stop using deprecated command helpers in the tests           #2640        6.9 ·
2023-09-21 Optimise startup performance by separating decoding f...     #2649        0.8 ·
2023-09-21 Improve accuracy of data on about page                       #2654        0.0
2023-09-22 Make annotations appear as comments rather than curat...     #2650       23.1 ···
2023-09-22 Fix visual bug for annotation content spilling over          #2655        0.0
2023-09-22 Make annotations look like speech bubbles                    #2658        1.2 ·
2023-09-22 Concierge annotations for the The Natural History of ...     #2659        1.3 ·
2023-09-22 Fix bug where users can't remove articles from list i...     #2660        0.8 ·
2023-09-22 Allow users on staging to annotate articles on a list        #2663        0.2 ·
2023-09-22 Do not stop the dev server when backstop is run              #2661        0.8 ·
```


## Usage
```
$ cycletime -h
Usage: cycletime [flags] [PATH]

Hours between first and last commit tagged with an issue number

PATH defaults to the current working directory

  -days int
        How many days to look back, -1 being infinity (default -1)
  -exclude string
        Exclude commits with authors that match this regex (default "^$")
  -gh
        Use gh cli to obtain issue titles
```

## Install

Two options:
  - clone the repo and `go install`
  - download prebuilt binaries from [releases page](https://github.com/erkannt/cycletime/releases)
