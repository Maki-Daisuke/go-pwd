package pwd

/*
#include <sys/types.h>
#include <pwd.h>
#include <stdlib.h>

// While getpwuid requires "uid_t" according to man page, it actually requires
// "__uit_t" in the source code, that causes cgo compile error ("uid_t" is
// actually aliased to __uid_t).
// Unlike Linux, getpwuid on Mac OS X requires uid_t properly. For compatibility,
// we use a C function as a bridge here.
struct passwd *getpwuid_aux(unsigned int uid) {
	return getpwuid((uid_t)uid);
}
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
		Shell: C.GoString(cpw.pw_shell),
	}
}

// Getpwuid searches the user database for an entry with a matching uid.
func Getpwuid(uid uint32) *Passwd {
	cpw := C.getpwuid_aux(C.uint(uid))
	if cpw == nil {
		return nil
	}
	return &Passwd{
		Name:  C.GoString(cpw.pw_name),
		UID:   uint32(cpw.pw_uid),
		GID:   uint32(cpw.pw_uid),
		Dir:   C.GoString(cpw.pw_dir),
		Shell: C.GoString(cpw.pw_shell),
	}
}
