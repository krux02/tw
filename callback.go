package tw

/*
#include "callback.h"
#include <string.h>
*/
import "C"

import "unsafe"

type SetVarCallback func(value unsafe.Pointer)
type GetVarCallback func(value unsafe.Pointer)
type ButtonCallback func()
type SummaryCallback func(value unsafe.Pointer) string

const MaxCallbacks = 32

var setVarCallbacks = make([]SetVarCallback, 0, MaxCallbacks)
var getVarCallbacks = make([]GetVarCallback, 0, MaxCallbacks)
var buttonCallbacks = make([]ButtonCallback, 0, MaxCallbacks)

//export goSetVarCallbackN
func goSetVarCallbackN(value unsafe.Pointer, N C.int) {
	setVarCallbacks[uint(N)](value)
}

//export goGetVarCallbackN
func goGetVarCallbackN(value unsafe.Pointer, N C.int) {
	getVarCallbacks[uint(N)](value)
}

//export goButtonCallbackN
func goButtonCallbackN(N C.int) {
	buttonCallbacks[uint(N)]()
}

func (bar *Bar) AddVarCB(name string, type_ Type, setCallback SetVarCallback, getCallback GetVarCallback, clientData unsafe.Pointer, def string) bool {
	currentSetVar := C.int(len(setVarCallbacks))
	currentGetVar := C.int(len(getVarCallbacks))

	setVarCallbacks = append(setVarCallbacks, setCallback)
	getVarCallbacks = append(getVarCallbacks, getCallback)

	return toBool(C.myAddVarCB((*C.TwBar)(bar), C.CString(name), C.TwType(type_), currentSetVar, currentGetVar, C.CString(def)))
}

func (bar *Bar) AddButton(name string, callback ButtonCallback, def string) bool {
	currentButton := C.int(len(buttonCallbacks))

	buttonCallbacks = append(buttonCallbacks, callback)

	return toBool(C.myAddButton((*C.TwBar)(bar), C.CString(name), currentButton, C.CString(def)))
}

type ErrorHandler func(errorMessage string)

var errorHandler0 ErrorHandler

//export goHandleError
func goHandleError(errorMessage *C.char) {
	errorHandler0(C.GoString(errorMessage))
}

func HandleErrors(errorHandler ErrorHandler) {
	errorHandler0 = errorHandler
	C.myHandleErrors()
}

var summaryCallback0 SummaryCallback

//export goSummary
func goSummary(summaryString *C.char, summaryMaxLength C.size_t, value unsafe.Pointer) {
	str := summaryCallback0(value)
	C.strncpy(summaryString, C.CString(str), summaryMaxLength);
}

func DefineStruct(name string, structMembers []StructMember, structSize uint, summaryCallback SummaryCallback) Type {
	c_name := C.CString(name)
	c_size := C.size_t(structSize)
	numMembers := C.uint(len(structMembers))
	memberPtr := (*C.TwStructMember)(&structMembers[0])

	if( summaryCallback == nil ) {
		return Type(C.TwDefineStruct(c_name, memberPtr, numMembers, c_size, nil, nil))
	} else {
		summaryCallback0 = summaryCallback
		return Type(C.myDefineStruct(c_name, memberPtr, numMembers, c_size))
	}
	panic("")
}
