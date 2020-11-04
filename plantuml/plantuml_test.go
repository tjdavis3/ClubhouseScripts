package plantuml

import (
	"reflect"
	"testing"
)

func Test_reEncodeB64(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name       string
		args       args
		wantRetval []byte
	}{
		{
			name:       "Test Full Charcter Set",
			args:       args{input: []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")},
			wantRetval: []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"),
		},
		{
			name:       "Test python lib results",
			args:       args{input: []byte("cyguSSwqKc3NUeByzMlMTlXQtVNwyk+yUnAsLclIzSvJTE4syczPUwhKLSxNLS5R4AJKKugCFYFVY1FWXJCfV5yqwOWQmpcCNBYA")},
			wantRetval: []byte("SoWkIImgAStDKU1opCbCJbNGjLDmoa-oKd0iBSb8pIl9J4uioSpFKmXABInDBIvHu09AAkW25O5LOr5MN92VLvogmEMGcfS2D1O0"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetval := reEncodeB64(tt.args.input); !reflect.DeepEqual(gotRetval, tt.wantRetval) {
				t.Errorf("reEncodeB64() = %v, want %v", gotRetval, tt.wantRetval)
			}
		})
	}
}

func TestEncodeDiagram(t *testing.T) {
	type args struct {
		diagram []byte
		level   int
	}
	tests := []struct {
		name               string
		args               args
		wantEncodedDiagram string
		wantErr            bool
	}{
		{
			name: "Test PlantUML Uncompressed Encoding Example",
			args: args{diagram: []byte("@startuml\n" +
				"Alice -> Bob: Authentication Request\n" +
				"Bob --> Alice: Authentication Response\n" +
				"@enduml"),
				level: 0,
			},
			// wantEncodedDiagram: "Syp9J4vLqBLJSCfFib9mB2t9ICqhoKnEBCdCprC8IYqiJIqkuGBAAUW2rO0LOr5LN92VLvpA1G00",
			wantEncodedDiagram: "~hG7DqON9qTMri2a5iQMDb82q-849lOZeWGNLqQ6LkT6bZONHfRsuWKcLnTMLpT0f2Rs8WBIq-845iQMDbEY11TNHePMvqQMDXT6blRY1IPNDmRsvpPGf0PMvaTMri",
			wantErr:            false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEncodedDiagram, err := EncodeDiagram(tt.args.diagram, tt.args.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeDiagram() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotEncodedDiagram != tt.wantEncodedDiagram {
				t.Errorf("EncodeDiagram() = %v, want %v", gotEncodedDiagram, tt.wantEncodedDiagram)
			}
		})
	}
}
