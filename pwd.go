package pwd

/*
#include <sys/types.h>
#include <pwd.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// Passwd represents an entry of the user database defined in <pwd.h>
type Passwd struct {
	Name  string
	UID   uint32
	GID   uint32
	Dir   string
	Shell string
}

// Getpwnam searches the user database for an entry with a matching name.
func Getpwnam(name string) *Passwd {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cpw := C.getpwnam(cname)
	if cpw == nil {
		return nil
	}
	return &Passwd{
		Name:  C.GoString(cpw.pw_name),
		UID:   uint32(cpw.pw_uid),
		GID:   uint32(cpw.pw_uid),
		Dir:   C.GoString(cpw.pw_dir),
		Shell: C.GoString(cpw.pw_shell)}
}
