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
	SetUp()

	os.Exit(m.Run())
}

func SetUp() {
	singoton.Register(singoton.RegisterOptions[MiniStruct0]{})
	singoton.Register(singoton.RegisterOptions[MiniStruct1]{})
	singoton.Register(singoton.RegisterOptions[MiniStruct2]{})
	singoton.Register(singoton.RegisterOptions[MiniStruct3]{})
	singoton.Register(singoton.RegisterOptions[MiniStruct4]{})
	singoton.Register(singoton.RegisterOptions[MiniStruct5]{})
	singoton.Register(singoton.RegisterOptions[MiniStruct6]{})
	singoton.Register(singoton.RegisterOptions[MiniStruct7]{})
	singoton.Register(singoton.RegisterOptions[MiniStruct8]{})
	singoton.Register(singoton.RegisterOptions[MiniStruct9]{})

	singoton.Register(singoton.RegisterOptions[MediumStruct0]{})
	singoton.Register(singoton.RegisterOptions[MediumStruct1]{})

	singoton.Register(singoton.RegisterOptions[LargeStruct0]{
		InitialValue: LargeStruct0{
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
		},
	})

	singoton.Register(singoton.RegisterOptions[LargeStruct1]{
		InitialValue: LargeStruct1{
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
		},
	})

	singoton.Register(singoton.RegisterOptions[LargeStruct2]{
		InitialValue: LargeStruct2{
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
		},
	})

	singoton.Register(singoton.RegisterOptions[LargeStruct3]{
		InitialValue: LargeStruct3{
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
		},
	})

	singoton.Register(singoton.RegisterOptions[LargeStruct4]{
		InitialValue: LargeStruct4{
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
		},
	})

	singoton.Register(singoton.RegisterOptions[LargeStruct5]{
		InitialValue: LargeStruct5{
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
		},
	})

	singoton.Register(singoton.RegisterOptions[LargeStruct6]{
		InitialValue: LargeStruct6{
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
		},
	})

	singoton.Register(singoton.RegisterOptions[LargeStruct7]{
		InitialValue: LargeStruct7{
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
		},
	})

	singoton.Register(singoton.RegisterOptions[LargeStruct8]{
		InitialValue: LargeStruct8{
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
		},
	})

	singoton.Register(singoton.RegisterOptions[LargeStruct9]{
		InitialValue: LargeStruct9{
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
		},
	})

	singoton.Register(singoton.RegisterOptions[LargeStruct10]{
		InitialValue: LargeStruct10{
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
		},
	})
}
