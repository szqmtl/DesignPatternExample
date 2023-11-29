package rshu.patterns.creational;

import rshu.components.Maze;
import rshu.components.games.BombedMazeGame;
import rshu.components.games.EnchantedMazeGame;
import rshu.components.games.MazeGame;

public class FactoryMethod {
    public static void main(String[] args) {
        Maze m1 = new MazeGame().createMaze();
        System.out.println("maze 1: " + m1);

        Maze m2 = new EnchantedMazeGame().createMaze();
        System.out.println("maze 2: " + m2);

        Maze m3 = new BombedMazeGame().createMaze();
        System.out.println("maze 3: " + m3);

    }

}
