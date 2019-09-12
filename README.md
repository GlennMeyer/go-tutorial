# Go Tutorial

## Requirements
- [Docker Engine - Community](https://docs.docker.com/install/)

## Setup

1. Register for a [News API key](https://newsapi.org/register)

1. Create a `.env` file in the root of the project enter values for the following variables:
    1. `NEWSAPI_KEY`
    1. `POSTGRES_USER`
    1. `POSTGRES_PASSWORD`
    1. `POSTGRES_DB`
    1. `POSTGRES_PORT`

1. `make compose`

## Usage

To view commonly used Docker commands, check out the [Makefile](https://github.com/GlennMeyer/go_tutorial/blob/master/Makefile)
