package mustard

//#include "parser.h"
//TSLanguage *tree_sitter_mustard();
import "C"
import (
	"unsafe"

	sitter "github.com/dcdc-io/go-tree-sitter"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_mustard())
	return sitter.NewLanguage(ptr)
}
