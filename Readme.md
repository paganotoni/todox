# TODOx

TODOx is a todo list app built with Go, HTMX, Tailwind CSS and SQLite to store the data. The original goal of this app is to try HTMX and see how it can be used to build a simple app.

## Features

- Create, Read, Update, and Delete TODOs
- Searching TODOs
- Complete/Uncomplete TODOs

## Stack

The stack is simple:

- Go for the backend
- HTMX for the frontend
- Tailwind CSS for styling
- SQLite for the persistence of the data.

I have set a Makefile to setup and run the app. 

## Architecture

The application is mostly written in Go and HTML. HTMX facilitates a lot of the interaction with the user on the frontend and the backend endpoints process requests and return HTML that will be then rendered by HTMX.

The Tailwind CSS Standalone CLI takes care of the styling by processing html files and adding resulting CSS to public/styles.css. Any CSS in the public folder is served by the Go server. The storage of the application is SQLite.

## Running in development

### Setup
In order to install the app in development you should download the tailwind css standalone CLI and air for hot reloading. You can do this by running the following command:

```
make setup
```

### Running 
To run the app in development you can run the following command:

```
make run
```

And visit http://localhost:3000 to see the app running.

## Deploying

The app contains a Dockerfile that can be used to build images and deploy to any container platform. 

A few environment variables to consider are:

```
GO_ENV                          - The environment the app is running in. Defaults to development. Use production in production.
PORT                            - The port the app will run on. Defaults to 3000.
```

Another consideration to keep in mind is that the build process assumes Linux and x64 architecture so if the docker image is being used on a different platform it may not work. This is because we pull the specific version of the tailwind standalone CLI.