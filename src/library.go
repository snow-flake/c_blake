/*
	Package main - transpiled by c2go version: v0.21.10 Zinc 2018-02-14

	If you have found any issues, please raise an issue at:
	https://github.com/elliotchance/c2go/
*/



package src

import "unsafe"
import "reflect"
import "github.com/elliotchance/c2go/darwin"

type int8_t int8
type int16_t int16
type int32_t int
type int64_t int64
type uint8_t uint8
type uint16_t uint16
type uint32_t uint32
type uint64_t uint64
type int_least8_t int8
type int_least16_t int16
type int_least32_t int
type int_least64_t int64
type uint_least8_t uint8
type uint_least16_t uint16
type uint_least32_t uint32
type uint_least64_t uint64
type int_fast8_t int8
type int_fast16_t int16
type int_fast32_t int
type int_fast64_t int64
type uint_fast8_t uint8
type uint_fast16_t uint16
type uint_fast32_t uint32
type uint_fast64_t uint64
type __int8_t int8
type __uint8_t uint8
type __int16_t int16
type __uint16_t uint16
type __int32_t int
type __uint32_t uint32
type __int64_t int64
type __uint64_t uint64
type __darwin_intptr_t int32
type __darwin_natural_t uint32
type __darwin_ct_rune_t darwin.CtRuneT
type __mbstate_t struct {
	value interface{}
	arr   [128]byte
}

func (unionVar *__mbstate_t) cast(t reflect.Type) reflect.Value {
	return reflect.NewAt(t, unsafe.Pointer(&unionVar.arr[0]))
}
func (unionVar *__mbstate_t) assign(v interface{}) {
	value := reflect.ValueOf(v).Elem()
	value.Set(unionVar.cast(value.Type()).Elem())
}
func (unionVar *__mbstate_t) UntypedSet(v interface{}) {
	value := reflect.ValueOf(v)
	unionVar.cast(value.Type()).Elem().Set(value)
}
func (unionVar *__mbstate_t) Get__mbstate8() (res []byte) {
	unionVar.assign(&res)
	return
}
func (unionVar *__mbstate_t) Set__mbstate8(v []byte) []byte {
	unionVar.value = v
	unionVar.UntypedSet(v)
	return v
}
func (unionVar *__mbstate_t) Get_mbstateL() (res int64) {
	unionVar.assign(&res)
	return
}
func (unionVar *__mbstate_t) Set_mbstateL(v int64) int64 {
	unionVar.value = v
	unionVar.UntypedSet(v)
	return v
}

type __darwin_mbstate_t int64
type __darwin_ptrdiff_t int32
type __darwin_size_t uint32
type __darwin_va_list int64
type __darwin_wchar_t int
type __darwin_rune_t int
type __darwin_wint_t int
type __darwin_clock_t uint32
type __darwin_socklen_t uint32
type __darwin_ssize_t int32
type __darwin_time_t int32
type __darwin_blkcnt_t int64
type __darwin_blksize_t int
type __darwin_dev_t int
type __darwin_fsblkcnt_t uint32
type __darwin_fsfilcnt_t uint32
type __darwin_gid_t uint32
type __darwin_id_t uint32
type __darwin_ino64_t uint64
type __darwin_ino_t uint64
type __darwin_mach_port_name_t uint32
type __darwin_mach_port_t uint32
type __darwin_mode_t uint16
type __darwin_off_t int64
type __darwin_pid_t int
type __darwin_sigset_t uint32
type __darwin_suseconds_t int
type __darwin_uid_t uint32
type __darwin_useconds_t uint32
type __darwin_uuid_t []uint8
type __darwin_uuid_string_t []byte
type __darwin_pthread_handler_rec struct {
	__routine func(interface{})
	__arg     interface{}
	__next    []int64
}
type _opaque_pthread_attr_t struct {
	__sig    int32
	__opaque []byte
}
type _opaque_pthread_cond_t struct {
	__sig    int32
	__opaque []byte
}
type _opaque_pthread_condattr_t struct {
	__sig    int32
	__opaque []byte
}
type _opaque_pthread_mutex_t struct {
	__sig    int32
	__opaque []byte
}
type _opaque_pthread_mutexattr_t struct {
	__sig    int32
	__opaque []byte
}
type _opaque_pthread_once_t struct {
	__sig    int32
	__opaque []byte
}
type _opaque_pthread_rwlock_t struct {
	__sig    int32
	__opaque []byte
}
type _opaque_pthread_rwlockattr_t struct {
	__sig    int32
	__opaque []byte
}
type _opaque_pthread_t struct {
	__sig           int32
	__cleanup_stack []int64
	__opaque        []byte
}
type __darwin_pthread_attr_t struct {
	__sig           int32
	__cleanup_stack []int64
	__opaque        []byte
}
type __darwin_pthread_cond_t _opaque_pthread_cond_t
type __darwin_pthread_condattr_t _opaque_pthread_condattr_t
type __darwin_pthread_key_t uint32
type __darwin_pthread_mutex_t _opaque_pthread_mutex_t
type __darwin_pthread_mutexattr_t _opaque_pthread_mutexattr_t
type __darwin_pthread_once_t _opaque_pthread_once_t
type __darwin_pthread_rwlock_t _opaque_pthread_rwlock_t
type __darwin_pthread_rwlockattr_t _opaque_pthread_rwlockattr_t
type __darwin_pthread_t []_opaque_pthread_t
type u_int8_t uint8
type u_int16_t uint16
type u_int32_t uint32
type u_int64_t uint64
type register_t int64
type uintptr_t uint32
type user_addr_t uint64
type user_size_t uint64
type user_ssize_t int64
type user_long_t int64
type user_ulong_t uint64
type user_time_t int64
type user_off_t int64
type syscall_arg_t uint64
type intptr_t int32
type intmax_t int32
type uintmax_t uint32
type __darwin_nl_item int
type __darwin_wctrans_t int
type __darwin_wctype_t uint32
type size_t uint32
type rsize_t uint32
type errno_t int
type ssize_t int32

var sigma [][]uint8 = [][]uint8{[]uint8{uint8((uint8(0))), uint8((uint8(1))), uint8((uint8(2))), uint8((uint8(3))), uint8((uint8(4))), uint8((uint8(5))), uint8((uint8(6))), uint8((uint8(7))), uint8((uint8(8))), uint8((uint8(9))), uint8((uint8(10))), uint8((uint8(11))), uint8((uint8(12))), uint8((uint8(13))), uint8((uint8(14))), uint8((uint8(15)))}, []uint8{uint8((uint8(14))), uint8((uint8(10))), uint8((uint8(4))), uint8((uint8(8))), uint8((uint8(9))), uint8((uint8(15))), uint8((uint8(13))), uint8((uint8(6))), uint8((uint8(1))), uint8((uint8(12))), uint8((uint8(0))), uint8((uint8(2))), uint8((uint8(11))), uint8((uint8(7))), uint8((uint8(5))), uint8((uint8(3)))}, []uint8{uint8((uint8(11))), uint8((uint8(8))), uint8((uint8(12))), uint8((uint8(0))), uint8((uint8(5))), uint8((uint8(2))), uint8((uint8(15))), uint8((uint8(13))), uint8((uint8(10))), uint8((uint8(14))), uint8((uint8(3))), uint8((uint8(6))), uint8((uint8(7))), uint8((uint8(1))), uint8((uint8(9))), uint8((uint8(4)))}, []uint8{uint8((uint8(7))), uint8((uint8(9))), uint8((uint8(3))), uint8((uint8(1))), uint8((uint8(13))), uint8((uint8(12))), uint8((uint8(11))), uint8((uint8(14))), uint8((uint8(2))), uint8((uint8(6))), uint8((uint8(5))), uint8((uint8(10))), uint8((uint8(4))), uint8((uint8(0))), uint8((uint8(15))), uint8((uint8(8)))}, []uint8{uint8((uint8(9))), uint8((uint8(0))), uint8((uint8(5))), uint8((uint8(7))), uint8((uint8(2))), uint8((uint8(4))), uint8((uint8(10))), uint8((uint8(15))), uint8((uint8(14))), uint8((uint8(1))), uint8((uint8(11))), uint8((uint8(12))), uint8((uint8(6))), uint8((uint8(8))), uint8((uint8(3))), uint8((uint8(13)))}, []uint8{uint8((uint8(2))), uint8((uint8(12))), uint8((uint8(6))), uint8((uint8(10))), uint8((uint8(0))), uint8((uint8(11))), uint8((uint8(8))), uint8((uint8(3))), uint8((uint8(4))), uint8((uint8(13))), uint8((uint8(7))), uint8((uint8(5))), uint8((uint8(15))), uint8((uint8(14))), uint8((uint8(1))), uint8((uint8(9)))}, []uint8{uint8((uint8(12))), uint8((uint8(5))), uint8((uint8(1))), uint8((uint8(15))), uint8((uint8(14))), uint8((uint8(13))), uint8((uint8(4))), uint8((uint8(10))), uint8((uint8(0))), uint8((uint8(7))), uint8((uint8(6))), uint8((uint8(3))), uint8((uint8(9))), uint8((uint8(2))), uint8((uint8(8))), uint8((uint8(11)))}, []uint8{uint8((uint8(13))), uint8((uint8(11))), uint8((uint8(7))), uint8((uint8(14))), uint8((uint8(12))), uint8((uint8(1))), uint8((uint8(3))), uint8((uint8(9))), uint8((uint8(5))), uint8((uint8(0))), uint8((uint8(15))), uint8((uint8(4))), uint8((uint8(8))), uint8((uint8(6))), uint8((uint8(2))), uint8((uint8(10)))}, []uint8{uint8((uint8(6))), uint8((uint8(15))), uint8((uint8(14))), uint8((uint8(9))), uint8((uint8(11))), uint8((uint8(3))), uint8((uint8(0))), uint8((uint8(8))), uint8((uint8(12))), uint8((uint8(2))), uint8((uint8(13))), uint8((uint8(7))), uint8((uint8(1))), uint8((uint8(4))), uint8((uint8(10))), uint8((uint8(5)))}, []uint8{uint8((uint8(10))), uint8((uint8(2))), uint8((uint8(8))), uint8((uint8(4))), uint8((uint8(7))), uint8((uint8(6))), uint8((uint8(1))), uint8((uint8(5))), uint8((uint8(15))), uint8((uint8(11))), uint8((uint8(9))), uint8((uint8(14))), uint8((uint8(3))), uint8((uint8(12))), uint8((uint8(13))), uint8((uint8(0)))}, []uint8{uint8((uint8(0))), uint8((uint8(1))), uint8((uint8(2))), uint8((uint8(3))), uint8((uint8(4))), uint8((uint8(5))), uint8((uint8(6))), uint8((uint8(7))), uint8((uint8(8))), uint8((uint8(9))), uint8((uint8(10))), uint8((uint8(11))), uint8((uint8(12))), uint8((uint8(13))), uint8((uint8(14))), uint8((uint8(15)))}, []uint8{uint8((uint8(14))), uint8((uint8(10))), uint8((uint8(4))), uint8((uint8(8))), uint8((uint8(9))), uint8((uint8(15))), uint8((uint8(13))), uint8((uint8(6))), uint8((uint8(1))), uint8((uint8(12))), uint8((uint8(0))), uint8((uint8(2))), uint8((uint8(11))), uint8((uint8(7))), uint8((uint8(5))), uint8((uint8(3)))}, []uint8{uint8((uint8(11))), uint8((uint8(8))), uint8((uint8(12))), uint8((uint8(0))), uint8((uint8(5))), uint8((uint8(2))), uint8((uint8(15))), uint8((uint8(13))), uint8((uint8(10))), uint8((uint8(14))), uint8((uint8(3))), uint8((uint8(6))), uint8((uint8(7))), uint8((uint8(1))), uint8((uint8(9))), uint8((uint8(4)))}, []uint8{uint8((uint8(7))), uint8((uint8(9))), uint8((uint8(3))), uint8((uint8(1))), uint8((uint8(13))), uint8((uint8(12))), uint8((uint8(11))), uint8((uint8(14))), uint8((uint8(2))), uint8((uint8(6))), uint8((uint8(5))), uint8((uint8(10))), uint8((uint8(4))), uint8((uint8(0))), uint8((uint8(15))), uint8((uint8(8)))}}
var cst []uint32 = []uint32{uint32((uint32(608135816))), 2242054355, uint32((uint32(320440878))), uint32((uint32(57701188))), 2752067618, uint32((uint32(698298832))), uint32((uint32(137296536))), 3964562569, uint32((uint32(1160258022))), uint32((uint32(953160567))), 3193202383, uint32((uint32(887688300))), 3232508343, 3380367581, uint32((uint32(1065670069))), 3041331479}
var padding []uint8 = []uint8{uint8((uint8(128))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0))), uint8((uint8(0)))}

type state struct {
	h      []uint32
	s      []uint32
	t      []uint32
	buflen int
	nullt  int
	buf    []uint8
}
// blake256_compress - transpiled function from  /Users/home/Development/c_blake/src/library.c:68
func blake256_compress(S []state, block []uint8) {
	var v []uint32 = make([]uint32, 16, 16)
	var m []uint32 = make([]uint32, 16, 16)
	var i uint32
	for i = uint32((uint32(0))); i < uint32((uint32(16))); i++ {
		m[i] = uint32((uint32(uint8((((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&block[0])) + (uintptr)(i*uint32((uint32(4))))*unsafe.Sizeof(block[0]))))[:])[0])))))<<uint64(24)|uint32((uint32(uint8((((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&block[0])) + (uintptr)(i*uint32((uint32(4))))*unsafe.Sizeof(block[0]))))[:])[1])))))<<uint64(16)|uint32((uint32(uint8((((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&block[0])) + (uintptr)(i*uint32((uint32(4))))*unsafe.Sizeof(block[0]))))[:])[2])))))<<uint64(8)|uint32((uint32(uint8((((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&block[0])) + (uintptr)(i*uint32((uint32(4))))*unsafe.Sizeof(block[0]))))[:])[3])))))
	}
	for i = uint32((uint32(0))); i < uint32((uint32(8))); i++ {
		v[i] = S[0].h[i]
	}
	v[8] = S[0].s[0]^uint32(608135816)
	v[9] = S[0].s[1]^2242054355
	v[10] = S[0].s[2]^uint32(320440878)
	v[11] = S[0].s[3]^uint32(57701188)
	v[12] = uint32((uint32(2752067618)))
	v[13] = uint32((uint32(698298832)))
	v[14] = uint32((uint32(137296536)))
	v[15] = uint32((uint32(3964562569)))
	if S[0].nullt == 0 {
		v[12] ^= S[0].t[0]
		v[13] ^= S[0].t[0]
		v[14] ^= S[0].t[1]
		v[15] ^= S[0].t[1]
	}
	for i = uint32((uint32(0))); i < uint32((uint32(14))); i++ {
		v[0] += m[sigma[i][0]]^cst[sigma[i][0+1]]+v[4]
		v[12] = (v[12]^v[0])<<uint64(32-16)|(v[12]^v[0])>>uint64(16)
		v[8] += v[12]
		v[4] = (v[4]^v[8])<<uint64(32-12)|(v[4]^v[8])>>uint64(12)
		v[0] += m[sigma[i][0+1]]^cst[sigma[i][0]]+v[4]
		v[12] = (v[12]^v[0])<<uint64(32-8)|(v[12]^v[0])>>uint64(8)
		v[8] += v[12]
		v[4] = (v[4]^v[8])<<uint64(32-7)|(v[4]^v[8])>>uint64(7)
		v[1] += m[sigma[i][2]]^cst[sigma[i][2+1]]+v[5]
		v[13] = (v[13]^v[1])<<uint64(32-16)|(v[13]^v[1])>>uint64(16)
		v[9] += v[13]
		v[5] = (v[5]^v[9])<<uint64(32-12)|(v[5]^v[9])>>uint64(12)
		v[1] += m[sigma[i][2+1]]^cst[sigma[i][2]]+v[5]
		v[13] = (v[13]^v[1])<<uint64(32-8)|(v[13]^v[1])>>uint64(8)
		v[9] += v[13]
		v[5] = (v[5]^v[9])<<uint64(32-7)|(v[5]^v[9])>>uint64(7)
		v[2] += m[sigma[i][4]]^cst[sigma[i][4+1]]+v[6]
		v[14] = (v[14]^v[2])<<uint64(32-16)|(v[14]^v[2])>>uint64(16)
		v[10] += v[14]
		v[6] = (v[6]^v[10])<<uint64(32-12)|(v[6]^v[10])>>uint64(12)
		v[2] += m[sigma[i][4+1]]^cst[sigma[i][4]]+v[6]
		v[14] = (v[14]^v[2])<<uint64(32-8)|(v[14]^v[2])>>uint64(8)
		v[10] += v[14]
		v[6] = (v[6]^v[10])<<uint64(32-7)|(v[6]^v[10])>>uint64(7)
		v[3] += m[sigma[i][6]]^cst[sigma[i][6+1]]+v[7]
		v[15] = (v[15]^v[3])<<uint64(32-16)|(v[15]^v[3])>>uint64(16)
		v[11] += v[15]
		v[7] = (v[7]^v[11])<<uint64(32-12)|(v[7]^v[11])>>uint64(12)
		v[3] += m[sigma[i][6+1]]^cst[sigma[i][6]]+v[7]
		v[15] = (v[15]^v[3])<<uint64(32-8)|(v[15]^v[3])>>uint64(8)
		v[11] += v[15]
		v[7] = (v[7]^v[11])<<uint64(32-7)|(v[7]^v[11])>>uint64(7)
		v[3] += m[sigma[i][14]]^cst[sigma[i][14+1]]+v[4]
		v[14] = (v[14]^v[3])<<uint64(32-16)|(v[14]^v[3])>>uint64(16)
		v[9] += v[14]
		v[4] = (v[4]^v[9])<<uint64(32-12)|(v[4]^v[9])>>uint64(12)
		v[3] += m[sigma[i][14+1]]^cst[sigma[i][14]]+v[4]
		v[14] = (v[14]^v[3])<<uint64(32-8)|(v[14]^v[3])>>uint64(8)
		v[9] += v[14]
		v[4] = (v[4]^v[9])<<uint64(32-7)|(v[4]^v[9])>>uint64(7)
		v[2] += m[sigma[i][12]]^cst[sigma[i][12+1]]+v[7]
		v[13] = (v[13]^v[2])<<uint64(32-16)|(v[13]^v[2])>>uint64(16)
		v[8] += v[13]
		v[7] = (v[7]^v[8])<<uint64(32-12)|(v[7]^v[8])>>uint64(12)
		v[2] += m[sigma[i][12+1]]^cst[sigma[i][12]]+v[7]
		v[13] = (v[13]^v[2])<<uint64(32-8)|(v[13]^v[2])>>uint64(8)
		v[8] += v[13]
		v[7] = (v[7]^v[8])<<uint64(32-7)|(v[7]^v[8])>>uint64(7)
		v[0] += m[sigma[i][8]]^cst[sigma[i][8+1]]+v[5]
		v[15] = (v[15]^v[0])<<uint64(32-16)|(v[15]^v[0])>>uint64(16)
		v[10] += v[15]
		v[5] = (v[5]^v[10])<<uint64(32-12)|(v[5]^v[10])>>uint64(12)
		v[0] += m[sigma[i][8+1]]^cst[sigma[i][8]]+v[5]
		v[15] = (v[15]^v[0])<<uint64(32-8)|(v[15]^v[0])>>uint64(8)
		v[10] += v[15]
		v[5] = (v[5]^v[10])<<uint64(32-7)|(v[5]^v[10])>>uint64(7)
		v[1] += m[sigma[i][10]]^cst[sigma[i][10+1]]+v[6]
		v[12] = (v[12]^v[1])<<uint64(32-16)|(v[12]^v[1])>>uint64(16)
		v[11] += v[12]
		v[6] = (v[6]^v[11])<<uint64(32-12)|(v[6]^v[11])>>uint64(12)
		v[1] += m[sigma[i][10+1]]^cst[sigma[i][10]]+v[6]
		v[12] = (v[12]^v[1])<<uint64(32-8)|(v[12]^v[1])>>uint64(8)
		v[11] += v[12]
		v[6] = (v[6]^v[11])<<uint64(32-7)|(v[6]^v[11])>>uint64(7)
	}
	for i = uint32((uint32(0))); i < uint32((uint32(16))); i++ {
		S[0].h[i%uint32((uint32(8)))] ^= v[i]
	}
	for i = uint32((uint32(0))); i < uint32((uint32(8))); i++ {
		S[0].h[i] ^= S[0].s[i%uint32((uint32(4)))]
	}
}
// blake256_init - transpiled function from  /Users/home/Development/c_blake/src/library.c:104
func blake256_init(S []state) {
	S[0].h[0] = uint32((uint32(1779033703)))
	S[0].h[1] = uint32((uint32(3144134277)))
	S[0].h[2] = uint32((uint32(1013904242)))
	S[0].h[3] = uint32((uint32(2773480762)))
	S[0].h[4] = uint32((uint32(1359893119)))
	S[0].h[5] = uint32((uint32(2600822924)))
	S[0].h[6] = uint32((uint32(528734635)))
	S[0].h[7] = uint32((uint32(1541459225)))
	S[0].nullt = 0
	S[0].t[1] = uint32((uint32(S[0].buflen = S[0].nullt)))
	S[0].t[0] = S[0].t[1]
	S[0].s[3] = uint32((uint32(0)))
	S[0].s[2] = S[0].s[3]
	S[0].s[1] = S[0].s[2]
	S[0].s[0] = S[0].s[1]
}
// blake256_update - transpiled function from  /Users/home/Development/c_blake/src/library.c:118
// datalen = number of bits
//
func blake256_update(S []state, data []uint8, datalen uint64) {
	var left int = S[0].buflen >> uint64(3)
	var fill int = 64 - left
	if left != 0 && datalen>>uint64(3)&uint64((uint64(63))) >= uint64(uint32(fill)) {
		__builtin___memcpy_chk(((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&S[0].buf[0])) + (uintptr)(left)*unsafe.Sizeof(S[0].buf[0]))))[:]), data, uint32(fill), uint32(darwin.BuiltinObjectSize(((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&S[0].buf[0])) + (uintptr)(left)*unsafe.Sizeof(S[0].buf[0]))))[:]).([]byte), 0)))
		S[0].t[0] += uint32(512)
		if S[0].t[0] == uint32((uint32(0))) {
			S[0].t[1] += 1
		}
		blake256_compress(S, S[0].buf)
		data = (*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&data[0])) + (uintptr)(fill)*unsafe.Sizeof(data[0]))))[:]
		datalen -= uint64(fill<<uint64(3))
		left = 0
	}
	for datalen >= uint64(512) {
		S[0].t[0] += uint32(512)
		if S[0].t[0] == uint32((uint32(0))) {
			S[0].t[1] += 1
		}
		blake256_compress(S, data)
		data = (*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&data[0])) + (uintptr)(64)*unsafe.Sizeof(data[0]))))[:]
		datalen -= uint64(512)
	}
	if datalen > uint64((uint64(0))) {
		__builtin___memcpy_chk(((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&S[0].buf[0])) + (uintptr)(left)*unsafe.Sizeof(S[0].buf[0]))))[:]), data, uint32(uint64((datalen >> uint64(3)))), uint32(darwin.BuiltinObjectSize(((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&S[0].buf[0])) + (uintptr)(left)*unsafe.Sizeof(S[0].buf[0]))))[:]).([]byte), 0)))
		S[0].buflen = int(uint64(left<<uint64(3))+uint64((datalen)))
	} else {
		S[0].buflen = 0
	}
}
// blake224_update - transpiled function from  /Users/home/Development/c_blake/src/library.c:149
// datalen = number of bits
//
func blake224_update(S []state, data []uint8, datalen uint64) {
	blake256_update(S, data, datalen)
}
// blake256_final_h - transpiled function from  /Users/home/Development/c_blake/src/library.c:153
/* one padding byte *///
/* enough space to fill the block  *///
/* need 2 compressions *///
//
func blake256_final_h(S []state, digest []uint8, pa uint8, pb uint8) {
	var msglen []uint8 = make([]uint8, 8, 8)
	var lo uint32 = S[0].t[0] + uint32((uint32(S[0].buflen)))
	var hi uint32 = S[0].t[1]
	if lo < uint32((uint32(S[0].buflen))) {
		hi += 1
	}
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&msglen[0])) + (uintptr)(0)*unsafe.Sizeof(msglen[0]))))[:])[0] = uint8((uint8(uint32((hi >> uint64(24))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&msglen[0])) + (uintptr)(0)*unsafe.Sizeof(msglen[0]))))[:])[1] = uint8((uint8(uint32((hi >> uint64(16))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&msglen[0])) + (uintptr)(0)*unsafe.Sizeof(msglen[0]))))[:])[2] = uint8((uint8(uint32((hi >> uint64(8))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&msglen[0])) + (uintptr)(0)*unsafe.Sizeof(msglen[0]))))[:])[3] = uint8((uint8(uint32((hi)))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&msglen[0])) + (uintptr)(4)*unsafe.Sizeof(msglen[0]))))[:])[0] = uint8((uint8(uint32((lo >> uint64(24))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&msglen[0])) + (uintptr)(4)*unsafe.Sizeof(msglen[0]))))[:])[1] = uint8((uint8(uint32((lo >> uint64(16))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&msglen[0])) + (uintptr)(4)*unsafe.Sizeof(msglen[0]))))[:])[2] = uint8((uint8(uint32((lo >> uint64(8))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&msglen[0])) + (uintptr)(4)*unsafe.Sizeof(msglen[0]))))[:])[3] = uint8((uint8(uint32((lo)))))
	if S[0].buflen == 440 {
		S[0].t[0] -= uint32(8)
		blake256_update(S, (*[1]uint8)(unsafe.Pointer(&pa))[:], uint64((uint64(8))))
	} else {
		if S[0].buflen < 440 {
			if S[0].buflen == 0 {
				S[0].nullt = 1
			}
			S[0].t[0] -= uint32(440-S[0].buflen)
			blake256_update(S, padding, uint64((uint64(440 - S[0].buflen))))
		} else {
			S[0].t[0] -= uint32(512-S[0].buflen)
			blake256_update(S, padding, uint64((uint64(512 - S[0].buflen))))
			S[0].t[0] -= uint32(440)
			blake256_update(S, (*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&padding[0])) + (uintptr)(1)*unsafe.Sizeof(padding[0]))))[:], uint64((uint64(440))))
			S[0].nullt = 1
		}
		blake256_update(S, (*[1]uint8)(unsafe.Pointer(&pb))[:], uint64((uint64(8))))
		S[0].t[0] -= uint32(8)
	}
	S[0].t[0] -= uint32(64)
	blake256_update(S, msglen, uint64((uint64(64))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(0)*unsafe.Sizeof(digest[0]))))[:])[0] = uint8((uint8(uint32((S[0].h[0] >> uint64(24))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(0)*unsafe.Sizeof(digest[0]))))[:])[1] = uint8((uint8(uint32((S[0].h[0] >> uint64(16))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(0)*unsafe.Sizeof(digest[0]))))[:])[2] = uint8((uint8(uint32((S[0].h[0] >> uint64(8))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(0)*unsafe.Sizeof(digest[0]))))[:])[3] = uint8((uint8(uint32((S[0].h[0])))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(4)*unsafe.Sizeof(digest[0]))))[:])[0] = uint8((uint8(uint32((S[0].h[1] >> uint64(24))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(4)*unsafe.Sizeof(digest[0]))))[:])[1] = uint8((uint8(uint32((S[0].h[1] >> uint64(16))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(4)*unsafe.Sizeof(digest[0]))))[:])[2] = uint8((uint8(uint32((S[0].h[1] >> uint64(8))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(4)*unsafe.Sizeof(digest[0]))))[:])[3] = uint8((uint8(uint32((S[0].h[1])))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(8)*unsafe.Sizeof(digest[0]))))[:])[0] = uint8((uint8(uint32((S[0].h[2] >> uint64(24))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(8)*unsafe.Sizeof(digest[0]))))[:])[1] = uint8((uint8(uint32((S[0].h[2] >> uint64(16))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(8)*unsafe.Sizeof(digest[0]))))[:])[2] = uint8((uint8(uint32((S[0].h[2] >> uint64(8))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(8)*unsafe.Sizeof(digest[0]))))[:])[3] = uint8((uint8(uint32((S[0].h[2])))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(12)*unsafe.Sizeof(digest[0]))))[:])[0] = uint8((uint8(uint32((S[0].h[3] >> uint64(24))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(12)*unsafe.Sizeof(digest[0]))))[:])[1] = uint8((uint8(uint32((S[0].h[3] >> uint64(16))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(12)*unsafe.Sizeof(digest[0]))))[:])[2] = uint8((uint8(uint32((S[0].h[3] >> uint64(8))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(12)*unsafe.Sizeof(digest[0]))))[:])[3] = uint8((uint8(uint32((S[0].h[3])))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(16)*unsafe.Sizeof(digest[0]))))[:])[0] = uint8((uint8(uint32((S[0].h[4] >> uint64(24))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(16)*unsafe.Sizeof(digest[0]))))[:])[1] = uint8((uint8(uint32((S[0].h[4] >> uint64(16))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(16)*unsafe.Sizeof(digest[0]))))[:])[2] = uint8((uint8(uint32((S[0].h[4] >> uint64(8))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(16)*unsafe.Sizeof(digest[0]))))[:])[3] = uint8((uint8(uint32((S[0].h[4])))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(20)*unsafe.Sizeof(digest[0]))))[:])[0] = uint8((uint8(uint32((S[0].h[5] >> uint64(24))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(20)*unsafe.Sizeof(digest[0]))))[:])[1] = uint8((uint8(uint32((S[0].h[5] >> uint64(16))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(20)*unsafe.Sizeof(digest[0]))))[:])[2] = uint8((uint8(uint32((S[0].h[5] >> uint64(8))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(20)*unsafe.Sizeof(digest[0]))))[:])[3] = uint8((uint8(uint32((S[0].h[5])))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(24)*unsafe.Sizeof(digest[0]))))[:])[0] = uint8((uint8(uint32((S[0].h[6] >> uint64(24))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(24)*unsafe.Sizeof(digest[0]))))[:])[1] = uint8((uint8(uint32((S[0].h[6] >> uint64(16))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(24)*unsafe.Sizeof(digest[0]))))[:])[2] = uint8((uint8(uint32((S[0].h[6] >> uint64(8))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(24)*unsafe.Sizeof(digest[0]))))[:])[3] = uint8((uint8(uint32((S[0].h[6])))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(28)*unsafe.Sizeof(digest[0]))))[:])[0] = uint8((uint8(uint32((S[0].h[7] >> uint64(24))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(28)*unsafe.Sizeof(digest[0]))))[:])[1] = uint8((uint8(uint32((S[0].h[7] >> uint64(16))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(28)*unsafe.Sizeof(digest[0]))))[:])[2] = uint8((uint8(uint32((S[0].h[7] >> uint64(8))))))
	((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&digest[0])) + (uintptr)(28)*unsafe.Sizeof(digest[0]))))[:])[3] = uint8((uint8(uint32((S[0].h[7])))))
}
// blake256_final - transpiled function from  /Users/home/Development/c_blake/src/library.c:191
func blake256_final(S []state, digest []uint8) {
	blake256_final_h(S, digest, uint8((uint8(129))), uint8((uint8(1))))
}
// blake256_hash - transpiled function from  /Users/home/Development/c_blake/src/library.c:196
// inlen = number of bytes
//
func blake256_hash(out []uint8, in []uint8, inlen uint64) {
	var S state
	blake256_init((*[1]state)(unsafe.Pointer(&S))[:])
	blake256_update((*[1]state)(unsafe.Pointer(&S))[:], in, inlen*uint64((uint64(8))))
	blake256_final((*[1]state)(unsafe.Pointer(&S))[:], out)
}
func init() {
}


func Sum256(in []uint8) [64]uint8 {
	output := [64]uint8 {}
	blake256_hash(output, in, len(in))
	return output
}