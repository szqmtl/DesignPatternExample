package rshu.components.doors;

import lombok.ToString;
import rshu.components.MapSite;
import rshu.components.rooms.Room;

@ToString(doNotUseGetters = true)
public abstract class Door implements MapSite {
    @ToString.Exclude
    private Room room1;
    @ToString.Exclude
    private Room room2;
    private boolean isOpen;
    
    public Door(Room room1, Room room2){
        this.room1 = room1;
        this.room2 = room2;
    }

    protected Room getRoom1(){
        return room1;
    }

    protected Room getRoom2(){
        return room2;
    }
}
