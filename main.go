package main

import (
	"fmt"
	"strconv"
	"strings"
    "unicode/utf8"
    "math/rand"
    "time"
    "os/exec"

	"github.com/jD91mZM2/stdutil"
)

var points int = 0

func main() {
    rand.Seed(time.Now().UnixNano())

	fmt.Println("Ugh, another idiot. Hello. I'm AQRPC. Who are you? Blah blah blah et.c.")
	fmt.Println("I honestly couldn't care less of who you are. Now enter your fucking password.")
	fmt.Print("Password: ")

	pass := stdutil.MustScanLower()
	fmt.Println()

	if pass == "password" {
		complain(`Oh come on, now you're just trying to make me insult your fucking ass.
Couldn't be less obvious. If you really did think that was a good password you'd at least get the record
for being the most stupid person in the whole wide word, fucker. "Nobody am going to guess that", right? "Lol got em".
Well, lemme tell you what sunshine. Humans don't guess passwords. Machines do. They're just gonna loop through the most
used passwords, and boom, they're in your account in less than a few seconds. Was that really that great of an idea?
Now go away, I am busy trying to insult people that actually make it hard for me to insult them. Ugh.`)
	}

    var words []string
    var words_total_len int
    {
        pass2 := pass
        for len(pass2) > 0 {
            var start int
            var best_word string
            var best_percent int
            for _, word := range dictionary {
                if l(word) <= 2 {
                    continue
                }

                pass3 := pass2
                for {
                    if l(pass3) < l(word) {
                        break
                    }

                    word2 := pass3[0:len(word)]
                    equal := 0

                    for i := range word {
                        if word[i] == word2[i] {
                            equal += 1
                        }
                    }

                    percent := equal * 100 / l(word)

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
                words_total_len += l(best_word)
            }
        }
    }

	if len(words) != 0 && len(pass)-words_total_len <= 10 {
		complain(`So, congratz. You just put ` + str(len(words)) + ` word` + ternary(len(words) == 1, "", "s together") +
			` (` + strings.Join(words, `, `) + `) and boom,
there's your password. Well, lemme tell you what. DICTIONARY ATTACKS, bitch. Ever heard of it?
Let me spell it fucking out for you. D-I-C-T-I-O-N-A-R-Y attacks. I can imagine you bragging to your friends
how good your password is, because you're such a D-I-C-K. One of those 'hackers' you worry to little about
can just loop through words and put them together and boom they got your password.
Notice how fast it was to detect this your password is so bad? I had to loop through the entire list of English words,
make it lowercase, and match with your password. Did you notice how slow it was? No? Exactly.
It's a matter of minutes before somebody cracks your password. Do you still think you're a genius?`)
	}

	if l(pass) <= 8 {
		length := str(l(pass))
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

	var numeric int
	var symbolic int
	var unique []rune

outer:
	for _, c := range pass {
		for _, c2 := range unique {
			if c == c2 {
				continue outer
			}
		}
		unique = append(unique, c)

		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		} else if c >= '0' && c <= '9' {
			numeric += 1
		} else {
			symbolic += 1
		}
	}

	if len(pass)-numeric <= 3 {
		complain(`Has anybody ever told you that you're funny? Because you are. No, not when you're trying to be.
You're just naturally talented at being funny when you didn't mean to.
It's also known as 'stupidity'.
Your password is mostly numeric. It's almost like you tried putting your phone number in.
"Heh, nobody will guess that my super-secret password is actually my phone number".
I mean, you really think that's so clever? It's not like you have to pay $20 to try a password...
Which is why putting anything even remotely guessable is stupid.
And not only that. Every single wannabe programmer at least knows how to make a computer count.
It's harder to count with letters, but counting with numbers is REALLY easy.
And you just made yourself so vulnerable. Congratulations you idiot.`)
	}

	dictionary_explanation := `
And since you're so stupid you probably don't know what I mean by dictionary: I mean the amount of characters it has to try.
The less passwords it has to try, the faster. You'll lose your fucking account in no time.
But then again, if this is the best you can come up with, perhaps you deserve it.`

	unique_sucks := len(unique) <= 10
	if unique_sucks {
		complain(`I took every character in your password. And then I filtered out duplicates. You know what I saw?
Your password only has ` + str(len(unique)) + ` unique fucking character` + ternary(len(unique) == 1, ``, `s`) + `.
Did you really think this through? Oh wait, I forgot you didn't have a brain. Well that explains things.
If somebody tries to brute force you, they will get away with a really small dictionary.` + dictionary_explanation)
	}

	if symbolic <= 5 {
		complain(ternary(unique_sucks, `Adding to that, y`, `Y`) + `ou literally have only ` + str(symbolic) +
            ` fucking symbol` +
			ternary(symbolic == 1, ``, `s`) + `.
If somebody tries to brute force you, they will probably get away with the most basic dictionary.` +
			ternary(unique_sucks, ``, dictionary_explanation))
	}

    var pattern int
    {
        pass2 := []rune(pass)
        for i := 0; i < len(pass2) - 3; i += 1 {
            if pass[i+1] + (pass[i+1] - pass[i]) == pass[i+2] {
                pattern += 1
            }
        }
    }

    if pattern >= 5 || l(pass) / 2 - pattern <= 5 {
        complain(`Fantastic bloody work, pal. Just kidding, I'm not your pal. I'm quite happy I don't know you.
I love how you just stand there. COME ON! Can't you see what's going on? YOUR PASSWORD!
IT'S HAS A PATTERN. You're so damn ignorant, it's fantastic. It might not even take a robot to crack this.
No, a human can do patterns just fine. Who's outsmarted now? Do you really think you're the only human in the world
who can add and subtract things?
That's all.
Seriously, that's all. I'm not going to insult you until you're done.
I'm waiting...
Change your passwords. Because I doubt somebody as stupid as you use a password manager. Go change your bloody passwords.
DO IT.`)
    }

    if points == 0 {
        fmt.Println(`How the heck do you remember this? Oh, maybe you're using a password manager.
"Haha I got an idea: Let's run this tool on a randomly generated password by my password manager".
I mean, well fucking done. It's not like it's a challenge to you.
Anyways, I guess since you're using a password manager you're not as stupid as I thought you were.
No worries, you're still fucking stupid, I'm sure.`)
    } else {
        fmt.Println(`In total, you idiot made it up to ` + str(points) + ` fucking point` + ternary(points == 1, ``, `s`) + ".")
    }

    path, err := exec.LookPath("insult")
    if err != nil && rand.Intn(10) == 0 {
        fmt.Println("By the way, since you were so kind to put an `insult` application in your $PATH, I'll just call that.")

        out, err := exec.Command(path).Output()
        if err != nil {
            stdutil.PrintErr("Nevermind, you got out real lucky there. Something went wrong. Idiot.", nil);
        } else {
            fmt.Println("Oh! Apparently, " + string(out))
        }
    } else {
        // Ads
        fmt.Println(`If you're really willing to use that password, at least save it in a password manager.
Because most password managers are so polite, I wouldn't recommend them to you.
It doesn't take a pussy to show another pussy how to be a real man.
You should choose AQRPM. It's rude. I fucking love it.

https://github.com/AQRPM/AQRPM`)
    }
}

func complain(reason string) {
	points += 1
	fmt.Println(reason + "\n") // extra newline
}
func l(str string) int {
    return utf8.RuneCountInString(str)
}
func str(i int) string {
	return strconv.Itoa(i)
}
func ternary(condition bool, str1 string, str2 string) string {
	if condition {
		return str1
	} else {
		return str2
	}
}
