package models

import(
    "encoding/json"
)

/**
 * Type definition for PeriodicityEnum enum
 */
type PeriodicityEnum int

/**
 * Value collection for PeriodicityEnum enum
 */
const (
    Periodicity_KEVERY PeriodicityEnum = 1 + iota
    Periodicity_KHOUR
    Periodicity_KDAY
    Periodicity_KWEEK
    Periodicity_KMONTH
    Periodicity_KYEAR
)

func (r PeriodicityEnum) MarshalJSON() ([]byte, error) { 
    s := PeriodicityEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *PeriodicityEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  PeriodicityEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts PeriodicityEnum to its string representation
 */
func PeriodicityEnumToValue(periodicityEnum PeriodicityEnum) string {
    switch periodicityEnum {
        case Periodicity_KEVERY:
    		return "kEvery"		
        case Periodicity_KHOUR:
    		return "kHour"		
        case Periodicity_KDAY:
    		return "kDay"		
        case Periodicity_KWEEK:
    		return "kWeek"		
        case Periodicity_KMONTH:
    		return "kMonth"		
        case Periodicity_KYEAR:
    		return "kYear"		
        default:
        	return "kEvery"
    }
}

/**
 * Converts PeriodicityEnum Array to its string Array representation
*/
func PeriodicityEnumArrayToValue(periodicityEnum []PeriodicityEnum) []string {
    convArray := make([]string,len( periodicityEnum))
    for i:=0; i<len(periodicityEnum);i++ {
        convArray[i] = PeriodicityEnumToValue(periodicityEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PeriodicityEnumFromValue(value string) PeriodicityEnum {
    switch value {
        case "kEvery":
            return Periodicity_KEVERY
        case "kHour":
            return Periodicity_KHOUR
        case "kDay":
            return Periodicity_KDAY
        case "kWeek":
            return Periodicity_KWEEK
        case "kMonth":
            return Periodicity_KMONTH
        case "kYear":
            return Periodicity_KYEAR
        default:
            return Periodicity_KEVERY
    }
}
