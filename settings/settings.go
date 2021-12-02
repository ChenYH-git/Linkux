package settings

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//全局变量,用来保存程序的所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name          string `mapstructure:"name"`
	Mode          string `mapstructure:"mode"`
	Version       string `mapstructure:"version"`
	StartTime     string `mapstructure:"start_time"`
	MachineID     int64  `mapstructure:"machine_id"`
	Port          int    `mapstructure:"port"`
	*LogConfig    `mapstructure:"log"`
	*MySQLConfig  `mapstructure:"mysql"`
	*RedisConfig  `mapstructure:"redis"`
	*WeiXinConfig `mapstructure:"weixin"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host        string        `mapstructure:"host"`
	Password    string        `mapstructure:"password"`
	Port        int           `mapstructure:"port"`
	DB          int           `mapstructure:"db"`
	PoolSize    int           `mapstructure:"pool_size"`
	IdleTimeOut time.Duration `mapstructure:"idle_time_out"`
}

type WeiXinConfig struct {
	Id        string `mapstructure:"id"`
	Secret    string `mapstructure:"secret"`
	GrantType string `mapstructure:"grant_type"`
}

func Init() (err error) {
	viper.SetConfigFile("./conf/config.yaml") // 指定配置文件
	//viper.SetConfigType("yaml") //专用于从远程获取配置信息时指定的配置
	viper.AddConfigPath(".")   // 指定查找配置文件的路径
	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		fmt.Println("viper.ReadInConfig failed, err:", err)
		return // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//把读取到的信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmashal failed, err:%v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config changed")
		//每次信息改变都把改变的信息反序列化到Conf中，实现动态更新
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})

	return
}
