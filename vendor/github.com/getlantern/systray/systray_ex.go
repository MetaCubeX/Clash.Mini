package systray

import (
	"os"
	"runtime"
)

type MenuItemEx struct {
	Item     *MenuItem
	Parent   *MenuItemEx
	Child    []*MenuItemEx
	Callback func(menuItemEx *MenuItemEx)
}

var (
	MenuList []*MenuItemEx
)

// RunEx SystrayEx入口 须在init()调用
func RunEx(onReady func(), onExit func()) {
	// use it on init
	go func() {
		runtime.LockOSThread()
		Run(onReady, func() {
			go onExit()
			os.Exit(1)
		})
		runtime.UnlockOSThread()
	}()
}

func getMenuItemEx(title string, tooltip string, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	menuItem := AddMenuItem(title, tooltip)
	menuItemEx = &MenuItemEx{
		Item:     menuItem,
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
	}
	subMenuItem.setExObj(subMenuItemEx)
	subMenuItemEx.Callback = func(e *MenuItemEx) {
		go f(subMenuItemEx)
	}
	return subMenuItemEx
}

func getSubMenuItemCheckboxEx(menuItem *MenuItem, title string, tooltip string, isChecked bool, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	subMenuItem := menuItem.AddSubMenuItemCheckbox(title, tooltip, isChecked)
	menuItemEx = &MenuItemEx{
		Item:     subMenuItem,
	}
	subMenuItem.setExObj(menuItemEx)
	menuItemEx.Callback = func(e *MenuItemEx) {
		go f(menuItemEx)
	}
	return menuItemEx
}

// AddMenuItemEx 添加增强版菜单项（同级）
func (mie *MenuItemEx) AddMenuItemEx(title string, tooltip string, f func(menuItem *MenuItemEx)) (menuItemEx *MenuItemEx) {
	menuItemEx = getSubMenuItemEx(mie.Parent.Item, title, tooltip, f)
	menuItemEx.Parent = mie.Parent
	mie.Parent.Child = append(mie.Parent.Child, menuItemEx)
	return
}

// AddMenuItemExBind 添加增强版菜单项（同级）并绑定到引用对象
func (mie *MenuItemEx) AddMenuItemExBind(title string, tooltip string, f func(menuItem *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	menuItemEx = getSubMenuItemEx(mie.Parent.Item, title, tooltip, f)
	menuItemEx.Parent = mie.Parent
	mie.Parent.Child = append(mie.Parent.Child, menuItemEx)
	*v = *menuItemEx
	return
}

// AddMenuItemExBind 添加增强版勾选框菜单项（同级）
func (mie *MenuItemEx) AddMenuItemCheckboxEx(title string, tooltip string, isChecked bool, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	menuItemEx = getSubMenuItemCheckboxEx(mie.Parent.Item, title, tooltip, isChecked, f)
	menuItemEx.Parent = mie.Parent
	mie.Parent.Child = append(mie.Parent.Child, menuItemEx)
	return
}

// AddMenuItemCheckboxExBind 添加增强版菜单项并绑定到引用对象
func (mie *MenuItemEx) AddMenuItemCheckboxExBind(title string, tooltip string, isChecked bool, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	menuItemEx = getSubMenuItemCheckboxEx(mie.Parent.Item, title, tooltip, isChecked, f)
	menuItemEx.Parent = mie.Parent
	mie.Parent.Child = append(mie.Parent.Child, menuItemEx)
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
	//mie.Child = append(mie.Child, subMenuItemEx)
	//return mie
	menuItemEx = getSubMenuItemEx(mie.Item, title, tooltip, f)
	menuItemEx.Parent = mie
	mie.Child = append(mie.Child, menuItemEx)
	return
}

// AddSubMenuItemExBind 添加增强版子菜单项并绑定到引用对象
func (mie *MenuItemEx) AddSubMenuItemExBind(title string, tooltip string, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	//subMenuItemEx := getMenuItemEx(title, tooltip, f)
	//mie.Child = append(mie.Child, subMenuItemEx)
	//return mie
	menuItemEx = getSubMenuItemEx(mie.Item, title, tooltip, f)
	menuItemEx.Parent = mie
	mie.Child = append(mie.Child, menuItemEx)
	*v = *menuItemEx
	return
}

// AddSubMenuItemCheckboxEx 添加增强版勾选框子菜单项
func (mie *MenuItemEx) AddSubMenuItemCheckboxEx(title string, tooltip string, isChecked bool, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	//subMenuItemEx := getMenuItemEx(title, tooltip, f)
	//mie.Child = append(mie.Child, subMenuItemEx)
	//return mie
	menuItemEx = getSubMenuItemCheckboxEx(mie.Item, title, tooltip, isChecked, f)
	menuItemEx.Parent = mie
	mie.Child = append(mie.Child, menuItemEx)
	return
}

// AddSubMenuItemCheckboxExBind 添加增强版勾选框子菜单项并绑定到引用对象
func (mie *MenuItemEx) AddSubMenuItemCheckboxExBind(title string, tooltip string, isChecked bool, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	//subMenuItemEx := getMenuItemEx(title, tooltip, f)
	//mie.Child = append(mie.Child, subMenuItemEx)
	//return mie
	menuItemEx = getSubMenuItemCheckboxEx(mie.Item, title, tooltip, isChecked, f)
	menuItemEx.Parent = mie
	mie.Child = append(mie.Child, menuItemEx)
	*v = *menuItemEx
	return
}

// NilCallback 空回调
func NilCallback(menuItem *MenuItemEx) {
	//log.Infoln("clicked %s, id: %d", menuItem.Item.GetTitle(), menuItem.Item.GetId())
}

// Hide hides a menu item
func (menuItemEx *MenuItemEx) Hide() {
	menuItemEx.Item.Hide()
}

// Show shows a previously hidden menu item
func (menuItemEx *MenuItemEx) Show() {
	menuItemEx.Item.Show()
}

// Checked returns if the menu item has a check mark
func (menuItemEx *MenuItemEx) Checked() bool {
	return menuItemEx.Item.Checked()
}

// Check a menu item regardless if it's previously checked or not
func (menuItemEx *MenuItemEx) Check() {
	menuItemEx.Item.Check()
}

// Uncheck a menu item regardless if it's previously unchecked or not
func (menuItemEx *MenuItemEx) Uncheck() {
	menuItemEx.Item.Uncheck()
}

// Get ID of a menu item
func (menuItemEx *MenuItemEx) GetId() uint32 {
	return menuItemEx.Item.GetId()
}

// Get title of a menu item
func (menuItemEx *MenuItemEx) GetTitle() string {
	return menuItemEx.Item.GetTitle()
}

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
