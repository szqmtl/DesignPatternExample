package rshu.components;

import java.util.HashMap;

import lombok.ToString;

@ToString
public class Maze {
    private HashMap<Integer, Room> rooms = new HashMap<>();

    public void addRoom(Room room){
        rooms.put(room.getRoomNumber(), room);
    }

    public Room getRoom(int roomNum){
        return rooms.get(roomNum);
    }
}