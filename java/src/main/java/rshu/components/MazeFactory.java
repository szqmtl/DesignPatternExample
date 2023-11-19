package rshu.components;

public class MazeFactory {
    public Maze makeMaze(){
        return new Maze();
    }

    public Wall makeWall(){
        return new Wall() {
            @Override
            public void enter() {
                System.out.println("crossing a wall");
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
            @Override
            public void enter() {
                System.out.println("switch a room " + this);
            }
            
        };
    }
}
