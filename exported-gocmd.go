// Public access to private functions from go/cmd
package pkglib

// Just exported internal type
type (
	VcsCmd      vcsCmd
	ImportStack importStack
)

// VcsByCmd returns the version control system for the given
// command name (hg, git, svn, bzr).
func VcsByCmd(cmd string) *VcsCmd {
	V := vcsByCmd(cmd)
	v := VcsCmd(*V)
	return &v
}

// Ping pings to determine scheme to use.
func (v *VcsCmd) Ping(scheme, repo string) error {
	V := vcsCmd(*v)
	return V.ping(scheme, repo)
}

// Create creates a new copy of repo in dir.
// The parent of dir must exist; dir must not.
func (v *VcsCmd) Create(dir, repo string) error {
	V := vcsCmd(*v)
	return V.create(dir, repo)
}

// Download downloads any new changes for the repo in dir.
func (v *VcsCmd) Download(dir string) error {
	V := vcsCmd(*v)
	return V.download(dir)
}

// Tags returns the list of available tags for the repo in dir.
func (v *VcsCmd) Tags(dir string) ([]string, error) {
	V := vcsCmd(*v)
	return V.tags(dir)
}

// TagSync syncs the repo in dir to the named tag,
// which either is a tag returned by tags or is v.tagDefault.
func (v *VcsCmd) TagSync(dir, tag string) error {
	V := vcsCmd(*v)
	return V.tagSync(dir, tag)
}

// DownloadPaths prepares the list of paths to pass to download.
// It expands ... patterns that can be expanded.  If there is no match
// for a particular pattern, downloadPaths leaves it in the result list,
// in the hope that we can figure out the repository from the
// initial ...-free prefix.
func DownloadPaths(args []string) []string {
	return downloadPaths(args)
}

// ImportPaths returns the import paths to use for the given command line.
func ImportPaths(args []string) []string {
	return importPaths(args)
}

// download runs the download half of the get command
// for the package named by the argument.
func Download(arg string, stk *ImportStack, getTestDeps bool) {
	Stk := importStack(*stk)
	download(arg, &Stk, getTestDeps)
}

func RunInstall(cmd *Command, args []string) {
	runInstall(cmd, args)
}
