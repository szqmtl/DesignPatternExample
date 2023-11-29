package rshu.components.doors;

import lombok.ToString;
import rshu.components.rooms.Room;

@ToString
public abstract class DoorNeedingSpell extends Door {
    public DoorNeedingSpell(Room room1, Room room2) {
        super(room1, room2);
    }
}
