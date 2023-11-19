package rshu.components;

import lombok.ToString;

@ToString
public class Spell {
    private String magicWord;

    public Spell(String words){
        this.magicWord = words;
    }

    public String get(){
        return magicWord;
    }
}
