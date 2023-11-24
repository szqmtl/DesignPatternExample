package rshu.components;

public class StandardMazeBuilder extends MazeBuilder {
    public void buildRoom(int n){
        var m = this.getMaze();
        if(m.getRoom(n) == null){
            Room r = new Room(n) {
                @Override
                public void enter() {
                    return;
                }
            };

            m.addRoom(r);

            r.setSide(Direction.North, new Wall() {
                @Override
                public void enter() {}
            });
            r.setSide(Direction.South, new Wall() {
                @Override
                public void enter() {}
            });
            r.setSide(Direction.East, new Wall() {
                @Override
                public void enter() {}
            });
            r.setSide(Direction.West, new Wall() {
                @Override
                public void enter() {}
            });
        }
    }
    public void buildDoor(int from, int to){
        var maze = this.getMaze();
        var r1 = maze.getRoom(from);
        var r2 = maze.getRoom(to);

        Door d = new Door(r1, r2) {
            @Override
            public void enter() { }
        };

        r1.setSide(commonWall(r1, r2), d);
        r2.setSide(commonWall(r2, r1), d);
    }

    private Direction commonWall(Room r1, Room r2){
        if(r1.getRoomNumber() > r2.getRoomNumber()){
            return Direction.West;
        }else{
            return Direction.East;
        }
    }
}
