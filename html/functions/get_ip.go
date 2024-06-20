package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Erro ao obter endereços de interface:", err)
		return
	}

	for _, addr := range addrs {
		// Verifica se o endereço é do tipo IPNet e se não é um endereço de loopback
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println("Meu IP é:", ipNet.IP.String())
				return
			}
		}
	}

	fmt.Println("Não foi possível encontrar um endereço IP válido.")
}
