class main {
    public static void main(String[] args) {
        File file = new File("../input_4.txt");
        BufferedReader br = new BufferedReader(new FileReader(file));
        String st;
        while((st=br.readLine()) != null){
            System.out.println(st);
        }
    }
}
