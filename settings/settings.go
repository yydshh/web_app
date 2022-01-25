package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")     // 指定配置文件名称（不需要带后缀）
	viper.SetConfigType("yaml")       // 指定配置文件类型
	viper.AddConfigPath("./settings") // 指定查找配置文件的路径，这里使用相对路径
	err = viper.ReadInConfig()        // 读取配置信息
	if err != nil {
		// 读取配置文件失败
		fmt.Printf("viper.ReadInConfig() failed ,err:%v \n", err)
		panic("配置文件读取失败")
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件修改了...")
	})
	return

}