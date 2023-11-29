package rshu.components.games;

import rshu.components.*;
import rshu.components.doors.Door;
import rshu.components.doors.DoorNeedingSpell;
import rshu.components.rooms.EnchantedRoom;
import rshu.components.rooms.Room;

public class EnchantedMazeGame extends MazeGame {
    public Room makeRoom(int n){
        return new EnchantedRoom(n, new Spell("enchanted words")) {
            @Override
            public void enter() {

            }
        };
    }
    public Door makeDoor(Room r1, Room r2){
        return new DoorNeedingSpell(r1, r2) {
            @Override
            public void enter() {

            }
        };
    }
}
