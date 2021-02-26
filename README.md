**Package for open/close cdrom tray under Go. Can be used for SATA devices.**



Example application:

```golang
package main

import (
	"fmt"
	eject "github.com/NerdDoc/goeject"
	"golang.org/x/sys/unix"
	"os"
)


func main()  {
	cd, err := eject.NewFile("/dev/cdrom")
	if err != nil {
		fmt.Println("ERROR:",err)
		os.Exit(1)
	}

	if _, err := unix.IoctlGetInt(int(cd.Fd()), eject.CDROMEJECT); err != nil {
		fmt.Println("ERROR:",err)
	}
	if _, err := unix.IoctlGetInt(int(cd.Fd()), eject.CDROMCLOSETRAY); err != nil {
		fmt.Println("ERROR:",err)
	}
}
```
