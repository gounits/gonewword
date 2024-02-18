// @Time  : 2024/2/18 17:22
// @Email: jtyoui@qq.com

package gonewword_test

import (
	"github.com/gounits/gonewword"
	"slices"
	"testing"
)

func TestNewMixCut(t *testing.T) {
	mix := gonewword.NewMixCut()
	values := mix.Cut("Good!阿Q这12年。Mr's 张")
	if slices.Equal(values, []string{"Good", "!", "阿", "Q", "这", "12", "年", "。", "Mr's", "张"}) {
		t.Error("切词测试错误")
	}
}
