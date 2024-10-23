package encrypter

import (
	"fmt"
	"strconv"
)

func main() {
	textoCriptografado := Criptografia("CRIPTOGRAFAR", "Text to encrypt")
	fmt.Println("Texto criptografado:", textoCriptografado, "\n")

	textoDescriptografado := Criptografia("DESCRIPTOGRAFAR", textoCriptografado)
	fmt.Println("Texto descriptografado:", textoDescriptografado, "\n")
}

func Criptografia(operacao, valor string) string {
	if len(valor) < 0 {
		return ""
	}

	var ultimoCod, tmpCodAsc, SrcPos, codAsc int
	chavePos := 0

	var valorFinal string
	chave := "AUQL23RL23DF90WI5E1EAS462NMCXEL1JAOMUMUMCL0EOMM4A4V3Y29KHJUI23476JHDYDB3434SKLKK3LAKDJ6L9YTIA7"

	if operacao == "CRIPTOGRAFAR" {
		ultimoCod = 1
		valorFinal = "01"

		for SrcPos = 0; SrcPos < len(valor); SrcPos++ {
			codAsc = (int(valor[SrcPos]) + ultimoCod) % 255

			if chavePos >= len(chave) {
				chavePos = 0
			}

			codAsc = codAsc ^ int(chave[chavePos])

			hexValue := fmt.Sprintf("%02x", codAsc)
			valorFinal += hexValue

			ultimoCod = codAsc
			chavePos++
		}
	} else if operacao == "DESCRIPTOGRAFAR" {
		ultimoCod, _ = decodeHexString("0x" + valor[0:2])
		SrcPos = 2

		for SrcPos < len(valor) {
			codAsc, _ = decodeHexString("0x" + valor[SrcPos:SrcPos+2])

			if chavePos >= len(chave) {
				chavePos = 0
			}

			tmpCodAsc = codAsc ^ int(chave[chavePos])

			if tmpCodAsc <= ultimoCod {
				tmpCodAsc = 255 + tmpCodAsc - ultimoCod
			} else {
				tmpCodAsc = tmpCodAsc - ultimoCod
			}

			valorFinal = addToFinalValue(tmpCodAsc, valorFinal)

			ultimoCod = codAsc
			SrcPos = SrcPos + 2
			chavePos += 1
		}
	}
	return valorFinal
}

func decodeHexString(hexString string) (int, error) {
	// Remove o prefixo "0x" se estiver presente
	if len(hexString) > 2 && hexString[:2] == "0x" {
		hexString = hexString[2:]
	}

	// Converte a string hexadecimal para um inteiro de 64 bits
	intValue, err := strconv.ParseInt(hexString, 16, 64)
	if err != nil {
		return 0, err
	}

	return int(intValue), nil
}

func addToFinalValue(value int, finalValue string) string {
	return finalValue + string(rune(value))
}

func calculateCodAsc(valor string, srcPos, ultimoCod int) int {
	// Obtém o código ASCII do caractere na posição SrcPos da string valor
	charCodeAtSrcPos := int(valor[srcPos])

	// Calcula o novo código ASCII somando o código do caractere na posição SrcPos
	// com o último código e aplicando a operação módulo 255
	codAsc := (charCodeAtSrcPos + ultimoCod) % 255

	return codAsc
}
