package models

import(
    "encoding/json"
)

/**
 * Type definition for ActionEnum enum
 */
type ActionEnum int

/**
 * Value collection for ActionEnum enum
 */
const (
    Action_KACTIVATE ActionEnum = 1 + iota
    Action_KDEACTIVATE
    Action_KPAUSE
    Action_KRESUME
)

func (r ActionEnum) MarshalJSON() ([]byte, error) { 
    s := ActionEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *ActionEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  ActionEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts ActionEnum to its string representation
 */
func ActionEnumToValue(actionEnum ActionEnum) string {
    switch actionEnum {
        case Action_KACTIVATE:
    		return "kActivate"		
        case Action_KDEACTIVATE:
    		return "kDeactivate"		
        case Action_KPAUSE:
    		return "kPause"		
        case Action_KRESUME:
    		return "kResume"		
        default:
        	return "kActivate"
    }
}

/**
 * Converts ActionEnum Array to its string Array representation
*/
func ActionEnumArrayToValue(actionEnum []ActionEnum) []string {
    convArray := make([]string,len( actionEnum))
    for i:=0; i<len(actionEnum);i++ {
        convArray[i] = ActionEnumToValue(actionEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ActionEnumFromValue(value string) ActionEnum {
    switch value {
        case "kActivate":
            return Action_KACTIVATE
        case "kDeactivate":
            return Action_KDEACTIVATE
        case "kPause":
            return Action_KPAUSE
        case "kResume":
            return Action_KRESUME
        default:
            return Action_KACTIVATE
    }
}
