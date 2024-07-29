package logging

import "testing"

func TestMyLogger_LogDebug(t *testing.T) {
	type fields struct {
		level int
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test LogDebug",
			fields: fields{
				level: 0,
			},
			args: args{
				msg: "Test Debug",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &MyLogger{}
			l.LogDebug().Msg(tt.args.msg)
		})
	}
}
