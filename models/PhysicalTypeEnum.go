// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PhysicalTypeEnum enum
 */
type PhysicalTypeEnum int

/**
 * Value collection for PhysicalTypeEnum enum
 */
const (
    PhysicalType_KHOST PhysicalTypeEnum = 1 + iota
    PhysicalType_KWINDOWSCLUSTER
)

func (r PhysicalTypeEnum) MarshalJSON() ([]byte, error) {
    s := PhysicalTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *PhysicalTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PhysicalTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PhysicalTypeEnum to its string representation
 */
func PhysicalTypeEnumToValue(physicalTypeEnum PhysicalTypeEnum) string {
    switch physicalTypeEnum {
        case PhysicalType_KHOST:
    		return "kHost"
        case PhysicalType_KWINDOWSCLUSTER:
    		return "kWindowsCluster"
        default:
        	return "kHost"
    }
}

/**
 * Converts PhysicalTypeEnum Array to its string Array representation
*/
func PhysicalTypeEnumArrayToValue(physicalTypeEnum []PhysicalTypeEnum) []string {
    convArray := make([]string,len( physicalTypeEnum))
    for i:=0; i<len(physicalTypeEnum);i++ {
        convArray[i] = PhysicalTypeEnumToValue(physicalTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PhysicalTypeEnumFromValue(value string) PhysicalTypeEnum {
    switch value {
        case "kHost":
            return PhysicalType_KHOST
        case "kWindowsCluster":
            return PhysicalType_KWINDOWSCLUSTER
        default:
            return PhysicalType_KHOST
    }
}
