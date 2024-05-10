// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package playback

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
	"github.com/saltosystems/winrt-go/windows/foundation/collections"
)

const SignatureMediaBreak string = "rc(Windows.Media.Playback.MediaBreak;{714be270-0def-4ebc-a489-6b34930e1558})"

type MediaBreak struct {
	ole.IUnknown
}

func (impl *MediaBreak) GetPlaybackList() (*MediaPlaybackList, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaBreak))
	defer itf.Release()
	v := (*iMediaBreak)(unsafe.Pointer(itf))
	return v.GetPlaybackList()
}

func (impl *MediaBreak) GetPresentationPosition() (*foundation.IReference, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaBreak))
	defer itf.Release()
	v := (*iMediaBreak)(unsafe.Pointer(itf))
	return v.GetPresentationPosition()
}

func (impl *MediaBreak) GetInsertionMethod() (MediaBreakInsertionMethod, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaBreak))
	defer itf.Release()
	v := (*iMediaBreak)(unsafe.Pointer(itf))
	return v.GetInsertionMethod()
}

func (impl *MediaBreak) GetCustomProperties() (*collections.ValueSet, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaBreak))
	defer itf.Release()
	v := (*iMediaBreak)(unsafe.Pointer(itf))
	return v.GetCustomProperties()
}

func (impl *MediaBreak) GetCanStart() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaBreak))
	defer itf.Release()
	v := (*iMediaBreak)(unsafe.Pointer(itf))
	return v.GetCanStart()
}

func (impl *MediaBreak) SetCanStart(value bool) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaBreak))
	defer itf.Release()
	v := (*iMediaBreak)(unsafe.Pointer(itf))
	return v.SetCanStart(value)
}

const GUIDiMediaBreak string = "714be270-0def-4ebc-a489-6b34930e1558"
const SignatureiMediaBreak string = "{714be270-0def-4ebc-a489-6b34930e1558}"

type iMediaBreak struct {
	ole.IInspectable
}

type iMediaBreakVtbl struct {
	ole.IInspectableVtbl

	GetPlaybackList         uintptr
	GetPresentationPosition uintptr
	GetInsertionMethod      uintptr
	GetCustomProperties     uintptr
	GetCanStart             uintptr
	SetCanStart             uintptr
}

func (v *iMediaBreak) VTable() *iMediaBreakVtbl {
	return (*iMediaBreakVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iMediaBreak) GetPlaybackList() (*MediaPlaybackList, error) {
	var out *MediaPlaybackList
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetPlaybackList,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out MediaPlaybackList
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaBreak) GetPresentationPosition() (*foundation.IReference, error) {
	var out *foundation.IReference
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetPresentationPosition,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.IReference
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaBreak) GetInsertionMethod() (MediaBreakInsertionMethod, error) {
	var out MediaBreakInsertionMethod
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetInsertionMethod,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out MediaBreakInsertionMethod
	)

	if hr != 0 {
		return MediaBreakInsertionMethodInterrupt, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaBreak) GetCustomProperties() (*collections.ValueSet, error) {
	var out *collections.ValueSet
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetCustomProperties,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out collections.ValueSet
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaBreak) GetCanStart() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetCanStart,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaBreak) SetCanStart(value bool) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetCanStart,
		uintptr(unsafe.Pointer(v)),                // this
		uintptr(*(*byte)(unsafe.Pointer(&value))), // in bool
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

const GUIDiMediaBreakFactory string = "4516e002-18e0-4079-8b5f-d33495c15d2e"
const SignatureiMediaBreakFactory string = "{4516e002-18e0-4079-8b5f-d33495c15d2e}"

type iMediaBreakFactory struct {
	ole.IInspectable
}

type iMediaBreakFactoryVtbl struct {
	ole.IInspectableVtbl

	MediaBreakCreate                         uintptr
	MediaBreakCreateWithPresentationPosition uintptr
}

func (v *iMediaBreakFactory) VTable() *iMediaBreakFactoryVtbl {
	return (*iMediaBreakFactoryVtbl)(unsafe.Pointer(v.RawVTable))
}

func MediaBreakCreate(insertionMethod MediaBreakInsertionMethod) (*MediaBreak, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.Playback.MediaBreak", ole.NewGUID(GUIDiMediaBreakFactory))
	if err != nil {
		return nil, err
	}
	v := (*iMediaBreakFactory)(unsafe.Pointer(inspectable))

	var out *MediaBreak
	hr, _, _ := syscall.SyscallN(
		v.VTable().MediaBreakCreate,
		0,                             // this is a static func, so there's no this
		uintptr(insertionMethod),      // in MediaBreakInsertionMethod
		uintptr(unsafe.Pointer(&out)), // out MediaBreak
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func MediaBreakCreateWithPresentationPosition(insertionMethod MediaBreakInsertionMethod, presentationPosition foundation.TimeSpan) (*MediaBreak, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.Playback.MediaBreak", ole.NewGUID(GUIDiMediaBreakFactory))
	if err != nil {
		return nil, err
	}
	v := (*iMediaBreakFactory)(unsafe.Pointer(inspectable))

	var out *MediaBreak
	hr, _, _ := syscall.SyscallN(
		v.VTable().MediaBreakCreateWithPresentationPosition,
		0,                                      // this is a static func, so there's no this
		uintptr(insertionMethod),               // in MediaBreakInsertionMethod
		uintptr(presentationPosition.Duration), // in foundation.TimeSpan
		uintptr(unsafe.Pointer(&out)),          // out MediaBreak
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
