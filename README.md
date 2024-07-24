# GO Mongo
This example project of Go-lang to manage data with MongoDB.

## Setup Project
1. Copy `.env.example` to `.env`
2. Edit value in `.env` to match your configuration
3. Setup mongo container by
   ```sh
   make compose-up
   ```
4. After mongo container is running, you can run project by
   ```sh
   go run main.go
   ```
   or you can build the project first and run builded file.
   