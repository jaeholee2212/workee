# workee

Go CLI와 재사용 가능한 Workee 스킬을 함께 개발하는 프로젝트입니다.
Workee 작업 절차를 실제 개발에 적용하는 테스트베드로도 사용합니다.

## Go CLI 시작하기

Go 1.22 이상이 필요합니다. Makefile 명령은 `make`가 설치된 환경에서 사용할 수 있습니다.
저장소 루트에서 실행합니다.

```sh
go run ./cmd/workee
go run ./cmd/workee --help
```

현재 CLI는 기본 실행 시 `workee`를 출력합니다. `-h` 또는 `--help`로 도움말을
표시하며, 아직 업무용 하위 명령은 제공하지 않습니다.
도움말과 오류는 stderr로 출력합니다. 정상 실행과 도움말의 종료 코드는 0,
알 수 없는 옵션이나 위치 인수의 종료 코드는 2입니다.

빌드 후 실행하려면:

```sh
go build -o bin/workee ./cmd/workee
./bin/workee
./bin/workee --help
```

같은 작업을 Makefile로 실행할 수 있습니다.

```sh
make build  # bin/workee 빌드
make run    # 소스에서 실행
make test   # 테스트
```

검사 명령:

```sh
go test ./...
go vet ./...
```

## 구조

```text
cmd/workee/
  main.go       # CLI 진입점과 인수 처리
  main_test.go  # 출력·도움말·실패 종료 코드 검사
go.mod          # Go 모듈 (외부 의존성 없음)
skills/         # Workee 스킬과 작업 절차
```

## Workee 스킬 설치

Node.js와 npm이 필요합니다. 기존 `install` 대상은 Go CLI가 아닌 Workee 스킬을 설치합니다.

```sh
make install
```

자세한 작업 흐름과 스킬 설명은 [skills/README.md](skills/README.md)를 참고하세요.
