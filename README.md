# rugby-scores

This is a survey of the current state of the Google Cloud Platform and how efficiently one can bang out a modern application.

Broad brushstrokes of what I'd like to put together (some of it is new to me):

- Continuous deployment with Jenkins
- GKE/Docker/Kubernetes for infrastructure provisioning
- "New" Google Cloud Endpoints
- golang
- GAE - flexible environment
- Datastore without Objectify (How do you live without an ORM? You're living like animals...)
- Node.js in GKE
- react/redux for web client
- react-native for mobile client

## Update 1 (4/21/17):
Configured a Google Compute Engine set of three cos-stable-56-9000-84-2 VMs on us-east1-d n1-standard-1 (1 vCPU, 3.75 GB memory)
  - Roughly followed this: [https://cloud.google.com/solutions/jenkins-on-container-engine-tutorial]
  - Two Instance Groups are created, with a separate network defined for the Jenkins install. While only one Node Pool is created in Kubernetes, we do run two different things on the same billable resources:
    - for Jenkins, which is going to be a relatively static environment, we push a volume image up. Jenkins also needs a persistent disk.
    - For the actual API we are supposed to be developing I used docker to build an image and push that up to the Container Registry (rather than the Compute Engine images) since they are going to be changing and we need to let Jenkins manage the images. Note that both Instance Groups map onto the same infrastructure. If we were running multiple client projects, it would be better to have separate hardware for the build server and applications server but I appreciate the abstraction of capability away from hardware. At any rate, it is a simple change to make if it becomes necessary.
  - The Container Engine hosts a single Container Cluster with just one Node Pool. If we needed capabilities hosted outside of us-east1-d, we would need more Node Pools in the cluster.
  - The boundary between Container Engine and Kubernetes is the Node Pool (GCE) and Nodes (K8S). There are three nodes that host a variety of pods, which handle cross-cutting things like the load balancing, DNS, logging and heap management as well as application functionality. The Jenkins pod exist in only one of the nodes. By scaling the scores-api-production deployment up to 3, we get two pods in one node, one in a second and none in the third. There is however a canary pod in the third node. I need to play around some more with the scaling to see how pods are mapped to nodes.

The API pods (scores-api-* ) are going to have two images. Since they won't ever be deployed independently (a pod is the smallest unit of deployment) we can configure our ReplicaSet with two images see [https://cloud.google.com/endpoints/docs/quickstart-container-engine]:
  - The Google Cloud Endpoint proxy (doesn't change during development)
  - The golang:onbuild-based image with the service implementation and unit tests (see [https://hub.docker.com/_/golang/])

CD is set up with Jenkins:
  - Roughly followed: [https://cloud.google.com/solutions/continuous-delivery-jenkins-container-engine]
  - Google Source Repository synchs with this github repo
  - pushing to git triggers Jenkins to:
    - pull from branch, build new image, tagged with app name, branch and environment
    - run tests with `go test`
    - on success, push image to Google Container Registry
    - deploy pod to appropriate node

# Update 2 (4/24/17):

Establishing a containerized REST api that can be deployed behind the Google Container Engine ESP proxy. The idea is pretty cool (the framework handles the scaling, edge server topology, security, etc.). Compared to the AWS counterpart there are some things lacking. For instance, the API key management doesn't provide any real reporting or throttling that would allow you to control and bill integration partners leveraging your services.

I've started with the `golang:onbuild` docker image but spent some time looking at `golang:<version>` and `golang:alpine` to understand how the golang ecosystem supports containerization. Package management and imports, particularly within your own golang app can be subtly tricky.

There is so much boilerplate code in standing up simple REST apis that I looked at a couple of options for jumpstarting the work. I looked at [https://github.com/tideland/gorest](tideland/gorest) for a bit but ended up digging into [http://editor.swagger.io/#!/](swagger) more. Rather than using the native codegen in Swagger I am going to work with [https://goswagger.io/](go-swagger) which seems more robust. Code generation has the pain points of creating more bloated source and having an up-front ramp-up cost, but once established correctly you have a much nicer development experience.

My challenge in go-swagger was that it wants to run in https by default and that requires X.509 certs to be installed in the image. GKE has it's own idea of "secrets" that I didn't spend the time to understand how we can propogate into the containers and then access from golang binaries. My thought was that SSL should be established between the clients and the ESP proxy and that we can just run over HTTP inside the cluster. It saves a bit of processing overhead and, since the api implementation containers aren't externally addressable, shouldn't reduce the effective security.

 
