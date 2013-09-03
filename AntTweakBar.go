//ant tweas bar wrapper
package tw

/*
#cgo linux LDFLAGS: -L/usr/local/lib -lAntTweakBar -lstdc++ -lGL
#include <AntTweakBar.h>
*/
import "C"

import "unsafe"
import "reflect"

func toBool(i C.int) bool {
	return i == 0
}

func ptr(v interface{}) unsafe.Pointer {
	if v == nil {
		return unsafe.Pointer(nil)
	}

	rv := reflect.ValueOf(v)
	var et reflect.Value
	switch rv.Type().Kind() {
	case reflect.Uintptr:
		offset, _ := v.(uintptr)
		return unsafe.Pointer(offset)
	case reflect.Ptr:
		et = rv.Elem()
	case reflect.Slice:
		et = rv.Index(0)
	default:
		panic("type must be a pointer, a slice, uintptr or nil")
	}

	return unsafe.Pointer(et.UnsafeAddr())
}

type Bar C.TwBar

func NewBar(name string) *Bar {
	return (*Bar)(C.TwNewBar(C.CString(name)))
}

func (bar *Bar) Delete() {
	C.TwDeleteBar((*C.TwBar)(bar))
}

func DeleteAllBars() bool {
	return toBool(C.TwDeleteAllBars())
}

func (bar *Bar) SetTopBar() bool {
	return toBool(C.TwSetTopBar((*C.TwBar)(bar)))
}

func GetTopBar() *Bar {
	return (*Bar)(C.TwGetTopBar())
}

func (bar *Bar) SetBottomBar() bool {
	return toBool(C.TwSetBottomBar((*C.TwBar)(bar)))
}

func GetBottomBar() *Bar {
	return (*Bar)(C.TwGetBottomBar())
}

func (bar *Bar) GetBarName() string {
	return C.GoString(C.TwGetBarName((*C.TwBar)(bar)))
}

func GetBarCount() int {
	return int(C.TwGetBarCount())
}

func GetBarByIndex(barIndex int) *Bar {
	return (*Bar)(C.TwGetBarByIndex(C.int(barIndex)))
}

func GetBarByName(barName string) *Bar {
	return (*Bar)(C.TwGetBarByName(C.CString(barName)))
}

func (bar *Bar) RefreshBar() bool {
	return toBool(C.TwRefreshBar((*C.TwBar)(bar)))
}

type Type C.TwType

const (
	TYPE_DIR3F    = C.TW_TYPE_DIR3F
	TYPE_DIR3D    = C.TW_TYPE_DIR3D
	TYPE_QUAT4F   = C.TW_TYPE_QUAT4F
	TYPE_QUAT4D   = C.TW_TYPE_QUAT4D
	TYPE_CDSTRING = C.TW_TYPE_CDSTRING
	// W_TYPE_CSSTRING(maxsize) = C.TW_TYPE_CSSTRING(maxsize)
	TYPE_COLOR4F = C.TW_TYPE_COLOR4F
	TYPE_COLOR3F = C.TW_TYPE_COLOR3F
	TYPE_COLOR32 = C.TW_TYPE_COLOR32
	TYPE_DOUBLE  = C.TW_TYPE_DOUBLE
	TYPE_FLOAT   = C.TW_TYPE_FLOAT
	TYPE_UINT32  = C.TW_TYPE_UINT32
	TYPE_INT32   = C.TW_TYPE_INT32
	TYPE_UINT16  = C.TW_TYPE_UINT16
	TYPE_INT16   = C.TW_TYPE_INT16
	TYPE_UINT8   = C.TW_TYPE_UINT8
	TYPE_INT8    = C.TW_TYPE_INT8
	TYPE_CHAR    = C.TW_TYPE_CHAR
	TYPE_BOOL32  = C.TW_TYPE_BOOL32
	TYPE_BOOL16  = C.TW_TYPE_BOOL16
	TYPE_BOOL8   = C.TW_TYPE_BOOL8
)

func (bar *Bar) AddVarRW(name string, type_ Type, var_ unsafe.Pointer, def string) bool {
	return toBool(C.TwAddVarRW((*C.TwBar)(bar), C.CString(name), C.TwType(type_), var_, C.CString(def)))
}

func (bar *Bar) AddVarRO(name string, type_ Type, var_ unsafe.Pointer, def string) bool {
	return toBool(C.TwAddVarRO((*C.TwBar)(bar), C.CString(name), C.TwType(type_), var_, C.CString(def)))
}

func (bar *Bar) AddSeparator(name string, def string) bool {
	return toBool(C.TwAddSeparator((*C.TwBar)(bar), C.CString(name), C.CString(def)))
}

func (bar *Bar) RemoveVar(name string) bool {
	return toBool(C.TwRemoveVar((*C.TwBar)(bar), C.CString(name)))
}

func (bar *Bar) RemoveAllVars() bool {
	return toBool(C.TwRemoveAllVars((*C.TwBar)(bar)))
}

type EnumVal C.TwEnumVal
type StructMember C.TwStructMember

func Define(def string) bool {
	return toBool(C.TwDefine(C.CString(def)))
}

func DefineEnum(name string, enumValues []EnumVal) Type {
	numValues := C.uint(len(enumValues))
	return Type(C.TwDefineEnum(C.CString(name), (*C.TwEnumVal)(&enumValues[0]), numValues))
}

func DefineEnumFromString(name string, enumString string) {
	C.TwDefineEnumFromString(C.CString(name), C.CString(enumString))
}

type ParamValueType C.TwParamValueType

const (
	PARAM_INT32   = C.TW_PARAM_INT32
	PARAM_FLOAT   = C.TW_PARAM_FLOAT
	PARAM_DOUBLE  = C.TW_PARAM_DOUBLE
	PARAM_CSTRING = C.TW_PARAM_CSTRING // Null-terminated array of char (ie, c-string)
)

func (bar *Bar) GetParam(varName string, paramName string, paramValueType ParamValueType, outValueMaxCount int, outValues unsafe.Pointer) bool {
	return toBool(C.TwGetParam((*C.TwBar)(bar), C.CString(varName), C.CString(paramName), C.TwParamValueType(paramValueType), C.uint(outValueMaxCount), unsafe.Pointer(outValues)))
}

func (bar *Bar) TwSetParam(varName string, paramName string, paramValueType ParamValueType, inValueCount int, inValues unsafe.Pointer) bool {
	return toBool(C.TwSetParam((*C.TwBar)(bar), C.CString(varName), C.CString(paramName), C.TwParamValueType(paramValueType), C.uint(inValueCount), unsafe.Pointer(inValues)))
}

// ----------------------------------------------------------------------------
//  Management functions and definitions
// ----------------------------------------------------------------------------

type GraphAPI C.TwGraphAPI

const (
	OPENGL_CORE = C.TW_OPENGL_CORE
	OPENGL      = C.TW_OPENGL
	DIRECT3D9   = C.TW_DIRECT3D9
	DIRECT3D10  = C.TW_DIRECT3D10
	DIRECT3D11  = C.TW_DIRECT3D11
)

func Init(graphAPI GraphAPI, device unsafe.Pointer) bool {
	return toBool(C.TwInit(C.TwGraphAPI(graphAPI), device))
}

func Terminate() {
	C.TwTerminate()
}

func Draw() bool {
	return toBool(C.TwDraw())
}

func WindowSize(width, height int) bool {
	return toBool(C.TwWindowSize(C.int(width), C.int(height)))
}

func SetCurrentWindow(windowID int) bool {
	return toBool(C.TwSetCurrentWindow(C.int(windowID)))
}

func GetCurrentWindow() int {
	return int(C.TwGetCurrentWindow())
}

func WindowExists(windowID int) bool {
	return toBool(C.TwWindowExists(C.int(windowID)))
}

type KeyModifier C.TwKeyModifier

const (
	KMOD_NONE  = C.TW_KMOD_NONE // same codes as SDL keysym.mod
	KMOD_SHIFT = C.TW_KMOD_SHIFT
	KMOD_CTRL  = C.TW_KMOD_CTRL
	KMOD_ALT   = C.TW_KMOD_ALT
	KMOD_META  = C.TW_KMOD_META
)

type KeySpecial C.TwKeySpecial

const (
	KEY_BACKSPACE = C.TW_KEY_BACKSPACE
	KEY_TAB       = C.TW_KEY_TAB
	KEY_CLEAR     = C.TW_KEY_CLEAR
	KEY_RETURN    = C.TW_KEY_RETURN
	KEY_PAUSE     = C.TW_KEY_PAUSE
	KEY_ESCAPE    = C.TW_KEY_ESCAPE
	KEY_SPACE     = C.TW_KEY_SPACE
	KEY_DELETE    = C.TW_KEY_DELETE
	KEY_UP        = C.TW_KEY_UP
	KEY_DOWN      = C.TW_KEY_DOWN
	KEY_RIGHT     = C.TW_KEY_RIGHT
	KEY_LEFT      = C.TW_KEY_LEFT
	KEY_INSERT    = C.TW_KEY_INSERT
	KEY_HOME      = C.TW_KEY_HOME
	KEY_END       = C.TW_KEY_END
	KEY_PAGE_UP   = C.TW_KEY_PAGE_UP
	KEY_PAGE_DOWN = C.TW_KEY_PAGE_DOWN
	KEY_F1        = C.TW_KEY_F1
	KEY_F2        = C.TW_KEY_F2
	KEY_F3        = C.TW_KEY_F3
	KEY_F4        = C.TW_KEY_F4
	KEY_F5        = C.TW_KEY_F5
	KEY_F6        = C.TW_KEY_F6
	KEY_F7        = C.TW_KEY_F7
	KEY_F8        = C.TW_KEY_F8
	KEY_F9        = C.TW_KEY_F9
	KEY_F10       = C.TW_KEY_F10
	KEY_F11       = C.TW_KEY_F11
	KEY_F12       = C.TW_KEY_F12
	KEY_F13       = C.TW_KEY_F13
	KEY_F14       = C.TW_KEY_F14
	KEY_F15       = C.TW_KEY_F15
	KEY_LAST      = C.TW_KEY_LAST
)

func KeyPressed(key int, modifiers int) bool {
	return toBool(C.TwKeyPressed(C.int(key), C.int(modifiers)))
}

func KeyTest(key int, modifiers int) bool {
	return toBool(C.TwKeyTest(C.int(key), C.int(modifiers)))
}

type MouseAction C.TwMouseAction

const (
	MOUSE_RELEASED = C.TW_MOUSE_RELEASED
	MOUSE_PRESSED  = C.TW_MOUSE_PRESSED
)

type MouseButtonID C.TwMouseButtonID

const (
	MOUSE_LEFT   = C.TW_MOUSE_LEFT   // same code as SDL_BUTTON_LEFT
	MOUSE_MIDDLE = C.TW_MOUSE_MIDDLE // same code as SDL_BUTTON_MIDDLE
	MOUSE_RIGHT  = C.TW_MOUSE_RIGHT  // same code as SDL_BUTTON_RIGHT
)

func MouseButton(action MouseAction, button MouseButtonID) bool {
	return toBool(C.TwMouseButton(C.TwMouseAction(action), C.TwMouseButtonID(button)))
}

func MouseMotion(mouseX, mouseY int) bool {
	return toBool(C.TwMouseMotion(C.int(mouseX), C.int(mouseY)))
}

func MouseWheel(pos int) bool {
	return toBool(C.TwMouseWheel(C.int(pos)))
}

func GetLastError() string {
	return C.GoString((*C.char)(C.TwGetLastError()))
}

func EventMouseButtonGLFW(glfwButton, glfwAction int) {
	C.TwEventMouseButtonGLFW(C.int(glfwButton), C.int(glfwAction))
}

func EventKeyGLFW(glfwKey, glfwAction int) {
	C.TwEventKeyGLFW(C.int(glfwKey), C.int(glfwAction))
}

func EventCharGLFW(glfwChar, glfwAction int) {
	C.TwEventCharGLFW(C.int(glfwChar), C.int(glfwAction))
}

func EventMousePosGLFW(mouseX, mouseY int) {
	C.TwEventMousePosGLFW(C.int(mouseX), C.int(mouseY))
}

func EventMouseWheelGLFW(wheelPos int) {
	C.TwEventMouseWheelGLFW(C.int(wheelPos))
}
