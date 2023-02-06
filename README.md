# Code Cracker Kata

<p>This is the <a href="https://codingdojo.org/kata/CodeCracker" target="_blank">coding dojo kata</a> resolution in golang.</p>

## Problem Description:

<p>Given an alphabet decryption key like the one below, create a program that can crack any message using the decryption key.</p>

**Alphabet**

```txt
a b c d e f g h i j k l m n o p q r s t u v w x y z
```

**Decryption Key**

```txt
! ) " ( £ * % & > < @ a b c d e f g h i j k l m n o
```

<p>
You could also create a encryption program that will encrypt any message you give it using the key.
</p>

## How to test it:

- If you already have golang installed:

```shel
$ git clone git@github.com:antunesgabriel/codecracker.git && cd ./codecracker
$ go test .
//Or with benchmark test

$ go test . -bench=.
```

- If you don't have golang installed:

> coming soon

## About My Resolutions:

### First resolution:

<p>
When I saw the problem I understood that it was just a replacement of characters both to encrypt and to decrypt, since we have defined which are the letters and the code that represents each letter and both are the same size. So the first solution I proposed was to do an iteration within another one for each character of the message to be encrypted or decrypted.
</p>

**so the first code looks like this:**

```go
package codecracker


var (
	ALPHABET = []rune("abcdefghijklmnopqrstuvwxyz")
	ALPHABET_KEY = []rune(`!)"(£*%&><@abcdefghijklmno`)
)

func Encrypt(message string) string {
	encrypted := ""

	for _, messageChar := range message {

		for idx, char := range ALPHABET {

			if string(messageChar) == string(char) {
				encrypted += string(ALPHABET_KEY[idx])

				break
			}
		}
	}

	return encrypted
}

func Decrypt(messageEncrypt string) string {
	decrypted := ""

	for _, messageChar := range messageEncrypt {

		for idx, char := range ALPHABET_KEY {

			if string(messageChar) == string(char) {
				decrypted += string(ALPHABET[idx])
			}

		}
	}

	return decrypted
}
```

<p>
The code above was the first resolution that came to my mind to solve the problem quickly, however it has a complexity of <b>O(Nˆ2)</b> since we have a linear process in which we iterate the input message and for each character we can go through the worst case all <b>ALPHABET</b> and <b>ALPHABET_KEY</b> which makes the complexity of this algorithm <b>O(Nˆ2)</b>.
</p>

### Second resolution:

<p>
I saw that to improve the code without having to iterate over our key set and key equivalent, we would need to use a dictionary that already contained this key formation and its equivalent value. The resulting code looked like this:
</p>

```go
package codecracker

var (
	EncryptDicionary = map[rune]rune{}
	DecryptDicionary = map[rune]rune{}
)

func init() {
	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	alphabetKey := []rune(`!)"(£*%&><@abcdefghijklmno`)

	for idx, charDecrypted := range alphabet {
		charEncrypted := alphabetKey[idx]

		EncryptDicionary[charDecrypted] = charEncrypted
		DecryptDicionary[charEncrypted] = charDecrypted
	}
}

func Encrypt(message string) string {
	encrypted := make([]rune, len(message))

	for idx, char := range message {
		encrypted[idx] = EncryptDicionary[char]
	}

	return string(encrypted)
}

func Decrypt(messageEncrypt string) string {
	decrypted := make([]rune, len(messageEncrypt))

	for idx, char := range messageEncrypt {
		decrypted[idx] = DecryptDicionary[char]
	}

	return string(decrypted)
}
```

<p>
In the code above, the init function (<a href="https://tutorialedge.net/golang/the-go-init-function/" target="_blank">which is automatically executed in golang as soon as this package is imported</a>) is used to create two key and value dictionaries for both decryption and encryption. With that we have only one we iterate only the input parameter in the functions of encrypting and decrypting, giving us an algorithm with complexity of <b>O(N)</n>.
</p>
