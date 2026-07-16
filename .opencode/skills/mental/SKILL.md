---
name: mental
description: >-
  Cross-session memory manager for LLM workflows. Persists context,
  tasks, and topic-indexed checkpoints across sessions and agents.
  Activate on: "save memory", "checkpoint", "load context", "what were
  we doing", "search memory", "add task", "mark done", "show tasks",
  "init memory", "start tracking", "mental mem", mental session search.
license: Apache-2.0
compatibility: >-
  Requires mental binary in PATH. Install:
  go install github.com/mrbrandao/mental@latest
  MENTAL_DIR defaults to ~/.local/share/mental.
allowed-tools: Bash
metadata:
  author: mrbrandao
  version: "0.1.0"
---

# mental

Cross-session memory management for LLM workflows.

## Trigger Vocabulary

| User says | Command |
|-----------|---------|
| "save memory", "checkpoint", "wrapping up" | `mental mem save` |
| "load context for X", "what were we doing on X" | `mental mem load X` |
| "search memory for X" | `mental mem search X --project <name>` |
| "add task", "create task for ..." | `mental mem task add --project <name> <title>` |
| "mark done", "finished X", "task X is done" | `mental mem task done --project <name> <id>` |
| "show tasks", "what's pending" | `mental mem task list --project <name>` |
| "init memory", "start tracking", "new project" | `mental mem init <name>` |
| "search my sessions for X" | `mental session search -s "X"` |

## Step 1 — Session Start: Load Context

When starting work on a known project, load its memory:

```bash
mental mem load <project>
```

Read the output. It contains MEMORY.md (current state) and the task list.
Use this to orient the session without asking the user to re-explain context.

## Step 2 — During Session: Manage Tasks

Add tasks as they are identified:

```bash
mental mem task add --project <name> "Implement rollback script"
# → Added task #t001: Implement rollback script
```

Mark tasks done when completed:

```bash
mental mem task done --project <name> t001
# → Marked #t001 as done
```

List all tasks at any time:

```bash
mental mem task list --project <name>
```

## Step 3 — Session End: Save Checkpoint

At session end (user says "wrapping up", "done for today", "save memory"):

Write the following block and pipe it to `mental mem save`:

```
project: <project-name>
session.id: <current-session-id>
session.client: opencode
session.model: <current-model>
session.dir: /path/to/project
topics: <topic1>, <topic2>, <topic3>
files: <file1>, <file2>
summary: <one sentence describing this session>
memory:
# <project-name>
status: active
updated: <ISO datetime>

## Context
<3-5 sentences: current state of the project>

## Decisions
- <decision>: <rationale>

## Next Steps
- [ ] <next action>

## Related
<related projects if any>

## Search
Topics and past sessions: topics.yaml
Command: mental mem search "<topic>" --project <name>
---
## What We Did
<paragraph summary of what happened this session>

## Decisions Made
- <decision and why>

## Open Questions
- <unresolved question>

## Handoff
<1-2 sentences: exactly where to pick up next session>
```

Then pipe to save:

```bash
# Write the block above to /tmp/mental-save.txt then:
mental mem save < /tmp/mental-save.txt
```

## Step 4 — Cross-Session Search

To find what was discussed in past sessions:

```bash
mental session search -a opencode -s "rollback strategy"

# Search memory topics:
mental mem search "rollback strategy" --project foo
```

## Step 5 — Init a New Project

When starting a new project for the first time:

```bash
mental mem init my-project
# Creates: ~/.local/share/mental/projects/my-project/
# Files:   MEMORY.md, tasks.yaml, topics.yaml, checkpoints/
```

## Gotchas

- Always load memory at session start before doing any work.
- Topics must be comma-separated in the save input.
- The `--project` flag is required for search, task, and list commands.
- MENTAL_DIR can be overridden: `MENTAL_DIR=/custom/path mental mem load foo`
- Use `mental extensions list` to see available extensions.
