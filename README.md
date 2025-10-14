Got it 👍 — here’s a concise, neutral version you can drop directly into your `README.md`:

---

````markdown
# Snippetbox

## Git Worktree & Stacked PR Workflow

This project uses **Git worktrees** to manage stacked pull requests.  
Each feature branch lives in its own directory so you can develop multiple branches in parallel.

### Create a new feature branch

```bash
git worktree add ../worktrees/snippetbox-feature-a -b feature-a main
```

### Create a stacked feature branch (depends on feature-a)

```bash
git worktree add ../worktrees/snippetbox-feature-b -b feature-b feature-a
```

### Rebase after an earlier PR merges

```bash
cd ../worktrees/snippetbox-feature-b
git fetch origin
git rebase origin/main
```

### Share the same `.env` file across worktrees

Run this from the main worktree to symlink the shared `.env`:

```bash
ln -s "$(pwd)/.env" ../worktrees/snippetbox-feature-b/.env
```
