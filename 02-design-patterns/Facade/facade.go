package Facade

import "fmt"

/*外观（Facade）模式又叫作门面模式，
是一种通过为多个复杂的子系统提供一个一致的接口，
而使这些子系统更加容易被访问的模式。
该模式对外有一个统一接口，
外部应用程序不用关心内部子系统的具体细节，
这样会大大降低应用程序的复杂度，提高了程序的可维护性。

外观（Facade）模式的结构比较简单，主要是定义了一个高层接口。它包含了对各个子系统的引用，客户端可以通过它访问各个子系统的功能。现在来分析其基本结构和实现方法。
1. 模式的结构
外观（Facade）模式包含以下主要角色。
外观（Facade）角色：为多个子系统对外提供一个共同的接口。
子系统（Sub System）角色：实现系统的部分功能，客户可以通过外观角色访问它。
客户（Client）角色：通过一个外观角色访问各个子系统的功能。

*/

type CarHeader struct {
}

func (header *CarHeader) SetHeader() {
	fmt.Println("Set CarHeader ok.")
}

type CarBody struct {
}

func (body *CarBody) SetBody() {
	fmt.Println("set body ok.")
}

type CarFacade struct {
	Header *CarHeader
	Body   *CarBody
}

func NewCarFacade() CarFacade {
	return CarFacade{
		Header: &CarHeader{},
		Body:   &CarBody{},
	}
}
func (car *CarFacade) CreateCompleteCar() {
	fmt.Println(" creating a car...")
	car.Header.SetHeader()
	car.Body.SetBody()
	fmt.Println("it is a new car.")
}
