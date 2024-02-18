// @Time  : 2024/2/18 15:52
// @Email: jtyoui@qq.com

package gonewword

import (
	"unicode"
)

// Cuter 切词接口，输入一段文本，对进行切词
type Cuter interface {
	Cut(string) []string
}

type mix struct{}

// NewMixCut
// 混合切词，切词的好坏直接影响到结果的好坏
// 包括对：中文、英文、符号、数字等切词
func NewMixCut() Cuter {
	return &mix{}
}

// Cut 切词的步骤
// 1. 先判断属于那个类型
// 2. 在判断是否需要切割
func (m *mix) Cut(text string) (cut []string) {
	c := make([]rune, 0)
	codes := []rune(text)
	length := len(codes)

	for i, r := range codes {

		if unicode.IsDigit(r) || unicode.IsNumber(r) { // 判断是否是数字
			c = append(c, r)
			continue
		}

		if unicode.Is(unicode.Scripts["Han"], r) { // 判断是否是中文
			c = append(c, r)
			cut = append(cut, string(c))
			// 如果是中文，则清空切词
			c = c[:0]
			continue
		}

		// 39的符号比较特殊，属于英文中的' ，比如在 Mr's 中不拆分
		if r == 39 {
			if i < length-1 {
				if IsEnLetter(codes[i-1]) && IsEnLetter(codes[i+1]) {
					c = append(c, r)
					continue
				}
			}
		}

		if unicode.IsPunct(r) || unicode.IsSymbol(r) { // 判断是特殊符号
			if len(c) > 0 {
				cut = append(cut, string(c))
			}
			cut = append(cut, string(r))
			c = c[:0]
			continue
		}

		if IsEnLetter(r) { // 判断是英文字母
			c = append(c, r)

			if i < length-1 {
				next := codes[i+1] // 如果下一位不是字母。那么就切词
				if !IsEnLetter(next) && next != 39 {
					cut = append(cut, string(c))
					c = c[:0]
				}
			}
			continue
		}

		if unicode.IsSpace(r) { // 判断是否是空白符
			if len(c) > 0 {
				cut = append(cut, string(c))
				c = c[:0]
			}
		}
	}

	if len(c) > 0 {
		cut = append(cut, string(c))
	}
	return
}

// IsEnLetter 判断是否是英文字母
func IsEnLetter(s rune) bool {
	return (s >= 97 && s <= 122) || (s >= 65 && s <= 90)
}
