package main

import (
	"fmt"
	"sync"
	"time"
)

// contact representa un contacto en la agenda
type Contact struct {
	nombre            string
	apellido          string
	correoElectronico string
	telefono          string
}

// Agenda maneja los contactos de forma concurrente
type Agenda struct {
	// mapa de contact con correo electronico como clave
	contacts map[string]Contact
	mu       sync.RWMutex
}

// Crear una nueva agenda
func NewAgenda() *Agenda { //  no recibe parametro de entrada, devuelve un puntero a la Agenda
	return &Agenda{
		contacts: make(map[string]Contact),
	}
}

// AgregarContacto anade un nuevo contacto de forma segura y concurrente
func (a *Agenda) AgregarContacto(c Contact) { // gorutine recibe puntero a la agenda, AgregarContacto recibe un contacto
	a.mu.Lock()                         // bloqueo exlusivo de recurso(nadie mas que esta gorutine puede leer o escribir)
	defer a.mu.Unlock()                 // desbloqueo recurso antes de salir
	a.contacts[c.correoElectronico] = c // asigno a la agenda con el correo electronico como clave el contacto
}

// EliminarContacto remueve un contacto (con email como clave)
func (a *Agenda) EliminarContacto(correo string) {
	a.mu.Lock()         // bloqueo recurso
	defer a.mu.Unlock() // desbloqueo recurso antes de salir
	delete(a.contacts, correo)
}

func (a *Agenda) BuscarContacto(correo string) (Contact, bool) {
	a.mu.RLock() // Bloqueo de lectura(operacion de consulta como buscar un contacto)(Múltiples goroutines puedan leer simultáneamente)
	defer a.mu.RUnlock()
	contact, exists := a.contacts[correo]
	return contact, exists
}

// MostrarAgenda imprime todos los contactos
func (a *Agenda) MostrarAgenda() {
	a.mu.RLock() // Bloqueo de lectura(operacion de consulta como leer la agenda)
	defer a.mu.RUnlock()
	fmt.Println("\n--- Agenda de Contactos ---")
	for _, contact := range a.contacts {
		fmt.Printf("Nombre: %s %s\nEmail: %s\nTeléfono: %s\n\n",
			contact.nombre, contact.apellido,
			contact.correoElectronico, contact.telefono)
	}
	fmt.Printf("Total contactos: %d\n", len(a.contacts))
}

func main8() {
	agenda := NewAgenda()
	var wg sync.WaitGroup

	agregarContacto := func(c Contact) { // funcion helper para agregar contacto concurrentemente
		defer wg.Done()
		agenda.AgregarContacto(c)
		fmt.Printf("Agrego: %s %s\n", c.nombre, c.apellido)
	}

	buscarContacto := func(correo string) { // funcion helper para buscar contactos concurrentemente
		defer wg.Done()
		if contacto, existe := agenda.BuscarContacto(correo); existe { // (if existe)
			fmt.Printf("Encontrado:‌ %s %s (%s)\n", contacto.nombre, contacto.apellido, correo)
		} else {
			fmt.Printf("No encontrado: %s\n", correo)
		}
	}

	eliminarContacto := func(correo string) {
		defer wg.Done()
		agenda.EliminarContacto(correo)
		fmt.Printf("Eliminado: %s\n", correo)
	}

	// Datos de prueba
	contactos := []Contact{
		{"Juan", "Pérez", "juan@example.com", "555-1000"},
		{"María", "Gómez", "maria@example.com", "555-2000"},
		{"Carlos", "López", "carlos@example.com", "555-3000"},
		{"Ana", "Martínez", "ana@example.com", "555-4000"},
		{"Luisa", "Fernández", "luisa@example.com", "555-5000"},
		{"Pedro", "Sánchez", "pedro@example.com", "555-6000"},
	}

	// Calculamos el número exacto de operaciones
	numAgregar := len(contactos) // 6 operaciones
	numBuscar := 3               // 3 búsquedas
	numEliminar := 2             // 2 eliminaciones
	totalOperaciones := numAgregar + numBuscar + numEliminar

	wg.Add(totalOperaciones) // Añadimos el total exacto de operaciones

	// Agregar contactos
	for _, c := range contactos {
		go agregarContacto(c)
	}

	// Pequeña pausa para permitir que algunos contactos se agreguen
	time.Sleep(100 * time.Millisecond)

	// Operaciones de búsqueda
	go buscarContacto("juan@example.com")
	go buscarContacto("maria@example.com")
	go buscarContacto("email@inexistente.com")

	// Operaciones de eliminación
	go eliminarContacto("juan@example.com")
	go eliminarContacto("carlos@example.com")

	wg.Wait() // Esperar a que todas las operaciones terminen

	// Mostrar estado final de la agenda
	agenda.MostrarAgenda()
}
