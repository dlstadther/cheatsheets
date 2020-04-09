# Git

---

## Terminology

| Term | Description |
|------|-------------|
| local | on the harddrive of the computer running your git commands |
| remote | central scm server (i.e. Github, Gitlab, Bitbucket, etc.) |
| origin | traditional tag associated with a repo's remote |
| commit | set of file changes with a message (think of this similar to a checkpoint save) |
| branch | ... |
| merge | ... |
| rebase | ... |
| push | upload your local commits to a remote |
| pull | download remote commits to your local |
| clone | download a remote repo to your local for the first time |
| staged | local files which have been prepared to be included in the next commit (the targets of `git add`) |
| unstaged | local files which have been included in a prior commit with changes not yet ready to be committed |
| untracked | local files within a repo's directory which have never been included in a commit |

---

## Commands

### Set up ssh key for use with ssh git auth
```shell
ssh-keygen -t rsa -b 4096 -C "yourEmail@domain.com"
# provide optional passphrase
eval "$(ssh-agent -s)"
ssh-add -K ~/.ssh/id_rsa

# copy public key to scm provider
# follow their instructions for where you set in your auth profile
cat ~/.ssh/id_rsa.pub
```

### Start a new local repo
```shell
mkdir localRepo
cd ./localRepo
git init
```

### Pull Remote Repo to Local (for the first time)
```shell
# assuming Github project fakerepo, owned by dlstadther
# with ssh auth
git clone git@github.com:dlstadther/fakerepo.git

# with https basic auth
git clone https://github.com/dlstadther/fakerepo.git

# ssh auth with custom directory name
git clone git@github.com:dlstadther/fakerepo.git foobar
```

### Update local references to remotes
```shell
git fetch origin
```

### Start new branch
```shell
# create and checkout a new branch "feature/my-feature" based off local master
git checkout -b feature/my-feature master
```

### View all locally tracked branches
```shell
git branch -a
```

### Change branches
```shell
# assumes already on master, and the existence of develop
git checkout develop

# switch back to previous branch
git checkout -
```

### Update branch with remote master
```shell
# without updating local master w/o merge commit
git fetch origin
git rebase origin/master

# with local update and merge commit
git checkout master
git pull origin master
git checkout -
git merge master
```

### Rename branch
```shell
git branch -m oldname newname

# rename current branch
git branch -m newname
```

### Trash all tracked (and uncommitted) changes
```shell
git reset --hard
```

### Trash uncommitted changes for a single tracked file
```shell
git checkout HEAD -- path/to/file
```

---
