package models

import(
    "encoding/json"
)

/**
 * Type definition for TaskTypeEnum enum
 */
type TaskTypeEnum int

/**
 * Value collection for TaskTypeEnum enum
 */
const (
    TaskType_RESTORE TaskTypeEnum = 1 + iota
    TaskType_CLONE
    TaskType_BACKUPNOW
    TaskType_FIELDMESSAGE
)

func (r TaskTypeEnum) MarshalJSON() ([]byte, error) { 
    s := TaskTypeEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *TaskTypeEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  TaskTypeEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts TaskTypeEnum to its string representation
 */
func TaskTypeEnumToValue(taskTypeEnum TaskTypeEnum) string {
    switch taskTypeEnum {
        case TaskType_RESTORE:
    		return "Restore"		
        case TaskType_CLONE:
    		return "Clone"		
        case TaskType_BACKUPNOW:
    		return "BackupNow"		
        case TaskType_FIELDMESSAGE:
    		return "FieldMessage"		
        default:
        	return "Restore"
    }
}

/**
 * Converts TaskTypeEnum Array to its string Array representation
*/
func TaskTypeEnumArrayToValue(taskTypeEnum []TaskTypeEnum) []string {
    convArray := make([]string,len( taskTypeEnum))
    for i:=0; i<len(taskTypeEnum);i++ {
        convArray[i] = TaskTypeEnumToValue(taskTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TaskTypeEnumFromValue(value string) TaskTypeEnum {
    switch value {
        case "Restore":
            return TaskType_RESTORE
        case "Clone":
            return TaskType_CLONE
        case "BackupNow":
            return TaskType_BACKUPNOW
        case "FieldMessage":
            return TaskType_FIELDMESSAGE
        default:
            return TaskType_RESTORE
    }
}
