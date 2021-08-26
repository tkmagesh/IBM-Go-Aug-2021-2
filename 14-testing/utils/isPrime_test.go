package utils

import "testing"

/* func TestIsPrime(t *testing.T) {
	//Arrange
	no := 11
	expectedResult := true

	//Act
	actualResult := IsPrime(no)

	//Assert
	if actualResult != expectedResult {
		t.Fail()
	}
}

func TestIsNotPrime(t *testing.T) {
	//Arrange
	no := 12
	expectedResult := false

	//Act
	actualResult := IsPrime(no)

	//Assert
	if actualResult != expectedResult {
		t.Fail()
	}
} */

//sub tests

/* func TestUtils(t *testing.T) {
	t.Run("Number Tests", func(t *testing.T) {
		t.Run("Prime Number Tests", func(t *testing.T) {
			t.Run("Test 11 is Prime", func(t *testing.T) {
				//Arrange
				no := 11
				expectedResult := true

				//Act
				actualResult := IsPrime(no)

				//Assert
				if actualResult != expectedResult {
					t.Fail()
				}
			})

			t.Run("Test 12 is Not Prime", func(t *testing.T) {
				//Arrange
				no := 12
				expectedResult := false

				//Act
				actualResult := IsPrime(no)

				//Assert
				if actualResult != expectedResult {
					t.Fail()
				}
			})
		})
	})
} */

//data driven tests
type PrimeTestCase struct {
	name           string
	no             int
	expectedResult bool
}

func TestUtils(t *testing.T) {
	t.Run("Number Tests", func(t *testing.T) {
		t.Run("Prime Test", func(t *testing.T) {
			testCases := []PrimeTestCase{
				{name: "Testing Prime for 11", no: 11, expectedResult: true},
				{name: "Testing Prime for 13", no: 13, expectedResult: true},
				{name: "Testing Prime for 15", no: 15, expectedResult: false},
				{name: "Testing Prime for 17", no: 17, expectedResult: true},
			}
			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					//Arrange
					//Act
					actualResult := IsPrime(testCase.no)

					//Assert
					if actualResult != testCase.expectedResult {
						t.Fail()
					}
				})
			}
		})
	})
}
