package p4

import (
	"fmt"
	"strings"
)

// ListDepotFiles runs "p4 fstat" and parses the results into a slice of DepotFile structs.
// Order of resulting slice is alphabetical by Path, ignoring case.
func (p *P4) ListDepotFiles() ([]DepotFile, error) {
	return p.runAndParseDepotFiles(
		fmt.Sprintf(`%s fstat -T depotFile,headAction,headChange,headType,digest -Ol `+
			`-F '^(headAction=move/delete | headAction=purge | headAction=archive | headAction=delete)' //%s/...`,
			p.cmd(), p.Client,
		),
	)
}

// ListDepotFilesForPaths runs "p4 fstat" and parses the results into a slice of DepotFile structs.
// Order of resulting slice is alphabetical by Path, ignoring case.
// Accepts a list of depot paths with wildcards, relative to the stream root.
func (p *P4) ListDepotFilesForPaths(paths []string) ([]DepotFile, error) {
	depotPaths := make([]string, len(paths))
	for i, _ := range paths {
		// make full depot path from stream relative path
		depotPaths[i] = fmt.Sprintf("//%s/%s", p.Client, paths[i])
	}
	pathStr := strings.Join(depotPaths, " ")

	return p.runAndParseDepotFiles(
		fmt.Sprintf(`%s fstat -T depotFile,headAction,headChange,headType,digest -Ol `+
			`-F '^(headAction=move/delete | headAction=purge | headAction=archive | headAction=delete)' %s`,
			p.cmd(), pathStr,
		),
	)
}
