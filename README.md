# Gossip Glommers

`fly.io` collaborated with `Jepsen` to create [a series of distributed systems challenges](https://fly.io/dist-sys). These challenges guide users through writing concurrent code to meet specific requirements which are then tested with the [maelstrom](https://github.com/jepsen-io/maelstrom) testing library.

## Development

Each directory corresponds to one of the challenges, the details of which can be found on the `fly.io` website. I use the `test.sh` convention to store the `maelstrom` command the program must past. Again, this command is given in the `fly.io` website.

### Running locally

1. Install [maelstrom](https://github.com/jepsen-io/maelstrom)
2. From the challenge you want to test, run `go install . && ./test.sh`

### Running in Docker

Coming soon
