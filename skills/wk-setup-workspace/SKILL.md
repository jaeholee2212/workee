---
name: wk-setup-workspace
description: 프로젝트 메타데이터를 처음 만들 때 사용한다.
disable-model-invocation: true
---
## 작업 순서
1. 기존 저장소와 `wk-factory.yaml`부터 확인한다.
2. 파일이 없으면 짧은 기본 설정을 제안한다.
3. 확인된 정보로 한 파일만 만든다. 프로젝트 ID를 추천하고 사용자로 부터 확인을 받는다.
4. 기존 파일을 허락 없이 덮어쓰지 않는다.
5. `wk-factory.yaml`에서 필요한 폴더를 생성하자.
6. 팩토리 설정 `wk-factory.yaml`에서 registry.root 폴더에 작업 개요와 공통 규칙을 저장하는 `project.md`를 만든다.

## `wk-factory.yaml` 예시
```
schema: 1
project:
  id: <PROJECT_ID>
registry:
  root: ~/.cache/rfactory/registries/<PROJECT_ID>
  guide: ~/.cache/rfactory/registries/<PROJECT_ID>/project.md
work:
  id_prefix: <TO_UPPERCASES_OF_THE_FIRST_THREE_CODE_OF_PROJECT_ID>
glossary:
  path: docs/GLOSSARY.md
```

## 작업 개요와 공통 규칙 `project.md`
작업 식별과 색인 항목은 [장부 규약](../wk-work/references/registry.md)을 따른다.
작업 생성 시 아래 색인에 항목을 추가한다.
단계와 상태의 원본은 각 작업의 `work.md`에 둔다.

```
# <PROJECT_ID>
- 파일 쓰기: 한 파일에 한 작성자

## 작업 색인
| 작업 ID | 제목 | 상위 작업 | work.md |
| --- | --- | --- | --- |
```
