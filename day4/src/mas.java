import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;

public class mas {

    static char[][] result;
    static StringBuilder resultLeft;
    static StringBuilder resultRight;
    static int xmasCount = 0;
    public static void main(String[] args) throws IOException {
        File file = new File("resources/input_4.txt");
        BufferedReader br = new BufferedReader(new FileReader(file));
        String st;
        ArrayList<String> lst = new ArrayList<>();
        while((st=br.readLine()) != null){
            lst.add(st);
        }

        result = new char[lst.size()][lst.getFirst().length()];
        for(int i = 0; i < lst.size(); i++) {
            char[] charArray= lst.get(i).toCharArray();
            result[i] = charArray;
        }
        //System.out.println(Arrays.deepToString(result));
        int xCount = 0;
        for(int i = 0; i < result.length; i++){
            for(int j = 0; j < result[i].length; j++){
                if ('A' == (result[i][j])){
                    checkForX_MAS(i, j);
                }
            }
        }
        System.out.println(xmasCount);

    }

    private static void checkForX_MAS(int first, int second) {
        resultLeft = new StringBuilder();
        resultRight = new StringBuilder();

        if(first - 1 < 0 || second - 1 < 0 || first + 1 >= result.length || second + 1 >= result[second].length){
           return;
        } else {
            resultLeft.append(result[first -1][second -1]);
            resultLeft.append(result[first][second]);
            resultLeft.append(result[first + 1][second + 1]);

            resultRight.append(result[first -1][second +1]);
            resultRight.append(result[first][second]);
            resultRight.append(result[first + 1][second - 1]);
        }
        if ((resultLeft.toString().equals("MAS") || resultLeft.toString().equals("SAM")) && (resultRight.toString().equals("MAS") || resultRight.toString().equals("SAM"))){
            xmasCount++;
        }
        resultLeft.setLength(0);
        resultRight.setLength(0);


    }


}
