# btool ⚙️

This is a simple tool to automate compressing and decompressing `.tar.br` files
concurrently. The `tar` command doesn't seem to have support for this, so I
wrote a simple wee tool for myself.

This may be a bit buggy, really only made it for my use. If you use the tool
in a path where a folder contains a dot, it'll probably break

## Usage

Decompress files

In the current working directory, these files exist:

- first.tar.br
- second.tar.br
- third.tar.br

```shell
btool -d *
```

Output:

- first/
- second/
- third/
- first.tar.br
- second.tar.br
- third.tar.br

Compress files

Without a flag set, the tool will compress files and folders that you specify
as arguments into `.tar.br` files. This may take some time, brotli is much slower
at compression than decompression.

I've attempted to speed this up by compressing each specified file in it's own
goroutine.

In the current working directory, these folders exist:

- first/
- second/
- third/

```shell
btool *
```

Output:

- first/
- second/
- third/
- first.tar.br
- second.tar.br
- third.tar.br
