package main

import (
	"reflect"
	"testing"
)

func Test_vent_getLevel(t *testing.T) {
	type fields struct {
		level int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Returns . if level equals 0", fields: fields{0}, want: "."},
		{name: "Returns string of level if level not equals 0", fields: fields{4}, want: "4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := vent{
				level: &tt.fields.level,
			}
			got := v.getLevelValue()
			if got != tt.want {
				t.Errorf("vent.getLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vent_increaseLevel(t *testing.T) {
	type fields struct {
		level int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "Inceases level", fields: fields{level: 0}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := vent{
				level: &tt.fields.level,
			}
			v.increaseLevel()
			if *v.level != tt.want {
				t.Errorf("vent.level = %v, want %v", *v.level, tt.want)
			}
		})
	}
}

func Test_convertEntry(t *testing.T) {
	type args struct {
		entry string
	}
	tests := []struct {
		name    string
		args    args
		wantRet []pair
		wantErr bool
	}{
		{name: "Converts on X axis", args: args{"0,0 -> 2,0"}, wantRet: []pair{{0, 0}, {1, 0}, {2, 0}}, wantErr: false},
		{name: "Converts on Y axis", args: args{"0,0 -> 0,2"}, wantRet: []pair{{0, 0}, {0, 1}, {0, 2}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := convertEntry(tt.args.entry)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("convertEntry() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_drawDiagram(t *testing.T) {
	tests := []struct {
		name string
		want [][]vent
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := drawDiagram(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("drawDiagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
