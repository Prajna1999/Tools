package main

import (
    "fmt"
    rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtctokenbuilder2"
)

func main() {
    appID := "970CA35de60c44645bbae8a215061b33"
    appCertificate := "5CFd2fd1755d40ecb72977518be15d3b"
    channelName := "7d72365eb983485397e3e3f9d460bdda"
    uid := uint32(2882341273)
    uidStr := "2882341273"
    expirationInSeconds := uint32(3600)

    result, err := rtctokenbuilder.BuildTokenWithUid(appID, appCertificate, channelName, uid, rtctokenbuilder.RoleSubscriber, expirationInSeconds)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Token with uid: %s\n", result)
    }

    result, err = rtctokenbuilder.BuildTokenWithUserAccount(appID, appCertificate, channelName, uidStr, rtctokenbuilder.RoleSubscriber, expirationInSeconds)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Token with userAccount: %s\n", result)
    }

    result, err1 := rtctokenbuilder.BuildTokenWithUidAndPrivilege(appID, appCertificate, channelName, uid,
        expirationInSeconds, expirationInSeconds, expirationInSeconds, expirationInSeconds, expirationInSeconds)
    if err1 != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Token with uid and privilege: %s\n", result)
    }

    result, err1 = rtctokenbuilder.BuildTokenWithUserAccountAndPrivilege(appID, appCertificate, channelName, uidStr,
        expirationInSeconds, expirationInSeconds, expirationInSeconds, expirationInSeconds, expirationInSeconds)

    if err1 != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Token with userAccount and privilege: %s\n", result)
    }
}