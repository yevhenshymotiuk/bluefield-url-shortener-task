# bluefield-url-shortener-task

Run the server using Makefile `make run`

## Tasks
### Required
- [X] The requests to shortened URLs should be redirected to their
  original URL (status 302) or return 404 for unknown URLs.
- [X] Simple HTML form should be served on the index page where users can
  input URL and retrieve the shortened version from server.
- [X] All of the implemented HTTP handlers should have unit tests.

### Optional
- [X] All shortened URLs should be persisted locally to a file using
  simple storage methods (SQLite, BoltDB, CSV..).
- [X] The redirect requests should be cached in memory for certain
  amount of time.
