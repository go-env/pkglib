package pkglib

type VcsCmd vcsCmd

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
