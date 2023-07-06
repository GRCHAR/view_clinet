package setting

type setting struct {
	videoOutPath string
}

var config = setting{}

func init() {

}

func getSettingInfo() {

}

func getSetting() *setting {
	return &config
}

func (set *setting) saveSetting() {

}
