package rshu.components.factories;

import rshu.components.Spell;
import rshu.components.doors.Door;
import rshu.components.doors.DoorNeedingSpell;
import rshu.components.rooms.EnchantedRoom;
import rshu.components.rooms.Room;

public class EnchantedMazeFactory extends MazeFactory {
    public EnchantedMazeFactory(){}

    public Room makeRoom(int n){
        return new EnchantedRoom(n, castSpell()){
            @Override
            public void enter() {
                System.out.println("entering enchanted room " + this);
            }
        };
    }

    public Door makeDoor(Room r1, Room r2){
        return new DoorNeedingSpell(r1, r2) {
            @Override
            public void enter() {
                System.out.println("crossing door needing spell between " + this.getRoom1() + ", " + this.getRoom2());
            }
            
        };
    }
    private Spell castSpell() {
        return new Spell("enchanted spell");
    }
}
