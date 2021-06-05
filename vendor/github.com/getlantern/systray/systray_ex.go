package systray

import (
	"container/list"
	"os"
	"runtime"
)

type MenuItemEx struct {
	Item     	*MenuItem
	Parent   	*MenuItemEx
	Children    *list.List
	Callback 	func(menuItemEx *MenuItemEx)
}

var (
	MenuList 	[]*MenuItemEx
)

// RunEx SystrayEx入口 须在init()调用
func RunEx(onReady func(), onExit func()) {
	// use it on init
	go func() {
		runtime.LockOSThread()
		Run(onReady, func() {
			onExit()
			os.Exit(1)
		})
		runtime.UnlockOSThread()
	}()
}

// NilCallback 空回调
func NilCallback(menuItem *MenuItemEx) {
	//log.Infoln("clicked %s, id: %d", menuItem.Item.GetTitle(), menuItem.Item.GetId())
}

// AddMenuItemEx 添加增强版菜单项（同级）
func (mie *MenuItemEx) AddMenuItemEx(title string, tooltip string, f func(menuItem *MenuItemEx)) (menuItemEx *MenuItemEx) {
	menuItemEx = getSubMenuItemEx(mie.Parent.Item, title, tooltip, f)
	menuItemEx.Parent = mie.Parent
	mie.Parent.Children.PushBack(menuItemEx)
	return
}

// AddMenuItemExBind 添加增强版菜单项（同级）并绑定到引用对象
func (mie *MenuItemEx) AddMenuItemExBind(title string, tooltip string, f func(menuItem *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	menuItemEx = getSubMenuItemEx(mie.Parent.Item, title, tooltip, f)
	menuItemEx.Parent = mie.Parent
	mie.Parent.Children.PushBack(menuItemEx)
	*v = *menuItemEx
	return
}

// AddMenuItemExBind 添加增强版勾选框菜单项（同级）
func (mie *MenuItemEx) AddMenuItemCheckboxEx(title string, tooltip string, isChecked bool, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	menuItemEx = getSubMenuItemCheckboxEx(mie.Parent.Item, title, tooltip, isChecked, f)
	menuItemEx.Parent = mie.Parent
	mie.Parent.Children.PushBack(menuItemEx)
	return
}

// AddMenuItemCheckboxExBind 添加增强版菜单项并绑定到引用对象
func (mie *MenuItemEx) AddMenuItemCheckboxExBind(title string, tooltip string, isChecked bool, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	menuItemEx = getSubMenuItemCheckboxEx(mie.Parent.Item, title, tooltip, isChecked, f)
	menuItemEx.Parent = mie.Parent
	mie.Parent.Children.PushBack(menuItemEx)
	*v = *menuItemEx
	return
}

// AddMainMenuItemEx 添加增强版主菜单项
func AddMainMenuItemEx(title string, tooltip string, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	menuItemEx = getMenuItemEx(title, tooltip, f)
	MenuList = append(MenuList, menuItemEx)
	return
}

// AddMainMenuItemEx 添加增强版子菜单项
func (mie *MenuItemEx) AddSubMenuItemEx(title string, tooltip string, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	//subMenuItemEx := getMenuItemEx(title, tooltip, f)
	//mie.Children = append(mie.Children, subMenuItemEx)
	//return mie
	menuItemEx = getSubMenuItemEx(mie.Item, title, tooltip, f)
	menuItemEx.Parent = mie
	mie.Children.PushBack(menuItemEx)
	//mie.setSubMenu()
	return
}

// AddSubMenuItemExBind 添加增强版子菜单项并绑定到引用对象
func (mie *MenuItemEx) AddSubMenuItemExBind(title string, tooltip string, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	menuItemEx = mie.AddSubMenuItemEx(title, tooltip, f)
	*v = *menuItemEx
	return
}

// AddSubMenuItemCheckboxEx 添加增强版勾选框子菜单项
func (mie *MenuItemEx) AddSubMenuItemCheckboxEx(title string, tooltip string, isChecked bool, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	//subMenuItemEx := getMenuItemEx(title, tooltip, f)
	//mie.Children = append(mie.Children, subMenuItemEx)
	//return mie
	menuItemEx = getSubMenuItemCheckboxEx(mie.Item, title, tooltip, isChecked, f)
	menuItemEx.Parent = mie
	mie.Children.PushBack(menuItemEx)
	//mie.setSubMenu()
	return
}

// AddSubMenuItemCheckboxExBind 添加增强版勾选框子菜单项并绑定到引用对象
func (mie *MenuItemEx) AddSubMenuItemCheckboxExBind(title string, tooltip string, isChecked bool, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	menuItemEx = mie.AddSubMenuItemCheckboxEx(title, tooltip, isChecked, f)
	*v = *menuItemEx
	return
}

//// AddSeparator adds a separator bar to the menu
//func AddSeparator(mie *MenuItemEx) *MenuItemEx {
//	menuItemEx := &MenuItemEx{
//	}
//	addSeparator(menuItemEx.GetId())
//	//addSeparator(atomic.AddUint32(&currentID, 1))
//	return menuItemEx
//}

// SwitchCheckboxGroup 切换增强版勾选框菜单项组 设置指定项勾选与否，组内其他项相反
func SwitchCheckboxGroup(newValue *MenuItemEx, checked bool, values []*MenuItemEx) {
	for _, value := range values {
		if value.GetId() == newValue.GetId() {
			if checked {
				value.Check()
			} else {
				value.Uncheck()
			}
		} else {
			if checked {
				value.Uncheck()
			} else {
				value.Check()
			}
		}
	}
}

// SwitchCheckboxBrother 切换增强版勾选框菜单项组 设置指定项勾选与否，其他兄弟项相反
func SwitchCheckboxBrother(newValue *MenuItemEx, checked bool) {
	SwitchCheckboxGroupByList(newValue, checked, newValue.Parent.Children)
}

// SwitchCheckboxGroupByList 切换增强版勾选框菜单项组 设置指定项勾选与否，组内其他项相反
func SwitchCheckboxGroupByList(newValue *MenuItemEx, checked bool, values *list.List) {
	if values == nil || values.Len() == 0 {
		newValue.Checked()
	}
	for e := values.Front(); e != nil; e = e.Next() {
		value := e.Value.(*MenuItemEx)
		if value.GetId() == newValue.GetId() {
			if checked {
				value.Check()
			} else {
				value.UncheckFull()
			}
		} else {
			if checked {
				value.UncheckFull()
			} else {
				value.Check()
			}
		}
	}
}

// UncheckFull uncheck with children
func (menuItemEx *MenuItemEx) UncheckFull() *MenuItemEx {
	for e := menuItemEx.Children.Front(); e != nil; e = e.Next() {
		e.Value.(*MenuItemEx).UncheckFull()
	}
	menuItemEx.Uncheck()
	return menuItemEx
}

// SetIcon sets the icon of a menu item. Only works on macOS and Windows.
// iconBytes should be the content of .ico/.jpg/.png
func (menuItemEx *MenuItemEx) SetIcon(iconBytes []byte) *MenuItemEx {
	menuItemEx.Item.SetIcon(iconBytes)
	return menuItemEx
}

// SetTemplateIcon sets the icon of a menu item as a template icon (on macOS). On Windows, it
// falls back to the regular icon bytes and on Linux it does nothing.
// templateIconBytes and regularIconBytes should be the content of .ico for windows and
// .ico/.jpg/.png for other platforms.
func (menuItemEx *MenuItemEx) SetTemplateIcon(templateIconBytes []byte, regularIconBytes []byte) *MenuItemEx {
	menuItemEx.Item.SetTemplateIcon(templateIconBytes, regularIconBytes)
	return menuItemEx
}

// SetTitle set the text to display on a menu item
func (menuItemEx *MenuItemEx) SetTitle(title string) *MenuItemEx {
	menuItemEx.Item.SetTitle(title)
	return menuItemEx
}

// SetTooltip set the tooltip to show when mouse hover
func (menuItemEx *MenuItemEx) SetTooltip(tooltip string) *MenuItemEx {
	menuItemEx.Item.SetTitle(tooltip)
	return menuItemEx
}

// Disabled checks if the menu item is disabled
func (menuItemEx *MenuItemEx) Disabled() bool {
	return menuItemEx.Item.Disabled()
}

// Enable a menu item regardless if it's previously enabled or not
func (menuItemEx *MenuItemEx) Enable() *MenuItemEx {
	menuItemEx.Item.Enable()
	return menuItemEx
}

// Disable a menu item regardless if it's previously disabled or not
func (menuItemEx *MenuItemEx) Disable() *MenuItemEx {
	menuItemEx.Item.Disable()
	return menuItemEx
}

// Hide hides a menu item
func (menuItemEx *MenuItemEx) Hide() *MenuItemEx {
	menuItemEx.Item.Hide()
	return menuItemEx
}

// Show shows a previously hidden menu item
func (menuItemEx *MenuItemEx) Show() *MenuItemEx {
	menuItemEx.Item.Show()
	return menuItemEx
}

// Checked returns if the menu item has a check mark
func (menuItemEx *MenuItemEx) Checked() bool {
	return menuItemEx.Item.Checked()
}

// Check a menu item regardless if it's previously checked or not
func (menuItemEx *MenuItemEx) Check() *MenuItemEx {
	menuItemEx.Item.Check()
	return menuItemEx
}

// Uncheck a menu item regardless if it's previously unchecked or not
func (menuItemEx *MenuItemEx) Uncheck() *MenuItemEx {
	menuItemEx.Item.Uncheck()
	return menuItemEx
}

// Get ID of a menu item
func (menuItemEx *MenuItemEx) GetId() uint32 {
	return menuItemEx.Item.GetId()
}

// Get title of a menu item
func (menuItemEx *MenuItemEx) GetTitle() string {
	return menuItemEx.Item.GetTitle()
}

// Delete a menu item with children
func (menuItemEx *MenuItemEx) Delete() {
	menuItemEx.ClearChildren()
	menuItemEx.Hide()
}

func (menuItemEx *MenuItemEx) ClearChildren() *MenuItemEx {
	if menuItemEx.Children.Len() > 0 {
		lChild := menuItemEx.Children
		var next *list.Element
		for e := lChild.Front(); e != nil; e = next {
			next = e.Next()
			child := lChild.Remove(e).(*MenuItemEx)
			child.ClearChildren()
			child.Hide()
		}
	}
	menuItemEx.unsetSubMenu()
	return menuItemEx
}

func getMenuItemEx(title string, tooltip string, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	menuItem := AddMenuItem(title, tooltip)
	menuItemEx = &MenuItemEx{
		Item:		menuItem,
		Children: 	list.New(),
	}
	menuItem.setExObj(menuItemEx)
	menuItemEx.Callback = func(e *MenuItemEx) {
		go f(menuItemEx)
	}
	return menuItemEx
}

func getSubMenuItemEx(menuItem *MenuItem, title string, tooltip string, f func(menuItemEx *MenuItemEx)) (subMenuItemEx *MenuItemEx) {
	subMenuItem := menuItem.AddSubMenuItem(title, tooltip)
	subMenuItemEx = &MenuItemEx{
		Item:     subMenuItem,
		Children: 	list.New(),
	}
	subMenuItem.setExObj(subMenuItemEx)
	subMenuItemEx.Callback = func(e *MenuItemEx) {
		go f(subMenuItemEx)
	}
	return subMenuItemEx
}

func getSubMenuItemCheckboxEx(menuItem *MenuItem, title string, tooltip string, isChecked bool, f func(menuItemEx *MenuItemEx)) (subMenuItemEx *MenuItemEx) {
	subMenuItem := menuItem.AddSubMenuItemCheckbox(title, tooltip, isChecked)
	subMenuItemEx = &MenuItemEx{
		Item:     subMenuItem,
		Children: 	list.New(),
	}
	subMenuItem.setExObj(subMenuItemEx)
	subMenuItemEx.Callback = func(e *MenuItemEx) {
		go f(subMenuItemEx)
	}
	return subMenuItemEx
}

func (menuItemEx *MenuItemEx) unsetSubMenu() *MenuItemEx {
	item := menuItemEx.Item
	_, err := wt.convertToNormalMenu(uint32(item.id))
	if err != nil {
		log.Errorf("Unable to unsetSubMenu: %v", err)
		return menuItemEx
	}
	return menuItemEx
}

func (menuItemEx *MenuItemEx) setSubMenu() *MenuItemEx {
	item := menuItemEx.Item
	_, err := wt.convertToSubMenu(uint32(item.id))
	if err != nil {
		log.Errorf("Unable to setSubMenu: %v", err)
		return menuItemEx
	}
	return menuItemEx
}
