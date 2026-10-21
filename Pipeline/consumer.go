package pipeline

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func EmailWorker(id int, queue chan Emails, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range queue {
		smtpHost := "localhost"
		smtpPort := "1025"

		// 		formatedstr:= fmt.Sprintf(
		// 			"To: %s\r\nSubject: test Subject Email\r\n\r\n%s\r\n",
		// 			data.Email,
		//     		"Just for testing",
		// )
		// 		msg:=[] byte(formatedstr)

		msg, err1 := ExecuteTemplate(data)

		if err1 != nil {
			fmt.Printf("Worker %d Error parsing template for %s",
				id, data.Email)
			continue
		}

		fmt.Printf("wroker %d Sending Email to %s ........\n",
			id, data.Email)
		err := smtp.SendMail(smtpHost+":"+smtpPort,
			nil, "123@gmail.com", //add your email 
			[]string{data.Email}, []byte(msg))

		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("wroker %d Sent Email to %s \n",
			id, data.Email)
	}
}
