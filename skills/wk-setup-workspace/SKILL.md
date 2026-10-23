---
name: wk-setup-workspace
description: 프로젝트 메타데이터를 처음 만들 때 사용한다.
disable-model-invocation: true
---
## 작업 순서
1. 기존 저장소와 `wk-factory.yaml`부터 확인한다.
2. 파일이 없으면 기존 저장소와 장부에서 프로젝트 식별 정보를 확인한다.
3. 새 프로젝트의 ID는 사용자 지정값이나 기존 식별 정보를 우선한다.
4. 지정값이 없으면 저장소 이름에서 충돌하지 않는 ID를 정한다.
5. 기존 장부가 있다는 근거는 있으나 원본 후보가 갈리면 후보와 확인에 필요한 정보를 반환한다.
6. 확인된 정보와 아래 기본값으로 `wk-factory.yaml`을 만든다.
7. 기존 파일을 허락 없이 덮어쓰지 않는다.
8. `wk-factory.yaml`에서 필요한 폴더를 생성한다.
9. `registry.root` 폴더에 작업 개요와 공통 규칙을 저장하는 `project.md`를 만든다.

기존 장부가 없는 새 프로젝트는 아래 기본값으로 초기화를 이어 간다.
기존 장부의 원본이 미확정이면 새 설정과 장부를 만들지 않는다.
필요한 사용자 판단은 [자율 진행과 사용자 판단](../wk-work/policies/progression.md#자율-진행과-사용자-판단)을 따른다.

## `wk-factory.yaml` 예시
새 설정의 기본 장부 경로는 `~/.cache/workee/registries/<PROJECT_ID>`다.

```
schema: 1
project:
  id: <PROJECT_ID>
registry:
  root: ~/.cache/workee/registries/<PROJECT_ID>
  guide: ~/.cache/workee/registries/<PROJECT_ID>/project.md
work:
  id_prefix: <TO_UPPERCASES_OF_THE_FIRST_THREE_CODE_OF_PROJECT_ID>
glossary:
  path: docs/GLOSSARY.md
```

## 작업 개요와 공통 규칙 `project.md`
작업 식별과 색인 항목은 [작업 식별](../wk-work/references/registry.md#작업-식별)과 [색인 관리](../wk-work/references/registry.md#색인-관리)를 따른다.
작업 생성 시 아래 색인에 항목을 추가한다.
단계와 상태의 원본은 각 작업의 `work.md`에 둔다.

```
# <PROJECT_ID>
- 파일 쓰기: 한 파일에 한 작성자

## 작업 색인
| 작업 ID | 제목 | 상위 작업 | work.md |
| --- | --- | --- | --- |
```

## 결과

- 설정 원본과 장부 가이드가 실제 존재하는지 확인한다.
- 프로젝트 ID·장부 경로·작업 접두사와 초기화 결과를 [스킬 실행 결과](../wk-work/references/work-model.md#스킬-실행-결과)로 반환한다.
