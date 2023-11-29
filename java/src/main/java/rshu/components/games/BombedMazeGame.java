package rshu.components.games;


import rshu.components.rooms.Room;
import rshu.components.rooms.RoomWithABomb;
import rshu.components.walls.BombedWall;
import rshu.components.walls.Wall;

public class BombedMazeGame extends MazeGame {
    public Wall makeWall(){
        return new BombedWall() {
            @Override
            public void enter() {
            }
        };
    }

    public Room makeRoom(int n){
        return new RoomWithABomb(n) {
            @Override
            public void enter() {
            }
        };
    }
}
