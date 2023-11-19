package rshu.components;

import java.util.HashMap;
import java.util.Map;

import lombok.ToString;

@ToString(doNotUseGetters = true)
public abstract class Room implements MapSite {
    private int roomNumber;
    private Map<Direction, MapSite> sides = new HashMap<>();

    public Room(int roomNo){
        this.roomNumber = roomNo;
    }

    public int getRoomNumber(){
        return this.roomNumber;
    }

    public MapSite getSide(Direction direction){
        return sides.get(direction);
    }

    public void setSide(Direction direction, MapSite side){
        sides.put(direction, side);
    }
}
