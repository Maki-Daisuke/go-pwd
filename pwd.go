/*
Package pwd is a thin wrapper of C library <pwd.h>.
This is designed as thin as possible, but aimed to be thread-safe.
*/
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
import (
	"sync"
	"unsafe"
)

// Passwd represents an entry of the user database defined in <pwd.h>
type Passwd struct {
	Name  string
	UID   uint32
	GID   uint32
	Gecos string
	Dir   string
	Shell string
}

func cpasswd2go(cpw *C.struct_passwd) *Passwd {
	return &Passwd{
		Name:  C.GoString(cpw.pw_name),
		UID:   uint32(cpw.pw_uid),
		GID:   uint32(cpw.pw_uid),
		Gecos: C.GoString(cpw.pw_gecos),
		Dir:   C.GoString(cpw.pw_dir),
		Shell: C.GoString(cpw.pw_shell),
	}
}

// Getpwnam searches the user database for an entry with a matching name.
func Getpwnam(name string) *Passwd {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cpw := C.getpwnam(cname)
	if cpw == nil {
		return nil
	}
	return cpasswd2go(cpw)
}

// Getpwuid searches the user database for an entry with a matching uid.
func Getpwuid(uid uint32) *Passwd {
	cpw := C.getpwuid_aux(C.uint(uid))
	if cpw == nil {
		return nil
	}
	return cpasswd2go(cpw)
}

// Getpwents returns all entries in the user databases.
// This is aimed to be thread-safe, that is, if a goroutine is executing this
// function, another goroutine is blocked until it completes.
func Getpwents() []*Passwd {
	pwentMutex.Lock()
	defer pwentMutex.Unlock()
	C.setpwent()
	defer C.endpwent()
	ents := make([]*Passwd, 0, 10)
	for {
		cpw := C.getpwent()
		if cpw == nil {
			break
		}
		ents = append(ents, cpasswd2go(cpw))
	}
	return ents
}

var pwentMutex = sync.Mutex{}
