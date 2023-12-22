# 001 - Code Generation

## Context

A simple URL shortener should create a unique up-to-6-chars code string and 
associate that with a target URL. That means there are 68,719,476,736 
possible combinations for short URLs.

## Options Considered

1. Create a base64 hash when the `POST` request is received
    * Pros
      * Easy implementation
    * Cons
      * Scaling issues when possible combinations are limited, user could 
        get stuck waiting until server manages to create a `hash` that 
        hadn't been used yet
2. Create all combinations possible and fill in when necessary
    * Pros
      * All combinations ready to go
    * Cons
      * Database would have 
3. Create a few combinations and keep creating only if needed
    * Pros
    * Cons

## Decision



## Consequences
What becomes easier or more difficult to do because of this change?

## Conclusion
What is the conclusion after the decision?

