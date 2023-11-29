package rshu.components.factories;

import rshu.components.Maze;
import rshu.components.doors.Door;
import rshu.components.rooms.Room;
import rshu.components.walls.Wall;

public class MazeFactory {
    public Maze makeMaze(){
        return new Maze();
    }

    public Wall makeWall(){
        return new Wall() {
            @Override
            public void enter() {
                System.out.println("crossing a walls");
            }
        };
    }

    public Room makeRoom(int n){
        return new Room(n) {
            @Override
            public void enter() {
                System.out.println("entering a room " + this.getRoomNumber());
            }
        };
    }

    public Door makeDoor(Room r1, Room r2){
        return new Door(r1, r2) {
            public void enter() {
                System.out.println("switch a room " + this);
            }
            
        };
    }
}
