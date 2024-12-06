import java.io.*;
import java.util.ArrayList;
import java.util.Arrays;

public class Main {
    static char[][] result;
    static StringBuilder resultString;
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
                if ('X' == (result[i][j])){
                    checkForXMAS(i, j);
                }
            }
        }
        System.out.println(xmasCount);

    }

    private static void checkForXMAS(int first, int second) {
        resultString = new StringBuilder();
        // Check result forwards
        for(int i = 0; i < 4; i++){
            if (second + i >= result[first].length){
                break;
            } else {
                resultString.append(result[first][second + i]);
            }
        }
        checkResult();

       // check result backwards
        for(int i = 0; i < 4; i++){
            if (second - i < 0){
                break;
            } else {
                resultString.append(result[first][second - i]);
            }
        }
        checkResult();

        // check result up
        for(int i = 0; i < 4; i++){
            if (first - i < 0){
                break;
            } else {
                resultString.append(result[first - i][second]);
            }
        }
        checkResult();

        // check result down
        for(int i = 0; i < 4; i++){
            if (first + i >= result.length){
                break;
            } else {
                resultString.append(result[first + i][second]);
            }
        }
        checkResult();

        // check result diagonal left up
        for(int i = 0; i < 4; i++){
            if (first - i < 0 || second - i < 0){
                break;
            } else {
                resultString.append(result[first - i][second - i]);
            }
        }
        checkResult();

        // check result diagonal right up
        for(int i = 0; i < 4; i++){
            if (first - i < 0 || second + i >= result[second].length){
                break;
            } else {
                resultString.append(result[first - i][second + i]);
            }
        }
        checkResult();

        // check result diagonal left down
        for(int i = 0; i < 4; i++){
            if (second - i < 0 || first + i >= result.length){
                break;
            } else {
                resultString.append(result[first + i][second - i]);
            }
        }
        checkResult();

        // check result diagonal right down
        for(int i = 0; i < 4; i++){
            if (first + i >= result.length || second + i >= result[second].length){
                break;
            } else {
                resultString.append(result[first + i][second + i]);
            }
        }
        checkResult();


    }

    private static void checkResult() {
        if (resultString.toString().equals("XMAS")){
            xmasCount++;
        }
        resultString.setLength(0);
    }
}