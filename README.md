# nzen-iot-accelerometer

MQTT 브로커를 통해 publish되는 데이터를 파싱하여 accelerometer의 데이터를 수집하는 모듈

<img width="1331" alt="image" src="https://github.com/suslmk-lee/nzen-iot-accelerometer/assets/67575226/fd727e18-74d0-4175-b49d-ff5757220d9a">

# nzen-iot-accelerometer

nzen-iot-accelerometer는 가속도계 데이터를 MQTT 브로커로 전송하는 IoT 애플리케이션입니다. 이 프로젝트는 Kubernetes 클러스터에서 Mosquitto MQTT 브로커와 함께 배포됩니다.

## 목차
- [소개](#소개)
- [기능](#기능)
- [설치](#설치)
- [구성](#구성)
- [실행](#실행)

## 소개

nzen-iot-accelerometer는 가속도계 데이터를 수집하여 MQTT 브로커로 전송하는 Go 언어로 작성된 애플리케이션입니다. 이 프로젝트는 Kubernetes 클러스터에 배포되며, Mosquitto를 MQTT 브로커로 사용합니다.  
Test용 Publisher 모듈은 [링크](https://github.com/suslmk-lee/nzen-iot-client-test) 참조하시기 바랍니다.

## 기능

- 가속도계 데이터 수집
- MQTT 브로커로 데이터 전송
- Kubernetes 클러스터에 배포

## 설치

### 사전 요구 사항

- Kubernetes 클러스터
- kubectl
- Docker or Podman
- Go 언어 환경

### 클론

프로젝트를 클론합니다:

```sh
git clone https://github.com/suslmk-lee/nzen-iot-accelerometer.git
cd nzen-iot-accelerometer
```

### Docker 이미지 빌드

```sh
## docker 이미지 빌드
docker build -t <Docker Hub UserName>/nzen-iot-accelerometer:latest .
## docker 이미지 푸시
docker push <Docker Hub UserName>/nzen-iot-accelerometer:latest
```

## 구성

main.go: MQTT 클라이언트 및 데이터 처리 로직  
common/common.go: 설정 파일 읽기  
constraints/constraints.go: 상수 정의  
deployment/deployment.yaml: Kubernetes Deployment 및 ConfigMap  
deployment/mqtt.yaml: Mosquitto MQTT 브로커 배포 파일  


## 실행

```shell
## Mosquitto 배포
kubectl apply -f deployment/mqtt.yaml

## nzen-iot-accelerometer 배포
kubectl apply -f deployment/deployment.yaml

## nzen-iot-accelerometer 배포 확인
$ kubectl get pod
NAME                                                 READY   STATUS    RESTARTS   AGE
nzen-iot-accelerometer-deployment-54f56d4465-gtt4b   1/1     Running   0          6h49m

## nzen-iot-accelerometer 로그 확인 / subcribe 데이터 확인
kubectl logs -f <nzen-iot-accelerometer-pod-name>
```

> EdgeNode로 배포하고자 할때 deployment.yaml에 nodeSelect 부분 추가 후 배포바랍니다.  
> EdgeNode에 배포된 파드의 데이터를 조회하고자 할 때는, 배포된 엣지 노드에 접속하여 아래 명렁어로 확인 가능합니다.

```shell
## deployment.yaml  
    spec:
      containers:
      ...
      nodeSelector:
        node-role.kubernetes.io/edge:
```
```shell
## 컨테이너 아이디 조회
$ sudo crictl ps
CONTAINER           IMAGE                                                                                                                                                                     CREATED             STATE               NAME                     ATTEMPT             POD ID              POD
f07df32bd5098       44ce789b-kr1-registry.container.nhncloud.com/container-platform-registry/nzen-iot-accelerometer@sha256:4c094e63b39a54a0a7af7b1af85151cd7cd32cbcfad474bdf8431c05b377b4a1   7 hours ago         Running             nzen-iot-accelerometer   0                   94b8d232891eb       nzen-iot-accelerometer-deployment-54f56d4465-gtt4b
b1fc386136517       docker.io/kubeedge/edgemesh-agent@sha256:a116703b06f59eb3a901034992be09dc41618f10f42bb3bbf03a08df69126ba8                                                                 7 hours ago         Running             edgemesh-agent           0                   90c012a177e9b       edgemesh-agent-2vt2x

## 컨테이너 로그 확인
$ sudo crictl logs -f f07df32bd5098
Received accelerometer data: X=9.57, Y=5.55, Z=2.19
Received accelerometer data: X=0.92, Y=0.98, Z=9.48
Received accelerometer data: X=7.35, Y=6.64, Z=5.41
Received accelerometer data: X=3.10, Y=0.42, Z=6.96
Received accelerometer data: X=3.91, Y=4.26, Z=4.71
Received accelerometer data: X=3.68, Y=3.65, Z=0.72
Received accelerometer data: X=7.59, Y=4.23, Z=0.63
Received accelerometer data: X=2.19, Y=7.17, Z=3.89
Received accelerometer data: X=7.50, Y=0.71, Z=2.84
Received accelerometer data: X=5.75, Y=4.01, Z=9.36
Received accelerometer data: X=4.29, Y=9.76, Z=0.34
Received accelerometer data: X=2.46, Y=5.26, Z=3.37
Received accelerometer data: X=3.47, Y=2.82, Z=0.18
```

