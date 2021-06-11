// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeHypervProtectionSourceEnum enum
 */
type TypeHypervProtectionSourceEnum int

/**
 * Value collection for TypeHypervProtectionSourceEnum enum
 */
const (
    TypeHypervProtectionSource_KSCVMMSERVER TypeHypervProtectionSourceEnum = 1 + iota
    TypeHypervProtectionSource_KSTANDALONEHOST
    TypeHypervProtectionSource_KSTANDALONECLUSTER
    TypeHypervProtectionSource_KHOSTGROUP
    TypeHypervProtectionSource_KHOST
    TypeHypervProtectionSource_KHOSTCLUSTER
    TypeHypervProtectionSource_KVIRTUALMACHINE
    TypeHypervProtectionSource_KNETWORK
    TypeHypervProtectionSource_KDATASTORE
)

func (r TypeHypervProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeHypervProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeHypervProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeHypervProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeHypervProtectionSourceEnum to its string representation
 */
func TypeHypervProtectionSourceEnumToValue(type_hyperv_ProtectionSourceEnum TypeHypervProtectionSourceEnum) string {
    switch type_hyperv_ProtectionSourceEnum {
        case TypeHypervProtectionSource_KSCVMMSERVER:
    		return "kSCVMMServer"
        case TypeHypervProtectionSource_KSTANDALONEHOST:
    		return "kStandaloneHost"
        case TypeHypervProtectionSource_KSTANDALONECLUSTER:
    		return "kStandaloneCluster"
        case TypeHypervProtectionSource_KHOSTGROUP:
    		return "kHostGroup"
        case TypeHypervProtectionSource_KHOST:
    		return "kHost"
        case TypeHypervProtectionSource_KHOSTCLUSTER:
    		return "kHostCluster"
        case TypeHypervProtectionSource_KVIRTUALMACHINE:
    		return "kVirtualMachine"
        case TypeHypervProtectionSource_KNETWORK:
    		return "kNetwork"
        case TypeHypervProtectionSource_KDATASTORE:
    		return "kDatastore"
        default:
        	return "kSCVMMServer"
    }
}

/**
 * Converts TypeHypervProtectionSourceEnum Array to its string Array representation
*/
func TypeHypervProtectionSourceEnumArrayToValue(type_hyperv_ProtectionSourceEnum []TypeHypervProtectionSourceEnum) []string {
    convArray := make([]string,len( type_hyperv_ProtectionSourceEnum))
    for i:=0; i<len(type_hyperv_ProtectionSourceEnum);i++ {
        convArray[i] = TypeHypervProtectionSourceEnumToValue(type_hyperv_ProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeHypervProtectionSourceEnumFromValue(value string) TypeHypervProtectionSourceEnum {
    switch value {
        case "kSCVMMServer":
            return TypeHypervProtectionSource_KSCVMMSERVER
        case "kStandaloneHost":
            return TypeHypervProtectionSource_KSTANDALONEHOST
        case "kStandaloneCluster":
            return TypeHypervProtectionSource_KSTANDALONECLUSTER
        case "kHostGroup":
            return TypeHypervProtectionSource_KHOSTGROUP
        case "kHost":
            return TypeHypervProtectionSource_KHOST
        case "kHostCluster":
            return TypeHypervProtectionSource_KHOSTCLUSTER
        case "kVirtualMachine":
            return TypeHypervProtectionSource_KVIRTUALMACHINE
        case "kNetwork":
            return TypeHypervProtectionSource_KNETWORK
        case "kDatastore":
            return TypeHypervProtectionSource_KDATASTORE
        default:
            return TypeHypervProtectionSource_KSCVMMSERVER
    }
}
