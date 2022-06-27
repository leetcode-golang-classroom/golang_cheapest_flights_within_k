package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	n := 4
	flights := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{2, 0, 100},
		{1, 3, 600},
		{2, 3, 200},
	}
	src := 0
	dst := 3
	k := 1
	for idx := 0; idx < b.N; idx++ {
		findCheapestPrice(n, flights, src, dst, k)
	}
}
func Test_findCheapestPrice(t *testing.T) {
	type args struct {
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0, dst = 3, k = 1",
			args: args{n: 4,
				flights: [][]int{
					{0, 1, 100},
					{1, 2, 100},
					{2, 0, 100},
					{1, 3, 600},
					{2, 3, 200},
				},
				src: 0,
				dst: 3,
				k:   1,
			},
			want: 700,
		},
		{
			name: "n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1",
			args: args{
				n: 3,
				flights: [][]int{
					{0, 1, 100},
					{1, 2, 100},
					{0, 2, 500},
				},
				src: 0,
				dst: 2,
				k:   1,
			},
			want: 200,
		},
		{
			name: "n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0",
			args: args{
				n: 3,
				flights: [][]int{
					{0, 1, 100},
					{1, 2, 100},
					{0, 2, 500},
				},
				src: 0,
				dst: 2,
				k:   0,
			},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCheapestPrice(tt.args.n, tt.args.flights, tt.args.src, tt.args.dst, tt.args.k); got != tt.want {
				t.Errorf("findCheapestPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
