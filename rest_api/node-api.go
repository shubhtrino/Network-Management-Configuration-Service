package rest_api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"crypto/rand"
	"encoding/binary"

	"github.com/gorilla/mux"
)

// swagger:model Node
type Node struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Latitude  int    `json:"latitude"`
	Longitude int    `json:"longitude"`
}

var Nodes []Node

func (s *Server) defaultResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var messageContent = "Welcome to network management configurtion !"
	json.NewEncoder(w).Encode(messageContent)
}

/* swagger:route GET /getAllNodes allNodes
This returns all nodes stored


produces:
- application/json
responses:
	200: Node

*/
func (s *Server) getAllNodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Nodes)
}

func (s *Server) getNodeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(Node{Id: -1, Name: "Unknown Node Requested"})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for _, node := range Nodes {
		if node.Id == id {
			json.NewEncoder(w).Encode(node)
		}
	}
}

/* swagger:route POST /node addNewNode
This is used to insert a new node, the node is part of the request body


Consumes:
- application/json
produces:
- application/json
responses:
	200: Node
	400: Node

*/

func (s *Server) addNewNode(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var node Node
	err := json.Unmarshal(reqBody, &node)
	if err != nil {
		fmt.Println("error:", err)
	}

	id, err := generateID(8) // generate an 8-byte ID
	if err != nil {
		log.Fatal(err)
	}
	node.Id = int(binary.BigEndian.Uint32(id))
	Nodes = append(Nodes, node)
	json.NewEncoder(w).Encode(node)
}

/* swagger:route DELETE /node{id} deleteNode
This is used to delete the node whose id is provided as the URL parameter id


Consumes:
- application/json
produces:
- application/json
responses:
	200:
	400: Node

*/
func (s *Server) deleteNode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(Node{Id: -1, Name: "Unknown Node Requested"})
	}
	for index, node := range Nodes {
		if node.Id == id {
			Nodes = append(Nodes[:index], Nodes[index+1:]...)
		}
	}
}

/* swagger:route PUT /node/{id} updateNode
This is used to update the node whose id is provided as the URL parameter id


Consumes:
- application/json
produces:
- application/json
responses:
	200: Node
	400: Node

*/

func (s *Server) updateNode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(Node{Id: -1, Name: "Unknown Node Requested"})
	}

	var index int

	for i, node := range Nodes {
		if node.Id == id {
			reqBody, _ := io.ReadAll(r.Body)
			err := json.Unmarshal(reqBody, &node)
			if err != nil {
				fmt.Println("error:", err)
			}
			index = i

			Nodes[index].Name = node.Name
			Nodes[index].Location = node.Location
			Nodes[index].Latitude = node.Latitude
			Nodes[index].Longitude = node.Longitude

			break

		}
	}
}

func generateID(length int) ([]byte, error) {
	id := make([]byte, length)
	_, err := rand.Read(id)
	if err != nil {
		return nil, err
	}
	return id, nil
}
