# K8S Jobs Watchdog

![GitHub top language](https://img.shields.io/github/languages/top/amirhnajafiz/job-watchdog)
![GitHub repo size](https://img.shields.io/github/repo-size/amirhnajafiz/job-watchdog)
![GitHub release (with filter)](https://img.shields.io/github/v/release/amirhnajafiz/job-watchdog)

In this project I implemented a monitoring system called watchdog. This operator purpose is
monitoring ```kubernetes``` jobs in intervals. Moreover, it sends their status and results
over ```Kafka``` message broker as ```JSON``` format.
This operator keeps track of ```kubernetes``` jobs and send push notification over Kafka cluster.

## Image

Operator's docker image is ```amirhossein21/job-watchdog```. Make sure to pull
the tags with ```stable``` prefix in their versions.

```shell
docker run -d -it \
  -v type=bind,source=~/.kube/config,dest=/app/config \
  amirhossein21/job-watchdog
```

### env variables

Image environment variables have ```jm_``` prefix. The list bellow displays all the
operator available env variables.

- ```jm_interval``` : jobs pulling interval in seconds (default 5)
- ```jm_kafka__host``` : kafka cluster host (default ```localhost:9292```)
- ```jm_kafka__topic``` : the topic which operator publishes on (default ```jm-jobs```)
- ```jm_kafka__partition``` : kafka partition (default 0)
- ```jm_cluster__kubeconfig``` : path to kubeconfig file (default ```config```)
- ```jm_cluster__namespace``` : kubernetes namespace (default ```default```)
