# Spacey

A Twitter Spaces search client
```
spacey career industry

┌──────────┬────────────────────────────────────────────────────────────────────┬──────┬──────────────┬──────────┬────────────────────────────────────────────┐
│ QUERY    │ SPACE                                                              │ LANG │ PARTICIPANTS │ SPEAKERS │ URL                                        │
├──────────┼────────────────────────────────────────────────────────────────────┼──────┼──────────────┼──────────┼────────────────────────────────────────────┤
│ career   │ #360Mentor:International Professional Career with @Kasabiiti       │ en   │          565 │        2 │ https://twitter.com/i/spaces/1ynJOZQgDRwGR │
│ career   │ LinkedIn, Resume and profile reviews. Career talk                  │ en   │          136 │        1 │ https://twitter.com/i/spaces/1yNGaYojpMRGj │
│ industry │ Are AAA games too safe? Are indies the only bastion of creativity? │ en   │           23 │        2 │ https://twitter.com/i/spaces/1BRJjnbNoagJw │
└──────────┴────────────────────────────────────────────────────────────────────┴──────┴──────────────┴──────────┴────────────────────────────────────────────┘
```

## Install

via `brew`

```
brew install rothgar/tap/spacey
```

via [`bin`](https://github.com/marcosnils/bin)
```
bin install rothgar/spacey
```

## Usage

You'll need to get a [developer API key](https://developer.twitter.com/en/docs/twitter-api/getting-started/getting-access-to-the-twitter-api) to use `spacey`.
Export the required fields into your environment.

```
export TWITTER_BEARER_TOKEN=AAAAAAAAAAAAAAAA...
```

Search for spaces based on multiple queries in the title of the space.
Queries are case insensitive and multiple words should be quoted (e.g. "fireside chat")
```
spacey $QUERY1 $QUERY2 $QUERY3
```

You can also filter output by a minimum number of participants and speakers.
You can also change the output type with the `--output` flag.
