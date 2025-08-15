
#include <stdio.h>
#include <limits.h>

//_____________________________________________________________________

void playWithCommandLineArgs(int argc, char *argv[]) {
	for ( int i = 0 ; i < argc ; i++ ) {
		printf("\n Argument Value At %d = %s", i, argv[i] );
	}
}

//_____________________________________________________________________
// DESIGN 01 : BAD CODE

int sum( int x, int y ) {
	return x + y;
}

void playWithSum() {
	int a, b, result = 0;

	a = 2147483647;
	b = 1;

	result = sum( a, b );
	printf("\nResult : %d", result ); // 2147483648

	a = -2147483648;
	b = -10;

	result = sum( a, b );
	printf("\nResult : %d", result ); // -2147483658
}

// Function : playWithSum
// Result : -2147483648
// Result : 2147483638

//_____________________________________________________________________

// GOOD DESIGN
//		DRIVEN BY FUNAMENTALS
int summation(int x, int y) {
	int result = 0;

	if ( ( y > 0 && x > INT_MAX - y ) || ( y < 0 && x < INT_MIN - y ) ) {
		printf("Cant't Calcualte Sum");
	} else {
		result = x + y;
		return result;
	}
}

//_____________________________________________________________________

void playWithTypes() {
	// char b = 900;
	
	char b = 90;
	printf("\nValue %d", b);

	// In C By Default
	//		int Is Signed By Default
	//			[ -32768, 32767 ]
	//		usinged int Is UnSigned int
	//			[ 0, 65535 ]
	
	int i = 100;
	unsigned ui = 900;
}

//_____________________________________________________________________

// In C
//		Strings Are NULL \0 Terminated 
void playWithString() {
	// In This Case NULL Character Is Stored As Last Character
	//		Size Of String Will Be Number Of Characters In String + 1
	//		+1 Is For Storing \0 Character
	// Compiler Will Calculate Size Of char Array With Above Logic
	char greeting[] = "Good Morning!";
	char username[] = "amarjit@icicibank.com";
	char password[] = "Good Morning!";

	// In The Following Case NULL Character Is NOT Stored As Last Character
	//		Because I Am Fixing Size Of String As Number Of Characters In String
	char dingdong[9] = "Ding Dong";  
	// char dingdong1[9] = "Ding Dong\0";
	char dingdongAgain[10] = "Ding Dong\0";
	char *dingdongAgain1 = "Ding Dong\0";
	
	printf("\nGreeting : %s", greeting);
	printf("\nDing Dong: %s", dingdong);
	printf("\nDing Dong Again: %s", dingdongAgain);
	printf("\nDing Dong Again 01: %s", dingdongAgain1);

// Function : playWithString
// Greeting : Good Morning!
// Ding Dong: Ding DongGood Morning!
}

//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________


int main(int argc, char *argv[]) {
	printf("\n\nFunction : playWithCommandLineArgs");
	playWithCommandLineArgs( argc, argv );

	printf("\n\nFunction : playWithSum");
	playWithSum();

	printf("\n\nFunction : playWithTypes");
	playWithTypes();

	printf("\n\nFunction : playWithString");
	playWithString();

	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");	
}

