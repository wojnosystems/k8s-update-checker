# Overview

A utility that scans your Kubernetes installation and tries to tell you everything you need to fix or change before updating in order to avoid breaking things.

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

# Writing rules

This is a community project, which means we need you, yup! you!, to help us write rules to detect problems in Kubernetes clusters that would break if upgraded blindly. Each and every Kubernetes release will have a rule set in the "pkg/rules" section. The idea isn't to get people to stop reading the CHANGELOGs, but to help create a checklist of known things to fix, in case they miss anything.

## Testing your rules

Local testing with native K8s are conducted using minikube. That means you'll need to install minikube and have the ability to create clusters locally, either through Docker or Virtual Box or however you decide to do it.

### Yea, but, how do I test GKE or EKS?

Yeah... about that... I'm not sure yet. I suppose we can just use the gcloud/aws-cli and work it into the framework somehow. I'm all ears on this one.

# No Warranty

No warranties are provided. Use this software at your own risk and after doing your own due diligence.

We are no responsible for any outages, downtime, or other damages for your use of this software.

This is a community-run project and we will not be able to report all problems or upgrade blockers. Kubernetes is complicated and extremely customizable. You must and are responsible for testing beyond this k8s-update-checker utility before upgrading. You have been warned.
