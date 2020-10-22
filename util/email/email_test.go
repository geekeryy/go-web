package email

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func init()  {
	viper.SetConfigName("./../../config.yaml")
	if err := viper.ReadInConfig();err!=nil{
		fmt.Println(err)
	}
}

func TestSendMail(t *testing.T) {
	err := SendMail([]string{"1126254578@qq.com"}, "subject", "你好")
	if err != nil {
		t.Error(err)
	}
}
