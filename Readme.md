# TODOx

TODOx is a simple todo list app built with Go and HTMX. It uses in memory storage and is meant to prove the concept of using HTMX with Go.

## Features

- Create, Read, Update, and Delete todos
- Searching TODOs

## Stack

The stack is simple:
- Go for the backend
- HTMX for the frontend
- Tailwind CSS for styling

I have set a Makefile to setup and run the app. 

## Architecture

The application is mostly written in Go and HTML. HTMX facilitates a lot of the interaction with the user on the frontend and the backend endpoints process requests and return HTML that will be then rendered by HTMX.

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