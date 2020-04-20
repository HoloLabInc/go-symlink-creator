# go-symlink-creator
Go-symlink-creator is a symbolic link creating tool.  
Settings file is written in YAML format.

## Usage
### Download
Download `go-symlink-creator.exe` from Releases page.

### Create settings file

```yaml
symlinks:
  # Single destination and target
  - src: \path\to\source\folder
    dest: \path\to\destination\folder
    target: targetFile

  # Multiple destinations and targets
  - src: \path\to\source\folder
    dest:
      - \path\to\destination\folder1
      - \path\to\destination\folder2
    target:
      - targetFile1
      - targetFile2

  # For Unity file or folder
  - include-meta-file: true    # Create symlink for .meta as well
    src: \path\to\source\folder
    dest: \path\to\destination\folder
    target: UnityFolder
```

### Execute
```
> go-symlink-creator.exe settings.yaml
```

## Author
Furuta, Yusuke ([@tarukosu](https://twitter.com/tarukosu))

## License
MIT