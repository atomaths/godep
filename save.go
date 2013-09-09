package main

import (
	"log"
	"os"
)

var cmdSave = &Command{
	Usage: "save [packages]",
	Short: "list current dependencies to a file",
	Long: `
Save writes a list of the dependencies of the named packages along
with the exact source control revision of each dependency.

Output goes to file Godeps.

For more about specifying packages, see 'go help packages'.
`,
	Run: runSave,
}

func runSave(cmd *Command, args []string) {
	g := &Godeps{
		ImportPath: MustLoadPackages(".")[0].ImportPath,
		GoVersion:  mustGoVersion(),
	}
	a := MustLoadPackages(args...)
	err := g.Load(a)
	if err != nil {
		log.Fatalln(err)
	}
	if g.Deps == nil {
		g.Deps = make([]Dependency, 0) // produce json [], not null
	}
	f, err := os.Create("Godeps")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = g.WriteTo(f)
	if err != nil {
		log.Fatalln(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
