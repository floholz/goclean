# Welcome to goclean 👋
[![GitHub tag (with filter)](https://img.shields.io/github/v/release/floholz/goclean?label=latest)](https://github.com/floholz/goclean/releases/latest)
[![GitHub License](https://img.shields.io/github/license/floholz/goclean)](./LICENSE)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/floholz/goclean?logo=go&labelColor=gray&label=%20)](https://go.dev/dl/)


> Goclean is a tool written in golang, to clean up specified directories on a schedule. 

## Usage

```yaml
## docker-compose.yml

version: "3"

services:
  goclean:
    container_name: "goclean"
    image: floholz/goclean
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

## Environment Variables

Configuring `goclean` is done with environment variables.

### `GO_CLEAN_PATHS` (_required_)
This specifies what directories goclean should monitor. The paths need to be provided in a `;` seperated list.

_e.g.:_
```bash
GO_CLEAN_PATHS="/mnt/goclean/"
GO_CLEAN_PATHS="/mnt/goclean/dir1;/mnt/goclean/dir2"
GO_CLEAN_PATHS="/mnt/goclean/dir1;/mnt/goclean/dir2;/mnt/also_goclean/dir3"
```
<br>

### `GO_CLEAN_SCHEDULE`
You can specify your own schedule for clean up, by setting this environment variable. The value should be a `crontab` 
schedule string with an optional `seconds` parameter.

By default, the schedule is set to `0 0 * * *` meaning `daily at 00:00`.

_e.g.:_
```bash
GO_CLEAN_SCHEDULE="0 3 * * *"             # “Daily at 03:00.”
GO_CLEAN_SCHEDULE="*/30 */10 * * * WED"   # “Every 30th second after every 10th minute on Wednesday.”
GO_CLEAN_SCHEDULE="55 15 14 1 8 *"        # “At 14:15:55 on 1st of August.”
GO_CLEAN_SCHEDULE="0 22 * * 1-5"          # “At 22:00 on every day-of-week from Monday through Friday.”
```
<br>

### `GO_CLEAN_MAX_AGE`
With this you can define the `maximum age` a file is allowed to have, before being deleted.
This uses the duration syntax, expanded by year, month and day.

By default, the maximum age for a file is set to `7d` meaning `7 days`.

_e.g.:_
```bash
GO_CLEAN_MAX_AGE="7d"               # “7 days.”
GO_CLEAN_MAX_AGE="1y6m"             # “1 year and 6 months.”
GO_CLEAN_MAX_AGE="11h59m59s"        # “11 hours, 59 minutes and 59 seconds.”
GO_CLEAN_MAX_AGE="2y8m7d12h5m30s"   # “2 years, 8 months, 7 days, 12 hours, 5 minutes and 30 seconds.”
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