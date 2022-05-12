//go:build !dynamic

package rapidsnark

// #cgo CFLAGS: -DUSE_VENDORED_RAPIDSNARK
// #cgo LDFLAGS: ${SRCDIR}/rapidsnark_vendor/librapidsnark-linux-arm64.a ${SRCDIR}/rapidsnark_vendor/libgmp-linux-arm64.a -lstdc++
import "C"
