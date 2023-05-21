# Uindr
Winder is a TailwindCLI wrapper written in Go it takes care of downloading and running the Tailwind CLI for you by considering the Architecture and OS of your machine.

## Installation

```bash
go install github.com/paganotoni/uinder@latest
```

## Usage
Once winder is installed you can run it and pass the Tailwind CLI arguments to it.

```bash
winder [args]

# An example
winder -i application.css --content "./*.html,./**/*.html" -o public/application.css -w

# Another one example
winder -i base.css --content "./*.html" -o base.css -m
```

If the CLI has not been downloaded it will be downloaded and cached in your home directory under `.winder` folder. Once is downloaded it will be used from the cache.