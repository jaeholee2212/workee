# Git 후보와 실제 커밋의 일치

후보 준비나 커밋 생성을 수행할 때 읽는다.
후보의 기록 항목은 [커밋 후보](../../wk-work/references/work-model.md#커밋-후보)를 따른다.

## 후보 구성

1. 소유 작업의 실제 구현 공간에서 `git rev-parse HEAD`로 부모 `P`를 확인한다.
2. 현재 index와 워킹트리의 변경·소유자를 기록한다.
3. 공용 index를 보존할 수 있는 별도 임시 index를 준비한다.
4. 임시 index를 `P`에서 시작한다.
5. 포함하기로 정한 변경만 임시 index에 반영한다.
6. `git write-tree`로 전체 tree `T`를 기록한다.
7. `P`와 `T`의 diff가 전달받은 범위와 일치하는지 확인한다.

아래 변수에는 확인한 실제 값을 넣는다.
`WK_CANDIDATE_INDEX`는 아직 존재하지 않는 임시 index 경로다.
임시 경로는 제품 소스 밖에 둔다.

```sh
GIT_INDEX_FILE="$WK_CANDIDATE_INDEX" git -C "$WK_REPO" read-tree "$WK_PARENT"
# 확인한 전체 파일만 포함할 때:
GIT_INDEX_FILE="$WK_CANDIDATE_INDEX" git -C "$WK_REPO" add -- "$WK_PATH"
# 파일 일부만 포함할 때는 위 add 대신 검토한 패치를 적용한다:
GIT_INDEX_FILE="$WK_CANDIDATE_INDEX" git -C "$WK_REPO" apply --cached --binary "$WK_PATCH"
GIT_INDEX_FILE="$WK_CANDIDATE_INDEX" git -C "$WK_REPO" write-tree
```

`add`와 패치 적용은 선택한 변경을 반영하는 두 가지 방법이다.
같은 변경에 두 방법을 연속 적용하지 않는다.
전체 파일을 포함할 때도 같은 파일의 다른 사용자 변경이 섞이지 않았는지 확인한다.
삭제·새 파일·실행 권한·심볼릭 링크도 후보 내용에 포함한다.
무관한 파일까지 포함하는 전역 staging은 사용하지 않는다.

## 검증 위치

다른 미커밋 변경이 있는 공유 공간의 검사 통과만으로 후보를 검증하지 않는다.
다음 중 후보 내용의 일치를 보장할 수 있는 방법을 사용한다.

- 공유 공간을 사용하려면 관련 쓰기가 끝났고 실제 소스 전체가 `T`와 같음을 확인한다.
- 공유 공간에 후보 밖의 변경이 있으면 비어 있는 별도 위치에 후보를 checkout한다.
- Git 메타데이터가 필요한 검사이면 같은 `T`를 재현한 별도의 Git 검증 공간을 준비한다.

임시 index에서 후보를 checkout하는 예시:

```sh
GIT_INDEX_FILE="$WK_CANDIDATE_INDEX" git -C "$WK_REPO" checkout-index --all --prefix="$WK_VERIFY_DIR/"
```

`WK_VERIFY_DIR`는 절대 경로다.
검증 위치에는 후보 밖의 소스나 이전 실행의 생성물을 복사하지 않는다.
필요한 의존성은 후보의 설정·잠금 파일에 맞춰 준비한다.
미추적·무시된 파일이 검사 결과에 영향을 주면 그 입력과 이유를 명시한다.
기존의 미커밋 구현을 의존성으로 복사해 후보의 누락을 보완하지 않는다.

Git 속성에 따른 변환이나 submodule 등이 있으면 checkout 결과와 필요한 고정 버전을 확인한다.
`git archive`의 `export-ignore`·`export-subst`는 후보 파일을 생략하거나 바꿀 수 있다.
내보내기 명령의 성공만으로 tree 일치를 판정하지 않는다.
검사 중 소스가 바뀌면 그 결과를 원래 `P`·`T`의 통과 근거로 사용하지 않는다.

## 커밋 생성

커밋 전후에는 지정된 한 주체만 Git 쓰기를 수행한다.
후보 파일을 읽는 방식이면 해당 파일의 쓰기 종료도 확인한다.
기존 훅과 서명 등 저장소의 커밋 설정을 유지한다.

| 현재 상태 | 생성 방법 |
| --- | --- |
| 공용 index에 후보 밖의 staged 변경이 없다. | 확인한 경로나 패치만 stage한다. `git write-tree`가 `T`인지 확인한 뒤 일반 `git commit`을 실행한다. |
| 다른 파일의 staged 변경이 있고 후보가 전체 파일 단위다. | 워킹트리의 후보 파일로 재구성한 tree가 `T`인지 확인한다. 확인한 경로만 `git commit --only -- <경로...>`로 커밋한다. |
| 같은 파일의 다른 staged 조각을 보존해야 하거나 선택한 방법으로 `T`를 보장할 수 없다. | 기존 staging을 유지한 채 필요한 범위 조정이나 실행 조건을 반환한다. |

`--only`는 워킹트리의 파일 내용을 읽는다.
일부 조각만 검증한 파일을 이 방법으로 커밋하지 않는다.
후보에 새 파일이 있으면 `--only` 실행 전에 확인한 새 경로만 공용 index에 등록한다.
등록 전후에 후보 밖의 staging이 유지됐는지 확인한다.
커밋 직전에 현재 HEAD가 `P`인지 다시 확인한다.
정확한 후보를 확인하기 위해 다른 소유자의 staging을 reset하거나 덮어쓰지 않는다.

생성 후 실제 커밋을 확인하는 명령:

```sh
git -C "$WK_REPO" rev-parse HEAD
git -C "$WK_REPO" rev-list --parents -n 1 HEAD
git -C "$WK_REPO" rev-parse 'HEAD^{tree}'
git -C "$WK_REPO" status --short
```

새 커밋의 유일한 부모가 `P`이고 tree가 `T`인지 확인한다.
후보 밖의 staged·unstaged 변경이 보존됐는지 대조한다.
커밋 후 남은 변경은 별도 후보로 다룰 수 있도록 반환한다.
실제 커밋과 후보가 다르면 SHA와 차이를 보존하고 보완 필요로 반환한다.
