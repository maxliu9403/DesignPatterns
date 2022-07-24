package gun

// 具体的产品
type AK47 struct {
	Gun
}

// 具体创建者
func newAK47() IGun {
	return &AK47{
		Gun{
			name:  "ak47",
			power: 10,
		},
	}
}
