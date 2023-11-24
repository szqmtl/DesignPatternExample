package rshu.patterns;

import rshu.components.CountingMazeBuilder;
import rshu.components.MazeBuilder;
import rshu.components.MazeGame;
import rshu.components.StandardMazeBuilder;

public class Builder {
    public static void main(String[] args){
        var game = new MazeGame();

        var b1 = new MazeBuilder();
        System.out.printf("maze 1: %s\n", game.buildeMaze(b1));

        var b2 = new StandardMazeBuilder();
        System.out.printf("maze 2: %s\n", game.buildeMaze(b2));

        var b3 = new CountingMazeBuilder();
        game.buildeMaze(b3);
        System.out.printf("maze 3 - rooms: %d, doors: %d\n", b3.getRooms(), b3.getDoors());

    }
}
