============
Triangulizor
============

Triangulizor is a little Python script to apply a “triangular pixel”
effect to images, like so:

.. image:: https://github.com/mccutchen/triangulizor/raw/master/examples/in.jpg

⇊☁⇊

.. image:: https://github.com/mccutchen/triangulizor/raw/master/examples/out.png


Installation
============

Use `pip`_ to install::

    pip install triangulizor


Usage
=====

Command Line Usage
------------------

First, find an image to triangulize! To generate the example above,
either of these commands will do the trick::

    $ triangulizor --show --tile-size=16 examples/in.jpg

The ``--show`` flag will cause the resulting image to be displayed
immediately instead of written to ``stdout`` or to disk. You can also
pass in the URL to an image that you want to process::

    $ triangulizor --show --tile-size=16 https://github.com/mccutchen/triangulizor/raw/master/examples/in.jpg

All command line options are given below::

    $ triangulizor --help

Outputs::

    usage: triangulizor [-h] [-t TILE_SIZE] [-v] [-vv] [-s] [infile] [outfile]

    Applies a "triangular pixel" effect to an image.

    positional arguments:
      infile                Image to process (path or URL; defaults to STDIN)
      outfile               Output file (defaults to STDOUT)

    optional arguments:
      -h, --help            show this help message and exit
      -t TILE_SIZE, --tile-size TILE_SIZE
                            Tile size (should be divisible by 2)
      -v, --verbose         Verbose output
      -vv                   Very verbose output
      -s, --show            Immediately display image instead of writing to
                            OUTFILE.

Library Usage
-------------

Triangulizor can also be used as a plain old Python library::

    >>> import triangulizor
    >>> triangulizor.triangulize('examples/in.jpg', 24)
    <Image._ImageCrop image mode=RGB size=384x216 at 0x10A5BA758>


Credits
=======

This was inspired entirely by `this awesomely helpful blog post`_ by
`@revdancatt`_.

.. _pip: http://www.pip-installer.org/
.. _this awesomely helpful blog post: http://revdancatt.com/2012/03/31/the-pxl-effect-with-javascript-and-canvas-and-maths/
.. _@revdancatt: http://twitter.com/revdancatt
