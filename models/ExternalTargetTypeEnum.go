// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ExternalTargetTypeEnum enum
 */
type ExternalTargetTypeEnum int

/**
 * Value collection for ExternalTargetTypeEnum enum
 */
const (
    ExternalTargetType_KS3COMPATIBLE ExternalTargetTypeEnum = 1 + iota
    ExternalTargetType_KQSTARTAPE
    ExternalTargetType_KAWSGOVCLOUD
    ExternalTargetType_KNAS
    ExternalTargetType_KAZUREGOVCLOUD
    ExternalTargetType_KAZURE
    ExternalTargetType_KGOOGLE
    ExternalTargetType_KAMAZON
    ExternalTargetType_KORACLE
    ExternalTargetType_KAMAZONC2S
)

func (r ExternalTargetTypeEnum) MarshalJSON() ([]byte, error) {
    s := ExternalTargetTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *ExternalTargetTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ExternalTargetTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ExternalTargetTypeEnum to its string representation
 */
func ExternalTargetTypeEnumToValue(externalTargetTypeEnum ExternalTargetTypeEnum) string {
    switch externalTargetTypeEnum {
        case ExternalTargetType_KS3COMPATIBLE:
    		return "kS3Compatible"
        case ExternalTargetType_KQSTARTAPE:
    		return "kQStarTape"
        case ExternalTargetType_KAWSGOVCLOUD:
    		return "kAWSGovCloud"
        case ExternalTargetType_KNAS:
    		return "kNAS"
        case ExternalTargetType_KAZUREGOVCLOUD:
    		return "kAzureGovCloud"
        case ExternalTargetType_KAZURE:
    		return "kAzure"
        case ExternalTargetType_KGOOGLE:
    		return "kGoogle"
        case ExternalTargetType_KAMAZON:
    		return "kAmazon"
        case ExternalTargetType_KORACLE:
    		return "kOracle"
        case ExternalTargetType_KAMAZONC2S:
    		return "kAmazonC2S"
        default:
        	return "kS3Compatible"
    }
}

/**
 * Converts ExternalTargetTypeEnum Array to its string Array representation
*/
func ExternalTargetTypeEnumArrayToValue(externalTargetTypeEnum []ExternalTargetTypeEnum) []string {
    convArray := make([]string,len( externalTargetTypeEnum))
    for i:=0; i<len(externalTargetTypeEnum);i++ {
        convArray[i] = ExternalTargetTypeEnumToValue(externalTargetTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ExternalTargetTypeEnumFromValue(value string) ExternalTargetTypeEnum {
    switch value {
        case "kS3Compatible":
            return ExternalTargetType_KS3COMPATIBLE
        case "kQStarTape":
            return ExternalTargetType_KQSTARTAPE
        case "kAWSGovCloud":
            return ExternalTargetType_KAWSGOVCLOUD
        case "kNAS":
            return ExternalTargetType_KNAS
        case "kAzureGovCloud":
            return ExternalTargetType_KAZUREGOVCLOUD
        case "kAzure":
            return ExternalTargetType_KAZURE
        case "kGoogle":
            return ExternalTargetType_KGOOGLE
        case "kAmazon":
            return ExternalTargetType_KAMAZON
        case "kOracle":
            return ExternalTargetType_KORACLE
        case "kAmazonC2S":
            return ExternalTargetType_KAMAZONC2S
        default:
            return ExternalTargetType_KS3COMPATIBLE
    }
}
