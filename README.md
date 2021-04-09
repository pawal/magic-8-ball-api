# Magic 8 ball API

Two endpoints for delivering the standard 20 Magic 8 ball answers as a REST API with JSON:

 * /magic - give the client one Magic 8 ball answer as a JSON object
 * /completemagic - give the client all 20 standard answer objects in an array

Example answer object:

```js
{
    "Id":13,
    "Answer":"Cannot predict now.",
    "Type":"neutral"
}
```

The Type attribite indicates if the answer is "affirmative", "neutral", or "negative".
