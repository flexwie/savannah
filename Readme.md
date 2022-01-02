# Savannah
Savannah is a CI/CD tool for [Hashicorp Nomad](https://www.nomadproject.io/). It automatically discovers changed deployment instructions in repositories and applies them, much like ArgoCD does for Kubernetes.

**This is a WIP! Do not expect this to be production-ready!**

## Getting started
You can run Savannah as a standalone application or container inside or outside of your cluster.

### Standalone
Grab the latest binaries from GitHub and extract them on the server, that Savannah is supposed to run on. Then start the application:
```bash
$ ./savannah --conf ./config.hcl
```

This will create the 

### Container
TODO