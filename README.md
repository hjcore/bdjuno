# gjuno
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/gotabit/gjuno/Tests)](https://github.com/gotabit/gjuno/actions?query=workflow%3ATests)
[![Go Report Card](https://goreportcard.com/badge/github.com/gotabit/gjuno)](https://goreportcard.com/report/github.com/gotabit/gjuno)
![Codecov branch](https://img.shields.io/codecov/c/github/gotabit/gjuno/cosmos/v0.40.x)

GJuno (shorthand for Gatabit Juno) is the [Juno](https://github.com/gotabit/gjuno) implementation
for [gotabit](https://github.com/gotabit/gotabit).

It extends the custom Juno behavior by adding different handlers and custom operations to make it easier for gotabit
showing the data inside the UI.

All the chains' data that are queried from the RPC and gRPC endpoints are stored inside
a [PostgreSQL](https://www.postgresql.org/) database on top of which [GraphQL](https://graphql.org/) APIs can then be
created using [Hasura](https://hasura.io/).


## Testing
If you want to test the code, you can do so by running

```shell
$ make test-unit
```

**Note**: Requires [Docker](https://docker.com).

This will:
1. Create a Docker container running a PostgreSQL database.
2. Run all the tests using that database as support.


