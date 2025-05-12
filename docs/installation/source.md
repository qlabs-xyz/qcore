# Build from Source

You can build OutBe Core Node on Linux, macOS, Windows.

## Dependencies

**install Golang** using [go](https://go.dev/)

## Build OutBe Core Node

With Go installed, you're ready to build OutBe Core Node. First, clone the repository:

```bash
git clone https://github.com/outbe/outbe-node/
cd outbe-node
```

Then, install OutBe Core Node into your `PATH` directly via:

```bash
make install
```

The binary will now be accessible as `outbe-noded` via the command line, and exist under your default `bin` folder.

Alternatively, you can build yourself with:

```bash
make build 
```

This will place the Q Core Node binary under `./build/outbe-noded`, and you can copy it to your directory of preference after that.

Compilation may take around ~5 minutes. Installation was successful if `outbe-noded --help` displays the help for outbe-noded
