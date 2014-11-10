// New public functions based on go/cmd
package pkglib

// Code we downloaded and all code that depends on it
// needs to be evicted from the package cache so that
// the information will be recomputed.  Instead of keeping
// track of the reverse dependency information, evict
// everything.
func CleanPackageCache() {
	for name := range packageCache {
		delete(packageCache, name)
	}
}
