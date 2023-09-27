// trunk-ignore-all(golangci-lint)
package test

type BaseStruct struct {
	Data0 string
	Data1 string
	Data2 string
	Data3 string
	Data4 string
	Data5 string
	Data6 string
	Data7 string
	Data8 string
	Data9 string

	Number0 int
	Number1 int
	Number2 int
	Number3 int
	Number4 int
	Number5 int
	Number6 int
	Number7 int
	Number8 int
	Number9 int
}

type MiniStruct0 struct {
	BaseStruct
	name0    string
	sirName0 string
	age0     int
	weight0  int
	height0  int
}
type MiniStruct1 struct {
	BaseStruct
	name1    string
	sirName1 string
	age1     int
	weight1  int
	height1  int
}
type MiniStruct2 struct {
	BaseStruct
	name2    string
	sirName2 string
	age2     int
	weight2  int
	height2  int
}
type MiniStruct3 struct {
	BaseStruct
	name3    string
	sirName3 string
	age3     int
	weight3  int
	height3  int
}
type MiniStruct4 struct {
	BaseStruct
	name4    string
	sirName4 string
	age4     int
	weight4  int
	height4  int
}
type MiniStruct5 struct {
	BaseStruct
	name5    string
	sirName5 string
	age5     int
	weight5  int
	height5  int
}
type MiniStruct6 struct {
	BaseStruct
	name6    string
	sirName6 string
	age6     int
	weight6  int
	height6  int
}
type MiniStruct7 struct {
	BaseStruct
	name7    string
	sirName7 string
	age7     int
	weight7  int
	height7  int
}
type MiniStruct8 struct {
	BaseStruct
	name8    string
	sirName8 string
	age8     int
	weight8  int
	height8  int
}
type MiniStruct9 struct {
	BaseStruct
	name9    string
	sirName9 string
	age9     int
	weight9  int
	height9  int
}

type MediumStruct0 struct {
	mini0 MiniStruct0
	mini1 MiniStruct1
	mini2 MiniStruct2
	mini3 MiniStruct3
	mini4 MiniStruct4
}

type MediumStruct1 struct {
	mini5 MiniStruct5
	mini6 MiniStruct6
	mini7 MiniStruct7
	mini8 MiniStruct8
	mini9 MiniStruct9
}

type LargeStruct0 struct {
	medium0 MediumStruct0
	medium1 MediumStruct1

	Mini0 MiniStruct0
	Mini1 MiniStruct1
	Mini2 MiniStruct2
	Mini3 MiniStruct3
	Mini4 MiniStruct4
	Mini5 MiniStruct5
	Mini6 MiniStruct6
	Mini7 MiniStruct7
	Mini8 MiniStruct8
	Mini9 MiniStruct9

	Data0 string
	Data1 string
	Data2 string
	Data3 string
	Data4 string
	Data5 string
	Data6 string
	Data7 string
	Data8 string
	Data9 string

	Number0 int
	Number1 int
	Number2 int
	Number3 int
	Number4 int
	Number5 int
	Number6 int
	Number7 int
	Number8 int
	Number9 int
}

type LargeStruct1 struct {
	medium0 MediumStruct0
	medium1 MediumStruct1

	Mini0 MiniStruct0
	Mini1 MiniStruct1
	Mini2 MiniStruct2
	Mini3 MiniStruct3
	Mini4 MiniStruct4
	Mini5 MiniStruct5
	Mini6 MiniStruct6
	Mini7 MiniStruct7
	Mini8 MiniStruct8
	Mini9 MiniStruct9

	Data0 string
	Data1 string
	Data2 string
	Data3 string
	Data4 string
	Data5 string
	Data6 string
	Data7 string
	Data8 string
	Data9 string

	Number0 int
	Number1 int
	Number2 int
	Number3 int
	Number4 int
	Number5 int
	Number6 int
	Number7 int
	Number8 int
	Number9 int
}

type LargeStruct2 struct {
	medium0 MediumStruct0
	medium1 MediumStruct1

	Mini0 MiniStruct0
	Mini1 MiniStruct1
	Mini2 MiniStruct2
	Mini3 MiniStruct3
	Mini4 MiniStruct4
	Mini5 MiniStruct5
	Mini6 MiniStruct6
	Mini7 MiniStruct7
	Mini8 MiniStruct8
	Mini9 MiniStruct9

	Data0 string
	Data1 string
	Data2 string
	Data3 string
	Data4 string
	Data5 string
	Data6 string
	Data7 string
	Data8 string
	Data9 string

	Number0 int
	Number1 int
	Number2 int
	Number3 int
	Number4 int
	Number5 int
	Number6 int
	Number7 int
	Number8 int
	Number9 int
}

type LargeStruct3 struct {
	medium0 MediumStruct0
	medium1 MediumStruct1

	Mini0 MiniStruct0
	Mini1 MiniStruct1
	Mini2 MiniStruct2
	Mini3 MiniStruct3
	Mini4 MiniStruct4
	Mini5 MiniStruct5
	Mini6 MiniStruct6
	Mini7 MiniStruct7
	Mini8 MiniStruct8
	Mini9 MiniStruct9

	Data0 string
	Data1 string
	Data2 string
	Data3 string
	Data4 string
	Data5 string
	Data6 string
	Data7 string
	Data8 string
	Data9 string

	Number0 int
	Number1 int
	Number2 int
	Number3 int
	Number4 int
	Number5 int
	Number6 int
	Number7 int
	Number8 int
	Number9 int
}

type LargeStruct4 struct {
	medium0 MediumStruct0
	medium1 MediumStruct1

	Mini0 MiniStruct0
	Mini1 MiniStruct1
	Mini2 MiniStruct2
	Mini3 MiniStruct3
	Mini4 MiniStruct4
	Mini5 MiniStruct5
	Mini6 MiniStruct6
	Mini7 MiniStruct7
	Mini8 MiniStruct8
	Mini9 MiniStruct9

	Data0 string
	Data1 string
	Data2 string
	Data3 string
	Data4 string
	Data5 string
	Data6 string
	Data7 string
	Data8 string
	Data9 string

	Number0 int
	Number1 int
	Number2 int
	Number3 int
	Number4 int
	Number5 int
	Number6 int
	Number7 int
	Number8 int
	Number9 int
}

type LargeStruct5 struct {
	medium0 MediumStruct0
	medium1 MediumStruct1

	Mini0 MiniStruct0
	Mini1 MiniStruct1
	Mini2 MiniStruct2
	Mini3 MiniStruct3
	Mini4 MiniStruct4
	Mini5 MiniStruct5
	Mini6 MiniStruct6
	Mini7 MiniStruct7
	Mini8 MiniStruct8
	Mini9 MiniStruct9

	Data0 string
	Data1 string
	Data2 string
	Data3 string
	Data4 string
	Data5 string
	Data6 string
	Data7 string
	Data8 string
	Data9 string

	Number0 int
	Number1 int
	Number2 int
	Number3 int
	Number4 int
	Number5 int
	Number6 int
	Number7 int
	Number8 int
	Number9 int
}

type LargeStruct6 struct {
	medium0 MediumStruct0
	medium1 MediumStruct1

	Mini0 MiniStruct0
	Mini1 MiniStruct1
	Mini2 MiniStruct2
	Mini3 MiniStruct3
	Mini4 MiniStruct4
	Mini5 MiniStruct5
	Mini6 MiniStruct6
	Mini7 MiniStruct7
	Mini8 MiniStruct8
	Mini9 MiniStruct9

	Data0 string
	Data1 string
	Data2 string
	Data3 string
	Data4 string
	Data5 string
	Data6 string
	Data7 string
	Data8 string
	Data9 string

	Number0 int
	Number1 int
	Number2 int
	Number3 int
	Number4 int
	Number5 int
	Number6 int
	Number7 int
	Number8 int
	Number9 int
}

type LargeStruct7 struct {
	medium0 MediumStruct0
	medium1 MediumStruct1

	Mini0 MiniStruct0
	Mini1 MiniStruct1
	Mini2 MiniStruct2
	Mini3 MiniStruct3
	Mini4 MiniStruct4
	Mini5 MiniStruct5
	Mini6 MiniStruct6
	Mini7 MiniStruct7
	Mini8 MiniStruct8
	Mini9 MiniStruct9

	Data0 string
	Data1 string
	Data2 string
	Data3 string
	Data4 string
	Data5 string
	Data6 string
	Data7 string
	Data8 string
	Data9 string

	Number0 int
	Number1 int
	Number2 int
	Number3 int
	Number4 int
	Number5 int
	Number6 int
	Number7 int
	Number8 int
	Number9 int
}

type LargeStruct8 struct {
	medium0 MediumStruct0
	medium1 MediumStruct1

	Mini0 MiniStruct0
	Mini1 MiniStruct1
	Mini2 MiniStruct2
	Mini3 MiniStruct3
	Mini4 MiniStruct4
	Mini5 MiniStruct5
	Mini6 MiniStruct6
	Mini7 MiniStruct7
	Mini8 MiniStruct8
	Mini9 MiniStruct9

	Data0 string
	Data1 string
	Data2 string
	Data3 string
	Data4 string
	Data5 string
	Data6 string
	Data7 string
	Data8 string
	Data9 string

	Number0 int
	Number1 int
	Number2 int
	Number3 int
	Number4 int
	Number5 int
	Number6 int
	Number7 int
	Number8 int
	Number9 int
}

type LargeStruct9 struct {
	medium0 MediumStruct0
	medium1 MediumStruct1

	Mini0 MiniStruct0
	Mini1 MiniStruct1
	Mini2 MiniStruct2
	Mini3 MiniStruct3
	Mini4 MiniStruct4
	Mini5 MiniStruct5
	Mini6 MiniStruct6
	Mini7 MiniStruct7
	Mini8 MiniStruct8
	Mini9 MiniStruct9

	Data0 string
	Data1 string
	Data2 string
	Data3 string
	Data4 string
	Data5 string
	Data6 string
	Data7 string
	Data8 string
	Data9 string

	Number0 int
	Number1 int
	Number2 int
	Number3 int
	Number4 int
	Number5 int
	Number6 int
	Number7 int
	Number8 int
	Number9 int
}

type LargeStruct10 struct {
	medium0 MediumStruct0
	medium1 MediumStruct1

	Mini0 MiniStruct0
	Mini1 MiniStruct1
	Mini2 MiniStruct2
	Mini3 MiniStruct3
	Mini4 MiniStruct4
	Mini5 MiniStruct5
	Mini6 MiniStruct6
	Mini7 MiniStruct7
	Mini8 MiniStruct8
	Mini9 MiniStruct9

	Data0 string
	Data1 string
	Data2 string
	Data3 string
	Data4 string
	Data5 string
	Data6 string
	Data7 string
	Data8 string
	Data9 string

	Number0 int
	Number1 int
	Number2 int
	Number3 int
	Number4 int
	Number5 int
	Number6 int
	Number7 int
	Number8 int
	Number9 int
}
