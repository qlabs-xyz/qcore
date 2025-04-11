# Build from Source

You can build Q Core Node on Linux, macOS, Windows.

## Dependencies

**install Golang** using [go](https://go.dev/)

## Build Q Core Node

With Go installed, you're ready to build Q Core Node. First, clone the repository:

```bash
git clone https://github.com/qlabs-xyz/qcore/
cd qcore
```

Then, install Q Core Node into your `PATH` directly via:

```bash
make install
```

The binary will now be accessible as `qcored` via the command line, and exist under your default `bin` folder.

Alternatively, you can build yourself with:

```bash
make build 
```

This will place the Q Core Node binary under `./build/qcored`, and you can copy it to your directory of preference after that.

Compilation may take around ~5 minutes. Installation was successful if `qcored --help` displays the help for qcored
