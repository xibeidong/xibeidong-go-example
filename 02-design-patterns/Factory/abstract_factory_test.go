package Factory

import "testing"

func TestXiaomiFactory(t *testing.T) {
	factory := XiaomiFactory{}
	factory.NewPhone().Start()
	factory.NewRouter().OpenWifi()
}
func TestHuaweiFactory(t *testing.T) {
	factory := HuaweiFactory{}
	factory.NewPhone().ShutDown()
	factory.NewRouter().Setting()
}
