FrameLog
A movie discovery website built with Vue, TypeScript, and Go. I'm building it as a personal project, with the idea of eventually having a place to keep track of what I watch and write about it.
Right now, you can browse popular movies and open a page with their synopsis, genres, runtime, and TMDB rating. Reviews and watchlists are next on the list, but aren't available yet.
Stack
- Vue 3, TypeScript, and Vue Router
- Vite
- Go with net/http
- TMDB for movie data and posters
The code and technical documentation are in English. The interface is in Spanish.
How it works
The frontend calls the Go API, which fetches movie data from TMDB. The TMDB token stays in the backend. During development, Vite forwards /api requests to Go on port 8080.
The backend is organized into a few small packages:
- cmd/api: starts the server and registers routes.
- internal/config: loads environment variables.
- internal/tmdb: handles requests to TMDB.
- internal/catalog: handles catalog and movie details endpoints.
There is no database yet. PostgreSQL is planned for users, reviews, and watchlists.
Running locally
You'll need Git, Node.js with npm, Go, and a TMDB API Read Access Token. Check frontend/package.json and backend/go.mod for the required runtime versions.
Clone the project:
git clone https://github.com/bCharry15/FrameLog.git
cd FrameLog
Create backend/.env using backend/.env.example as a reference:
TMDB_READ_ACCESS_TOKEN=YOUR_TMDB_READ_ACCESS_TOKEN
Use the full API Read Access Token, not the shorter API Key. Don't add Bearer before the value. Keep the real token out of Git and leave .env.example without credentials.
Start the backend from its directory so it can find .env:
cd backend
go mod download
go run ./cmd/api
In another terminal, from the project root:
cd frontend
npm ci
npm run dev
Open the URL printed by Vite, usually http://localhost:5173. Both processes need to stay running. If you change .env, restart Go.
API routes
Route	What it does
GET /api/v1/health/live	Checks that the API is running
GET /api/v1/movies	Returns the first page of popular movies
GET /api/v1/movies/{id}	Returns details for a TMDB movie ID


Movie requests use Spanish (es-CO), although some text may be missing or available only in another language. The popular list uses region CO; it doesn't include local cinema showtimes.
If the catalog fails to load, check the Go terminal. A TMDB 401 usually means the token is incorrect. If the browser cannot reach the backend, check that Go is running on port 8080 and that the Vite proxy is configured.
What's next
- Movie search and pagination
- User accounts
- Reviews, personal ratings, and spoiler controls
- Watchlists
- Tests for the main application flows
Vitest and Playwright were included in the initial setup. The starter tests still need to be adapted to the application.
Movie data
Data and artwork come from TMDB. The ratings shown are from TMDB, not from FrameLog users.
This product uses the TMDB API but is not endorsed or certified by TMDB.
Adding the approved TMDB logo to the app's credits is still pending. Data and image use is subject to TMDB's terms.
