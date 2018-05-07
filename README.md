# **아주대학교 2018-1학기 캡스톤디자인 Connection팀**
[![Build Status](https://travis-ci.org/AJOU-Connection/StopBus_Server.svg?branch=master)](https://travis-ci.org/AJOU-Connection/StopBus_Server)
[![Go Report Card](https://goreportcard.com/badge/github.com/AJOU-Connection/StopBus_Server)](https://goreportcard.com/report/github.com/AJOU-Connection/StopBus_Server)
[![codecov](https://codecov.io/gh/AJOU-Connection/StopBus_Server/branch/master/graph/badge.svg)](https://codecov.io/gh/AJOU-Connection/StopBus_Server)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## **프로젝트 팀**
- 학과: 정보통신대학 소프트웨어학과
- 담당 교수: 윤대균
- 학부생: 13 박명진, 13 김범근, 14 김지선, 15 김희연

## **프로젝트 개요**
버스는 대중교통 이용객 40%가 이용할 정도로 많은 사람들에게 사랑을 받고 있는 대중교통이다. 다른 교통 수단에 비해서 가격이 저렴하고, 많은 목적지를 제공한다. 하지만 버스는 여러 문제점을 가지고 있다.
![서울 시내버스 불편 민원 1위](http://news.tongplus.com/site/data/img_dir/2016/11/04/2016110401982_0.jpg)

여러 문제점들 중에서 가장 많은 부분을 차지하는 것은 **무정차 통과**이다. 이 문제를 해결하기 위해서 StopBus는 승하차벨 개념을 도입하여 승객과 버스기사의 편의를 향상시키고자 한다.

## **StopBus_Server 기능**
1. GBUS API를 사용하여 User, Diver, Panel로 실시간 버스정보 제공
2. FCM(Firebase Cloud Messaging)을 활용한 Push 알람 제공

## **사용 기술**
1. GitHub: 프로젝트 소스코드 버전 관리 
2. Trello: 팀 프로젝트 관리
3. Slack: 팀 프로젝트 간 채팅 프로그램 - Trello & GitHub와 연동
4. Golang: 구글이 개발한 프로그래밍 언어로 가비지 컬렉션 기능이 있고, 병행성(concurrent)을 잘 지원하는 컴파일 언어
5. FCM(Firebase Cloud Messaging): 무료로 메시지를 안정적으로 전송할 수 있는 교차 플랫폼 메시징 솔루션
6. MySQL: 버스의 승하차 유무를 저장하기 위한 데이터베이스