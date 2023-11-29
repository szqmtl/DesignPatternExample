package rshu.components.builders;

import lombok.Getter;
import rshu.components.builders.MazeBuilder;

@Getter
public class CountingMazeBuilder extends MazeBuilder {
    private int rooms = 0;
    private int doors = 0;

    public void buildRoom(int n){
        this.rooms++;
    }
    public void buildDoor(int from, int to){
        this.doors++;
    }
}
