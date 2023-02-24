package Factory

import "fmt"

//抽象工厂模式
//例 小米和华为都可以生产phone和router

type Phone interface {
	Start()
	ShutDown()
}

type Router interface {
	OpenWifi()
	Setting()
}

type HuaweiPhone struct {
}

func (phone *HuaweiPhone) Start() {
	fmt.Println("huawei phone start.")
}
func (phone *HuaweiPhone) ShutDown() {
	fmt.Println("huawei phone Shutdown.")
}

type XiaomiPhone struct {
}

func (phone *XiaomiPhone) Start() {
	fmt.Println("Xiaomi phone start.")
}
func (phone *XiaomiPhone) ShutDown() {
	fmt.Println("Xiaomi phone Shutdown.")
}

type HuaweiRouter struct {
}

func (router *HuaweiRouter) OpenWifi() {
	fmt.Println("huawei router open wifi.")
}
func (router *HuaweiRouter) Setting() {
	fmt.Println("huawei router setting.")
}

type XiaomiRouter struct {
}

func (router *XiaomiRouter) OpenWifi() {
	fmt.Println("xiaomi router open wifi.")
}
func (router *XiaomiRouter) Setting() {
	fmt.Println("xiaomi router setting.")
}

type ProductFactory interface {
	NewPhone()
	NewRouter()
}

type HuaweiFactory struct {
}

func (factory *HuaweiFactory) NewPhone() Phone {
	return &HuaweiPhone{}
}

func (factory *HuaweiFactory) NewRouter() Router {
	return &HuaweiRouter{}
}

type XiaomiFactory struct {
}

func (factory *XiaomiFactory) NewPhone() Phone {
	return &XiaomiPhone{}
}

func (factory *XiaomiFactory) NewRouter() Router {
	return &XiaomiRouter{}
}
