# TODOx

TODOx is a todo list app built with Go, HTMX, Tailwind CSS and SQLite to store the data. The original goal of this app is to try HTMX and see how it can be used to build a simple app.

## Features

- Create, Read, Update, and Delete TODOs
- Searching TODOs
- Complete/Reopen TODOs

## Stack

The stack aims to be as simple as possible. It consists of:

- Go is the main language.
- Gomponents to build the HTML.
- HTMX for the interactions with the user.
- Tailwind CSS for styling (CDN).
- SQLite as the database.

## Architecture

The application is mostly written in Go and HTML. HTMX facilitates a lot of the interaction with the user on the frontend and the backend endpoints process requests and return HTML served by Go into the client side.
The storage of the application is a shiny SQLite database.

## Running in development

### Setup
In order to run the app in development you should run the setup cli, it will download the standalone CLI. You can do this by running the following command:

```
go run ./cmd/setup
```

### Running
To run the app in development mode (with hot recompile) you can run the following command:

```
kit dev
```

And visit http://localhost:3000 to see the app running.

## Deploying

The app contains a Dockerfile that can be used to build images and deploy to any container platform.

A few environment variables to consider are:

```
GO_ENV  - The environment the app is running in (def: development)
PORT    - The port the app will run on. Defaults "3000"
```
