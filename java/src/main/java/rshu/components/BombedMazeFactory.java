package rshu.components;

public class BombedMazeFactory extends MazeFactory {
    public Wall makeWall(){
        return new BombedWall() {
            @Override
            public void enter() {
                System.out.println("crossing bombed wall");
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
