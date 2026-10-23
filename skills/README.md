# 소프트웨어 팩토리

로컬 개발자 한 명이 여러 에이전트·작업공간과 일하기 위한 재사용 가능한 한국어 스킬과 작업 절차다.
이 스킬 묶음은 런타임·데이터베이스·스케줄러를 제공하지 않는다.

## 설치

Node.js와 npm(`npx` 포함)이 필요하다. 저장소 루트에서 실행한다.

```sh
make install
```

Skills CLI의 `universal` 대상으로 모든 스킬을 `~/.agents/skills`에 확인 프롬프트 없이 설치한다.
다시 실행하면 현재 저장소 내용으로 재설치한다.

## 시작은 wk-work

[wk-work](wk-work/SKILL.md)에 목표와 원하는 진행 범위를 말한다.

- “이 요청을 논의하자. 문제와 대안을 조사하고 방향을 제안해줘.”
- “합의한 방향을 설계하고 실행 계획으로 만들어줘.”
- “이 작업의 승인된 계획을 구현하고 검증해줘.”
- “이 마스터 작업이 실제로 끝났는지 확인해줘.”
- “이 작업을 이어서 하자.”

| 단계 | 다루는 질문 | 주로 사용하는 스킬 |
| --- | --- | --- |
| [논의](wk-work/workflows/discuss.md) | 무엇을 왜 해결하고 어떤 방향을 택할까? | wk-describe, wk-diagnose, wk-delimit, wk-compare, wk-decide |
| [계획](wk-work/workflows/plan.md) | 어떤 설계로 해결하고 어떻게 구현·검증할까? | wk-plan, wk-compare, wk-diagnose, wk-decide, wk-contract |
| [구현](wk-work/workflows/implement.md) | 어떤 변경 단위를 만들고 커밋 후보로 준비할까? | wk-worktree, wk-coordinate, wk-build, wk-commit |
| [검증](wk-work/workflows/verify.md) | 변경 단위와 전체 목표의 완료 조건을 충족했나? | wk-verify, wk-review, wk-decide, wk-commit, wk-integrate |

## 함께 검토하고 실행을 맡기기

```text
목표와 요청 범위
  → 논의: 문제·대안을 사용자와 검토하며 조사
  → 계획: 설계·비용·검증 방법을 함께 검토
  → 실행 위임
  → 구현·단위 검증·로컬 커밋 반복
  → 전체 목표 검증·요청 범위의 통합
```

`wk-work`는 기존 결과를 재사용하고 필요한 단계부터 시작한다.
중요한 전제가 바뀌면 해당 논의나 계획으로 돌아간다.
논의·계획만 요청했으면 그 결과를 반환하며, 초안의 완료와 실행을 맡긴 상태는 구분한다.
실행을 맡긴 범위에서는 구현·검증·로컬 커밋을 이어 간다. 커밋 제외 등 사용자가 정한 제한을 우선한다.
기존 합의의 재사용, 직접 실행 가능한 요청, 통합·푸시의 권한은 [진행 정책](wk-work/policies/progression.md#자율-진행과-사용자-판단)을 따른다.

작은 목표는 하나의 작업으로 진행하고, 분담이 필요하면 결과 단위로 하위 작업을 둔다.
구현은 최상위 작업의 전용 브랜치·워크트리를 공유하며, 독립적인 결과는 서브에이전트에 맡길 수 있다.
개별 결과가 통과해도 합쳐진 결과가 원래 목표를 충족하는지는 따로 검증한다.
현재 상태와 산출물·이력을 작업 기록에 남겨 다음 실행에 사용한다.

## 필요한 규칙 찾기

| 알고 싶은 내용 | 원본 문서 |
| --- | --- |
| 자율 진행·권한·질문·종료 | [진행 정책](wk-work/policies/progression.md) |
| 논의·계획에서 주고받을 피드백 | [공동 탐구](wk-work/policies/collaboration.md) |
| 설정·작업 ID·산출물 저장·이력 | [장부 규약](wk-work/references/registry.md) |
| 결과 형식·작업 관계·상태·완료 판정 | [공통 작업 모델](wk-work/references/work-model.md) |
| 구현 공간 생성·재사용 | [wk-worktree](wk-worktree/SKILL.md) |
| 위임·진행 추적·결과 회수 | [구현 조율](wk-work/workflows/implement.md#실행자-선택) |
| 커밋 후보와 실제 커밋의 일치 | [wk-commit](wk-commit/SKILL.md) |

## 개별 스킬 사용과 작성

하위 스킬도 직접 요청할 수 있다. 맡긴 결과를 반환하며 다른 단계로 자동 확장하지 않는다.
최초 설정은 [wk-setup-workspace](wk-setup-workspace/SKILL.md), 작성 범위·자원 조정은 [wk-coordinate](wk-coordinate/SKILL.md),
별도 검토는 [wk-review](wk-review/SKILL.md), 중단·인계는 [wk-handoff](wk-handoff/SKILL.md)를 사용한다.

일부 스킬만 설치할 때도 공통 규약을 제공하는 `wk-work`는 함께 설치한다.
전체 흐름을 사용하려면 단계 표의 스킬과 설정·조정·검토·인계 스킬을 함께 설치한다.
필요한 스킬이 없으면 누락을 알리고 해당 절차를 임의로 대체하지 않는다.
스킬을 수정할 때는 [작성 규칙](AUTHORING.md)을 따른다.
