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
func (bc BlockChain) ObtenerBalance(IDB string) float64 { // receiver BlockChain, parametro input (string), retorna float64
	balance := 0.0
	act := bc.Genesis
	for act != nil { // recorro blockchain (while act <> nil)
		tx := act.Transaccion
		if tx.IDHacia == IDB { // si IdHacia(transaccion) == IDB ingresado por parametro
			balance += tx.monto // sumo monto de la transaccion
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
	return hex.EncodeToString(hash[:]) // retorno hash como un string encriptado
}

// crear una blockchain  vacia
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
	// No validar saldo si es una transacción de creación (desde " ") (transaccion inicial)
	if desde != "" && bc.ObtenerBalance(desde) < monto {
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
		if act.sig != nil && act.sig.HashAnt != act.Hash {
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

	usuarios := []Billetera{
		NuevaBilletera("alvaP14", "Alvaro", "Pagano"),
		NuevaBilletera("Agus073", "Agustin", "Servin"),
		NuevaBilletera("Mari22", "Maria", "Gomez"),
		NuevaBilletera("Jose_45", "Jose", "Perez"),
	}

	fmt.Println("\n--- Minado inicial ---")
	// Transacciones de creación deben tener desde = ""
	bc.EnviarTransaccion("", usuarios[0].ID, 1000) // Minado para Alvaro
	bc.EnviarTransaccion("", usuarios[1].ID, 500)  // Minado para Agustin
	bc.EnviarTransaccion("", usuarios[2].ID, 750)  // Minado para Maria

	fmt.Println("\n--- Saldos después del minado ---")
	for _, usuario := range usuarios {
		fmt.Printf("%s: %.2f\n", usuario.ID, bc.ObtenerBalance(usuario.ID))
	}

	fmt.Println("\n--- Transacciones normales ---")
	transacciones := []struct {
		desde, hacia string
		monto        float64
	}{
		{usuarios[0].ID, usuarios[1].ID, 100},
		{usuarios[2].ID, usuarios[0].ID, 50},
		{usuarios[1].ID, usuarios[3].ID, 75},
		{usuarios[0].ID, usuarios[2].ID, 200},
		{usuarios[2].ID, usuarios[3].ID, 100},
	}

	for _, tx := range transacciones {
		err := bc.EnviarTransaccion(tx.desde, tx.hacia, tx.monto)
		if err != nil {
			fmt.Printf("Error: %s -> %s: %v\n", tx.desde, tx.hacia, err)
		} else {
			fmt.Printf("Éxito: %s -> %s (%.2f)\n", tx.desde, tx.hacia, tx.monto)
		}
	}

	fmt.Println("\n--- Saldos finales ---")
	for _, usuario := range usuarios {
		fmt.Printf("%s: %.2f\n", usuario.ID, bc.ObtenerBalance(usuario.ID))
	}

	fmt.Println("\n--- Validación ---")
	fmt.Println("Blockchain válida:", bc.esValida())

	// Mostrar blockchain
	fmt.Println("\n--- Blockchain completa ---")
	contador := 0
	for act := bc.Genesis; act != nil; act = act.sig {
		tx := act.Transaccion
		fmt.Printf("Bloque %d: %s -> %s (%.2f) | Hash: %s\n",
			contador, tx.IDDesde, tx.IDHacia, tx.monto, act.Hash[:8]+"...")
		contador++
	}
}

/* Librerias
crypto/sha256
functionality to compute SHA-224 and SHA-256 cryptographic hash algorithms, as defined in FIPS 180-4.
These algorithms are widely used for generating fixed-size, unique checksums for data, ensuring data
integrity and providing short identities for binary or text blobs


encoding/hex
Hexadecimal encoding and decoding of data. It allows conversion between arbitrary binary data and
 its hexadecimal string representation

time
for measuring and displaying time

*/
