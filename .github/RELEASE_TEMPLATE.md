## Installation

You can use `go` to build SSH Manager locally with:

```shell
go install github.com/omegion/db-backup@latest
```

Or, you can use the usual commands to install or upgrade:

On OS X

```shell
sudo curl -fL https://github.com/omegion/db-backup/releases/download/{{.Env.VERSION}}/db-backup-darwin-amd64 -o /usr/local/bin/db-backup \
&& sudo chmod +x /usr/local/bin/db-backup
```

On Linux

```shell
sudo curl -fL https://github.com/omegion/db-backup/releases/download/{{.Env.VERSION}}/db-backup-linux-amd64 -o /usr/local/bin/db-backup \
&& sudo chmod +x /usr/local/bin/db-backup
```

On Windows (Powershell)

```powershell
Invoke-WebRequest -Uri https://github.com/omegion/db-backup/releases/download/{{.Env.VERSION}}/db-backup-windows-amd64 -OutFile $home\AppData\Local\Microsoft\WindowsApps\db-backup.exe
```

Otherwise, download one of the releases from the [release page](https://github.com/omegion/db-backup/releases/)
directly.

See the install [docs](https://github.com/omegion/db-backup) for more install options and instructions.

## Changelog
