package rshu.components;

import lombok.ToString;

@ToString
public abstract class EnchantedRoom extends Room {
    private Spell spell;
    public EnchantedRoom(int n, Spell s){
        super(n);
        this.spell = s;
    }
    
}
