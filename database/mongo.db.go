package database

import (
	c "admin-server/config"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbManager interface {
	// Devuelve un handle a una colección del nombre name.
	GetCollection(collection string) *mongo.Collection

	// Se desconecta de la base de datos, usualmente llamada al inicio del main.
	//
	//	func main() {
	// 		defer db.GetManager().Disconnect()
	//		//El resto del código aquí
	// 	}
	Disconnect()
}

// Tipo que contiene el apuntador al cliente y el contexto.
type dbHandler struct {
	db  *mongo.Client
	ctx context.Context
}

// Variable manejador de la base de datos, el cual se le
// asigna un 'dbHandler'
var mgr dbManager

var (
	user     = c.Config.Mongo.User
	password = c.Config.Mongo.Pass
	host     = c.Config.Mongo.Host
	port     = c.Config.Mongo.Port
	dbName   = c.Config.Mongo.DbName
	timeOut  = c.Config.Mongo.TimeOut
)

func init() {

	// Genera el string de conexión de una base de datos de mongo.
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", user, password, host, port)

	// Crea un nuevo cliente para la base de datos.
	db, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error al crear cliente de mongodb.\n", err)
	}

	// Establece el contexto para relizar la conexión.
	ctx, cancel := context.WithTimeout(context.TODO(), timeOut*time.Second)
	if err := db.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer cancel()

	// Verifica que se haya conectado a la base de datos.
	if err := db.Ping(ctx, nil); err != nil {
		log.Fatal("Error al establecer conexión con mongodb.\n", err)
	}

	mgr = &dbHandler{db: db, ctx: ctx}

}

// Devuelve el manejador de la conexión a la base de datos.
// Para evitar que se le asigne un nuevo valor.
func GetManager() dbManager {
	return mgr
}

// Implementación de la función 'Disconnect' de la interface 'dbManager'
func (mgr *dbHandler) Disconnect() {
	fmt.Println("Cerrando conexión con mongodb...")
	if mgr != nil {
		mgr.db.Disconnect(mgr.ctx)
	}
}

// Implementación de la función 'Disconnect' de la interface 'dbManager'
// Nota: En caso que el gestor no esté inicializado (nil), imprime el error
// y termina la aplicación (en teoria no debería de pasar).
func (mgr *dbHandler) GetCollection(name string) *mongo.Collection {
	if mgr == nil {
		err := fmt.Errorf("no se puede obtener la collección '%s', debido a que la base de datos no está inicializada", name)
		log.Fatal(err)
	}
	return mgr.db.Database(dbName).Collection(name)
}
