# Docker

There are two ways to obtain a OutBe Core Docker image:

1. [GitHub](#github)
2. [Building from Source](#building-the-docker-image)

Once you have the Docker image, proceed to [Using the Docker image](#using-the-docker-image).

## GitHub

OutBe Core Docker images for both x86_64 and ARM64 architectures are published with every OutBe Core release on GitHub Container Registry.

Pull the latest image:

```bash
docker pull ghcr.io/outbe/outbe-node
```

Or a specific version (e.g., v1.0.0):

```bash
docker pull ghcr.io/outbe/outbe-node:v1.0.0
```

Test the image:

```bash
docker run --rm ghcr.io/outbe/outbe-node --version
```

## Building the Docker Image

To build the Docker image from source, navigate to the repository root and run:

```bash
make local-image
```

The build may take several minutes. Once complete, test the image:

```bash
docker run outbe-node:local --version
```

## Using the Docker Image

You can run the Docker image using either:

1. [Plain Docker](#using-plain-docker)
2. [Docker Compose](#using-docker-compose)

### Using Plain Docker

Run OutBe Core with Docker:

```bash
docker run \
    -v ./outbe-node:outbe-node/devnet \
    -p 8545:8545 \
    -p 8586:8586 \
    -p 9091:9091 \
    --name outbe-node \
    outbe-node:local
```

The above command creates a container named `outbe-node`. It exposes port `8485` for Ethereum JSON-RPC, 
`8486` for Ethereum Websocket and port `9091` for Cosmos REST (gRPC-Gateway).

To use the remote image from GitHub Container Registry, replace `outbe-node:local` with `ghcr.io/outbe/outbe-node` and your chosen tag.

## Interacting with Q Core Inside Docker

To interact with Q Core inside the Docker container, open a shell:

```bash
docker exec -it outbe-node bash
```