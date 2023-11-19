package rshu.components;

import lombok.ToString;

@ToString
public abstract class RoomWithABomb extends Room {

    public RoomWithABomb(int roomNo) {
        super(roomNo);
    }
    
}
