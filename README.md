# rss-reader

A service that handles JSON API requests containing URLs, uses an internal RSS reader package to parse the rss feed of the URLs and returns a JSON containing the data.

## Set up

Run `go build` in the project directory

## Using the service

The service will start on port `:6060`. Using Postman sample requests can be sent for the service to process

Example: POST `http://localhost:6060/`

```
{
"urls": [
"https://rss.art19.com/apology-line",
"http://feeds.feedburner.com/experiment_podcast"
]
}
```

## Further work

Time formats: I saw that different rss sources can use different time formats, so in the future more formats can be handled by the RssReader package

Error handling: The errors inside the go routines are only logged, maybe they can be returned to the user to be used for debugging