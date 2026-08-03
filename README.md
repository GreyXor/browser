
# browser

Helpers to open URLs, readers, or files in the system default web browser.

This fork adds:

- `OpenReader` error wrapping;
- `ErrNotFound` error wrapping on BSD;
- Go 1.25+ support;
- reaping of spawned browser processes in the background, so long-running
  programs that call these helpers repeatedly don't accumulate zombies.

## Usage

``` go
import "github.com/cli/browser"

err = browser.OpenURL(url)
err = browser.OpenFile(path)
err = browser.OpenReader(reader)
```
