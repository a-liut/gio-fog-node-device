/*
 * Devices service
 *
 * Microservice for managing Giò Plants devices
 *
 * API version: 1.0.0
 * Contact: andrea.liut@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package repository

import (
	"gio-device-ms/pkg/model"
)

type RoomRepository struct {
	rooms map[string]*model.Room
}

func (r *RoomRepository) Get(id string) (*model.Room, error) {
	room, _ := r.rooms[id]

	return room, nil
}

func (r *RoomRepository) GetAll() ([]*model.Room, error) {
	res := make([]*model.Room, len(r.rooms))

	i := 0
	for _, d := range r.rooms {
		res[i] = d
		i++
	}

	return res, nil
}

func (r *RoomRepository) Insert(room *model.Room) (*model.Room, error) {
	room.ID = newID()

	r.rooms[room.ID] = room

	return room, nil
}

func (r *RoomRepository) GetByName(name string) (*model.Room, error) {
	for _, room := range r.rooms {
		if room.Name == name {
			return room, nil
		}
	}

	return nil, nil
}

var roomRepository *RoomRepository

func NewRoomRepository() (*RoomRepository, error) {
	if roomRepository == nil {
		roomRepository = &RoomRepository{make(map[string]*model.Room)}
	}

	return roomRepository, nil
}
