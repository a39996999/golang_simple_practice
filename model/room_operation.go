package model

import "chatroom/utils"

type Room struct {
	Room_id     int    `json:"Room_id"`
	Owner_id    int    `json:"Owner_id"`
	Owner_name  string `json:"Owner_name"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Create_time string `json:"Create_time"`
}

func CreateRoom(name string, owner_id int, description string) error {
	time := utils.GetCurrentTime()
	createSql := "insert into room(name, owner_id, description, create_time) values(?, ?, ?, ?)"
	_, err := db.Exec(createSql, name, owner_id, description, time)

	if err != nil {
		return err
	}
	return nil
}

func GetOwnerId(roomId int) (int, error) {
	querySql := "select owner_id from room where id = ?"
	var ownerId int
	err := db.QueryRow(querySql, roomId).Scan(&ownerId)
	if err != nil {
		return ownerId, err
	}
	return ownerId, err
}

func DeleteRoom(roomId int) error {
	deleteSql := "delete from room where id = ?"

	_, err := db.Exec(deleteSql, roomId)
	if err != nil {
		return err
	}
	return nil
}

func GetRoomList() ([]Room, error) {
	var roomlist []Room
	querySql := `select r.*, u.username from chating.room as r
				inner join chating.users as u on r.owner_id = u.id`
	row, err := db.Query(querySql)
	if err != nil {
		return roomlist, err
	}
	for row.Next() {
		var room Room
		if err := row.Scan(&room.Room_id, &room.Name, &room.Owner_id, &room.Description, &room.Create_time, &room.Owner_name); err != nil {
			return roomlist, err
		}
		roomlist = append(roomlist, room)
	}

	return roomlist, nil
}
