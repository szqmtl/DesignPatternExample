package rshu.components;

import java.util.HashMap;
import java.util.Map;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

@ToString(doNotUseGetters = true)
public abstract class Room implements MapSite {
    @Getter @Setter
    private int roomNumber;
    private final Map<Direction, MapSite> sides = new HashMap<>();

    public Room(int roomNo){
        this.roomNumber = roomNo;
    }

    public MapSite getSide(Direction direction){
        return sides.get(direction);
    }

    public void setSide(Direction direction, MapSite side){
        sides.put(direction, side);
    }
}
