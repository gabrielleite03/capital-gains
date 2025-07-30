// Pacote models define estruturas e tipos relacionados aos dados de negócios da aplicação,
// incluindo regras específicas de serialização para valores monetários.
package models

import (
	"fmt"
	"strconv"
)

// CapitalGains representa os ganhos de capital, contendo o valor do imposto calculado.
//
// O campo Tax é do tipo MyFloat64, que possui uma serialização JSON personalizada
// para garantir precisão e consistência na representação de valores monetários.
type CapitalGains struct {
	Tax MyFloat64 `json:"tax"`
}

// MyFloat64 é um alias para o tipo float64 com uma implementação personalizada do método MarshalJSON.
//
// Essa serialização garante que:
//   - Valores inteiros sejam representados com pelo menos uma casa decimal (ex: "10.0").
//   - Valores decimais mantenham sua precisão usando notação de ponto fixo.
//
// Esse comportamento é útil em APIs que exigem consistência na formatação de números decimais,
// especialmente para representar valores monetários.
type MyFloat64 float64

// MarshalJSON implementa a interface json.Marshaler para o tipo MyFloat64.
//
// Ele garante que valores inteiros sejam formatados com uma casa decimal (ex: "5.0"),
// enquanto valores com parte fracionária são mantidos com sua precisão atual.
func (mf MyFloat64) MarshalJSON() ([]byte, error) {
	// Formata o float64 para garantir notação de ponto fixo com pelo menos uma casa decimal.
	s := strconv.FormatFloat(float64(mf), 'f', -1, 64)

	// Se o valor for inteiro, força a presença de uma casa decimal.
	if float64(mf) == float64(int64(mf)) {
		return []byte(fmt.Sprintf("%.1f", mf)), nil
	}

	// Retorna o valor original formatado.
	return []byte(s), nil
}
