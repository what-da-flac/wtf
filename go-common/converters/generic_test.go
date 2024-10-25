package converters

import "testing"

func TestInterfaceToString(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path",
			args: args{
				i: "hello test",
			},
			want: "hello test",
		},
		{
			name: "not a string",
			args: args{
				i: 17.85,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceToString(tt.args.i); got != tt.want {
				t.Errorf("InterfaceToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceToInt(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy path",
			args: args{
				i: 1822,
			},
			want: 1822,
		},
		{
			name: "not int",
			args: args{
				i: "test",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceToInt(tt.args.i); got != tt.want {
				t.Errorf("InterfaceToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceToString1(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceToString(tt.args.i); got != tt.want {
				t.Errorf("InterfaceToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceToBool(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy path",
			args: args{
				i: true,
			},
			want: true,
		},
		{
			name: "false",
			args: args{
				i: false,
			},
			want: false,
		},
		{
			name: "not bool",
			args: args{
				i: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceToBool(tt.args.i); got != tt.want {
				t.Errorf("InterfaceToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
