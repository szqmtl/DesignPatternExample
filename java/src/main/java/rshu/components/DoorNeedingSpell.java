package rshu.components;

import lombok.ToString;

@ToString
public abstract class DoorNeedingSpell extends Door {
    public DoorNeedingSpell(Room room1, Room room2) {
        super(room1, room2);
    }
}
