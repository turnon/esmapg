package esmapg

// // table -> column -> index -> paths
// type tableFields map[string]map[string]map[string][]string

// [table] -> [column] -> [paths in index]
type tableFields map[string]map[string][]string

func (m *Map) Paths() {
	tfs := make(tableFields)
	m.Fields.paths([]string{m.Name}, []string{m.Name}, tfs)
}

func (fs *fields) paths(prefix []string, tableStack []string, tfs tableFields) {
	// currentTable := tableStack[len(tableStack)-1]

	// for _, attr := range fs.Only {
	// 	_, pathStr := concatPath(prefix, attr)
	// 	fmt.Println(pathStr)
	// }

	// for attr, subfields := range fs.BelongsTo {
	// 	attr = inflection.Singular(attr)
	// 	path, pathStr := concatPath(prefix, attr)
	// 	fmt.Println(pathStr)
	// 	subfields.paths(path, tfs)
	// }

	// for attr, subfields := range fs.HasOne {
	// 	attr = inflection.Singular(attr)
	// 	path, pathStr := concatPath(prefix, attr)
	// 	fmt.Println(pathStr)
	// 	subfields.paths(path, tfs)
	// }

	// for attr, subfields := range fs.HasMany {
	// 	attr = inflection.Plural(attr)
	// 	path, pathStr := concatPath(prefix, attr)
	// 	fmt.Println(pathStr)
	// 	subfields.paths(path, tfs)
	// }
}

// func concatPath(prefix []string, name string) ([]string, string) {
// 	// if len(prefix) == 1 {
// 	// 	return []string{name}, name
// 	// }
// 	path := append(prefix, name)
// 	return path, strings.Join(path, ".")
// }
