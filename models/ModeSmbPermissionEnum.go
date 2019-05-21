// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ModeSmbPermissionEnum enum
 */
type ModeSmbPermissionEnum int

/**
 * Value collection for ModeSmbPermissionEnum enum
 */
const (
    ModeSmbPermission_KFOLDERSUBFOLDERSANDFILES ModeSmbPermissionEnum = 1 + iota
    ModeSmbPermission_KFOLDERANDSUBFOLDERS
    ModeSmbPermission_KFOLDERANDFILES
    ModeSmbPermission_KFOLDERONLY
    ModeSmbPermission_KSUBFOLDERSANDFILESONLY
    ModeSmbPermission_KSUBFOLDERSONLY
    ModeSmbPermission_KFILESONLY
)

func (r ModeSmbPermissionEnum) MarshalJSON() ([]byte, error) {
    s := ModeSmbPermissionEnumToValue(r)
    return json.Marshal(s)
}

func (r *ModeSmbPermissionEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ModeSmbPermissionEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ModeSmbPermissionEnum to its string representation
 */
func ModeSmbPermissionEnumToValue(modeSmbPermissionEnum ModeSmbPermissionEnum) string {
    switch modeSmbPermissionEnum {
        case ModeSmbPermission_KFOLDERSUBFOLDERSANDFILES:
    		return "kFolderSubFoldersAndFiles"
        case ModeSmbPermission_KFOLDERANDSUBFOLDERS:
    		return "kFolderAndSubFolders"
        case ModeSmbPermission_KFOLDERANDFILES:
    		return "kFolderAndFiles"
        case ModeSmbPermission_KFOLDERONLY:
    		return "kFolderOnly"
        case ModeSmbPermission_KSUBFOLDERSANDFILESONLY:
    		return "kSubFoldersAndFilesOnly"
        case ModeSmbPermission_KSUBFOLDERSONLY:
    		return "kSubFoldersOnly"
        case ModeSmbPermission_KFILESONLY:
    		return "kFilesOnly"
        default:
        	return "kFolderSubFoldersAndFiles"
    }
}

/**
 * Converts ModeSmbPermissionEnum Array to its string Array representation
*/
func ModeSmbPermissionEnumArrayToValue(modeSmbPermissionEnum []ModeSmbPermissionEnum) []string {
    convArray := make([]string,len( modeSmbPermissionEnum))
    for i:=0; i<len(modeSmbPermissionEnum);i++ {
        convArray[i] = ModeSmbPermissionEnumToValue(modeSmbPermissionEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ModeSmbPermissionEnumFromValue(value string) ModeSmbPermissionEnum {
    switch value {
        case "kFolderSubFoldersAndFiles":
            return ModeSmbPermission_KFOLDERSUBFOLDERSANDFILES
        case "kFolderAndSubFolders":
            return ModeSmbPermission_KFOLDERANDSUBFOLDERS
        case "kFolderAndFiles":
            return ModeSmbPermission_KFOLDERANDFILES
        case "kFolderOnly":
            return ModeSmbPermission_KFOLDERONLY
        case "kSubFoldersAndFilesOnly":
            return ModeSmbPermission_KSUBFOLDERSANDFILESONLY
        case "kSubFoldersOnly":
            return ModeSmbPermission_KSUBFOLDERSONLY
        case "kFilesOnly":
            return ModeSmbPermission_KFILESONLY
        default:
            return ModeSmbPermission_KFOLDERSUBFOLDERSANDFILES
    }
}
