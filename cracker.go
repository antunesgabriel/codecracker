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
