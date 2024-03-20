import java.util.List;
import java.util.ArrayList;

public class JeffR_GoldMine
{
    static boolean FIND_PATHS = true;

    static boolean USE_ANSI = true;
    static String ANSI_GOLD_PREFIX = USE_ANSI ? "\u001b[1m\u001b[103m\u001b[91m" : "";
    static String ANSI_GOLD_SUFFIX = USE_ANSI ? "\u001b[0m" : "";
    
    static void dig( int[][] mine ) 
    {
        int rows = mine.length;
        int lr = rows - 1;
        int cols = mine[0].length;
        int lc = cols - 1;

        node[][] nodes = new node[rows][cols];

        for( int r = 0; r < rows; r++ ) {
            for( int c = 0; c < cols; c++ ) {

                nodes[r][c] = new node();

                if( c < lc ) {
                    nodes[r][c].up = r >  0  ? mine[r-1][c+1] : -1;
                    nodes[r][c].rt = r >= 0  ? mine[r  ][c+1] : -1;
                    nodes[r][c].dn = r <  lr ? mine[r+1][c+1] : -1;
                }
                else {
                    nodes[r][c].up = nodes[r][c].rt = nodes[r][c].dn = 0;
                }
            }
        }

        // dumpNodes(mine,nodes);

        for( int c = cols; c-- > 0; ) {
            for( int r = rows; r-- > 0; ) {
                if( c < lc) {
                    nodes[r][c].up += 
                        r > 0 
                        ?   0 // ( c == 0 ? mine[r][c] : 0 ) 
                            + Math.max( Math.max( nodes[r-1][c+1].up, nodes[r-1][c+1].rt ), nodes[r-1][c+1].dn )
                        : 0
                        ;

                    nodes[r][c].rt += 
                            0 // ( c == 0 ? mine[r][c] : 0 ) 
                            + Math.max( Math.max( nodes[r  ][c+1].up, nodes[r  ][c+1].rt ), nodes[r  ][c+1].dn )
                          ;

                    nodes[r][c].dn +=
                        r < lr
                        ?   0 // ( c == 0 ? mine[r][c] : 0 )     
                            +  Math.max( Math.max( nodes[r+1][c+1].up, nodes[r+1][c+1].rt ), nodes[r+1][c+1].dn )  
                        : 0
                        ;                        

                }
                else {
                    // nodes
                }

                /*
                System.out.println( "r=" + r + " c=" + c + ":");

                dumpNodes(mine,nodes);
                */
            }
        }

        /*
        for( int r = 0; r < rows; r++ ) {
            for( int c = 0; c < cols; c++ ) {
                if( r > 0 )
                    nodes[r][c].up += mine[r][c];
                
                nodes[r][c].rt += mine[r][c];

                if( r < lr)
                    nodes[r][c].dn += mine[r][c];
            }
        }
        */

        dumpNodes(mine,nodes);

        int maxGold = -1;
        List<List<coord>> paths = new ArrayList<>();
        for( int r = 0; r < rows; r++ ) {
            int maxStep = mine[r][0] + Math.max( Math.max( nodes[r][0].up, nodes[r][0].rt ), nodes[r][0].dn );
            if( maxStep > maxGold ) {
                maxGold = maxStep;
                if( FIND_PATHS ) {
                    paths.clear();
                    paths.add(new ArrayList<>());
                    paths.get(0).add(new coord(r,0));
                }
            }
            else if( maxStep == maxGold ) {
                if( FIND_PATHS ) {
                    paths.add(new ArrayList<>());
                    paths.getLast().add(new coord(r,0));
                }
            }
        }

        // dumpNodes(mine,nodes);

        int maxPaths = 0;
        if( FIND_PATHS ) {
            if( maxGold > 0 ) {
                maxPaths = 0;
                for( int p = 0; p < paths.size(); p++ ) {
                    List<coord> path = paths.get(p);
                    int sum = 0;
                    for( int c = 0; c < cols; c++ ) {
                        coord coords = path.get(c);
                        if( c < path.size()) {
                            sum += mine[coords.r][coords.c];    
                        }
                        if( c + 1 < cols && c + 1 == path.size()) {
                            node nOde = nodes[coords.r][coords.c];
                            int maxStep = Math.max( Math.max( nOde.up, nOde.rt ), nOde.dn );
                            boolean up = maxStep == nOde.up;
                            boolean rt = maxStep == nOde.rt;
                            boolean dn = maxStep == nOde.dn;
                            int subPaths = 0;
                            subPaths += up ? 1 : 0;
                            subPaths += rt ? 1 : 0;
                            subPaths += dn ? 1 : 0;
                            if( subPaths > 1 ) {
                                if( up ) {
                                    if( rt ) {
                                        List<coord> newPath = new ArrayList<>(path);
                                        newPath.add(new coord(coords.r ,coords.c+1));
                                        paths.add(newPath);
                                    }

                                    if( dn ) {
                                        List<coord> newPath = new ArrayList<>(path);
                                        newPath.add(new coord(coords.r+1,coords.c+1));
                                        paths.add(newPath);

                                    }

                                    path.add(new coord(coords.r-1,coords.c+1));
                                }
                                else if( rt ) {
                                    // dn:
                                    List<coord> newPath = new ArrayList<>(path);
                                    newPath.add(new coord(coords.r+1,coords.c+1));
                                    paths.add(newPath);

                                    path.add(new coord(coords.r ,coords.c+1));


                                }
                            }
                            else {
                                if( up )
                                    path.add(new coord(coords.r-1,coords.c+1));
                                else if( rt ) 
                                    path.add(new coord(coords.r ,coords.c+1));
                                else // dn
                                    path.add(new coord(coords.r+1,coords.c+1));
                            }
                        }
                    }
                    assert(sum == maxGold);
                }
            }
        }

        maxPaths = paths.size();

        if( maxGold == 0 )
            System.out.println( "The mine is devoid of gold??");
        else
            System.out.println("Max gold " 
                + ANSI_GOLD_PREFIX + maxGold + ANSI_GOLD_SUFFIX 
                + ( FIND_PATHS ? " in " + maxPaths + " path(s)." : ""));

        if( FIND_PATHS )                
            for( int p = 0; p < maxPaths; p++ ) {
                System.out.print("Path #" + p + ":");
                List<coord> path = paths.get(p);
                for( int s = 0; s < path.size(); s++ ) {
                    coord rc = path.get(s);
                    System.out.print( " [" + rc.r + ", " + rc.c + "]");
                }
                System.out.println();
                for( int r = 0; r < rows; r++ ) {
                    for( int c = 0; c < cols; c++ ) {
                        int value = mine[r][c];
                        if( path.contains(new coord(r,c)))
                            System.out.printf( "%s%5d%s, ", ANSI_GOLD_PREFIX, value, ANSI_GOLD_SUFFIX );
                        else
                            System.out.printf("%5d, ", value);
                    }
                    System.out.println();
                }
                System.out.println();
                    
            }

    }



    static class node 
    {
        int up; // up-and-right
        int rt; // straight-right
        int dn; // down-and-right
    }

    static class coord
    {
        int r, c;
        coord(int r, int c) {
            this.r = r;
            this.c = c;
        }

        @Override
        public boolean equals(Object that) {
            if( that instanceof coord ) {
                coord thatCoord = (coord)that;
                return thatCoord.r == this.r && thatCoord.c == this.c;
            }

            return false;
        }

        @Override
        public int hashCode() {
            return Integer.valueOf(Integer.valueOf(r) * 10000 + Integer.valueOf(c)).hashCode();
        }
    }

    static void dumpNodes( int[][] mine, node[][] nodes ) {
        System.out.println();
        int rows = mine.length;
        int cols = mine[0].length;
        int lc = cols - 1;
        for( int r = 0; r < rows; r++ ) {
            System.out.print("{ ");
            for( int c = 0; c < cols; c++ ) {
                String dir = "|";
                String delim = ", ";
                int max = Math.max( Math.max( nodes[r][c].up, nodes[r][c].rt), nodes[r][c].dn );
                boolean up = max == nodes[r][c].up;
                boolean rt = max == nodes[r][c].rt;
                boolean dn = max == nodes[r][c].dn;
                if( c < lc )
                {

                }
                else {
                    delim = "";
                    up = rt = dn = false;
                }
                System.out.printf( "%3d [%3d %3d %3d: %5d %c%c%c]%s", mine[r][c], nodes[r][c].up, nodes[r][c].rt, nodes[r][c].dn, max, 
                    up ? '/' : ' ', 
                    rt ? '-' : ' ', 
                    dn ? '\\' : ' ', 
                    delim);
            }
            System.out.println(" }");
        }
        System.out.println();

    }

    public static void main(String[] args) {

        int[][] mineAllOnes = 
        { 
            { 1, 1, 1 }, 
            { 1, 1, 1 }, 
            { 1, 1, 2 }
        };
        dig( mineAllOnes );

        int[][] mineSample = 
            { 
            { 0, 0, 0, 10 },
            { 0, 0, 0, 9 },
            { 0, 0, 0, 8 },
            { 1, 1, 1, 8 }
            };

        dig(mineSample);

        int[][] justOne = 
            {
                {0,0,1},
                {0,0,0},
                {1,0,0}
            };

        dig(justOne);
    }
}