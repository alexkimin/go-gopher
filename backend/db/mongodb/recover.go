package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/AlexMin314/go-gopher/backend/todo/constant"
	"github.com/AlexMin314/go-gopher/backend/todo/schema"
)

func recoverTodo(m *Mongo) {
	collection := m.DB.Collection("todos")

	// init data
	todo1 := schema.Todo{Title: "Learning go concurrency", Checked: constant.No}
	todo2 := schema.Todo{Title: "Learning mongodb in depth", Checked: constant.No}
	todo3 := schema.Todo{Title: "Learning MQTT and more", Checked: constant.No}
	// todo1 := schema.Todo{ID: primitive.NewObjectID(), Title: "Learning go concurrency", Checked: constant.No}
	// todo2 := schema.Todo{ID: primitive.NewObjectID(), Title: "Learning mongodb in depth", Checked: constant.No}
	// todo3 := schema.Todo{ID: primitive.NewObjectID(), Title: "Learning MQTT and more", Checked: constant.No}

	insertManyResult, err := collection.InsertMany(context.TODO(), []interface{}{todo1, todo2, todo3})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}