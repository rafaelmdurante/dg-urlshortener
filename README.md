# URL Shortener

My solution for the [DevGym][devgym]'s URL Shortener challenge.

## Challenge

In this challenge you must create a server to shorten URL and redirect short 
URLs to their original link.

| HTTP Method | Endpoint | Description                                |
|-------------|----------|--------------------------------------------|
| `POST`      | `/`      | Takes the URL and returns a unique code    |
| `GET`       | `/:code` | Uses the `code` to return the original URL |

Note that `code` is unique `6-chars string`. The same original URL should 
generate different `codes`.

## Solution

Solution based on filhodanuvem's [standard go api][ytgoapi] and Alex Xu's [URL shortener][system-design-interview] suggested architecture.

### App

To run the project on your machine.

```bash
make up
```

Now you should be able to create a short url and get redirected to the target url.

Read `docs/http` files for more details.

### Tests

To run all the tests:

```bash
make test
```

[devgym]: https://app.devgym.com.br/challenges/3ecd0771-981d-44dc-9eee-5ec69791a745
[ytgoapi]: https://github.com/filhodanuvem/ytgoapi/tree/main
[system-design-interview]: https://www.amazon.com/System-Design-Interview-insiders-Second/dp/B08CMF2CQF