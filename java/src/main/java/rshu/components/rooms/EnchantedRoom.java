package rshu.components.rooms;

import lombok.ToString;
import rshu.components.Spell;

@ToString
public abstract class EnchantedRoom extends Room {
    private Spell spell;
    public EnchantedRoom(int n, Spell s){
        super(n);
        this.spell = s;
    }
    
}
