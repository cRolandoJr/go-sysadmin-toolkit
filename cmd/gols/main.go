package main

import (
	"fmt"
	"os"
	"os/user"
	"slices"
	"strings"
	"syscall"
	"text/tabwriter"
)

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	slices.SortFunc(files, func(a, b os.DirEntry) int {
		// Convertimos a minúsculas para comparar sin importar mayúsculas
		return strings.Compare(strings.ToLower(a.Name()), strings.ToLower(b.Name()))
	})

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, "PERMISOS\tPROPIETARIOS\tTAMAÑO\tNOMBRE\n")

	for _, f := range files {
		info, _ := f.Info()
		stat, _ := info.Sys().(*syscall.Stat_t)

		uid := fmt.Sprintf("%d", stat.Uid)
		u, err := user.LookupId(uid)
		if err == nil {
			uid = u.Username
		}

		gid := fmt.Sprintf("%d", stat.Gid)
		g, err := user.LookupGroupId(gid)
		if err == nil {
			gid = g.Name
		}

		fmt.Fprintf(w, "%s\t%s:%s\t%d\t%s\n",
			info.Mode(),
			uid,
			gid,
			info.Size(),
			f.Name(),
		)

	}
	w.Flush()
}
