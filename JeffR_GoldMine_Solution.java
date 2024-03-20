

public class JeffR_GoldMine
{
    public static void dig( int[][] mine ) 
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
                        ?   ( c == 0 ? mine[r][c] : 0 ) 
                            + Math.max( Math.max( nodes[r-1][c+1].up, nodes[r-1][c+1].rt ), nodes[r-1][c+1].dn )
                        : 0
                        ;

                    nodes[r][c].rt += 
                            ( c == 0 ? mine[r][c] : 0 ) 
                            + Math.max( Math.max( nodes[r  ][c+1].up, nodes[r  ][c+1].rt ), nodes[r  ][c+1].dn )
                          ;

                    nodes[r][c].dn +=
                        r < lr
                        ?   ( c == 0 ? mine[r][c] : 0 )     
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

        // dumpNodes(mine,nodes);

        int maxGold = -1;
        int maxPaths = 0;
        for( int r = 0; r < rows; r++ ) {
            int maxStep = Math.max( Math.max( nodes[r][0].up, nodes[r][0].rt ), nodes[r][0].dn );
            if( maxStep > maxGold ) {
                maxGold = maxStep;
                maxPaths = 1;
            }
            else if( maxStep == maxGold ) {
                maxPaths++;
            }
        }

        if( maxGold == 0 )
            System.out.println( "The mine is devoid of gold??");
        else
            System.out.println("Max gold \u001b[1m" + maxGold + "\u001b[0m in " + maxPaths + " path(s).");

    }



    static class node 
    {
        int up; // up-and-right
        int rt; // straight-right
        int dn; // down-and-right
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

        // int[][] mineAllOnes = { { 1, 1, 1}, {1, 1, 1}, {1, 1, 1}};
        // dig( mineAllOnes );

        int[][] mineSample = 
            { { 0, 0, 0, 9 },
            { 0, 0, 0, 0 },
            { 0, 0, 0, 0 },
            { 1, 1, 1, 8 }
            };

        dig(mineSample);
    }
}