# MongoDB Sharded

[MongoDB Sharded](https://www.mongodb.com/) is a cross-platform document-oriented database. Classified as a NoSQL database, MongoDB eschews the traditional table-based relational database structure in favor of JSON-like documents with dynamic schemas, making the integration of data in certain types of applications easier and faster.

This chart uses the [sharding method](https://docs.mongodb.com/manual/sharding/) for distributing data across multiple machines. This is meant for deployments with very large data sets and high throughput operations.

## TL;DR

```bash
$ helm repo add bitnami https://charts.bitnami.com/bitnami
$ helm install my-release bitnami/mongodb-sharded
```

## Introduction

This chart bootstraps a [MongoDB Sharded](https://github.com/bitnami/bitnami-docker-mongodb-sharded) deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

Bitnami charts can be used with [Kubeapps](https://kubeapps.com/) for deployment and management of Helm Charts in clusters. This chart has been tested to work with NGINX Ingress, cert-manager, fluentd and Prometheus on top of the [BKPR](https://kubeprod.io/).

## Prerequisites

- Kubernetes 1.12+
- Helm 2.12+ or Helm 3.0-beta3+
- PV provisioner support in the underlying infrastructure
- ReadWriteMany volumes for deployment scaling

## Installing the Chart

To install the chart with the release name `my-release`:

```bash
$ helm install my-release bitnami/mongodb-sharded
```

The command deploys MongoDB on the Kubernetes cluster in the default configuration. The [Parameters](#parameters) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `my-release` deployment:

```bash
$ helm delete my-release
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Parameters

The following table lists the configurable parameters of the MongoDB chart and their default values.

### Global Configuration

| Parameter                                     | Description                                                                                                                                               | Default                                                  |
|-----------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------|
| `global.imageRegistry`                        | Global Docker image registry                                                                                                                              | `nil`                                                    |
| `global.imagePullSecrets`                     | Global Docker registry secret names as an array                                                                                                           | `[]` (does not add image pull secrets to deployed pods)  |
| `global.storageClass`                         | Global storage class for dynamic provisioning                                                                                                             | `nil`                                                    |
| `image.registry`                              | MongoDB image registry                                                                                                                                    | `docker.io`                                              |
| `image.repository`                            | MongoDB Image name                                                                                                                                        | `bitnami/mongodb-sharded`                                |
| `image.tag`                                   | MongoDB Image tag                                                                                                                                         | `{TAG_NAME}`                                             |
| `image.pullPolicy`                            | Image pull policy                                                                                                                                         | `IfNotPresent`                                           |
| `image.pullSecrets`                           | Specify docker-registry secret names as an array                                                                                                          | `[]` (does not add image pull secrets to deployed pods)  |
| `image.debug`                                 | Specify if debug logs should be enabled                                                                                                                   | `false`                                                  |
| `nameOverride`                                | String to partially override mongodb.fullname template with a string (will prepend the release name)                                                      | `nil`                                                    |
| `fullnameOverride`                            | String to fully override mongodb.fullname template with a string                                                                                          | `nil`                                                    |
| `volumePermissions.enabled`                   | Enable init container that changes volume permissions in the data directory (for cases where the default k8s `runAsUser` and `fsUser` values do not work) | `false`                                                  |
| `volumePermissions.image.registry`            | Init container volume-permissions image registry                                                                                                          | `docker.io`                                              |
| `volumePermissions.image.repository`          | Init container volume-permissions image name                                                                                                              | `bitnami/minideb`                                        |
| `volumePermissions.image.tag`                 | Init container volume-permissions image tag                                                                                                               | `buster`                                                 |
| `volumePermissions.image.pullPolicy`          | Init container volume-permissions image pull policy                                                                                                       | `Always`                                                 |
| `volumePermissions.resources`                 | Init container resource requests/limit                                                                                                                    | `nil`                                                    |
| `clusterDomain`                               | Default Kubernetes cluster domain                                                                                                                         | `cluster.local`                                          |
| `existingSecret`                              | Existing secret with MongoDB credentials                                                                                                                  | `nil`                                                    |
| `mongodbRootPassword`                         | MongoDB admin password                                                                                                                                    | `random alphanumeric string (10)`                        |
| `shards`                                      | Number of shards to be created                                                                                                                            | `2`                                                      |
| `common.mongodbEnableIPv6`                    | Switch to enable/disable IPv6 on MongoDB                                                                                                                  | `false`                                                  |
| `common.mongodbDirectoryPerDB`                | Switch to enable/disable DirectoryPerDB on MongoDB                                                                                                        | `false`                                                  |
| `common.mongodbSystemLogVerbosity`            | MongoDB systen log verbosity level                                                                                                                        | `0`                                                      |
| `common.mongodbDisableSystemLog`              | Whether to disable MongoDB system log or not                                                                                                              | `false`                                                  |
| `common.mongodbMaxWaitTimeout`                | Maximum time (in seconds) for MongoDB nodes to wait for another MongoDB node to be ready                                                                  | `120`                                                    |
| `common.podLabels`                            | Extra labels for all pods in the cluster (evaluated as a template)                                                                                        | `{}`                                                     |
| `common.podAnnotations`                       | Extra annotations for all pods in the cluster (evaluated as a template)                                                                                   | `{}`                                                     |
| `common.serviceAccount.name`                  | Name of a Service Account to be used by all Pods                                                                                                          | `nil`                                                    |
| `common.serviceAccount.create`                | Whether to create a Service Account for all pods automatically                                                                                            | `false`                                                  |
| `common.sidecars`                             | Attach additional containers to all pods in the cluster (evaluated as a template)                                                                         | `nil`                                                    |
| `common.useHostnames`                         | Enable DNS hostnames in the replica set config                                                                                                            | `true`                                                   |
| `common.initContainers`                       | Add additional init containers to all pods in the cluster (evaluated as a template)                                                                       | `nil`                                                    |
| `common.extraEnvVars`                         | Array containing extra env vars to be added to all pods in the cluster (evaluated as a template)                                                          | `nil`                                                    |
| `common.extraEnvVarsCM`                       | ConfigMap containing extra env vars to be added to all pods in the cluster (evaluated as a template)                                                      | `nil`                                                    |
| `common.extraEnvVarsSecret`                   | Secret containing extra env vars to be added to all pods in the cluster (evaluated as a template)                                                         | `nil`                                                    |
| `common.extraVolumes`                         | Array of extra volumes to be added to all pods in the cluster  (evaluated as template). Requires setting `common.extraVolumeMounts`                       | `nil`                                                    |
| `common.extraVolumeMounts`                    | Array of extra volume mounts to be added to all pods in the cluster (evaluated as template). Normally used with `common.extraVolumes`.                    | `nil`                                                    |
| `common.initScriptsCM`                        | ConfigMap containing `/docker-entrypoint-initdb.d` scripts to be executed at initialization time (evaluated as a template)                                | `nil`                                                    |
| `common.initScriptsSecret`                    | Secret containing `/docker-entrypoint-initdb.d` scripts to be executed at initialization time (that contain sensitive data). Evaluated as a template.     | `nil`                                                    |
| `service.name`                                | Kubernetes service name                                                                                                                                   | `nil`                                                    |
| `service.annotations`                         | Kubernetes service annotations                                                                                                                            | `{}`                                                     |
| `service.type`                                | Kubernetes Service type                                                                                                                                   | `ClusterIP`                                              |
| `service.clusterIP`                           | Static clusterIP or None for headless services                                                                                                            | `nil`                                                    |
| `service.port`                                | MongoDB service port                                                                                                                                      | `27017`                                                  |
| `service.extraPorts`                          | Extra ports to expose in the service (normally used with the `sidecar` value). Evaluated as a template.                                                   | `nil`                                                    |
| `service.externalTrafficPolicy`               | Enable client source IP preservation                                                                                                                      | `Cluster`                                                |
| `service.nodePort`                            | Port to bind to for NodePort service type                                                                                                                 | `nil`                                                    |
| `service.loadBalancerIP`                      | Static IP Address to use for LoadBalancer service type                                                                                                    | `nil`                                                    |
| `service.externalIPs`                         | External IP list to use with ClusterIP service type                                                                                                       | `[]`                                                     |
| `service.loadBalancerSourceRanges`            | List of IP ranges allowed access to load balancer (if supported)                                                                                          | `[]` (does not add IP range restrictions to the service) |
| `service.sessionAffinity`                     | Session Affinity for Kubernetes service, can be "None" or "ClientIP"                                                                                      | `None`                                                   |
| `replicaSetKey`                               | Key used for authentication in the replica sets                                                                                                           | `random alphanumeric string (10)`                        |
| `usePasswordFile`                             | Have the secrets mounted as a file instead of env vars                                                                                                    | `false`                                                  |
| `securityContext.enabled`                     | Enable security context                                                                                                                                   | `true`                                                   |
| `securityContext.fsGroup`                     | Group ID for the container                                                                                                                                | `1001`                                                   |
| `securityContext.runAsUser`                   | User ID for the container                                                                                                                                 | `1001`                                                   |
| `securityContext.runAsNonRoot`                | Run containers as non-root users                                                                                                                          | `true`                                                   |
| `livenessProbe.enabled`                       | Enable/disable the Liveness probe                                                                                                                         | `true`                                                   |
| `livenessProbe.initialDelaySeconds`           | Delay before liveness probe is initiated                                                                                                                  | `30`                                                     |
| `livenessProbe.periodSeconds`                 | How often to perform the probe                                                                                                                            | `10`                                                     |
| `livenessProbe.timeoutSeconds`                | When the probe times out                                                                                                                                  | `5`                                                      |
| `livenessProbe.successThreshold`              | Minimum consecutive successes for the probe to be considered successful after having failed.                                                              | `1`                                                      |
| `livenessProbe.failureThreshold`              | Minimum consecutive failures for the probe to be considered failed after having succeeded.                                                                | `6`                                                      |
| `readinessProbe.enabled`                      | Enable/disable the Readiness probe                                                                                                                        | `true`                                                   |
| `readinessProbe.initialDelaySeconds`          | Delay before readiness probe is initiated                                                                                                                 | `5`                                                      |
| `readinessProbe.periodSeconds`                | How often to perform the probe                                                                                                                            | `10`                                                     |
| `readinessProbe.timeoutSeconds`               | When the probe times out                                                                                                                                  | `5`                                                      |
| `readinessProbe.failureThreshold`             | Minimum consecutive failures for the probe to be considered failed after having succeeded.                                                                | `6`                                                      |
| `readinessProbe.successThreshold`             | Minimum consecutive successes for the probe to be considered successful after having failed.                                                              | `1`                                                      |

### Config Server configuration

| Parameter                                     | Description                                                                                                                                               | Default                                                  |
|-----------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------|
| `configsvr.replicas`                          | Number of nodes in the replica set (the first node will be primary)                                                                                       | `1`                                                      |
| `configsvr.podAnnotations`                    | Annotations to be added to the deployment or statefulsets                                                                                                 | `{}`                                                     |
| `configsvr.podLabels`                         | Annotations to be added to the deployment or statefulsets                                                                                                 | `{}`                                                     |
| `configsvr.resources`                         | Pod resources                                                                                                                                             | `{}`                                                     |
| `configsvr.priorityClassName`                 | Pod priority class name                                                                                                                                   | ``                                                       |
| `configsvr.nodeSelector`                      | Node labels for pod assignment (evaluated as a template)                                                                                                  | `{}`                                                     |
| `configsvr.affinity`                          | Affinity for pod assignment (evaluated as a template)                                                                                                     | `{}`                                                     |
| `configsvr.tolerations`                       | Toleration labels for pod assignment (evaluated as a template)                                                                                            | `{}`                                                     |
| `configsvr.podManagementPolicy`               | Statefulsets pod management policy (evaluated as a template)                                                                                              | `OrderedReady`                                           |
| `configsvr.updateStrategy`                    | Statefulsets update strategy policy (evaluated as a template)                                                                                             | `RollingUpdate`                                          |
| `configsvr.schedulerName`                     | Name of the k8s scheduler (other than default)                                                                                                            | `nil`                                                    |
| `configsvr.pdb.enabled`                       | Enable pod disruption budget                                                                                                                              | `false`                                                  |
| `configsvr.pdb.minAvailable`                  | Minimum number of available config pods allowed (`0` to disable)                                                                                          | `0`                                                      |
| `configsvr.pdb.maxUnavailable`                | Maximum number of unavailable config pods allowed (`0` to disable)                                                                                        | `1`                                                      |
| `configsvr.sidecars`                          | Attach additional containers (evaluated as a template)                                                                                                    | `nil`                                                    |
| `configsvr.initContainers`                    | Add additional init containers (evaluated as a template)                                                                                                  | `nil`                                                    |
| `configsvr.config`                            | MongoDB configuration                                                                                                                                     | `nil`                                                    |
| `configsvr.configCM`                          | ConfigMap name with MongoDB configuration (cannot be used with configsvr.config)                                                                          | `nil`                                                    |
| `configsvr.mongodbExtraFlags`                 | MongoDB additional command line flags                                                                                                                     | `[]`                                                     |
| `configsvr.extraEnvVars`                      | Array containing extra env vars (evaluated as a template)                                                                                                 | `nil`                                                    |
| `configsvr.extraEnvVarsCM`                    | ConfigMap containing extra env vars (evaluated as a template)                                                                                             | `nil`                                                    |
| `configsvr.extraEnvVarsSecret`                | Secret containing extra env vars (evaluated as a template)                                                                                                | `nil`                                                    |
| `configsvr.extraVolumes`                      | Array of extra volumes (evaluated as template). Requires setting `common.extraVolumeMounts`                                                               | `nil`                                                    |
| `configsvr.extraVolumeMounts`                 | Array of extra volume mounts (evaluated as template). Normally used with `common.extraVolumes`.                                                           | `nil`                                                    |
| `configsvr.persistence.enabled`               | Use a PVC to persist data                                                                                                                                 | `true`                                                   |
| `configsvr.persistence.mountPath`             | Path to mount the volume at                                                                                                                               | `/bitnami/mongodb`                                       |
| `configsvr.persistence.subPath`               | Subdirectory of the volume to mount at                                                                                                                    | `""`                                                     |
| `configsvr.persistence.storageClass`          | Storage class of backing PVC                                                                                                                              | `nil` (uses alpha storage class annotation)              |
| `configsvr.persistence.accessModes`           | Use volume as ReadOnly or ReadWrite                                                                                                                       | `[ReadWriteOnce]`                                        |
| `configsvr.persistence.size`                  | Size of data volume                                                                                                                                       | `8Gi`                                                    |
| `configsvr.persistence.annotations`           | Persistent Volume annotations                                                                                                                             | `{}`                                                     |
| `configsvr.external.host`           | Primary node of an external config server replicaset                                                                                                                              | `nil`                                                     |
| `configsvr.external.rootPassword`           | Root passworrd of the external config server replicaset                                                                                                                              | `nil`                                                     |
| `configsvr.external.replicasetName`           | Replicaset name of an external config server                                                                                                                              | `nil`                                                     |
| `configsvr.external.replicasetKey`           | Replicaset key of an external config server                                                                                                                              | `nil`                                                     |
| `configsvr.serviceAccount.name`                 | Name of a Service Account to be used by configsvr                                                                                                            | `nil`                                                    |
| `configsvr.serviceAccount.create`               | Whether to create a Service Account for configsvr automatically                                                                                              | `false`                                                  |

### Mongos configuration

| Parameter                                     | Description                                                                                                                                               | Default                                                  |
|-----------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------|
| `mongos.replicas`                             | Number of Mongos nodes to create                                                                                                                          | `1`                                                      |
| `mongos.podAnnotations`                       | Annotations to be added to the deployment or statefulsets                                                                                                 | `{}`                                                     |
| `mongos.podLabels`                            | Annotations to be added to the deployment or statefulsets                                                                                                 | `{}`                                                     |
| `mongos.resources`                            | Pod resources                                                                                                                                             | `{}`                                                     |
| `mongos.priorityClassName`                    | Pod priority class name                                                                                                                                   | ``                                                       |
| `mongos.nodeSelector`                         | Node labels for pod assignment (evaluated as a template)                                                                                                  | `{}`                                                     |
| `mongos.affinity`                             | Affinity for pod assignment (evaluated as a template)                                                                                                     | `{}`                                                     |
| `mongos.tolerations`                          | Toleration labels for pod assignment (evaluated as a template)                                                                                            | `{}`                                                     |
| `mongos.updateStrategy`                       | Statefulsets update strategy policy (evaluated as a template)                                                                                             | `RollingUpdate`                                          |
| `mongos.schedulerName`                        | Name of the k8s scheduler (other than default)                                                                                                            | `nil`                                                    |
| `mongos.pdb.enabled`                          | Enable pod disruption budget                                                                                                                              | `false`                                                  |
| `mongos.pdb.minAvailable`                     | Minimum number of available mongo pods allowed (`0` to disable)                                                                                           | `0`                                                      |
| `mongos.pdb.maxUnavailable`                   | Maximum number of unavailable mongo pods allowed (`0` to disable)                                                                                         | `1`                                                      |
| `mongos.sidecars`                             | Attach additional containers (evaluated as a template)                                                                                                    | `nil`                                                    |
| `mongos.initContainers`                       | Add additional init containers (evaluated as a template)                                                                                                  | `nil`                                                    |
| `mongos.config`                               | MongoDB configuration                                                                                                                                     | `nil`                                                    |
| `mongos.configCM`                             | ConfigMap name with MongoDB configuration (cannot be used with mongos.config)                                                                             | `nil`                                                    |
| `mongos.mongodbExtraFlags`                    | MongoDB additional command line flags                                                                                                                     | `[]`                                                     |
| `mongos.extraEnvVars`                         | Array containing extra env vars (evaluated as a template)                                                                                                 | `nil`                                                    |
| `mongos.extraEnvVarsCM`                       | ConfigMap containing extra env vars (evaluated as a template)                                                                                             | `nil`                                                    |
| `mongos.extraEnvVarsSecret`                   | Secret containing extra env vars (evaluated as a template)                                                                                                | `nil`                                                    |
| `mongos.extraVolumes`                         | Array of extra volumes (evaluated as template). Requires setting `common.extraVolumeMounts`                                                               | `nil`                                                    |
| `mongos.extraVolumeMounts`                    | Array of extra volume mounts (evaluated as template). Normally used with `common.extraVolumes`.                                                           | `nil`                                                    |
| `mongos.serviceAccount.name`                  | Name of a Service Account to be used by mongos                                                                                                            | `nil`                                                    |
| `mongos.serviceAccount.create`                | Whether to create a Service Account for mongos automatically                                                                                              | `false`                                                  |

### Shard configuration: Data nodes

| Parameter                                     | Description                                                                                                                                               | Default                                                  |
|-----------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------|
| `shardsvr.dataNode.replicas`                  | Number of nodes in each shard replica set (the first node will be primary)                                                                                | `1`                                                      |
| `shardsvr.dataNode.podAnnotations`            | Annotations to be added to the deployment or statefulsets                                                                                                 | `{}`                                                     |
| `shardsvr.dataNode.podLabels`                 | Annotations to be added to the deployment or statefulsets                                                                                                 | `{}`                                                     |
| `shardsvr.dataNode.resources`                 | Pod resources                                                                                                                                             | `{}`                                                     |
| `shardsvr.dataNode.priorityClassName`         | Pod priority class name                                                                                                                                   | ``                                                       |
| `shardsvr.dataNode.nodeSelector`              | Node labels for pod assignment (evaluated as a template)                                                                                                  | `{}`                                                     |
| `shardsvr.dataNode.affinity`                  | Affinity for pod assignment (evaluated as a template). Will include `.arbiterLoopId` which identifies the shard.                                          | `{}`                                                     |
| `shardsvr.dataNode.tolerations`               | Toleration labels for pod assignment (evaluated as a template)                                                                                            | `{}`                                                     |
| `shardsvr.dataNode.podManagementPolicy`                | Statefulsets pod management policy (evaluated as a template)                                                                                              | `OrderedReady`                                           |
| `shardsvr.dataNode.updateStrategy`            | Statefulsets update strategy policy (evaluated as a template)                                                                                             | `RollingUpdate`                                          |
| `shardsvr.dataNode.schedulerName`             | Name of the k8s scheduler (other than default)                                                                                                            | `nil`                                                    |
| `shardsvr.dataNode.pdb.enabled`               | Enable pod disruption budget                                                                                                                              | `false`                                                  |
| `shardsvr.dataNode.pdb.minAvailable`          | Minimum number of available data pods allowed (`0` to disable)                                                                                            | `0`                                                      |
| `shardsvr.dataNode.pdb.maxUnavailable`        | Maximum number of unavailable data pods allowed (`0` to disable)                                                                                          | `1`                                                      |
| `shardsvr.dataNode.sidecars`                  | Attach additional containers (evaluated as a template)                                                                                                    | `nil`                                                    |
| `shardsvr.dataNode.initContainers`            | Add additional init containers (evaluated as a template)                                                                                                  | `nil`                                                    |
| `shardsvr.dataNode.config`                    | MongoDB configuration                                                                                                                                     | `nil`                                                    |
| `shardsvr.dataNode.configCM`                  | ConfigMap name with MongoDB configuration (cannot be used with shardsvr.dataNode.config)                                                                  | `nil`                                                    |
| `shardsvr.dataNode.mongodbExtraFlags`         | MongoDB additional command line flags                                                                                                                     | `[]`                                                     |
| `shardsvr.dataNode.extraEnvVars`              | Array containing extra env vars (evaluated as a template)                                                                                                 | `nil`                                                    |
| `shardsvr.dataNode.extraEnvVarsCM`            | ConfigMap containing extra env vars (evaluated as a template)                                                                                             | `nil`                                                    |
| `shardsvr.dataNode.extraEnvVarsSecret`        | Secret containing extra env vars (evaluated as a template)                                                                                                | `nil`                                                    |
| `shardsvr.dataNode.extraVolumes`              | Array of extra volumes (evaluated as template). Requires setting `common.extraVolumeMounts`                                                               | `nil`                                                    |
| `shardsvr.dataNode.extraVolumeMounts`         | Array of extra volume mounts (evaluated as template). Normally used with `common.extraVolumes`.                                                           | `nil`                                                    |
| `shardsvr.dataNode.serviceAccount.name`       | Name of a Service Account to be used by shardsvr data pods                                                                                                | `nil`                                                    |
| `shardsvr.dataNode.serviceAccount.create`     | Whether to create a Service Account for shardsvr data pods automatically                                                                                  | `false`                                                  |
| `shardsvr.persistence.enabled`                | Use a PVC to persist data                                                                                                                                 | `true`                                                   |
| `shardsvr.persistence.mountPath`              | Path to mount the volume at                                                                                                                               | `/bitnami/mongodb`                                       |
| `shardsvr.persistence.subPath`                | Subdirectory of the volume to mount at                                                                                                                    | `""`                                                     |
| `shardsvr.persistence.storageClass`           | Storage class of backing PVC                                                                                                                              | `nil` (uses alpha storage class annotation)              |
| `shardsvr.persistence.accessModes`            | Use volume as ReadOnly or ReadWrite                                                                                                                       | `[ReadWriteOnce]`                                        |
| `shardsvr.persistence.size`                   | Size of data volume                                                                                                                                       | `8Gi`                                                    |
| `shardsvr.persistence.annotations`            | Persistent Volume annotations                                                                                                                             | `{}`                                                     |

### Shard configuration: Arbiters

| Parameter                                     | Description                                                                                                                                               | Default                                                  |
|-----------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------|
| `shardsvr.arbiter.replicas`                   | Number of arbiters in each shard replica set (the first node will be primary)                                                                             | `1`                                                      |
| `shardsvr.arbiter.podAnnotations`             | Annotations to be added to the deployment or statefulsets                                                                                                 | `{}`                                                     |
| `shardsvr.arbiter.podLabels`                  | Annotations to be added to the deployment or statefulsets                                                                                                 | `{}`                                                     |
| `shardsvr.arbiter.resources`                  | Pod resources                                                                                                                                             | `{}`                                                     |
| `shardsvr.arbiter.priorityClassName`          | Pod priority class name                                                                                                                                   | ``                                                       |
| `shardsvr.arbiter.nodeSelector`               | Node labels for pod assignment (evaluated as a template)                                                                                                  | `{}`                                                     |
| `shardsvr.arbiter.affinity`                   | Affinity for pod assignment (evaluated as a template). Will include `.arbiterLoopId` which identifies the shard.                                          | `{}`                                                     |
| `shardsvr.arbiter.tolerations`                | Toleration labels for pod assignment (evaluated as a template)                                                                                            | `{}`                                                     |
| `shardsvr.arbiter.podManagementPolicy`        | Statefulsets pod management policy (evaluated as a template)                                                                                              | `OrderedReady`                                           |
| `shardsvr.arbiter.updateStrategy`             | Statefulsets update strategy policy (evaluated as a template)                                                                                             | `RollingUpdate`                                          |
| `shardsvr.arbiter.schedulerName`              | Name of the k8s scheduler (other than default)                                                                                                            | `nil`                                                    |
| `shardsvr.arbiter.sidecars`                   | Attach additional containers (evaluated as a template)                                                                                                    | `nil`                                                    |
| `shardsvr.arbiter.initContainers`             | Add additional init containers (evaluated as a template)                                                                                                  | `nil`                                                    |
| `shardsvr.arbiter.config`                     | MongoDB configuration                                                                                                                                     | `nil`                                                    |
| `shardsvr.arbiter.configCM`                   | ConfigMap name with MongoDB configuration (cannot be used with shardsvr.arbiter.config)                                                                   | `nil`                                                    |
| `shardsvr.arbiter.mongodbExtraFlags`          | MongoDB additional command line flags                                                                                                                     | `[]`                                                     |
| `shardsvr.arbiter.extraEnvVars`               | Array containing extra env vars (evaluated as a template)                                                                                                 | `nil`                                                    |
| `shardsvr.arbiter.extraEnvVarsCM`             | ConfigMap containing extra env vars (evaluated as a template)                                                                                             | `nil`                                                    |
| `shardsvr.arbiter.extraEnvVarsSecret`         | Secret containing extra env vars (evaluated as a template)                                                                                                | `nil`                                                    |
| `shardsvr.arbiter.extraVolumes`               | Array of extra volumes (evaluated as template). Requires setting `common.extraVolumeMounts`                                                               | `nil`                                                    |
| `shardsvr.arbiter.extraVolumeMounts`          | Array of extra volume mounts (evaluated as template). Normally used with `common.extraVolumes`.                                                           | `nil`                                                    |
| `shardsvr.arbiter.serviceAccount.name`        | Name of a Service Account to be used by shardsvr arbiter pods                                                                                             | `nil`                                                    |
| `shardsvr.arbiter.serviceAccount.create`      | Whether to create a Service Account for shardsvr arbiter pods automatically                                                                               | `false`                                                  |

### Metrics exporter

| Parameter                                     | Description                                                                                                                                               | Default                                                  |
|-----------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------|
| `metrics.enabled`                             | Start a side-car prometheus exporter                                                                                                                      | `false`                                                  |
| `metrics.image.registry`                      | MongoDB exporter image registry                                                                                                                           | `docker.io`                                              |
| `metrics.image.repository`                    | MongoDB exporter image name                                                                                                                               | `bitnami/mongodb-exporter`                               |
| `metrics.image.tag`                           | MongoDB exporter image tag                                                                                                                                | `{TAG_NAME}`                                             |
| `metrics.image.pullPolicy`                    | Image pull policy                                                                                                                                         | `Always`                                                 |
| `metrics.image.pullSecrets`                   | Specify docker-registry secret names as an array                                                                                                          | `[]` (does not add image pull secrets to deployed pods)  |
| `metrics.podAnnotations.prometheus.io/scrape` | Additional annotations for Metrics exporter pod                                                                                                           | `true`                                                   |
| `metrics.podAnnotations.prometheus.io/port`   | Additional annotations for Metrics exporter pod                                                                                                           | `"9216"`                                                 |
| `metrics.extraArgs`                           | String with extra arguments for the MongoDB Exporter                                                                                                      | ``                                                       |
| `metrics.resources`                           | Exporter resource requests/limit                                                                                                                          | `{}`                                                     |
| `metrics.livenessProbe.enabled`               | Enable/disable the Liveness Check of Prometheus metrics exporter                                                                                          | `false`                                                  |
| `metrics.livenessProbe.initialDelaySeconds`   | Initial Delay for Liveness Check of Prometheus metrics exporter                                                                                           | `15`                                                     |
| `metrics.livenessProbe.periodSeconds`         | How often to perform Liveness Check of Prometheus metrics exporter                                                                                        | `5`                                                      |
| `metrics.livenessProbe.timeoutSeconds`        | Timeout for Liveness Check of Prometheus metrics exporter                                                                                                 | `5`                                                      |
| `metrics.livenessProbe.failureThreshold`      | Failure Threshold for Liveness Check of Prometheus metrics exporter                                                                                       | `3`                                                      |
| `metrics.livenessProbe.successThreshold`      | Success Threshold for Liveness Check of Prometheus metrics exporter                                                                                       | `1`                                                      |
| `metrics.readinessProbe.enabled`              | Enable/disable the Readiness Check of Prometheus metrics exporter                                                                                         | `false`                                                  |
| `metrics.readinessProbe.initialDelaySeconds`  | Initial Delay for Readiness Check of Prometheus metrics exporter                                                                                          | `5`                                                      |
| `metrics.readinessProbe.periodSeconds`        | How often to perform Readiness Check of Prometheus metrics exporter                                                                                       | `5`                                                      |
| `metrics.readinessProbe.timeoutSeconds`       | Timeout for Readiness Check of Prometheus metrics exporter                                                                                                | `1`                                                      |
| `metrics.readinessProbe.failureThreshold`     | Failure Threshold for Readiness Check of Prometheus metrics exporter                                                                                      | `3`                                                      |
| `metrics.readinessProbe.successThreshold`     | Success Threshold for Readiness Check of Prometheus metrics exporter                                                                                      | `1`                                                      |
| `metrics.serviceMonitor.enabled`              | if `true`, creates a Prometheus Operator ServiceMonitor (also requires `metrics.kafka.enabled` or `metrics.jmx.enabled` to be `true`)                     | `false`                                                  |
| `metrics.serviceMonitor.namespace`            | Namespace which Prometheus is running in                                                                                                                  | `monitoring`                                             |
| `metrics.serviceMonitor.interval`             | How frequently to scrape metrics (use by default, falling back to Prometheus' default)                                                                    | `nil`                                                    |
| `metrics.serviceMonitor.selector`             | Default to kube-prometheus install (CoreOS recommended), but should be set according to Prometheus install                                                | `{ prometheus: kube-prometheus }`                        |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example,

```bash
$ helm install my-release \
  --set shards=4,configsvr.replicas=3,shardsvr.dataNode.replicas=2 \
    bitnami/mongodb-sharded
```

The above command sets the number of shards to 4, the number of replicas for the config servers to 3 and number of replicas for data nodes to 2.

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,

```bash
$ helm install my-release -f values.yaml bitnami/mongodb-sharded
```

> **Tip**: You can use the default [values.yaml](values.yaml)

## Configuration and installation details

### [Rolling VS Immutable tags](https://docs.bitnami.com/containers/how-to/understand-rolling-tags-containers/)

It is strongly recommended to use immutable tags in a production environment. This ensures your deployment does not change automatically if the same tag is updated with a different image.

Bitnami will release a new chart updating its containers if a new version of the main container, significant changes, or critical vulnerabilities exist.

### Production configuration and horizontal scaling

This chart includes a `values-production.yaml` file where you can find some parameters oriented to production configuration in comparison to the regular `values.yaml`. You can use this file instead of the default one.

- Increase shards to 4:
```diff
- shards: 2
+ shards: 4
```

- Increase config server replicaset sive:
```diff
- configsvr.replicas: 1
+ configsvr.replicas: 3
```

- Increase data nodes per shard:
```diff
- shardsvr.dataNode.replicas: 1
+ shardsvr.dataNode.replicas: 2
```

- Enable arbiter node on each shard:
```diff
- shardsvr.arbiter.replicas: 0
+ shardsvr.arbiter.replicas: 1
```

- Start a side-car prometheus exporter:
```diff
- metrics.enabled: false
+ metrics.enabled: true
```

### Change MongoDB version

To modify the MongoDB version used in this chart you can specify a [valid image tag](https://hub.docker.com/r/bitnami/mongodb-sharded/tags/) using the `image.tag` parameter. For example, `image.tag=X.Y.Z`. This approach is also applicable to other images like exporters.

### Sharding

This chart deploys a sharded cluster by default. Some characteristics of this chart are:

- It allows HA by enabling replication on the shards and the config servers. The mongos instances can be scaled horizontally as well.
- The number of secondary and arbiter nodes can be scaled out independently.

### Initialize a fresh instance

The [Bitnami MongoDB](https://github.com/bitnami/bitnami-docker-mongodb-sharded) image allows you to use your custom scripts to initialize a fresh instance. You can create a custom config map and give it via `initScriptsCM`(check options for more details).

The allowed extensions are `.sh`, and `.js`.

### Sidecars and Init Containers

If you have a need for additional containers to run within the same pod as Kibana (e.g. an additional metrics or logging exporter), you can do so via the `sidecars` config parameter (available in the `mongos`, `shardsvr.dataNode`, `shardsvr.arbiter`, `configsvr` and `common` sections). Simply define your container according to the Kubernetes container spec.

```yaml
sidecars:
- name: your-image-name
  image: your-image
  imagePullPolicy: Always
  ports:
  - name: portname
   containerPort: 1234
```

Similarly, you can add extra init containers using the `initContainers` parameter.

```yaml
initContainers:
- name: your-image-name
  image: your-image
  imagePullPolicy: Always
  ports:
  - name: portname
   containerPort: 1234
```

### Adding extra environment variables

In case you want to add extra environment variables (useful for advanced operations like custom init scripts), you can use the `extraEnvVars` (available in the `mongos`, `shardsvr.dataNode`, `shardsvr.arbiter`, `configsvr` and `common` sections) property.

```yaml
extraEnvVars:
  - name: MONGODB_VERSION
    value: 4.0
```

Alternatively, you can use a ConfigMap or a Secret with the environment variables. To do so, use the `extraEnvVarsCM` or the `extraEnvVarsSecret` values.

### Using an external config server

It is possible to not deploy any shards or a config server. For example, it is possible to simply deploy `mongos` instances that point to an external MongoDB sharded database. If that is the case, set the `configsvr.external.host` and `configsvr.external.replicasetName` for the mongos instances to connect. For authentication, set the `configsvr.external.rootPassword` and `configsvr.external.replicasetKey` values.

## Persistence

The [Bitnami MongoDB](https://github.com/bitnami/bitnami-docker-mongodb-sharded) image stores the MongoDB data and configurations at the `/bitnami/mongodb` path of the container.

The chart mounts a [Persistent Volume](http://kubernetes.io/docs/user-guide/persistent-volumes/) at this location. The volume is created using dynamic volume provisioning.

### Adjust permissions of persistent volume mountpoint

As the image run as non-root by default, it is necessary to adjust the ownership of the persistent volume so that the container can write data into it.

By default, the chart is configured to use Kubernetes Security Context to automatically change the ownership of the volume. However, this feature does not work in all Kubernetes distributions.
As an alternative, this chart supports using an initContainer to change the ownership of the volume before mounting it in the final destination.

You can enable this initContainer by setting `volumePermissions.enabled` to `true`.

### Adding extra volumes

The Bitnami Kibana chart supports mounting extra volumes (either PVCs, secrets or configmaps) by using the `extraVolumes` and `extraVolumeMounts` properties (available in the `mongos`, `shardsvr.dataNode`, `shardsvr.arbiter`, `configsvr` and `common` sections). This can be combined with advanced operations like adding extra init containers and sidecars.

## Upgrading

If authentication is enabled, it's necessary to set the `mongodbRootPassword` and `replicaSetKey` when upgrading for readiness/liveness probes to work properly. When you install this chart for the first time, some notes will be displayed providing the credentials you must use. Please note down the password, and run the command below to upgrade your chart:

```bash
$ helm upgrade my-release bitnami/mongodb-sharded --set mongodbRootPassword=[PASSWORD] (--set replicaSetKey=[REPLICASETKEY])
```

> Note: you need to substitute the placeholders [PASSWORD] and [REPLICASETKEY] with the values obtained in the installation notes.

### To 2.0.0

MongoDB container images were updated to `4.4.x` and it can affect compatibility with older versions of MongoDB. Refer to the following guide to upgrade your applications:

- [Upgrade a Sharded Cluster to 4.4](https://docs.mongodb.com/manual/release-notes/4.4-upgrade-sharded-cluster/)
