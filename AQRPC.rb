#!/usr/bin/ruby

puts "Ugh, another idiot. Hello. I'm AQRPC. Who are you? Blah blah blah et.c."
puts "I honestly couldn't care less of who you are. Now enter your fucking password."
print "Password: "
pass = gets.strip.downcase
puts

points = 0

if pass == "password"
	points += 1
	puts "Oh come on, now you're just trying to make me insult your fucking ass.
Couldn't be less obvious. If you really did think that was a good password you'd at least get the record
for being the most stupid person in the whole wide word, fucker. \"Nobody am going to guess that\", right? \"Lol got em\".
Well, lemme tell you what sunshine. Humans don't guess passwords. Machines do. They're just gonna loop through the most
used passwords, and boom, they're in your account in less than a few seconds. Was that really that great of an idea?
Now go away, I am busy trying to insult people that actually make it hard for me to insult them. Ugh."
	puts
end

pass2 = pass.clone
words = []
lines = File.read(File.dirname(__FILE__) + "/words").lines
until pass2.empty?
	start = 0
	best_word = nil
	best_percent = 0
	lines.each do |line|
		word = line.strip.downcase
		if word.length <= 2
			next
		end

		pass3 = pass2.clone
		loop do
			if pass3.length < word.length
				break
			end
			word2 = pass3[0..word.length]
			equal = 0

			for i in 0...word.length
				if word[i] == word2[i]
					equal += 1
				end
			end

			percent = equal * 100 / word.length

			if percent > best_percent && percent > 70
				start = pass2.length - pass3.length
				best_word = word
				best_percent = percent
			end
			pass3.slice!(0)
		end
	end

	if best_word.nil?
		pass2.slice!(0, start+1)
	else
		pass2.slice!(0, start+best_word.length)
		words.push(best_word)
	end
end

if !words.empty? && pass.length - words.map{|w| w.length}.reduce(:+) <= 10
	points += 1
	puts "So, congratz. You just put " + words.length.to_s + " word" + if words.length == 1 then "" else "s togher" end +
" (" + words.join(", ") + ") and boom,
there's your password. Well, lemme tell you what. DICTIONARY ATTACKS, bitch. Ever heard of it?
Let me spell it fucking out for you. D-I-C-T-I-O-N-A-R-Y attacks. I can imagine you bragging to your friends
how good your password is, because you're such a D-I-C-K. One of those 'hackers' you worry to little about
can just loop through words and put them together and boom they got your password.
Notice how fast it was to detect this your password is so bad? I had to loop through the entire list of English words,
make it lowercase, and match with your password. Did you notice how slow it was? No? Exactly.
It's a matter of minutes before somebody cracks your password. Do you still think you're a genius?"
	puts
end

if pass.length <= 8
	points += 1
	puts "Let's mention how much your password sucks.
It's " + pass.length.to_s + " characters. __" + pass.length.to_s + "_fucking_characters__.
You know what's 8 characters? 'horrible'. That and anything with smaller length SUCKS.
You can count, right? I bet you can't, but if you can... Realize this: You can count with characters.
Woah, can you believe it? a + 1 = b. You realize that fucking much, right?
Now imagine a computer simply counting a, b, c. And when it hits the end it just adds another character.
Until it finds your motherfucking password. You know, I really hope your computer screen shows your reflection.
Wanna know why? Becaue you should really take a long look at yourself. How worthless you are.
You can't even come up with a secure password. Good day."
	puts
end

numeric = 0
symbolic = 0
unique = []
pass.split("").each do |c|
	if !unique.include? c
		unique.push(c)
	end
	ord = c.ord
	if (ord >= 'a'.ord && ord <= 'z'.ord) || (ord >= 'A'.ord && ord <= 'Z'.ord)
	elsif ord >= '0'.ord && ord <= '9'.ord
		numeric += 1
	else
		symbolic += 1
	end
end

if pass.length - numeric <= 3
	points += 1
	puts "Has anybody ever told you that you're funny? Because you are. No, not when you're trying to be.
You're just naturally talented at being funny when you didn't mean to.
It's also known as 'stupidity'.
Your password is mostly numeric. It's almost like you tried putting your phone number in.
\"Heh, nobody will guess that my super-secret password is actually my phone number\".
I mean, you really think that's so clever? It's not like you have to pay $20 to try a password...
Which is why putting anything even remotely guessable is stupid.
And not only that. Every single wannabe programmer at least knows how to make a computer count.
It's harder to count with letters, but counting with numbers is REALLY easy.
And you just made yourself so vulnerable. Congratulations you idiot."
	puts
end

unique_sucks = unique.length <= 10
if unique_sucks
	points += 1
	puts "I took every character in your password. And then I filtered out duplicates. You know what I saw?
Your password only has " + unique.length.to_s + " unique fucking characters.
Did you really think this through? Oh wait, I forgot you didn't have a brain. Well that explains things.
If somebody tries to brute force you, they will get away with a really small dictionary.
And since you're so stupid you probably don't know what I mean by dictionary: I mean the amount of characters it has to try.
The less passwords it has to try, the faster. You'll lose your fucking account in no time.
But then again, if this is the best you can come up with, perhaps you deserve it."
	puts
end

if symbolic <= 5
	points += 1
	puts "" + if unique_sucks then "Adding to that, y" else "Y" end +
"ou literally have only " + symbolic.to_s + " fucking symbol" + if symbolic == 1 then "" else "s" end + ".
If somebody tries to brute force you, they will probably get away with the most basic dictionary." +
if unique_sucks then "" else "
And since you're so stupid you probably don't know what I mean by dictionary: I mean the amount of characters it has to try.
The less passwords it has to try, the faster. You'll lose your fucking account in no time.
But then again, if this is the best you can come up with, perhaps you deserve it." end
	puts
end

pattern = 0
for i in 0..(pass.length - 3)
	if pass[i+1].ord + (pass[i+1].ord - pass[i].ord) == pass[i+2].ord
		pattern += 1
	end
end

if pattern >= 5 || pass.length / 2 - pattern <= 5
	points += 1
	puts "Fantastic bloody work, pal. Just kidding, I'm not your pal. I'm quite happy I don't know you.
I love how you just stand there. COME ON! Can't you see what's going on? YOUR PASSWORD!
IT'S HAS A PATTERN. You're so damn ignorant, it's fantastic. It might not even take a robot to crack this.
No, a human can do patterns just fine. Who's outsmarted now? Do you really think you're the only human in the world
who can add and subtract things?
That's all.
Seriously, that's all. I'm not going to insult you until you're done.
I'm waiting...
Change your passwords. Because I doubt somebody as stupid as you use a password manager. Go change your bloody passwords.
DO IT."
	puts
end

if points == 0
	puts "How the heck do you remember this? Oh, maybe you're using a password manager.
\"Haha I got an idea: Let's run this tool on a randomly generated password by my password manager\".
I mean, well fucking done. It's not like it's a challenge to you.
Anyways, I guess since you're using a password manager you're not as stupid as I thought you were.
No worries, you're still fucking stupid, I'm sure."
else
	puts "In total, you idiot made it up to " + points.to_s + " fucking point" + if points == 1 then "" else "s" end + "."
end

if rand(10) == 0
	insult=`insult`
	if $?.exitstatus == 0
		puts
		puts "By the way, since you were so kind to put an `insult` application in your $PATH, I'll just call that."
		puts "Oh! Apparently, " + insult
	end
else
	# Ads
	puts "If you're really willing to use that password, at least save it in a password manager.
Because most password managers are so polite, I wouldn't recommend them to you.
It doesn't take a pussy to show another pussy how to be a real man.
You should chose AQRPM. It's rude. I fucking love it."
	puts "https://github.com/Mnpn03/AQRPM"
	puts "https://mnpn.me/software"
end
