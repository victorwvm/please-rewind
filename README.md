# Please Rewind

A page that brings together what I have been listening to and what I have been watching, in one place.

The data comes from Last.fm and Letterboxd. The idea is that anyone can enter their own two usernames and see their own page, and also compare it with someone else's.

## Status

Nothing implemented yet. So far there is only the screen design and the choice of stack.

## Stack

- Go
- Vercel

## Where the data comes from

| Source | How | Needs a key |
|---|---|---|
| Last.fm | Official API | yes |
| Letterboxd | Profile RSS feed (`letterboxd.com/<username>/rss/`) | no |

Letterboxd has no open public API, so the way in is the RSS feed. It returns the 50 most recent entries.

## Planned URLs

```
victorwvm.com.br/rewind/<lastfm>/<letterboxd>
victorwvm.com.br/rewind/<lastfm>/<letterboxd>/vs/<lastfm2>/<letterboxd2>
```

## Time windows

30 days, 90 days, 365 days and everything. They match the Last.fm `period` values: `1month`, `3month`, `12month` and `overall`.
