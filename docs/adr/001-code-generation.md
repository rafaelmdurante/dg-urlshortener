# 001 - Code Generation

## Context

A simple URL shortener should create a unique up-to-6-chars code string and 
associate that with a target URL. That means there are 56,800,235,583
possible combinations for short URLs.

## Options Considered

1. Create a random base64 hash when the `POST` request is received
    * Pros
      * Easy implementation
    * Cons
      * Scaling issues when possible combinations are limited, user could 
        get stuck waiting until server manages to create a `hash` that 
        hadn't been used yet - aka hash collision
2. Create all combinations possible and fulfill the url details only when a `POST` request is received
    * Pros
      * All combinations ready to go
    * Cons
      * Database would be too big from start
3. Create a base62 hash from the equivalent Postgres ID when `POST` request is received
    * Pros
      * Easy to implement
      * Avoids hash collision 
    * Cons
      * Two database requests per each new URL - first to get the ID, second to update the URL Code
      * Potential security concern as IDs are auto increment

## Decision

Option 3 was decided as it is easy to implement and avoids hash collision.

## Conclusion

Although IDs are auto increment, that does not necessarily mean it is a threat. To avoid any security concern,
There should be an external unique ID generator but this is beyond the scope of this challenge.
