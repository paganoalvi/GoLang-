package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Billetera struct {
	ID       string
	Nombre   string
	Apellido string
}

type Transaccion struct {
	Desde     string
	Hacia     string
	monto     float64
	timeStamp time.Time
}

type Bloque struct {
	HashAnt     string
	Transaccion Transaccion
	timeStamp   time.Time
	Hash        string
	sig         *Bloque // para la lista enlazada
}

type BlockChain struct {
	Genesis *Bloque
	Cola    *Bloque
}

// calcular hash de un bloque

func calcularHash(block Block) string {
	registro := block.HashAnt + block.Transaccion.Desde + block.Transaccion.Hacia + fmt.Sprintf("%f", block.Transaccion.monto) + block.Transaccion.timeStamp.String()
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}
