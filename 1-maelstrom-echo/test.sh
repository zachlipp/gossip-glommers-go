#!/bin/bash

set -xuo pipefail

maelstrom test -w echo --bin ~/go/bin/maelstrom-echo --node-count 1 --time-limit 10
