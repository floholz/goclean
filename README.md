# Welcome to goclean 👋
[![GitHub tag (with filter)](https://img.shields.io/github/v/release/floholz/goclean?label=latest)](https://github.com/floholz/goclean/releases/latest)
[![GitHub License](https://img.shields.io/github/license/floholz/goclean)](./LICENSE)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/floholz/goclean?logo=go&labelColor=gray&label=%20)](https://go.dev/dl/)


> Goclean is a tool written in golang, to clean up specified directories on a schedule. 

## Usage

```yaml
version: "3"

services:
  goclean:
    container_name: "goclean"
    build:
      context: .
      dockerfile: Dockerfile
    environment:
      - GO_CLEAN_PATHS="/mnt/goclean/dir1;/mnt/goclean/dir2"
      - GO_CLEAN_SCHEDULE="0 3 * * *"
      - GO_CLEAN_MAX_AGE=7d
    volumes:
      - ~/dir1:/mnt/goclean/dir1
      - ~/dir2:/mnt/goclean/dir2
```

---

### 🤝 Contributing

Contributions, issues and feature requests are welcome!

Feel free to check [issues page](https://github.com/floholz/goclean/issues).


### 📝 License

Copyright © 2024 [floholz](https://github.com/floholz).

This project is [MIT](./LICENSE) licensed.

---

### Show your support

Give a ⭐ if this project helped you!