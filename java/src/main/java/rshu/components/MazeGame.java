package rshu.components;

public class MazeGame {
    public Maze createMaze(MazeFactory factory){
        Maze aMaze = factory.makeMaze();
        Room r1 = factory.makeRoom(1);
        Room r2 = factory.makeRoom(2);

        Door aDoor = factory.makeDoor(r1, r2);

        aMaze.addRoom(r1);
        aMaze.addRoom(r2);

        r1.setSide(Direction.North, factory.makeWall());
        r1.setSide(Direction.East, aDoor);
        r1.setSide(Direction.South, factory.makeWall());
        r1.setSide(Direction.West, factory.makeWall());

        r2.setSide(Direction.North, factory.makeWall());
        r2.setSide(Direction.East, factory.makeWall());
        r2.setSide(Direction.South, factory.makeWall());
        r2.setSide(Direction.West, aDoor);

        return aMaze;
    }

    public Maze buildeMaze(MazeBuilder builder){
        builder.buildMaze();
        builder.buildRoom(1);
        builder.buildRoom(2);
        builder.buildDoor(1, 2);

        return builder.getMaze();
    }
}
