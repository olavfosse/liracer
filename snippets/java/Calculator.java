package conditionalStatement;
import java.io.*;

public class Calculator {

	public static void main(String[] args) throws IOException {

		
		BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
		System.out.println("Enter First number");
		double num1 = Double.parseDouble(reader.readLine());
		System.out.println("Enter second nuber");
		double num2 = Double.parseDouble(reader.readLine());
		System.out.println("Enter + for Addition \n"+"- for Subraction\n"+"* for Multiplication\n"+"/ for Division");
		String input = reader.readLine();
		char operation = input.charAt(0);
		Condition(num1,num2,operation);
	}

	private static void Condition(double num1, double num2,char operation) {
		Operation oop = new Operation();

		switch(operation) {
		
		case '+' :System.out.println("The sum of "+num1 +" and "+num2+" is "+DoAddition(num1,num2));
			break;
		
		case '-' :System.out.println("The sum of "+num1 +" and "+num2+" is "+DoSubract(num1,num2));
		break;
		
		case '*' :System.out.println("The sum of "+num1 +" and "+num2+" is "+DoMultiply(num1,num2));
		break;
		
		case '/' :System.out.println("The sum of "+num1 +" and "+num2+" is "+DoDivision(num1,num2));
		break;
		}
		
	}

	static double DoDivision(double a,double b) {return a/b;}
	static double DoMultiply(double a,double b) {return a*b;}
	static double DoAddition(double a,double b) {return a+b;}
	static double DoSubract(double a,double b) {return a-b;}
}