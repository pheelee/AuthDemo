# AuthDemo

This is a simple Golang webserver which provides basic auth and forms auth for evaluation scenarios.
It generates a bunch of users to test authentication with.

I wrote this little app to test Microsoft Azure Application Proxy Password-based SSO.

## Usage

```bash
Usage of ./authdemo:
  -Port int
        Listening port (default 8091)
```

## Sample Output

```bash
Created the following demo users:
- username: Account1
  password: 7727f5768a5783e3a23d6ef5e23aa621
- username: Account2
  password: ddce769e5761b1c979162ba5585708c9
- username: Account3
  password: bdaf87954c87bf97c4148a8ecf522e95
- username: Account4
  password: 55c52556befa55c867b5dc8981261e07

created endpoints /basic and /forms
Listening on port 8091
```

# License
MIT License

Copyright (c) 2021 Philipp Ritter

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