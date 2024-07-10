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
kubectl apply -f k8s/mosquitto-deployment.yaml

## nzen-iot-accelerometer 배포
kubectl apply -f k8s/nzen-iot-accelerometer-deployment.yaml

## nzen-iot-accelerometer 배포 확인
kubectl get pods

## nzen-iot-accelerometer 로그 확인
kubectl logs <nzen-iot-accelerometer-pod-name>
```

