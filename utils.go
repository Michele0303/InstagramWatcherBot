package InstagramWatcherBot

import "fmt"

func banner() {
	ascii := fmt.Sprintf(`
░▀█▀░█▀▀░░░█░█░█▀█░▀█▀░█▀▀░█░█░█▀▀░█▀▄░░░█▀▄░█▀█░▀█▀
░░█░░█░█░░░█▄█░█▀█░░█░░█░░░█▀█░█▀▀░█▀▄░░░█▀▄░█░█░░█░
░▀▀▀░▀▀▀░░░▀░▀░▀░▀░░▀░░▀▀▀░▀░▀░▀▀▀░▀░▀░░░▀▀░░▀▀▀░░▀░`)
	fmt.Println(ascii, "\n")
}
