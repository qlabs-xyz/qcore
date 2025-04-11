# Running Q Core Node

## Running the Q Core Node Local

First, ensure that you have Q Core Node installed by following the [installation instructions][installation].

Now, to start the local node, run:

```bash
CLEAN=true sh scripts/test_node.sh
```

A genesis and default configuration will be created.

 If you want to run node without creating genesis and the default configuration, use the following command:

```bash
CLEAN=false sh scripts/test_node.sh
```

At this point, our Q Core node has started. 
