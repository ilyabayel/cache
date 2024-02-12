# Cache

It's an in-memory cache created only for learning purposes.
By doing this I learn how to publish a package, nothing else.

## Usage

Here is an example of the code with all exising methods for the `Cache` struct.

```go
	c := cache.New()

	c.Set("key", "value")

	v := c.Get("key") // "value"

	c.Delete("key")

	v := c.Get("key") // nil
```
