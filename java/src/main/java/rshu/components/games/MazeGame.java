package rshu.components.games;

import rshu.components.*;
import rshu.components.doors.Door;
import rshu.components.builders.MazeBuilder;
import rshu.components.factories.MazeFactory;
import rshu.components.rooms.Room;
import rshu.components.walls.Wall;

public class MazeGame {
    // For Abstract factory
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

    // for Builder
    public Maze buildeMaze(MazeBuilder builder){
        builder.buildMaze();
        builder.buildRoom(1);
        builder.buildRoom(2);
        builder.buildDoor(1, 2);

        return builder.getMaze();
    }

    // For factory method
    public Maze createMaze(){
        Maze m = makeMaze();

        var r1 = makeRoom(1);
        var r2 = makeRoom(2);
        var d = makeDoor(r1, r2);

        m.addRoom(r1);
        m.addRoom(r2);

        r1.setSide(Direction.North, makeWall());
        r1.setSide(Direction.East, d);
        r1.setSide(Direction.South, makeWall());
        r1.setSide(Direction.West, makeWall());

        r2.setSide(Direction.North, makeWall());
        r2.setSide(Direction.East, makeWall());
        r2.setSide(Direction.South, makeWall());
        r2.setSide(Direction.West, d);

        return m;
    }

    public Maze makeMaze(){
        return new Maze();
    }

    public Room makeRoom(int n){
        return new Room(n) {
            @Override
            public void enter() {
            }
        };
    }
    public Wall makeWall(){
        return new Wall() {
            @Override
            public void enter() {
            }
        };
    }
    public Door makeDoor(Room r1, Room r2){
        return new Door(r1, r2) {
            @Override
            public void enter() {

            }
        };
    }
}
