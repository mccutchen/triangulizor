import os
from distutils.core import setup


def read(fname):
    return open(os.path.join(os.path.dirname(__file__), fname)).read()


setup(
    name='triangulizor',
    version='1.0.0',
    description='Triangulize your images!',
    long_description=read('README.rst'),
    url='https://github.com/mccutchen/triangulizor',
    license='MIT',
    author='Will McCutchen',
    author_email='will@mccutch.org',
    classifiers = [
        'Development Status :: 4 - Beta',
        'Environment :: Console',
        'Intended Audience :: Developers',
        'Intended Audience :: Information Technology',
        'License :: OSI Approved :: MIT License',
        'Operating System :: OS Independent',
        'Programming Language :: Python',
        'Topic :: Artistic Software',
        'Topic :: Multimedia :: Graphics',
        'Topic :: Software Development :: Libraries :: Python Modules',
    ],
    packages=['triangulizor'],
    scripts=['bin/triangulizor'],
    install_requires=[
        "PIL >= 1.1.7",
    ],
)
