# Overview

This is a small program to solve a very specific problem I was facing with helm deployments. On new installs I wanted a chart that requires storage to not install it's PVC before the storage provider existed. If that happens the provider will process new claims but not the existing claims.

This program uses client-go to check for a default storage provider. It will attempts to find one every 5 seconds until it does, logging it's status each loop, and exiting once it exists. I used this in a `pre-install` hook with Helm to protect the chart that uses storage. Be aware that checking for the storageclass requires giving this pod a service account with permissions to list storageclasses. If that permission is missing the pod will log it, I simply added the `pre-install` to the permissions and weighted them to run before this pod.

This all works as described for me, and I thought it might be useful to have it hosted where I can pull it into clusters in a similar situation. If you use this and want to expand it some I'm fine with PRs just recognize it's intentionally small and single purpose just fork it if you want something expansive.
