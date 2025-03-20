# wsavel

This tool allows you to periodically save your WSL2 in order to roll back if necessary.

Prerequisites:

  - Only works on Windows (because during a WSL export it is unavailable)
  - WSL2 installed
  - gzip on WSL2

## How to use it?

Download the latest binary in [realease](https://github.com/mazama923/wsavel/releases).

Binary works with arguments:

  - `<wsl_name>` - Name of your wsl this parameter is a **string**
  - `<backup_path>` - Path to where to drop your backu, this parameter is a **string**
  - `<max_keep>` - Number of backups to keep, the parameter is an **int**
  - `<min_days>` - Minimum number of days that the last backup must be to trigger a new one this parameter is an **int**
  - `<compress>` - **Bool** parameter for whether or not you want to compress the backup (default false)

```bash
.\wsavel.exe Ubuntu c:\path\of\my\backup 2 14 true
```

It should therefore be triggered when your PC starts up via a shortcut in your shell:startup or by scheduled task.

## How do I roll back a backup?

For compress true:

```bash
wsl.exe --import Ubuntu C:\MyDistros\Ubuntu c:\path\of\my\backup\Ubuntu-backup-2006-03-06.tar.gz
```

For compress false:

```bash
wsl.exe --import Ubuntu C:\MyDistros\Ubuntu c:\path\of\my\backup\Ubuntu-backup-2006-03-06.tar
```

## Tech Stack

**TUI:** [bubbletea](https://github.com/charmbracelet/bubbletea) by CHARM
