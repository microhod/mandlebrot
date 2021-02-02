# Mandlebrot Renderer

This generate a `png` render of the [mandlebrot set](https://en.wikipedia.org/wiki/Mandelbrot_set) (see example below).

You can tinker with the variables in [main.go](./main.go) to view different parts of the set etc.

![mandlebrot](./mandlebrot_large.png)

## Requirements

* [golang](https://golang.org)

## Running

* clone the repo
* build the code

    ```sh
    # this will output the executable
    go build
    ```

* run the executable (on windows it would be `mandlebrot.exe`)

    ```sh
    # <WIDTH> is an optional parameter to set the width of the desired render in pixels
    ./mandlebrot <WIDTH>
    ```
