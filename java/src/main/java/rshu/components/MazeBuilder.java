package rshu.components;

public class MazeBuilder{
    private Maze maze;
    public void buildMaze(){
        this.maze = new Maze();
    }

    public void buildRoom(int n){
        this.maze.addRoom(new Room(n) {
            @Override
            public void enter() {
                return;
            }
        });
    }

    public void buildDoor(int from, int to){
        Room r1 = maze.getRoom(from);
        Room r2 = maze.getRoom(to);

        Door d = new Door(r1, r2) {
            @Override
            public void enter() {
                return;
            }
        };

        r1.setSide(Direction.South, d);
        r2.setSide(Direction.North, d);
    }

    public Maze getMaze(){
        return maze;
    }
}