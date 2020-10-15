package practice;

import java.util.Scanner;

public class SimpleInterest {
	
	public static Scanner sc = new Scanner(System.in);

	public static void main(String[] args) {
		
		double principle, rate, time, interest;
		
		System.out.println("Enter Principle amount :");
		principle = sc.nextDouble();
		
		System.out.println("Enter Rate: ");
		rate = sc.nextDouble();
		
		System.out.println("Enter time period:");
		time = sc.nextDouble();
		
		System.out.println("Your Simple interest is "+ ((principle * rate * time )/ 100));
	}

}
