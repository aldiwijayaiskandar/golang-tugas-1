package matematika

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCekGanjilGenap(t *testing.T) {
	tests :=
		[]struct {
			name     string
			expected string
			request  []int
		}{
			{
				name:     "Jika Kosong Maka Return String Kosong",
				expected: "",
				request:  []int{},
			},

			{
				name:     "Jika Ada Satu Number Maka Return String GanjilGenap Nomer Teresbut Saja",
				expected: "ganjil",
				request:  []int{1},
			},

			{
				name:     "Jika Ada Lebih Dari Number Maka Return String GanjilGenap Ada , Setelah Nomer",
				expected: "ganjil, genap, ganjil, genap, ganjil",
				request:  []int{1, 2, 3, 4, 5},
			},
		}

	assert := assert.New(t)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CekGanjilGenap(test.request...)
			assert.Equal(result, test.expected)
		})
	}
}

func BenchmarkCekGanjilGenap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CekGanjilGenap()
	}
}
