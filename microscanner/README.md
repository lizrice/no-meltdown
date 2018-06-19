### Microscanner demo

Generates an HTML-format vulnerability report for the Wordpress image.

You need a (https://github.com/aquasecurity/microscanner)[MicroScanner token].

```bash
# token in MICROSCANNER_TOKEN
docker build -f Dockerfile.wp --build-arg=token=$MICROSCANNER_TOKEN --no-cache .
# Get the container ID of the built container
./copy.sh <container ID> # Copies the ms-out.html file out of the container image
```

In browser open ms-out.html file to see vulnerability info

If there's a high-sev vulnerability, microscanner returns non-zero, which fails the Docker build.