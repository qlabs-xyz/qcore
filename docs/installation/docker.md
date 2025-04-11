# Docker

There are two ways to obtain a Q Core Docker image:

1. [GitHub](#github)
2. [Building from Source](#building-the-docker-image)

Once you have the Docker image, proceed to [Using the Docker image](#using-the-docker-image).

## GitHub

Q Core Docker images for both x86_64 and ARM64 architectures are published with every Q Core release on GitHub Container Registry.

Pull the latest image:

```bash
docker pull ghcr.io/glabs-xyz/qcore
```

Or a specific version (e.g., v1.0.0):

```bash
docker pull ghcr.io/glabs-xyz/qcore:v1.0.0
```

Test the image:

```bash
docker run --rm ghcr.io/glabs-xyz/qcore --version
```

## Building the Docker Image

To build the Docker image from source, navigate to the repository root and run:

```bash
make local-image
```

The build may take several minutes. Once complete, test the image:

```bash
docker run qcore:local --version
```

## Using the Docker Image

You can run the Docker image using either:

1. [Plain Docker](#using-plain-docker)
2. [Docker Compose](#using-docker-compose)

### Using Plain Docker

Run Q Core with Docker:

```bash
docker run \
    -v .qcore:qcore/devnet \
    -p 8545:8545 \
    -p 8586:8586 \
    -p 9091:9091 \
    --name qcore \
    qcore:local
```

The above command creates a container named `qcore`. It exposes port `8485` for Ethereum JSON-RPC, 
`8486` for Ethereum Websocket and port `9091` for Cosmos REST (gRPC-Gateway).

To use the remote image from GitHub Container Registry, replace `qcore:local` with `ghcr.io/glabs-xyz/qcore` and your chosen tag.

## Interacting with Q Core Inside Docker

To interact with Q Core inside the Docker container, open a shell:

```bash
docker exec -it qcore bash
```