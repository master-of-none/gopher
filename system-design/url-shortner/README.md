# URL Shortner

A Go URL shortener backed by PostgreSQL and Redis, with a React + Vite frontend in `url-shortener-ui/`.

## What it does

- Creates short codes from URLs
- Redirects short codes to the stored destination
- Caches redirect lookups in Redis
- Removes expired URLs during cleanup

## Project Layout

- `main.go` - server startup, middleware, shutdown, and cleanup worker
- `handler/` - HTTP handlers for shortening and redirecting
- `db/` - PostgreSQL connection and cleanup helpers
- `redis/` - Redis client wrapper
- `cron/` - periodic expired-link cleanup worker
- `method/` - Base62 encoding helpers
- `mw/` - CORS and rate limit middleware
- `model/` - shared data structures
- `url-shortener-ui/` - frontend app

## Requirements

- Go 1.22 or newer
- PostgreSQL
- Redis
- Node.js 18+ for the frontend

## Configuration

The backend loads `.env` and expects `DATABASE_URL`.

Example:

```env
DATABASE_URL="postgres://user:password@localhost:5432/url_shortener_db?sslmode=disable"
```

Redis is currently configured in code to use `localhost:6379`.

## Run the Backend

From the repository root:

```bash
go run .
```

The server listens on `http://localhost:8080`.

For local development, you can also use `air` for live reload if you have it installed:

```bash
air
```

## Run the Frontend

From `url-shortener-ui/`:

```bash
npm install
npm run dev
```

Available scripts:

- `npm run build`
- `npm run lint`
- `npm run preview`

## Testing

You can test the API with any HTTP client, including YAAK or Postman.

Useful endpoints:

- `POST http://localhost:8080/shorten`
- `GET http://localhost:8080/{code}`

You can also run the frontend locally from `url-shortener-ui/` with `npm run dev` and use it against the backend running on `localhost:8080`.

## API

### Create a short URL

`POST /shorten`

Request body:

```json
{
  "url": "example.com"
}
```

Optional query parameter:

- `expiry` - expiration time in minutes

Example:

```bash
curl -X POST "http://localhost:8080/shorten?expiry=60" \
  -H "Content-Type: application/json" \
  -d '{"url":"example.com"}'
```

Success response:

- Status: `201 Created`
- Headers: `Location: http://localhost:8080/{code}`
- Body:

```json
{
  "short_url": "http://localhost:8080/{code}",
  "code": "{code}"
}
```

Behavior:

- If the URL does not start with `http://` or `https://`, the backend prepends `https://`
- If `expiry` is omitted, the link expires after 1 minute
- If `expiry` is invalid, the handler returns `400 Bad Request`
- If the JSON body is invalid, the handler returns `400 Bad Request`

### Redirect a short URL

`GET /{code}`

Example:

```bash
curl -i "http://localhost:8080/abc123"
```

Success response:

- Status: `302 Found`
- Redirects to the stored long URL

Error responses:

- `404 Not Found` with JSON body `{"error":"URL Not Found"}`
- `410 Gone` with JSON body `{"error":"Link has been Expired"}`

## Notes

- Redirects are cached in Redis after the first lookup.
- Expired URLs are deleted by the cleanup worker and again on shutdown.
- The backend uses CORS, request logging, recovery, and rate limiting middleware.
