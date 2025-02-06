# Gophercises

## Exercise 2: URL Shortener

### Initial Thoughts
**Total Time**: 3 hours 42 mins

Dealing with libararies and types is still tricky but easier as I learn more about standard libraries and how they work. I started thinking that structs were just types, describing the shape and type of data... I happened upon this interesting (for me) structure:

https://github.com/Nostromos/url-shortener/blob/0ff8a57d23fe2a137c76e888b0300dd4bea5e6e7/handler.go#L113-L117

We create a map, then loop through the array of `pathURLs` and assign keys and values in `pathsToUrls` *by referencing the struct keys and values*. It isn't possible to do it this way in Javascript (lacking types) and I'm pretty sure there's no way to do it in Typescript either. I'm still confused why/how structs can determine access but I'm too used to thinking in JS/TS, where I'd have had to create the object with `Path:` and `URL:` keys. Anyway - cool!

https://github.com/Nostromos/url-shortener/blob/0ff8a57d23fe2a137c76e888b0300dd4bea5e6e7/handler.go#L85-L88

Another thing I'm still wrapping my head around is how low level things are. `parseJSON` takes a slice of bytes and the `encoding/json` package is working with raw bytes. In JS, JSON would be parsed into an object (iirc) and you'd be accessing that instead. 

Lastly, pointers. Still interesting and a touch confusing. I like that I can directly access memory addresses of variables (as in the above) and understand the use cases but dereferencing is confusing. I've become used to small projects and not having to do my own memory management so I can understand how this would be incredibly powerful for applications needing performance.

### Final Thoughts
Golang is tricky and I'm still in the phase where simple things are a struggle. It's frustrating to know that better patterns/syntax/methods/approaches exist but I don't know enough to use them. Turning off Copilot suggestions was a good idea but its slow going. 
