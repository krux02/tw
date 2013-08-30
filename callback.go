package tw

/*
#include "callback.h"
*/
import "C"

import "unsafe"

type SetVarCallback func(value unsafe.Pointer)
type GetVarCallback func(value unsafe.Pointer)
type ButtonCallback func()

type SummaryCallback func(summaryString string, summaryMaxLength uint, value unsafe.Pointer)

const MaxCallbacks = 32

var setVarCallbacks [MaxCallbacks]SetVarCallback
var currentSetVar = 0

var getVarCallbacks [MaxCallbacks]GetVarCallback
var currentGetVar = 0

var buttonCallbacks [MaxCallbacks]ButtonCallback
var currentButton = 0

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
	if currentGetVar >= MaxCallbacks || currentSetVar >= MaxCallbacks {
		panic("maximum number of callbacks reached (max 10)")
	}

	setVarCallbacks[currentSetVar] = setCallback
	getVarCallbacks[currentGetVar] = getCallback

	r := toBool(C.myAddVarCB((*C.TwBar)(bar), C.CString(name), C.TwType(type_), C.int(currentSetVar), C.int(currentGetVar), C.CString(def)))

	currentGetVar += 1
	currentSetVar += 1

	return r
}

func (bar *Bar) AddButton(name string, callback ButtonCallback, def string) bool {
	if currentButton >= MaxCallbacks {
		panic("maximum number of callbacks reached (max 10)")
	}

	buttonCallbacks[currentButton] = callback

	r := toBool(C.myAddButton((*C.TwBar)(bar), C.CString(name), C.int(currentButton), C.CString(def)))

	currentButton += 1

	return r
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
