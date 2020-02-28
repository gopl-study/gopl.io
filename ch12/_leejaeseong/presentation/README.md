# 리플렉션

**리플렉션**이 뭔데?   
**컴파일**시 알 수 없는 타입에 대해
- 변수를 **갱신**.
- 값을 **조사**.
- 메소드를 **호출**.
- 표현 방식에 따른 고유 작업(**태깅등**)



**태그**는 대표적으로 struct타입을 정의 후 필드 뒤에 붙는 `json:"{fieldname}"` 이 이에 해당함.

```golang
type SomeType struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}
```

## 왜 리플렉션을 사용하는가?

밑에 3가지를 충족하지 못했을 경우 
- 공용(빌트인) 및 해당 패키지의 특정 **인터페이스**를 충족 X
- 알려진 **표현 방식** X
- 함수 설계 시에는 해당 **타입 존재** X 

## 그래서 리플렉션이 왜 필요함 두번째? // Sprint 예제

해당 Sprint 예제는  
타입 스위치문으로 시작하여 **String 메소드**가 정의 되어있을 경우 String 메소드를 호출  
이후 값의 **기본 타입**을 이용하여 스위치 케이스를 추가하고 적절한 포맷팅 작업을 수행


하지만 byte, short, int8(16,32,64), uint, uint8(16,32,64), float32(64), slice, channel, func, map, pointer, slice, struct 등등
다양한 케이스를 다 추가하기엔 무한정 많기 때문에 불가능에 가까움. 

```golang
package main

import (
	"strconv"
)

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}

	switch x := x.(type) {
	case stringer: return x.String()
	case string: return x
	case int: return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return "???"
	}
}

type noStringerType struct {
	Id int
	Name string
}

type StringerType struct {
	Id int
	Name string
}

func (s StringerType) String() string {
	return strconv.Itoa(s.Id) + " / " + s.Name
}

func main() {
	println(Sprint("string")) // string
	println(Sprint(123456)) // 123456
	println(Sprint(true)) // true
	println(Sprint(false)) // false
	println(Sprint(StringerType{
		Id:   135,
		Name: "Stringer Type",
	})) // 135 / Stringer Type

	println(Sprint(123.0)) // ???
	println(Sprint(int64(1234))) // ???
	println(Sprint(int32(12))) // ???
	println(Sprint(noStringerType{
		Id:   321,
		Name: "No Stringer Type",
	})) // ???
}
```

위 코드는 [이곳 Go Playground](https://play.golang.org/p/afbE-odcdTf)로 들어가면 있다.
## 그래서 배운다 리플렉션(reflect) 

리플렉션은 reflect 패키지에 의해 제공해주며   
`reflect.Type`와 `reflect.Value` 두가지를 주의 깊게 알아야한다.


### reflect.Type
고언어의 타입을 나타냄  
항상 동적 타입을 반환  

#### 예시 
```golang
package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String()) // int
	fmt.Println(t) // int
  
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // *os.File
}
```

위 코드는 [이곳 Go Playground](https://play.golang.org/p/bDLEY_jqhGT)로 들어가면 있다.

해당 예시에서 눈치 챈 사람이 있듯이 **reflect.Type**은 **fmt.Stringer**를 충족하기때문에 `fmt.Println(t)`를 하더라도 같은 결과 인것을 볼 수 있음.

추가적으로 **fmt.Printf**는 내부적으로 reflect.TypeOf를 사용하는 포맷이 있다 표기법은 **%T**이다.

```golang
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%T\n", 3) // int
	fmt.Printf("%T\n", os.Stdout) // *os.File
}
```

### reflect.Value
모든 타입의 값도 저장 가능하며  
**fmt.Stringer**를 충족하지만 별도로 **string**타입이지 않으면 **String()** 메소드는 타입을 보여준다.

```golang
package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(3)
	fmt.Println(v) // 3
	fmt.Printf("%v\n", v) // 3
	fmt.Println(v.String()) <int Value>
  
	v = reflect.ValueOf("str str")
	fmt.Println(v) // str str
	fmt.Printf("%v\n", v) // str str
	fmt.Println(v.String()) // str str
}
```

위 코드는 [이곳 Go Playground](https://play.golang.org/p/BXU6rpmD06b)로 들어가면 있다.


**reflect.Value.Type** 메서드를 사용하여 **reflect.Type**을 받을수도 있다.
```golang
package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(3)
	fmt.Println(v.Type())
}
```

위 코드는 [이곳 Go Playground](https://play.golang.org/p/Zdew0OpYHPO)로 들어가면 있다.

reflect.ValueOf의 역연산은 **reflect.Value.Interface** 메소드이다.

```golang
package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(3)
	x := v.Interface()
	i := x.(int)
	fmt.Printf("%d\n", x) // 3
	fmt.Printf("%d\n", i) // 3
	fmt.Printf("%T\n", x) // int
	fmt.Printf("%T\n", i) // int
}
```

위 코드는 [이곳 Go Playground](https://play.golang.org/p/UXYcweWYVSm)로 들어가면 있다.

#### reflect.Value 정리
**reflect.Value**와 **interface{}** 는 둘다 임의의 값을 가질 수 있다.   
단 **interface{}** 는 타입을 알려면 타입을 강제화 하여 스위치 케이스를 태우는 방법 밖에 없지만   
**reflect.Value** 는 타입과 무관하게 내용을 조사가 가능하다.

## 결론 // Sprint 업그레이드

reflect.Value의 Kind 메소드를 사용하여 스위치 케이스를 태워 포맷팅 하는 예제

```golang
package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func main() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(Any(x)) // 1
	fmt.Println(Any(d)) // 1
	fmt.Println(Any([]int64{x})) // []int64 0x8202b87b0
	fmt.Println(Any([]time.Duration{d})) // []time.Duration 0x8202b87e0
}
```

위 코드는 [이곳 Go Playground](https://play.golang.org/p/BgGFJEePFyF)로 들어가면 있다.
