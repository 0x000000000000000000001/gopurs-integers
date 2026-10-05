import "gopurs/output/gopurs_runtime"

func And(n1 int, n2 int) int { return int(gopurs_runtime.IntBitAnd(int64(n1), int64(n2))) }
func Or(n1 int, n2 int) int { return int(gopurs_runtime.IntBitOr(int64(n1), int64(n2))) }
func Xor(n1 int, n2 int) int { return int(gopurs_runtime.IntBitXor(int64(n1), int64(n2))) }
func Shl(n1 int, n2 int) int { return int(gopurs_runtime.IntShl(int64(n1), int64(n2))) }
func Shr(n1 int, n2 int) int { return int(gopurs_runtime.IntShr(int64(n1), int64(n2))) }
func Zshr(n1 int, n2 int) int { return int(gopurs_runtime.IntZshr(int64(n1), int64(n2))) }
func Complement(n int) int { return int(gopurs_runtime.IntBitNot(int64(n))) }
