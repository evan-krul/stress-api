<div id="top"></div>

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

[comment]: <> ([![Contributors][contributors-shield]][contributors-url])

[comment]: <> ([![Forks][forks-shield]][forks-url])

[comment]: <> ([![Stargazers][stars-shield]][stars-url])

[comment]: <> ([![Issues][issues-shield]][issues-url])

[comment]: <> ([![MIT License][license-shield]][license-url])

[comment]: <> ([![LinkedIn][linkedin-shield]][linkedin-url])



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/github_username/repo_name">
    <img src="https://golang.org/lib/godoc/images/go-logo-blue.svg" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Stress API Pod</h3>

  <p align="center">

[comment]: <> (    CgroupV2 PSI Sidecar can be deployed on any kubernetes pod with access to cgroupv2 PSI metrics.)
  </p>
</div>



[comment]: <> (<!-- TABLE OF CONTENTS -->)

[comment]: <> (  <summary>Table of Contents</summary>)

[comment]: <> (  <ol>)

[comment]: <> (    <li>)

[comment]: <> (      <a href="#about-the-project">About The Project</a>)

[comment]: <> (      <ul>)

[comment]: <> (        <li><a href="#built-with">Built With</a></li>)

[comment]: <> (      </ul>)

[comment]: <> (    </li>)

[comment]: <> (    <li>)

[comment]: <> (      <a href="#getting-started">Getting Started</a>)

[comment]: <> (      <ul>)

[comment]: <> (        <li><a href="#prerequisites">Prerequisites</a></li>)

[comment]: <> (        <li><a href="#installation">Installation</a></li>)

[comment]: <> (      </ul>)

[comment]: <> (    </li>)

[comment]: <> (    <li><a href="#usage">Usage</a></li>)

[comment]: <> (    <li><a href="#roadmap">Roadmap</a></li>)

[comment]: <> (    <li><a href="#contributing">Contributing</a></li>)

[comment]: <> (    <li><a href="#license">License</a></li>)

[comment]: <> (    <li><a href="#contact">Contact</a></li>)

[comment]: <> (    <li><a href="#acknowledgments">Acknowledgments</a></li>)

[comment]: <> (  </ol>)



<!-- ABOUT THE PROJECT -->
## About

This is a docker container that can be deployed to run any stress-ng stress tests on the container over an exposed api.


### Built With

* [Go Lang](https://golang.org/)
* [Gorilla Mux](https://github.com/gorilla/mux)
* [Stress-ng](https://wiki.ubuntu.com/Kernel/Reference/stress-ng)


<!-- GETTING STARTED -->

[comment]: <> (## Getting Started)

## Build Image
There are two docker files one for regular deployment and the other for debugging.

#### Regular image
1. `docker build -f ./Dockerfile . -t evankrul/stress-api:v1`
2. `docker push evankrul/stress-api:v1`
#### Debug image
1. `docker build -f ./Dockerfile.debug . -t evankrul/stress-api:v1`
2. `docker push evankrul/stress-api:v1`
#### Port
Set `PORT` env var to specify the api port.

<!-- USAGE EXAMPLES -->
## Usage
#### Example kubernetes yaml
```yaml
---
#Namespace
apiVersion: v1
kind: Namespace
metadata:
  name: stress-ng
---
#Stressor 1 Deployment
apiVersion: apps/v1
kind: Deployment
metadata:
  name: stress-ng
  namespace: default
spec:
  selector:
    matchLabels:
      app: stress-ng
  replicas: 1
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 0
  template:
    metadata:
      labels:
        app: stress-ng
    spec:
      containers:
        - name: stress-api
          image: evankrul/stress-api:v1
          imagePullPolicy: Always
          ports:
            - containerPort: 2335
              name: stress
          env:
            - name: PORT
              value: "2335"
          resources:
            requests:
              cpu: 1
              memory: "500Mi"
            limits:
              cpu: 1
              memory: "500Mi"
---
apiVersion: v1
kind: Service
metadata:
  name: cgroup-monitor-sc #this will be the Domain name
  namespace: default
spec:
  selector:
    app: stress-ng
  ports:
    - name: stress
      port: 2335
      targetPort: 2335
  type: LoadBalancer
```
### API
There are a few endpoints:
- `/` Homepage
- `/health` K8s health endpoint
- `/stress` Stress API endpoint (POST)
#### Stress API
To start a stressor send a POST request to `/stress` with a JSON body following this schema:

```json
{
  "args": [
    "..."
  ],
  "timeout": "..."
}
```

When the server receives a request it creates a new async go routine calling the following:

`stress-ng [args] --timeout=[timeout]`

Due to the async nature of the go routines multiple request can be sent one-after-another to stress the container over 100%*

_* Assuming the container has fixed resources deploying one stressor will saturate the resource to 100%. Running more stressors will add more pressure to the resource. Beyond some linux scheduler fluctuations you will not actually go over 100%._ 

#### EG Calls
`stress-ng --cpu 0 --timeout 5m`
```json
{
	"args": [
		"--cpu",
		"0"
	],
	"timeout": "5m"
}
```
`stress-ng --memrate 3 --memrate-bytes 5G --timeout 5m`
```json
{
	"args": [
		"--memrate",
		"3",
		"--memrate-bytes",
		"5G"
	],
	"timeout": "5m"
}
```

<!-- CONTACT -->
## Contact

Evan Krul - [Website](https://krul.ca)


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/github_username/repo_name.svg?style=for-the-badge
[contributors-url]: https://github.com/github_username/repo_name/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/github_username/repo_name.svg?style=for-the-badge
[forks-url]: https://github.com/github_username/repo_name/network/members
[stars-shield]: https://img.shields.io/github/stars/github_username/repo_name.svg?style=for-the-badge
[stars-url]: https://github.com/github_username/repo_name/stargazers
[issues-shield]: https://img.shields.io/github/issues/github_username/repo_name.svg?style=for-the-badge
[issues-url]: https://github.com/github_username/repo_name/issues
[license-shield]: https://img.shields.io/github/license/github_username/repo_name.svg?style=for-the-badge
[license-url]: https://github.com/github_username/repo_name/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: https://golang.org/lib/godoc/images/go-logo-blue.svg