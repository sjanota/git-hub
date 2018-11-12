# git-pr

git-pr is commandline tool to help you deal with your pull request routine. I came to the idea when I was under heavy 
load of many pull requests at the same time. It was designed as a tool to help track status of allow those PRs as dozen 
 tabs in browser made it difficult to have overview of what's going on.

## Commands

### `git pr`

Lists all locally available PRs. Output includes PR number, branch and title. If there is PR for current branch
it's row will be preceded with `*`.

Examples:
```
$ git pr
  1      pr                               Test commit
* 2      my-awesome-feature               I'm testing now PRs
```

```
$ git pr --details
  1      pr                               Test commit
  └── Awaits CI build
* 2      my-awesome-feature               I'm testing now PRs
```

### `git pr fetch`

Downloads information about pull requests assigned to you. Use it to synchronise local data of PRs, just like 
you would use `git fetch` for branches.

Usage:
```
git pr fetch # fetch PRs' data from origin
git pr fetch other # fetch PRs' data from remote 'other'
git pr fetch --all # fetch PRs' data from all remotes
```

### `git pr clean`

It is a reverse of `git pr fetch`: it will remove all locally stored data about PRs.

### `git pr open`

Opens PR for currently checked out branch in your default browser.

### `git pr comment`

Allows you to edit comment locally before publishing it. The idea is to make notes with progress on 
your local machine, where you develop and once you're ready (e.g. before leaving office) publish them to GitHub.

Usage:
```
git pr comment # will open text editor (as per 'git commit') to modify comment
git pr comment -m "Awaits CI build" # will replace already existing comment with given text
git pr comment -m "Awaits CI build" --append # will append given text to already existing comment in new line
```

### `git pr status`

It's counterpart of `git status` and shows information about PR related to current branch.

Example:
```
$ git pr status 
On pull request sjanota/git-pr#1
    Test commit

Pull request sjanota/git-pr#1 is out-of-sync
    (use "git pr push" to push comment to GitHub)
Comment:
    Awaits CI build
    Awaits review from @sjanota
```

### `git pr push`

Publishes comment on current PR to GitHub.