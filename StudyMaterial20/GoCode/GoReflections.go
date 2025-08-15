
package main

import (
	"reflect"
	"strconv"
	"fmt"
)

//____________________________________________________________
//____________________________________________________________

func Any( value interface{} ) string {
	return formatValue( reflect.ValueOf(value) )
}

func formatValue( v reflect.Value ) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "INVALID"
	
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "Number " + strconv.FormatInt( v.Int(), 10 )

    case reflect.Uint, reflect.Uint8, reflect.Uint16,
        reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return "Unsigned Number " + strconv.FormatUint(v.Uint(), 10)

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

//____________________________________________________________

func playWithIntData() {
	var something int64 	= 100;
	var something1 int32 	= 111;

	var something2 float64 	= 100.90;
	var something3 float32 	= 111.99;

	fmt.Println("Value: ", Any( something ) )
	fmt.Println("Value: ", Any( something1 ) )
	fmt.Println("Value: ", Any( something2 ) )
	fmt.Println("Value: ", Any( something3 ) )
}

//____________________________________________________________
//____________________________________________________________
//____________________________________________________________
//____________________________________________________________

func main() {
	fmt.Println("\nFunction : playWithIntData")
	playWithIntData()

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}

