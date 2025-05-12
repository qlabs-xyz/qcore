# Running OutBe Core Node

## Running the OutBe Core Node Local

First, ensure that you have OutBe Core Node installed by following the [installation instructions][installation].

Now, to start the local node, run:

```bash
CLEAN=true sh scripts/test_node.sh
```

A genesis and default configuration will be created.

 If you want to run node without creating genesis and the default configuration, use the following command:

```bash
CLEAN=false sh scripts/test_node.sh
```

At this point, our OutBe Core node has started. 
