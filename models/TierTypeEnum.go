package models

import(
    "encoding/json"
)

/**
 * Type definition for TierTypeEnum enum
 */
type TierTypeEnum int

/**
 * Value collection for TierTypeEnum enum
 */
const (
    TierType_KAMAZONS3STANDARD TierTypeEnum = 1 + iota
    TierType_KAMAZONS3STANDARDIA
    TierType_KAMAZONGLACIER
)

func (r TierTypeEnum) MarshalJSON() ([]byte, error) { 
    s := TierTypeEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *TierTypeEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  TierTypeEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts TierTypeEnum to its string representation
 */
func TierTypeEnumToValue(tierTypeEnum TierTypeEnum) string {
    switch tierTypeEnum {
        case TierType_KAMAZONS3STANDARD:
    		return "kAmazonS3Standard"		
        case TierType_KAMAZONS3STANDARDIA:
    		return "kAmazonS3StandardIA"		
        case TierType_KAMAZONGLACIER:
    		return "kAmazonGlacier"		
        default:
        	return "kAmazonS3Standard"
    }
}

/**
 * Converts TierTypeEnum Array to its string Array representation
*/
func TierTypeEnumArrayToValue(tierTypeEnum []TierTypeEnum) []string {
    convArray := make([]string,len( tierTypeEnum))
    for i:=0; i<len(tierTypeEnum);i++ {
        convArray[i] = TierTypeEnumToValue(tierTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TierTypeEnumFromValue(value string) TierTypeEnum {
    switch value {
        case "kAmazonS3Standard":
            return TierType_KAMAZONS3STANDARD
        case "kAmazonS3StandardIA":
            return TierType_KAMAZONS3STANDARDIA
        case "kAmazonGlacier":
            return TierType_KAMAZONGLACIER
        default:
            return TierType_KAMAZONS3STANDARD
    }
}
