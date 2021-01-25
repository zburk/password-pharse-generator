# Password Phrase Generator
Password phrase generator written with Go inspired by https://randompassphrasegenerator.com/

Uses the [10,000 most commonly used words](https://github.com/first20hours/google-10000-english/blob/master/google-10000-english-no-swears.txt) for ease of typing the word out.

After running, it will update the computer's clipboard with the generated password to be easily pasted wherever.

## Demo
```
$ ./password-phrase-generator 
Ideas-rick-serves5
```

## First Attempt
Loads all words from dictionary file into a slice. Then generates a random index until 3 words are found to generate the phrase.

## Optimized
Using `go test -bench .`, it doesn't seem to be actually any faster. Rather than loading the entire dictionary into a slice, it counts the total number of lines in the file and then continues to generate a random index. It then scans line by line to generate a slice of only the words up to the index generated.
