package models

import(
    "encoding/json"
)

/**
 * Type definition for ModeEnum enum
 */
type ModeEnum int

/**
 * Value collection for ModeEnum enum
 */
const (
    Mode_KWHITELIST ModeEnum = 1 + iota
    Mode_KBLACKLIST
)

func (r ModeEnum) MarshalJSON() ([]byte, error) { 
    s := ModeEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *ModeEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  ModeEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts ModeEnum to its string representation
 */
func ModeEnumToValue(modeEnum ModeEnum) string {
    switch modeEnum {
        case Mode_KWHITELIST:
    		return "kWhitelist"		
        case Mode_KBLACKLIST:
    		return "kBlacklist"		
        default:
        	return "kWhitelist"
    }
}

/**
 * Converts ModeEnum Array to its string Array representation
*/
func ModeEnumArrayToValue(modeEnum []ModeEnum) []string {
    convArray := make([]string,len( modeEnum))
    for i:=0; i<len(modeEnum);i++ {
        convArray[i] = ModeEnumToValue(modeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ModeEnumFromValue(value string) ModeEnum {
    switch value {
        case "kWhitelist":
            return Mode_KWHITELIST
        case "kBlacklist":
            return Mode_KBLACKLIST
        default:
            return Mode_KWHITELIST
    }
}
