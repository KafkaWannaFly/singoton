package test

import (
	"os"
	"testing"

	"github.com/KafkaWannaFly/singoton"
)

func BenchmarkGetBigStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = singoton.Get[LargeStruct0]()
		_, _ = singoton.Get[LargeStruct1]()
		_, _ = singoton.Get[LargeStruct2]()
		_, _ = singoton.Get[LargeStruct3]()
		_, _ = singoton.Get[LargeStruct4]()
		_, _ = singoton.Get[LargeStruct5]()
		_, _ = singoton.Get[LargeStruct6]()
		_, _ = singoton.Get[LargeStruct7]()
		_, _ = singoton.Get[LargeStruct8]()
		_, _ = singoton.Get[LargeStruct9]()
		_, _ = singoton.Get[LargeStruct10]()
	}
}

func TestMain(m *testing.M) {
	setUp()

	os.Exit(m.Run())
}

func setUp() {
	singoton.Register(MiniStruct0{})
	singoton.Register(MiniStruct1{})
	singoton.Register(MiniStruct2{})
	singoton.Register(MiniStruct3{})
	singoton.Register(MiniStruct4{})
	singoton.Register(MiniStruct5{})
	singoton.Register(MiniStruct6{})
	singoton.Register(MiniStruct7{})
	singoton.Register(MiniStruct8{})
	singoton.Register(MiniStruct9{})

	singoton.Register(MediumStruct0{})
	singoton.Register(MediumStruct1{})

	singoton.Register(LargeStruct0{
		Data0: "Data0",
		Data1: "Data1",
		Data2: "Data2",
		Data3: "Data3",
		Data4: "Data4",
		Data5: "Data5",
		Data6: "Data6",
		Data7: "Data7",
		Data8: "Data8",
		Data9: "Data9",
	})

	singoton.Register(LargeStruct1{
		Data0: "Data0",
		Data1: "Data1",
		Data2: "Data2",
		Data3: "Data3",
		Data4: "Data4",
		Data5: "Data5",
		Data6: "Data6",
		Data7: "Data7",
		Data8: "Data8",
		Data9: "Data9",
	})

	singoton.Register(LargeStruct2{
		Data0: "Data0",
		Data1: "Data1",
		Data2: "Data2",
		Data3: "Data3",
		Data4: "Data4",
		Data5: "Data5",
		Data6: "Data6",
		Data7: "Data7",
		Data8: "Data8",
		Data9: "Data9",
	})

	singoton.Register(LargeStruct3{
		Data0: "Data0",
		Data1: "Data1",
		Data2: "Data2",
		Data3: "Data3",
		Data4: "Data4",
		Data5: "Data5",
		Data6: "Data6",
		Data7: "Data7",
		Data8: "Data8",
		Data9: "Data9",
	})

	singoton.Register(LargeStruct4{
		Data0: "Data0",
		Data1: "Data1",
		Data2: "Data2",
		Data3: "Data3",
		Data4: "Data4",
		Data5: "Data5",
		Data6: "Data6",
		Data7: "Data7",
		Data8: "Data8",
		Data9: "Data9",
	})

	singoton.Register(LargeStruct5{
		Data0: "Data0",
		Data1: "Data1",
		Data2: "Data2",
		Data3: "Data3",
		Data4: "Data4",
		Data5: "Data5",
		Data6: "Data6",
		Data7: "Data7",
		Data8: "Data8",
		Data9: "Data9",
	})

	singoton.Register(LargeStruct6{
		Data0: "Data0",
		Data1: "Data1",
		Data2: "Data2",
		Data3: "Data3",
		Data4: "Data4",
		Data5: "Data5",
		Data6: "Data6",
		Data7: "Data7",
		Data8: "Data8",
		Data9: "Data9",
	})

	singoton.Register(LargeStruct7{
		Data0: "Data0",
		Data1: "Data1",
		Data2: "Data2",
		Data3: "Data3",
		Data4: "Data4",
		Data5: "Data5",
		Data6: "Data6",
		Data7: "Data7",
		Data8: "Data8",
		Data9: "Data9",
	})

	singoton.Register(LargeStruct8{
		Data0: "Data0",
		Data1: "Data1",
		Data2: "Data2",
		Data3: "Data3",
		Data4: "Data4",
		Data5: "Data5",
		Data6: "Data6",
		Data7: "Data7",
		Data8: "Data8",
		Data9: "Data9",
	})

	singoton.Register(LargeStruct9{
		Data0: "Data0",
		Data1: "Data1",
		Data2: "Data2",
		Data3: "Data3",
		Data4: "Data4",
		Data5: "Data5",
		Data6: "Data6",
		Data7: "Data7",
		Data8: "Data8",
		Data9: "Data9",
	})

	singoton.Register(LargeStruct10{
		Data0: "Data0",
		Data1: "Data1",
		Data2: "Data2",
		Data3: "Data3",
		Data4: "Data4",
		Data5: "Data5",
		Data6: "Data6",
		Data7: "Data7",
		Data8: "Data8",
		Data9: "Data9",
	})
}
