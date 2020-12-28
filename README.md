# Overview

A utility that scans your Kubernetes installation and tries  to tell you everything you need to fix before updating in order to upgrade without breaking things.

# Example of use

```shell script
k8supdatechecker --context-name=CONTEXT
```

```
Scanning Kubernetes Cluster: CONTEXT
Running Kubernetes v1.15. Looking for deprecated features for v1.16.
......F..F....
Deprecated APIs found and are in use:
 1. kubectl --namespace production get svc/myservice
   1. Uses type: k8s.io/beta, but beta is deprcated, replace the type with X instead.
   1. Suggested fix: kubectl --namespace production apply -f PATCH_CONTENT
```

or


```shell script
k8supdatechecker --context-name=CONTEXT
```

```
Scanning Kubernetes Cluster: CONTEXT
Running Kubernetes v1.15. Looking for deprecated features for v1.16.
..............
OK: No Changes are recommended at this time.
```

# How to use it

Download the version that matches your current version of kubernetes. Since Kubernetes guarantees at most 2 minor versions of compatibility, this tool is intended to help you migrate your cluster one minor version at a time. For example, if you're currently running Kubernetes 1.19, and you download the 1.19.X version of this library, this will help you migrate to 1.20.

You can find out which version of Kubernetes you're running by using:

```shell
kubeclt 
```

# No Warranty

No warantee provided. Use this software at your own discretion and after doing your own due dilligence.

We are no responsible for any outages, downtime, or other damages for your use of this software.

This is a community-run project and we will not be able to report all problems. You must and are resonsible for testing beyond this k8s-update-checker utility before upgrading. You have been warned.
