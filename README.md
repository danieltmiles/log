# Package log

```Go
import "github/monsooncommerce/log"
```

## Usage

```Go
package main

import (
    "github/monsooncommerce/log"
)

func main() {
    log.SetTag("myapp")
    log.Error("This is an error message")
}
```

Output:

```
2013-10-02T13:41:43-07:00 host.monsooncommerce.com myapp[25702]: ERROR This is an error message
```
