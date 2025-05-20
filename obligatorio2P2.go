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
	IDDesde   string
	IDHacia   string
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

// obtener balance => metodo de BlockChain
func (bc BlockChain) ObtenerBalance(IDB string) float64 {
	balance := 0.0
	act := bc.Genesis
	for act != nil {
		tx := act.Transaccion
		if tx.IDHacia == IDB {
			balance += tx.monto
		}
		if tx.IDDesde == IDB {
			balance -= tx.monto
		}
		act = act.sig
	}
	return balance
}

// calcular hash de un bloque
func calcularHash(block Bloque) string {
	registro := block.HashAnt + block.Transaccion.IDDesde + block.Transaccion.IDHacia + fmt.Sprintf("%f", block.Transaccion.monto) + block.Transaccion.timeStamp.String()
	hash := sha256.Sum256([]byte(registro))
	return hex.EncodeToString(hash[:])
}

// crear una blockchain
func NuevaBlockChain() BlockChain {
	genesisTx := Transaccion{"", "", 0, time.Now()}
	genesisBlock := Bloque{
		HashAnt:     "",
		Transaccion: genesisTx,
		timeStamp:   time.Now(),
	}
	genesisBlock.Hash = calcularHash(genesisBlock)
	return BlockChain{
		Genesis: &genesisBlock,
		Cola:    &genesisBlock,
	}
}

// crear billetera
func NuevaBilletera(id, nombre, apellido string) Billetera {
	return Billetera{
		ID:       id,
		Nombre:   nombre,
		Apellido: apellido,
	}
}

// Enviar transaccion si es posible
func (bc *BlockChain) EnviarTransaccion(desde, hacia string, monto float64) error {
	if (bc.ObtenerBalance(desde)) < monto { // invocacion a ObtenerBalance
		return fmt.Errorf("Saldo insuficiente")
	}

	tx := Transaccion{
		IDDesde:   desde,
		IDHacia:   hacia,
		monto:     monto,
		timeStamp: time.Now(),
	}
	nuevoBloque := Bloque{
		HashAnt:     bc.Cola.Hash,
		Transaccion: tx,
		timeStamp:   time.Now(),
	}
	nuevoBloque.Hash = calcularHash(nuevoBloque)
	bc.Cola.sig = &nuevoBloque
	bc.Cola = &nuevoBloque
	return nil
}

func (bc BlockChain) esValida() bool {
	act := bc.Genesis
	for act != nil {
		if act.sig.HashAnt != act.Hash {
			return false
		}
		if calcularHash(*act) != act.Hash {
			return false
		}
		act = act.sig
	}
	return true
}

func mainO2() {
	bc := NuevaBlockChain()

	alvaro := NuevaBilletera("alvaP14", "Alvaro", "Pagano")
	agustin := NuevaBilletera("Agus073", "Agustin", "Servin")

	// incializo con una transaccion de fondos a alvaro
	bc.EnviarTransaccion("", alvaro.ID, 1000)

	err := bc.EnviarTransaccion(alvaro.ID, agustin.ID, 30)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Transaccion realizada con exito")
	}
	fmt.Println("Saldo alva: ", bc.ObtenerBalance(alvaro.ID))
	fmt.Println("Saldo agus: ", bc.ObtenerBalance(agustin.ID))

	fmt.Println("Blockchain valida?", bc.esValida())

}
