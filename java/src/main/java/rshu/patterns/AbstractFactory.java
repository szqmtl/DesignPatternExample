package rshu.patterns;

import rshu.components.BombedMazeFactory;
import rshu.components.EnchantedMazeFactory;
import rshu.components.Maze;
import rshu.components.MazeFactory;
import rshu.components.MazeGame;

public class AbstractFactory {
    public static void main(String[] args){
        MazeGame game = new MazeGame();

        MazeFactory f1 = new MazeFactory();
        Maze m = game.createMaze(f1);

        System.out.println("maze 1: " + m);

        MazeFactory f2 = new EnchantedMazeFactory();
        m = game.createMaze(f2);

        System.out.println("maze 2: " + m);

        MazeFactory f3 = new BombedMazeFactory();
        m = game.createMaze(f3);

        System.out.println("maze 3: " + m);

    }
}
