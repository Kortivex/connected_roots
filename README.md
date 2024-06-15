# Connected Roots

## Table of Contents

1. [General Info](#general-info)
2. [Dependencies](#dependencies)
3. [Environment Variables](#environment-variables)
4. [Documentation](#documentation)
5. [Others](#others)

## General Info

> _Go VERSION_ to use: [_v1.22.4_](https://golang.org/doc/devel/release.html#go1.22)

Connected Roots is a service to store data about farms and the data recollected by the sensors.

## Dependencies

- **PostgreSQL**: The database where data are stored. Kind of data to save.

## Environment Variables

This service is configured using environment variables.

The next table shows all the environment variables related to Connected Root configuration.

| **Name** | **Default value** | **Description** |
|----------|-------------------|-----------------|
|          |                   |                 |

## Documentation

- **[OpenAPI](docs/api/connected_roots_openapi.yaml)**

## Others

### pre-commit

You need [pre-commit](https://pre-commit.com/) to committing in the project. Execute this command:

- `pip install pre-commit`

To test the pre-commit works run `pre-commit --version` and you get the tool version installed previously.

All the config related to **pre-commit** is in `.pre-commit-config.yaml` file.

Then you need to install the pre-commit hook in the repository executing:

- `pre-commit install` and get something like `pre-commit installed at .git/hooks/pre-commit`
- `pre-commit install --hook-type commit-msg` and get something like `pre-commit installed at .git/hooks/commit-msg`
