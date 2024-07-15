
## Husky Golang Version control hook for git

Requires the [Pre-Commit.com](https://pre-commit.com/) Hook Management Framework.

## Using the hooks
You need to first install the binary from here, 

```bash
go install github.com/icehuntmen/husky`
```
In the root directory of your project, create a VERSION file and specify the initial version, for example:

**VERSION**
```
1.0.0
```

You can copy/paste the following snippet into your `.pre-commit-config.yaml` file.

**.pre-commit-config.yaml**
```yaml
repos:
  - repo: https://github.com/icehuntmen/husky
    rev: v1.0.3
    hooks:
      - id: increment-patch-version
        name: Increment Patch Version
        entry: go run github.com/icehuntmen/husky@v1.0.3 patch
        language: golang
        pass_filenames: false
        always_run: true
```

## You project
```yaml
your-project-name/
├── .pre-commit-config.yaml
├── VERSION
├── go.mod
└── main.go
```
