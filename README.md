## QRServe - HTTP microservice for QR Code generation

[![Build Status](https://travis-ci.org/dobarkod/qrserve.svg?branch=master)](https://travis-ci.org/dobarkod/qrserve)

An implementation of a simple HTTP microservice that generates
QR Code upon request. All the QR heavy lifting is done by the
[qrcode](https://github.com/skip2/go-qrcode/) Go package.

## Installation

Install the latest stable version from GitHub:

    go get github.com/dobarkod/qrserve

## Running the service

To run the service, specify the IP address and TCP port to listen on:

    qrserve address:port

The address is optional and can be omitted, in which case QRServe will listen
on all interfaces. To listen only on a specific IP address, specify it like
this (example for localhost and port 8000):

    qrserve 127.0.0.1:8000

## Using the service

To generate a QR code, make a GET request to the `http://address:port/` URL
with the following query string parameters:

* `data` - data (text) to be encoded, eg. an URL
* `size` - size (in pixels), of the generated QR code (if less than required, will be automatically increased)
* `q` - error correction level (optional); possible values are `L` (low), `M` (medium, default), `Q` (high), `H` (highest)

The service will respond with a status code of 200 and the generated QR
code in PNG image format. In case of error, the service will return an
error code (400 in case of bad request data, 500 if image could not be generated
for other reasons).

To avoid impacting the service stability, the maximum allowed size is 4096
(that is, the generated image will be at most 4096x4096 px in size).


For example:

    curl -o qr.png http://127.0.0.1:8000/?data=Hello+world&size=100

## License (MIT)

Copyright (c) 2016 Good Code and contributors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
