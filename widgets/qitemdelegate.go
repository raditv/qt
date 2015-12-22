package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QItemDelegate struct {
	QAbstractItemDelegate
}

type QItemDelegate_ITF interface {
	QAbstractItemDelegate_ITF
	QItemDelegate_PTR() *QItemDelegate
}

func PointerFromQItemDelegate(ptr QItemDelegate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemDelegate_PTR().Pointer()
	}
	return nil
}

func NewQItemDelegateFromPointer(ptr unsafe.Pointer) *QItemDelegate {
	var n = new(QItemDelegate)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QItemDelegate_") {
		n.SetObjectName("QItemDelegate_" + qt.Identifier())
	}
	return n
}

func (ptr *QItemDelegate) QItemDelegate_PTR() *QItemDelegate {
	return ptr
}

func (ptr *QItemDelegate) HasClipping() bool {
	defer qt.Recovering("QItemDelegate::hasClipping")

	if ptr.Pointer() != nil {
		return C.QItemDelegate_HasClipping(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QItemDelegate) SetClipping(clip bool) {
	defer qt.Recovering("QItemDelegate::setClipping")

	if ptr.Pointer() != nil {
		C.QItemDelegate_SetClipping(ptr.Pointer(), C.int(qt.GoBoolToInt(clip)))
	}
}

func NewQItemDelegate(parent core.QObject_ITF) *QItemDelegate {
	defer qt.Recovering("QItemDelegate::QItemDelegate")

	return NewQItemDelegateFromPointer(C.QItemDelegate_NewQItemDelegate(core.PointerFromQObject(parent)))
}

func (ptr *QItemDelegate) CreateEditor(parent QWidget_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *QWidget {
	defer qt.Recovering("QItemDelegate::createEditor")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QItemDelegate_CreateEditor(ptr.Pointer(), PointerFromQWidget(parent), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	defer qt.Recovering("QItemDelegate::itemEditorFactory")

	if ptr.Pointer() != nil {
		return NewQItemEditorFactoryFromPointer(C.QItemDelegate_ItemEditorFactory(ptr.Pointer()))
	}
	return nil
}

func (ptr *QItemDelegate) ConnectSetEditorData(f func(editor *QWidget, index *core.QModelIndex)) {
	defer qt.Recovering("connect QItemDelegate::setEditorData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEditorData", f)
	}
}

func (ptr *QItemDelegate) DisconnectSetEditorData() {
	defer qt.Recovering("disconnect QItemDelegate::setEditorData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEditorData")
	}
}

//export callbackQItemDelegateSetEditorData
func callbackQItemDelegateSetEditorData(ptrName *C.char, editor unsafe.Pointer, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QItemDelegate::setEditorData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEditorData"); signal != nil {
		signal.(func(*QWidget, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

}

func (ptr *QItemDelegate) SetItemEditorFactory(factory QItemEditorFactory_ITF) {
	defer qt.Recovering("QItemDelegate::setItemEditorFactory")

	if ptr.Pointer() != nil {
		C.QItemDelegate_SetItemEditorFactory(ptr.Pointer(), PointerFromQItemEditorFactory(factory))
	}
}

func (ptr *QItemDelegate) ConnectSetModelData(f func(editor *QWidget, model *core.QAbstractItemModel, index *core.QModelIndex)) {
	defer qt.Recovering("connect QItemDelegate::setModelData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setModelData", f)
	}
}

func (ptr *QItemDelegate) DisconnectSetModelData() {
	defer qt.Recovering("disconnect QItemDelegate::setModelData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setModelData")
	}
}

//export callbackQItemDelegateSetModelData
func callbackQItemDelegateSetModelData(ptrName *C.char, editor unsafe.Pointer, model unsafe.Pointer, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QItemDelegate::setModelData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModelData"); signal != nil {
		signal.(func(*QWidget, *core.QAbstractItemModel, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQAbstractItemModelFromPointer(model), core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

}

func (ptr *QItemDelegate) SizeHint(option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QItemDelegate::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QItemDelegate_SizeHint(ptr.Pointer(), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QItemDelegate) DestroyQItemDelegate() {
	defer qt.Recovering("QItemDelegate::~QItemDelegate")

	if ptr.Pointer() != nil {
		C.QItemDelegate_DestroyQItemDelegate(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}