// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeGcpProtectionSourceEnum enum
 */
type TypeGcpProtectionSourceEnum int

/**
 * Value collection for TypeGcpProtectionSourceEnum enum
 */
const (
    TypeGcpProtectionSource_KIAMUSER TypeGcpProtectionSourceEnum = 1 + iota
    TypeGcpProtectionSource_KPROJECT
    TypeGcpProtectionSource_KREGION
    TypeGcpProtectionSource_KAVAILABILITYZONE
    TypeGcpProtectionSource_KVIRTUALMACHINE
    TypeGcpProtectionSource_KVPC
    TypeGcpProtectionSource_KSUBNET
    TypeGcpProtectionSource_KNETWORKSECURITYGROUP
    TypeGcpProtectionSource_KINSTANCETYPE
)

func (r TypeGcpProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeGcpProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeGcpProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeGcpProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeGcpProtectionSourceEnum to its string representation
 */
func TypeGcpProtectionSourceEnumToValue(typeGcpProtectionSourceEnum TypeGcpProtectionSourceEnum) string {
    switch typeGcpProtectionSourceEnum {
        case TypeGcpProtectionSource_KIAMUSER:
    		return "kIAMUser"
        case TypeGcpProtectionSource_KPROJECT:
    		return "kProject"
        case TypeGcpProtectionSource_KREGION:
    		return "kRegion"
        case TypeGcpProtectionSource_KAVAILABILITYZONE:
    		return "kAvailabilityZone"
        case TypeGcpProtectionSource_KVIRTUALMACHINE:
    		return "kVirtualMachine"
        case TypeGcpProtectionSource_KVPC:
    		return "kVPC"
        case TypeGcpProtectionSource_KSUBNET:
    		return "kSubnet"
        case TypeGcpProtectionSource_KNETWORKSECURITYGROUP:
    		return "kNetworkSecurityGroup"
        case TypeGcpProtectionSource_KINSTANCETYPE:
    		return "kInstanceType"
        default:
        	return "kIAMUser"
    }
}

/**
 * Converts TypeGcpProtectionSourceEnum Array to its string Array representation
*/
func TypeGcpProtectionSourceEnumArrayToValue(typeGcpProtectionSourceEnum []TypeGcpProtectionSourceEnum) []string {
    convArray := make([]string,len( typeGcpProtectionSourceEnum))
    for i:=0; i<len(typeGcpProtectionSourceEnum);i++ {
        convArray[i] = TypeGcpProtectionSourceEnumToValue(typeGcpProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeGcpProtectionSourceEnumFromValue(value string) TypeGcpProtectionSourceEnum {
    switch value {
        case "kIAMUser":
            return TypeGcpProtectionSource_KIAMUSER
        case "kProject":
            return TypeGcpProtectionSource_KPROJECT
        case "kRegion":
            return TypeGcpProtectionSource_KREGION
        case "kAvailabilityZone":
            return TypeGcpProtectionSource_KAVAILABILITYZONE
        case "kVirtualMachine":
            return TypeGcpProtectionSource_KVIRTUALMACHINE
        case "kVPC":
            return TypeGcpProtectionSource_KVPC
        case "kSubnet":
            return TypeGcpProtectionSource_KSUBNET
        case "kNetworkSecurityGroup":
            return TypeGcpProtectionSource_KNETWORKSECURITYGROUP
        case "kInstanceType":
            return TypeGcpProtectionSource_KINSTANCETYPE
        default:
            return TypeGcpProtectionSource_KIAMUSER
    }
}
