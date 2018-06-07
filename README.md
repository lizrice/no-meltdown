# Resources for "Don't Have A Meltdown"

Six practical tips for securing your container-based deployment

## Static analysis
Protecting against common attack vectors like SQL injection.

* Golang static analysis - [GoASTScanner/gas](https://github.com/GoASTScanner/gas) 
* [List of static analysis tools](https://en.wikipedia.org/wiki/List_of_tools_for_static_code_analysis)

## TLS scanning
Protecting you from leaving your connections unsecured.

## Image scanning
Protecting you from known exploits.

* [Docker Enterprise Edition image scanning](https://docs.docker.com/ee/#secure-supply-chain)
* [MicroScanner](https://github.com/aquasecurity/microscanner)

## Container OS
Minimizing your attack surface to reduce likelihood of both known and unknown vulnerabilities being present.

* A [Container OS comparison](https://blog.codeship.com/container-os-comparison/)
* [CIS Docker benchmark](https://www.cisecurity.org/benchmark/docker/) and [docker-bench](https://github.com/docker/docker-bench-security)
* [CIS Kubernetes benchmark](https://www.cisecurity.org/benchmark/kubernetes/) and [kube-bench](https://github.com/aquasecurity/kube-bench)

## Limit bind mounts & privileges
Limiting the potential effect of an attack. 

* Avoid running as root: [canihasnonprivilegedcontainers](http://canihaznonprivilegedcontainers.info/)
* [Jess Frazelle on the privileged flag](https://twitter.com/jessfraz/status/985947353981976576)

## Runtime protection 
Protecting you from attacks that cause your containers to behave in unexpected ways. 

* Default Docker [AppArmor](https://docs.docker.com/engine/security/apparmor/) and [seccomp](https://docs.docker.com/engine/security/seccomp/) profiles
* Aqua's [runtime protection](https://www.aquasec.com/use-cases/container-runtime-protection/)

