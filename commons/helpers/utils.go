package helpers

import (
    "os"
)

func IsProduction() bool {
    environment := os.Getenv("ENVIRONMENT")
    isProduction := true
    if environment == "development" {
        isProduction = false
    }
    return isProduction
}

//func CreatePrefixUUID(nameTable string) string {
//    now := time.Now()
//    newRandom, err := uuid.NewRandom()
//    if err != nil {
//        log.Fatal(err)
//        return ""
//    }
//    return nameTable + "." + strings.Replace(newRandom.String(), "-", "", -1) + strconv.Itoa(now.Nanosecond())
//}
