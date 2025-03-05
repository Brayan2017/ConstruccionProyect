public class Main {
    public static void main(String[] args) {
        AveVoladora gorrion = new Gorrion();
        gorrion.volar();


        AveNoVoladora pinguino = new Pinguino();
        //pinguino.volar() - este codigo no compilara
    }
}