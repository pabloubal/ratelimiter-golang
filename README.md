# Rate Limiter
HTTP Rate limiter. Sits in front of http endpoints and limits the amount of requests to them.

## Requirements
- Add a remote endpoint with a global limit of requests per second
- Receive an http request for any endpoint. If it's present in the rate limiter config, limit it, if not allow passthrough.


# Things to improve
- Domain shouldn't be pegged to RestAPIErrors. They should be generic errors and get their translation to rest api errors in the infra layer