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

func Test_getVentLevels(t *testing.T) {
	type args struct {
		coord []pair
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{name: "Works for single-node vent", args: args{[]pair{{0, 0}}}, want: map[int]int{1: 1}},
		{name: "Works for multi-node vent", args: args{[]pair{{0, 0}, {0, 1}}}, want: map[int]int{1: 2}},
		{name: "Works for vent overlap", args: args{[]pair{{0, 0}, {0, 1}, {0, 0}, {1, 0}}}, want: map[int]int{1: 2, 2: 1}},
		{name: "Works for for diagonal top left to bottom right", args: args{[]pair{{0, 0}, {1, 1}, {2, 2}}}, want: map[int]int{1: 3}},
		{name: "Works for for diagonal bottom left to top right", args: args{[]pair{{0, 2}, {1, 1}, {2, 0}}}, want: map[int]int{1: 3}},
		{name: "Works for for diagonal top right to bottom left", args: args{[]pair{{2, 0}, {1, 1}, {0, 2}}}, want: map[int]int{1: 3}},
		{name: "Works for for diagonal top left to bottom right", args: args{[]pair{{2, 2}, {1, 1}, {0, 0}}}, want: map[int]int{1: 3}},
		{name: "Works for for diagonal overlap", args: args{[]pair{{2, 2}, {1, 1}, {0, 0}, {0, 2}, {1, 1}, {2, 0}}}, want: map[int]int{1: 4, 2: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVentLevels(tt.args.coord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getVentLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
