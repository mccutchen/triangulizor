# Triangulizor

Triangulizor is a little Python script to apply a "triangular pixel" effect to
images, like so:

<img src="https://github.com/mccutchen/triangulizor/raw/master/examples/in.jpg" align="middle">
➡☁➡
<img src="https://github.com/mccutchen/triangulizor/raw/master/examples/out.png" align="middle">

## Usage

Triangulizor requires Python 2.7+ and the [Python Imaging Library][3]. First,
install `PIL` (usually this is as easy as `pip install pil`). Next, find an
image to triangulize! To generate the example above, either of these commands
will do the trick:

```bash
./triangulizor.py --show --tile-size=16 examples/in.jpg
```

The `--show` flag will cause the resulting image to be displayed immediately
instead of written to `stdout` or to disk. You can also pass in the URL to an
image that you want to process:

```bash
./triangulizor.py --show --tile-size=16 https://github.com/mccutchen/triangulizor/raw/master/examples/in.jpg
```

All command line options are given below:

```bash
$ ./triangulizor.py --help
```

```
usage: triangulizor.py [-h] [-t TILE_SIZE] [-v] [-vv] [-s] [infile] [outfile]

Applies a "triangular pixel" effect to an image.

positional arguments:
  infile                Image to process (path or URL; defaults to STDIN)
  outfile               Output file (defaults to STDOUT)

optional arguments:
  -h, --help            show this help message and exit
  -t TILE_SIZE, --tile-size TILE_SIZE
                        Tile size (must be divisible by 2; defaults to 20)
  -v, --verbose         Verbose output
  -vv                   Very verbose output
  -s, --show            Immediately display image instead of writing to
                        OUTFILE.
```

## Credits

This was inspired entirely by [this awesomely helpful blog post][1] by
[@revdancatt][2].

[1]: http://revdancatt.com/2012/03/31/the-pxl-effect-with-javascript-and-canvas-and-maths/
[2]: http://twitter.com/revdancatt
[3]: http://pypi.python.org/pypi/PIL
