# taskTracker

## Install

### Use as app
```bash
git clone git@github.com:dejitarudemon/taskTracker.git
```

### Use as package
```bash
go get github.com/dejitarudemon/taskTracker
```

## Usage

### Build
```bash
cd taskTracker
go build
```

### Init a datafile

```bash
./taskTracker.exe init
```

### Create a new task

```bash
./taskTracker.exe add <task_description>
```

For exemple:
```bash
./taskTracker.exe add "new task"
```

### Update the task

```bash
./taskTracker.exe update <task_id> <task_description>
```

For exemple:
```bash
./taskTracker.exe update 1 "The task with a new description"
```

### Mark the task

#### As In-Progress
```bash
./taskTracker.exe markInProgress <task_id>
```

For exemple:
```bash
./taskTracker.exe markInProgress 1
```

#### As Done
```bash
./taskTracker.exe markDone <task_id>
```

For exemple:
```bash
./taskTracker.exe markDone 1
```

### Delete the task

```bash
./taskTracker.exe delete <task_id>
```

For exemple:
```bash
./taskTracker.exe delete 1
```

### List the tasks

```bash
./taskTracker.exe list --status <status>
```

--status - is an optional flag