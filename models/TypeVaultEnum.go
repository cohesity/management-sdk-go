// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeVaultEnum enum
 */
type TypeVaultEnum int

/**
 * Value collection for TypeVaultEnum enum
 */
const (
    TypeVault_KNEARLINE TypeVaultEnum = 1 + iota
    TypeVault_KCOLDLINE
    TypeVault_KGLACIER
    TypeVault_KS3
    TypeVault_KAZURESTANDARD
    TypeVault_KS3COMPATIBLE
    TypeVault_KQSTARTAPE
    TypeVault_KGOOGLESTANDARD
    TypeVault_KGOOGLEDRA
    TypeVault_KAWSGOVCLOUD
    TypeVault_KNAS
    TypeVault_KAZUREGOVCLOUD
)

func (r TypeVaultEnum) MarshalJSON() ([]byte, error) {
    s := TypeVaultEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeVaultEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeVaultEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeVaultEnum to its string representation
 */
func TypeVaultEnumToValue(typeVaultEnum TypeVaultEnum) string {
    switch typeVaultEnum {
        case TypeVault_KNEARLINE:
    		return "kNearline"
        case TypeVault_KCOLDLINE:
    		return "kColdline"
        case TypeVault_KGLACIER:
    		return "kGlacier"
        case TypeVault_KS3:
    		return "kS3"
        case TypeVault_KAZURESTANDARD:
    		return "kAzureStandard"
        case TypeVault_KS3COMPATIBLE:
    		return "kS3Compatible"
        case TypeVault_KQSTARTAPE:
    		return "kQStarTape"
        case TypeVault_KGOOGLESTANDARD:
    		return "kGoogleStandard"
        case TypeVault_KGOOGLEDRA:
    		return "kGoogleDRA"
        case TypeVault_KAWSGOVCLOUD:
    		return "kAWSGovCloud"
        case TypeVault_KNAS:
    		return "kNAS"
        case TypeVault_KAZUREGOVCLOUD:
    		return "kAzureGovCloud"
        default:
        	return "kNearline"
    }
}

/**
 * Converts TypeVaultEnum Array to its string Array representation
*/
func TypeVaultEnumArrayToValue(typeVaultEnum []TypeVaultEnum) []string {
    convArray := make([]string,len( typeVaultEnum))
    for i:=0; i<len(typeVaultEnum);i++ {
        convArray[i] = TypeVaultEnumToValue(typeVaultEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeVaultEnumFromValue(value string) TypeVaultEnum {
    switch value {
        case "kNearline":
            return TypeVault_KNEARLINE
        case "kColdline":
            return TypeVault_KCOLDLINE
        case "kGlacier":
            return TypeVault_KGLACIER
        case "kS3":
            return TypeVault_KS3
        case "kAzureStandard":
            return TypeVault_KAZURESTANDARD
        case "kS3Compatible":
            return TypeVault_KS3COMPATIBLE
        case "kQStarTape":
            return TypeVault_KQSTARTAPE
        case "kGoogleStandard":
            return TypeVault_KGOOGLESTANDARD
        case "kGoogleDRA":
            return TypeVault_KGOOGLEDRA
        case "kAWSGovCloud":
            return TypeVault_KAWSGOVCLOUD
        case "kNAS":
            return TypeVault_KNAS
        case "kAzureGovCloud":
            return TypeVault_KAZUREGOVCLOUD
        default:
            return TypeVault_KNEARLINE
    }
}
