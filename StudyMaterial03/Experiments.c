
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
	char b = 900;
	printf("\nValue %d", b);
}

//_____________________________________________________________________
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
	// printf("\n\nFunction : ");	
}

