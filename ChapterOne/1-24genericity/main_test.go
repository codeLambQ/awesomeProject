package main

import "testing"

// 测试文件
// 单元测试的文件必须是 _test.go 结尾
// 单元测试的方法必须是 Test_ 开头，入参里必须要有 t *testing.T

func Test_min1_1(t *testing.T) {
	//if min1(ChapterOne, 2) != ChapterOne {
	//	t.Errorf("min1(ChapterOne,2) == %d; want ChapterOne", min1(ChapterOne, 2))
	//}

	// 使用结构体来改造测试
	type args struct {
		a int
		b int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"ChapterOne", args{1, 2}, 1},
		{"2", args{2, 3}, 2},
		{"3", args{12, 23}, 12},
		{"4", args{4, 5}, 4},
	}

	for _, tt := range tests {
		if got := min1(tt.args.a, tt.args.b); got != tt.want {
			t.Errorf("min1() = %v, want %v", got, tt.want)
		}
	}
}

// 标准的 go 语言测试方法
func Test_min1_2(t *testing.T) {
	type args struct {
		a int
		b int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"ChapterOne", args{1, 2}, 1},
		{"2", args{2, 3}, 2},
		{"3", args{12, 23}, 12},
		{"4", args{4, 5}, 4},
	}

	for _, tt := range tests {
		// 可以并发的运行测试
		t.Run(tt.name, func(t *testing.T) {
			if got := min1(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min1() = %v, want %v", got, tt.want)
			}
		})
	}
}
