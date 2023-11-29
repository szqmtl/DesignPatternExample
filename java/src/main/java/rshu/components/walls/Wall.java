package rshu.components.walls;

import lombok.ToString;
import rshu.components.MapSite;

@ToString(doNotUseGetters = true)
public abstract class Wall implements MapSite {
    private final String name = "Wall";
}
