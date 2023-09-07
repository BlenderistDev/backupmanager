# backupmanager
CLI tool for backup managing

## commands

### replace
replaces old backups from source dir to storage dir

##### flags:
--source - path to source dir

--storage - path to storage dir

--source-keep - days to keep backups in source dir. Default 14

### delete
deletes old backups from storage dir
At least one backup will be left for each week

##### flags:
--storage - path to storage dir

--storage-keep - days to keep backups in storage dir. Default 30

### emptydir
deletes empty directories

##### flags:
--storage - path to find empty directories

### serve
serves source and storage paths with replace, delete and emptydir commands

##### flags:
--source - path to source dir

--storage - path to storage dir

--source-keep - days to keep backups in source dir. Default 14

--storage-keep - days to keep backups in storage dir. Default 30

--sleep - hours to sleep between serve iterations. Default 1

### help
shows help
