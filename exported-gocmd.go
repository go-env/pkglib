package pkglib

// VcsByCmd returns the version control system for the given
// command name (hg, git, svn, bzr).
func VcsByCmd(cmd string) *vcsCmd {
	return vcsByCmd(cmd)
}

// ping pings to determine scheme to use.
func (v *vcsCmd) Ping(scheme, repo string) error {
	return v.ping(scheme, repo)
}

// create creates a new copy of repo in dir.
// The parent of dir must exist; dir must not.
func (v *vcsCmd) Create(dir, repo string) error {
	return v.create(dir, repo)
}

// download downloads any new changes for the repo in dir.
func (v *vcsCmd) Download(dir string) error {
	return v.download(dir)
}

// tags returns the list of available tags for the repo in dir.
func (v *vcsCmd) Tags(dir string) ([]string, error) {
	return v.tags(dir)
}

// tagSync syncs the repo in dir to the named tag,
// which either is a tag returned by tags or is v.tagDefault.
func (v *vcsCmd) TagSync(dir, tag string) error {
	return v.tagSync(dir, tag)
}
