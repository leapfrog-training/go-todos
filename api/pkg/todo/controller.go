package todo

import (
	"encoding/json"
	"log"
	"net/http"
	database "todos/api/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
 * To structure the response given by the application
 * @type {struct} Response
 */
type Response struct {
	Status     string      `json:"status"`
	StatucCode int         `json:"statusCode"`
	Data       interface{} `json:"data,omitempty"`
}

/**
 * To get all list present in the to-do application
 * @function Index
 * @param {http.ResponseWriter} w
 * @param {*http.Request} r
 * @return {json}
 */
func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/todo" {
		http.NotFound(w, r)
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		var todo Todo
		var todo_all []Todo
		cursor, err := database.Collection.Find(database.Ctx, bson.D{})
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(&Response{"failure", 204, "No Content"})
			return
		}

		for cursor.Next(database.Ctx) {
			err := cursor.Decode(&todo)
			if err != nil {
				w.WriteHeader(http.StatusNoContent)
				json.NewEncoder(w).Encode(&Response{"failure", 204, "No Content"})
				return
			}
			todo_all = append(todo_all, todo)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&Response{"success", 200, todo_all})

	} else {
		http.Error(w, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
	}
}

/**
 * To create new entry to to-do application
 * @function Store
 * @param {http.ResponseWriter} w
 * @param {*http.Request} r
 * @return {json}
 */
func Store(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/todo/create" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {

		var data Todo

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(&Response{"failure", 204, "No Content"})
			return
		}

		result, err := database.Collection.InsertOne(database.Ctx, data)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(&Response{"failure", 204, "No Content"})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&Response{"success", 201, result})
	} else {
		http.Error(w, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
	}

}

/**
 * To update existing entry to to-do application
 * @function Store
 * @param {http.ResponseWriter} w
 * @param {*http.Request} r
 * @param {int} id
 * @return {json}
 */
func Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "application/json")
	todoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{"failure", 400, "Bad Request"})
		return
	}
	if len(id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{"failure", 400, "Bad Request"})
		return
	}
	if r.URL.Path != "/todo/edit" {
		http.NotFound(w, r)
		return
	}
	if r.Method == "PUT" {
		var data Todo

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(&Response{"failure", 204, "No Content"})
			return
		}

		result, err := database.Collection.UpdateOne(database.Ctx, bson.M{"_id": todoID}, bson.D{{Key: "$set", Value: data}})
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(&Response{"failure", 204, "No Content"})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&Response{"success", 200, result})
	} else {
		http.Error(w, "Only PUT requests are allowed!", http.StatusMethodNotAllowed)
	}
}

/**
 * To delete existing entry to to-do application
 * @function Store
 * @param {http.ResponseWriter} w
 * @param {*http.Request} r
 * @param {int} id
 * @return {json}
 */
func Destory(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "application/json")
	todoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{"failure", 400, "Bad Request"})
		return
	}
	if len(id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{"failure", 400, "Bad Request"})
		return
	}
	if r.URL.Path != "/todo/delete" {
		http.NotFound(w, r)
		log.Println("delete error")
		return
	}
	if r.Method == "DELETE" {
		result, err := database.Collection.DeleteOne(database.Ctx, bson.D{{Key: "_id", Value: todoID}})
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(&Response{"failure", 204, "No Content"})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&Response{"success", 200, result})
	} else {
		http.Error(w, "Only DELETE requests are allowed!", http.StatusMethodNotAllowed)
	}
}

/**
 * To mark all completed in existing entry to to-do application
 * @function Store
 * @param {http.ResponseWriter} w
 * @param {*http.Request} r
 * @param {int} id
 * @return {json}
 */
func MarkAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path != "/todo/mark-all" {
		http.NotFound(w, r)
		return
	}
	if r.Method == "PUT" {
		var data Todo

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(&Response{"failure", 204, "No Content"})
			return
		}

		result, err := database.Collection.UpdateMany(database.Ctx, bson.M{}, bson.D{{Key: "$set", Value: data}})
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(&Response{"failure", 204, "No Content"})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&Response{"success", 200, result})
	} else {
		http.Error(w, "Only PUT requests are allowed!", http.StatusMethodNotAllowed)
	}
}