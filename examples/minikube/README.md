Kuma Control Plane inside Minikube
====================

## Pre-requirements

- `minikube`
- `kubectl`

## Usage

### Start Minikube

```bash
minikube start
```

### Build Control Plane image

```bash
make build/example/minikube -C ../.. -f Makefile.e2e.mk
```

### Deploy demo setup into Minikube

```bash
make deploy/example/minikube -C ../.. -f Makefile.e2e.mk
```

### Make test requests

```bash
make wait/example/minikube -C ../.. -f Makefile.e2e.mk
make curl/example/minikube -C ../.. -f Makefile.e2e.mk
```

### Verify Envoy stats

```bash
make verify/example/minikube -C ../.. -f Makefile.e2e.mk
```

### Observe Envoy stats

```bash
make stats/example/minikube -C ../.. -f Makefile.e2e.mk
```

E.g.,
```
# TYPE envoy_cluster_upstream_rq_total counter
envoy_cluster_upstream_rq_total{envoy_cluster_name="localhost_8000"} 11
envoy_cluster_upstream_rq_total{envoy_cluster_name="ads_cluster"} 1
envoy_cluster_upstream_rq_total{envoy_cluster_name="pass_through"} 3
```

where

* `envoy_cluster_upstream_rq_total{envoy_cluster_name="localhost_8000"}` is a number of `inbound` requests
* `envoy_cluster_upstream_rq_total{envoy_cluster_name="pass_through"}` is a number of `outbound` requests

### Enable mTLS

```bash
make apply/example/minikube/mtls -C ../.. -f Makefile.e2e.mk
```

### Wait until Envoy is configured for mTLS

```bash
make wait/example/minikube/mtls -C ../.. -f Makefile.e2e.mk
```

### Make test requests via Envoy with mTLS

```bash
make curl/example/minikube -C ../.. -f Makefile.e2e.mk
```

### Verify Envoy mTLS stats

```bash
make verify/example/minikube/mtls -C ../.. -f Makefile.e2e.mk
```

### Verify kumactl workflow

```bash
make kumactl/example/minikube -C ../.. -f Makefile.e2e.mk
```
