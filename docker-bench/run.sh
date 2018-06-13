#!/bin/sh

docker run -it --net host --pid host --userns host --cap-add audit_control \
     -e DOCKER_CONTENT_TRUST=$DOCKER_CONTENT_TRUST \
     -v /var/lib:/var/lib \
     -v /var/run/docker.sock:/var/run/docker.sock \
     -v /etc:/etc --label docker_bench_security \
     docker/docker-bench-security
