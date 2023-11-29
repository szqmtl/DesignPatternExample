package rshu.components.factories;

import rshu.components.rooms.Room;
import rshu.components.rooms.RoomWithABomb;
import rshu.components.walls.BombedWall;
import rshu.components.walls.Wall;

public class BombedMazeFactory extends MazeFactory {
    public Wall makeWall(){
        return new BombedWall() {
            @Override
            public void enter() {
                System.out.println("crossing bombed walls");
            }
        };
    }

    public Room makeRoom(int n){
        return new RoomWithABomb(n) {
            @Override
            public void enter() {
                System.out.println("entering a room with a bomb");
            }
        };
    }
    
}
