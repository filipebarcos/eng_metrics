# Engineering Metrics

The goal of this project is to be able to easily fetch GH data in order to have an idea
of how your engineering team is performing. This tool won't give you those metrics for
free, but the data.

## System design

This project was built to run in Google Cloud Functions (it was an easy way to deploy in
the company I work), this means each function can be triggered independently.

Functions will fetch the data from GH and store into PostgreSQL. The idea is to leverage
Metabase for querying this data and creating dashboards to give some insights.

## Before running

### PEM File

To run this app, you need to [create a GitHub App](https://docs.github.com/en/developers/apps/creating-a-github-app) and
[create a private key for authentication](https://docs.github.com/en/developers/apps/authenticating-with-github-apps#generating-a-private-key) and
add it to `./certs` folder. Ensure you name your private key file as `github.engmetrics.private-key.pem`. This filename
is set as a `const` in `github/config.go`.

### Environment Variables
Ensure your environment variables are properly configured. Run `cp cmd/.env{.dev,}` and edit the,
newly created, `.env` file.

* `DATABASE_URL`: You don't need to change this, as it points to the docker container;
* `GITHUB_APP_ID`: This ID you can find once you create your app, it's a numeric value;
* `GITHUB_INSTALLATION_ID`: This ID you can find once you create your app, it's a numeric value;
* `GITHUB_ORGANIZATION`: Use this to set which org you want to fetch repo info from;
* `GITHUB_REPOSITORIES`: This is a comman separated list of repositories to find information from;
* `MIGRATION_RELATIVE_PATH`: This is important for GCP, since the path where the code lives and the function changes from what we have here, so using `.` if fine for local dev, but dependending where you deploying these functions, you might need to change this value.

## How to run

Once your variables are configure, you can run docker-compose to ensure you have postgres and metabase running.
```sh
rake up # Rake does the honors of running `docker-compose up -d` for you
```

Once docker is running, build and run the app. We have all the functions "puggled" on `./cmd` folder. Each
function lives on the project root folder, but they're added/configure on `./cmd/cmg.go` so we can run them
locally by doing:
```sh
$ go build
./cmd/cmd
```

## Trying GitHub API queries

Go to [GitHub explorer](https://developer.github.com/v4/explorer/), there you can try the queries we're doing
so you can know what type of data we're fetching from GH. You can find queries in `github/github.go` comments.

## Dependencies
- [GitHub v4 API](https://github.com/shurcooL/githubv4)
- [GitHub Installation](https://github.com/shurcooL/githubv4)
- [Tern](https://github.com/jackc/tern): PostgreSQL migration tool
- [PGX](https://github.com/jackc/pgx): PostgreSQL driver

## Dev Dependencies
- Ruby 2.7.0
- Rake`gem install rake`

## Developing

### Docker commands

We have docker compose with a single container, for now. If we want to have up and down, we can leverage rake for it.
```sh
$ rake up
$ rake down
```

### Creating migrations

Run `rake "db:migrate:create[<migration_name>]"` (replace `<migration_name>` with the file name you want). A new file
will be created on `./migrate/`. It's SQL file with a magic comment, check https://github.com/jackc/tern#migrations
in case of questions.

#### Migrating

Run `rake db:migrate` to run all pending migratioins and if you want to know the stats, you can run
`rake db:migrate:status`.

### Get list of commands available
```sh
$ rake -T
rake app                      # Build and run app
rake app:build                # Builds the app
rake app:build_and_run        # Build and run app
rake app:run                  # Run built app
rake db:console               # Open psql console
rake db:migrate:create[name]  # Creates a new migration
rake db:migrate:rollback      # Rolls back latest migration
rake db:migrate:status        # Prints information about migration status
rake db:migrate:up            # Migrate up to latest version
rake db:nuke                  # Nukes DB data
rake down                     # Stops docker so it does not consume your computer resources
rake migrate                  # Migrate up to latest version
rake up                       # Runs docker-compose to spin up dependencies
```
