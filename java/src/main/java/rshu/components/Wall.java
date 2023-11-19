package rshu.components;

import lombok.ToString;

@ToString(doNotUseGetters = true)
public abstract class Wall implements MapSite {
    private String name = "Wall";
}
