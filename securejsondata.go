package securejsondata

import (
	"log"
)

// SecureJsonData is used to store encrypted data (for example in data_source table). Only values are separately
// encrypted.
type SecureJsonData map[string][]byte

// DecryptedValue returns single decrypted value from SecureJsonData. Similar to normal map access second return value
// is true if the key exists and false if not.
func (s SecureJsonData) DecryptedValue(key string) (string, bool) {
	if value, ok := s[key]; ok {
		decryptedData, err := Decrypt(value, SecretKey())
		if err != nil {
			log.Fatalln(err)
		}
		return string(decryptedData), true
	}
	return "", false
}

// Decrypt returns map of the same type but where the all the values are decrypted. Opposite of what
// GetEncryptedJsonData is doing.
func (s SecureJsonData) Decrypt() map[string]string {
	decrypted := make(map[string]string)
	for key, data := range s {
		decryptedData, err := Decrypt(data, SecretKey())
		if err != nil {
			log.Fatalln(err)
		}

		decrypted[key] = string(decryptedData)
	}
	return decrypted
}

// GetEncryptedJsonData returns map where all keys are encrypted.
func GetEncryptedJsonData(sjd map[string]string) SecureJsonData {
	encrypted := make(SecureJsonData)
	for key, data := range sjd {
		encryptedData, err := Encrypt([]byte(data), SecretKey())
		if err != nil {
			log.Fatalln(err)
		}

		encrypted[key] = encryptedData
	}
	return encrypted
}
