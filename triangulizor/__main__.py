import argparse
from cStringIO import StringIO
import logging
import os
import re
import sys
import urllib2

from .triangulizor import triangulize

if __name__ == '__main__':

    def path_or_url(x):
        if re.match(r'^https?://', x):
            try:
                return urllib2.urlopen(x)
            except urllib2.URLError, e:
                raise argparse.ArgumentTypeError(str(e))
        elif os.path.isfile(x):
            return open(x, 'rb')
        else:
            msg = '%r not found' % x
            raise argparse.ArgumentTypeError(msg)

    arg_parser = argparse.ArgumentParser(
        prog='triangulizor',
        description='Applies a "triangular pixel" effect to an image.')
    arg_parser.add_argument(
        'infile', nargs='?', default=sys.stdin, type=path_or_url,
        help='Image to process (path or URL; defaults to STDIN)')
    arg_parser.add_argument(
        'outfile', nargs='?', default=sys.stdout,
        type=argparse.FileType('wb'),
        help='Output file (defaults to STDOUT)')
    arg_parser.add_argument(
        '-t', '--tile-size', type=int, default=0,
        help='Tile size (should be divisible by 2)')
    arg_parser.add_argument(
        '-f', '--format', type=str, default='PNG',
        help='Output file format (defaults to PNG; must be supported by PIL)')
    arg_parser.add_argument(
        '-v', '--verbose', default=False, action='store_const', const=True,
        help='Verbose output')
    arg_parser.add_argument(
        '-vv', default=False, action='store_const', const=True,
        help='Very verbose output')
    arg_parser.add_argument(
        '-s', '--show', default=False, action='store_const', const=True,
        help='Immediately display image instead of writing to OUTFILE.')

    args = arg_parser.parse_args()

    if args.verbose or args.vv:
        logger = logging.getLogger()
        logger.setLevel(logging.DEBUG if args.vv else logging.INFO)

    # We need to buffer the input because neither sys.stdin nor urllib2
    # response objects support seek()
    inbuffer = StringIO(args.infile.read())
    try:
        image = triangulize(inbuffer, args.tile_size)
    except IOError, e:
        logging.error('Unable to open image: %s', e)
    except KeyboardInterrupt:
        logging.info('Interrupted by user, exiting...')
        sys.exit(1)
    else:
        if args.show:
            image.show()
        else:
            image.save(args.outfile, args.format)
