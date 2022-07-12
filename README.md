# ebest

## 소개
[ebest-sdk](https://github.com/sangx2/ebest-sdk) 를 이용한 윈도우 전용 api 서버

## 주의사항
- golang windows x86(x86-64는 동작하지 않음) 파일 설치
- xing api 설치
- 모의투자 도메인 : demo.ebestsec.co.kr
- 실투자 도메인 : hts.ebestsec.co.kr

## 프로세스 구조
ebest에서는 07시에서 07시 30분 사이에 재접속이 요구된다. 하지만 제공하는 shutdown 및 connect 호출 시에 오류를 발생한다.

따라서 2개의 프로세스가 존재하고 그 역할은 다음과 같다.

###manager.exe
trader 프로세스의 시작&종료 관리
- manager.exe 실행 시 기존 trader 프로세스 종료, config에 설정된 지연시간(10s) 이후 trader 프로세스 시작
- 오전 7시 trader 프로세스 종료
- 오전 7시 30분 trader 프로세스 시작

#### 설정 파일
```yaml
{
  "ManagerSettings": {
    "LogLevel": "info",
    "LogFileName": ".\\logs\\manager.log"
  }
}
```

### trader.exe
거래에 필요한 기능을 구현

#### 설정 파일
```yaml
{
  "AppSettings": {
    "ID": "",
    "Passwd": "",
    "CertPasswd": "",
    "Server":"demo.ebestsec.co.kr",
    "ResPath": "C:\\eBEST\\xingAPI\\Res\\",
    "DataPath": ".\\data",
    "QueueSize": 1000
  },
  "AccountSettings": {
    "Accounts": {
      "00000000000": "0000"
    }
  },
  "SQLSettings": {
    "Enable": false,
    "DataSource": ""
  },
  "APISettings": {
    "KeyAuth": {
      "Enable": false,
      "Key": "authKey"
    },
    "TLS": {
      "Enable": false,
      "CertPEMPath": ".\\certs\\cert.pem",
      "KeyPEMPath": ".\\certs\\privkey.pem"
    },
    "Port": "8000"
  },
  "TraderSettings": {
    "LogLevel": "info",
    "LogFileName": ".\\logs\\trader.log"
  }
}
```

## 기능
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/10764121-c42a84fb-d12c-40eb-bf27-1354038e460a?action=collection%2Ffork&collection-url=entityId%3D10764121-c42a84fb-d12c-40eb-bf27-1354038e460a%26entityType%3Dcollection%26workspaceId%3D94a5776e-50ae-43ac-a5b3-83c72a7d2b4a)

[계좌]
- 전체 계좌 조회

[자산]
- 전체/계좌별 자산 조회

[잔고]
- 계좌별 잔고 조회

[종목]
- 전체 종목 조회
- 종목 조회
- 매수/매도/정정/취소 요청

[주문]
- 매수/매도/정정/취소 요청 목록 조회
- 매수/매도/정정/취소 요청 조회
- 접수/체결/정정/취소/거부 조회

## 개발 예정
현재 모의투자 예외처리(callback.go)만 되어 있음. 실거래 테스트 요구됨.

[자산]
- 기간별 자산 정보

[잔고]
- 기간별 잔고 정보

[종목]
- 실시간 호가 잔량(HA/H1)
- 실시간 체결(S3/K3)
- 실시간 거래원(OK/K1)

[주문]
- 실시간 접수/체결/정정/취소/거부

[기타]
- 실시간 뉴스 정보

[설정]
- 데이터베이스 연동