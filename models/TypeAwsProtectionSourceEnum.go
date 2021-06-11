// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeAwsProtectionSourceEnum enum
 */
type TypeAwsProtectionSourceEnum int

/**
 * Value collection for TypeAwsProtectionSourceEnum enum
 */
const (
    TypeAwsProtectionSource_KIAMUSER TypeAwsProtectionSourceEnum = 1 + iota
    TypeAwsProtectionSource_KREGION
    TypeAwsProtectionSource_KAVAILABILITYZONE
    TypeAwsProtectionSource_KEC2INSTANCE
    TypeAwsProtectionSource_KVPC
    TypeAwsProtectionSource_KSUBNET
    TypeAwsProtectionSource_KNETWORKSECURITYGROUP
    TypeAwsProtectionSource_KINSTANCETYPE
    TypeAwsProtectionSource_KKEYPAIR
)

func (r TypeAwsProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeAwsProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeAwsProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeAwsProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeAwsProtectionSourceEnum to its string representation
 */
func TypeAwsProtectionSourceEnumToValue(typeAwsProtectionSourceEnum TypeAwsProtectionSourceEnum) string {
    switch typeAwsProtectionSourceEnum {
        case TypeAwsProtectionSource_KIAMUSER:
    		return "kIAMUser"
        case TypeAwsProtectionSource_KREGION:
    		return "kRegion"
        case TypeAwsProtectionSource_KAVAILABILITYZONE:
    		return "kAvailabilityZone"
        case TypeAwsProtectionSource_KEC2INSTANCE:
    		return "kEC2Instance"
        case TypeAwsProtectionSource_KVPC:
    		return "kVPC"
        case TypeAwsProtectionSource_KSUBNET:
    		return "kSubnet"
        case TypeAwsProtectionSource_KNETWORKSECURITYGROUP:
    		return "kNetworkSecurityGroup"
        case TypeAwsProtectionSource_KINSTANCETYPE:
    		return "kInstanceType"
        case TypeAwsProtectionSource_KKEYPAIR:
    		return "kKeyPair"
        default:
        	return "kIAMUser"
    }
}

/**
 * Converts TypeAwsProtectionSourceEnum Array to its string Array representation
*/
func TypeAwsProtectionSourceEnumArrayToValue(typeAwsProtectionSourceEnum []TypeAwsProtectionSourceEnum) []string {
    convArray := make([]string,len( typeAwsProtectionSourceEnum))
    for i:=0; i<len(typeAwsProtectionSourceEnum);i++ {
        convArray[i] = TypeAwsProtectionSourceEnumToValue(typeAwsProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeAwsProtectionSourceEnumFromValue(value string) TypeAwsProtectionSourceEnum {
    switch value {
        case "kIAMUser":
            return TypeAwsProtectionSource_KIAMUSER
        case "kRegion":
            return TypeAwsProtectionSource_KREGION
        case "kAvailabilityZone":
            return TypeAwsProtectionSource_KAVAILABILITYZONE
        case "kEC2Instance":
            return TypeAwsProtectionSource_KEC2INSTANCE
        case "kVPC":
            return TypeAwsProtectionSource_KVPC
        case "kSubnet":
            return TypeAwsProtectionSource_KSUBNET
        case "kNetworkSecurityGroup":
            return TypeAwsProtectionSource_KNETWORKSECURITYGROUP
        case "kInstanceType":
            return TypeAwsProtectionSource_KINSTANCETYPE
        case "kKeyPair":
            return TypeAwsProtectionSource_KKEYPAIR
        default:
            return TypeAwsProtectionSource_KIAMUSER
    }
}
