# URL Shortener

My solution for the [DevGym][devgym]'s URL Shortener challenge.

## Requesites

In this challenge you must create a server to shorten URL and redirect short 
URLs to their original link.

| HTTP Method | Endpoint | Description                                |
|-------------|----------|--------------------------------------------|
| `POST`      | `/`      | Takes the URL and returns a unique code    |
| `GET`       | `/:code` | Uses the `code` to return the original URL |

Note that `code` is unique `6-chars string`. The same original URL should 
generate different `codes`.

## 

[devgym]: https://devgym.com.br