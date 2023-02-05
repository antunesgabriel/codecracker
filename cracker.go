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