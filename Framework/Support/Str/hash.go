package Str

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"os"
)

// Md5 calculate the md5 hash of a string
func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// Md5File calculates the md5 hash of a given file
func Md5File(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()

	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// Crc32 calculates the crc32 polynomial of a string
func Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

// Bin2hex converts binary data into hexadecimal representation
func Bin2hex(src []byte) string {
	return hex.EncodeToString(src)
}

// Hex2bin decodes a hexadecimally encoded binary string
func Hex2bin(str string) []byte {
	s, _ := hex.DecodeString(str)
	return s
}

// Sha1 calculates the sha1 hash of a string
func Sha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha1File calculates the sha1 hash of a file
func Sha1File(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := sha1.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

//
//// Md5 md5()
//func Md5(str string) string {
//	hash := md5.New()
//	hash.Write([]byte(str))
//	return hex.EncodeToString(hash.Sum(nil))
//}
//
//// Md5File md5_file()
//func Md5File(path string) (string, error) {
//	f, err := os.Open(path)
//	if err != nil {
//		return "", err
//	}
//	defer f.Close()
//
//	fi, err := f.Stat()
//	if err != nil {
//		return "", err
//	}
//
//	var size int64 = 1048576 // 1M
//	hash := md5.New()
//
//	if fi.Size() < size {
//		data, err := ioutil.ReadFile(path)
//		if err != nil {
//			return "", err
//		}
//		hash.Write(data)
//	} else {
//		b := make([]byte, size)
//		for {
//			n, err := f.Read(b)
//			if err != nil {
//				break
//			}
//
//			hash.Write(b[:n])
//		}
//	}
//
//	return hex.EncodeToString(hash.Sum(nil)), nil
//}

//
//// Sha1 sha1()
//func Sha1(str string) string {
//	hash := sha1.New()
//	hash.Write([]byte(str))
//	return hex.EncodeToString(hash.Sum(nil))
//}
//
//// Sha1File sha1_file()
//func Sha1File(path string) (string, error) {
//	data, err := ioutil.ReadFile(path)
//	if err != nil {
//		return "", err
//	}
//	hash := sha1.New()
//	hash.Write([]byte(data))
//	return hex.EncodeToString(hash.Sum(nil)), nil
//}
//
//// Crc32 crc32()
//func Crc32(str string) uint32 {
//	return crc32.ChecksumIEEE([]byte(str))
//}
