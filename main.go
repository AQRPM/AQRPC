package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jD91mZM2/stdutil"
)

var points int = 0

func main() {
	fmt.Println("Ugh, another idiot. Hello. I'm AQRPC. Who are you? Blah blah blah et.c.")
	fmt.Println("I honestly couldn't care less of who you are. Now enter your fucking password.")
	fmt.Print("Password: ")

	pass := stdutil.MustScanLower()
	fmt.Println()

	points := 0

	if pass == "password" {
		complain(`Oh come on, now you're just trying to make me insult your fucking ass.
Couldn't be less obvious. If you really did think that was a good password you'd at least get the record
for being the most stupid person in the whole wide word, fucker. "Nobody am going to guess that", right? "Lol got em".
Well, lemme tell you what sunshine. Humans don't guess passwords. Machines do. They're just gonna loop through the most
used passwords, and boom, they're in your account in less than a few seconds. Was that really that great of an idea?
Now go away, I am busy trying to insult people that actually make it hard for me to insult them. Ugh.`)
	}

	pass2 := pass
	var words []string
	var words_total_len int
	for len(pass2) > 0 {
		var start int
		var best_word string
		var best_percent int
		for _, word := range dictionary {
			if len(word) <= 2 {
				continue
			}

			pass3 := pass2
			for {
				if len(pass3) < len(word) {
					break
				}

				word2 := pass3[0:len(word)]
				equal := 0

				for i := range word {
					if word[i] == word2[i] {
						equal += 1
					}
				}

				percent := equal * 100 / len(word)

				if percent > best_percent && percent > 70 {
					start = len(pass2) - len(pass3)
					best_word = word
					best_percent = percent
				}
				pass3 = pass3[1:]
			}
		}

		if len(best_word) == 0 {
			pass2 = pass2[1:]
		} else {
			pass2 = pass2[start+len(best_word):]
			words = append(words, best_word)
			words_total_len += len(best_word)
		}
	}

	if len(words) != 0 && len(pass)-words_total_len <= 10 {
		plural := ""
		if len(words) != 1 {
			plural = "s together"
		}
		complain(`So, congratz. You just put ` + strconv.Itoa(len(words)) + ` word` + plural +
			`(` + strings.Join(words, `, `) + `) and boom,
there's your password. Well, lemme tell you what. DICTIONARY ATTACKS, bitch. Ever heard of it?
Let me spell it fucking out for you. D-I-C-T-I-O-N-A-R-Y attacks. I can imagine you bragging to your friends
how good your password is, because you're such a D-I-C-K. One of those 'hackers' you worry to little about
can just loop through words and put them together and boom they got your password.
Notice how fast it was to detect this your password is so bad? I had to loop through the entire list of English words,
make it lowercase, and match with your password. Did you notice how slow it was? No? Exactly.
It's a matter of minutes before somebody cracks your password. Do you still think you're a genius?`)
	}

	if len(pass) <= 8 {
		length := strconv.Itoa(len(pass))
		complain(`Let's mention how much your password sucks.
It's ` + length + ` characters. __` + length + `_fucking_characters__.
You know what's 8 characters? 'horrible'. That and anything with smaller length SUCKS.
You can count, right? I bet you can't, but if you can... Realize this: You can count with characters.
Woah, can you believe it? a + 1 = b. You realize that fucking much, right?
Now imagine a computer simply counting a, b, c. And when it hits the end it just adds another character.
Until it finds your motherfucking password. You know, I really hope your computer screen shows your reflection.
Wanna know why? Becaue you should really take a long look at yourself. How worthless you are.
You can't even come up with a secure password. Good day.`)
	}

	fmt.Println(points) // because Go can't even compile if something is unused
	// TODO: Implement the rest of the thing
}

func complain(reason string) {
	points += 1
	fmt.Println(reason + "\n") // extra newline
}
