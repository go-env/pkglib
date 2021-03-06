// Public access to private functions from go/cmd
package pkglib

import "go/token"

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

// LoadImport scans the directory named by path, which must be an import path,
// but possibly a local import path (an absolute file system path or one beginning
// with ./ or ../).  A local relative path is interpreted relative to srcDir.
// It returns a *Package describing the package found in that directory.
func LoadImport(path string, srcDir string, stk *ImportStack, importPos []token.Position) *Package {
	Stk := importStack(*stk)
	return loadImport(path, srcDir, &Stk, importPos)
}

// repoRoot represents a version control system, a repo, and a root of
// where to put it on disk.
type RepoRoot struct {
	Vcs *VcsCmd

	// repo is the repository URL, including scheme
	Repo string

	// root is the import path corresponding to the root of the
	// repository
	Root string
}

// RepoRootForImportPath analyzes importPath to determine the
// version control system, and code repository to use.
func RepoRootForImportPath(importPath string) (*RepoRoot, error) {
	if rr, err := repoRootForImportPath(importPath); err == nil {
		vcs := VcsCmd(*rr.vcs)
		return &RepoRoot{&vcs, rr.repo, rr.root}, nil
	} else {
		return &RepoRoot{}, err
	}
}
