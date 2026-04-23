// Given an array of integers arr, move all the zeros to the end of the array
// while maintaining the relative order of all non-zero elements.
//
// Examples:
//   Input:  [1, 2, 0, 4, 3, 0, 5, 0]
//   Output: [1, 2, 4, 3, 5, 0, 0, 0]
//
//   Input:  [10, 20, 30]
//   Output: [10, 20, 30]
//
//   Input:  [0, 0]
//   Output: [0, 0]

static void MoveZerosToEnd(int[] arr)
{
    // TODO: Implement your solution here
}

int[] arr1 = { 1, 2, 0, 4, 3, 0, 5, 0 };
Console.WriteLine($"Input:  [{string.Join(", ", arr1)}]");
MoveZerosToEnd(arr1);
Console.WriteLine($"Output: [{string.Join(", ", arr1)}]");
Console.WriteLine("Expected: [1, 2, 4, 3, 5, 0, 0, 0]\n");

int[] arr2 = { 10, 20, 30 };
Console.WriteLine($"Input:  [{string.Join(", ", arr2)}]");
MoveZerosToEnd(arr2);
Console.WriteLine($"Output: [{string.Join(", ", arr2)}]");
Console.WriteLine("Expected: [10, 20, 30]\n");

int[] arr3 = { 0, 0 };
Console.WriteLine($"Input:  [{string.Join(", ", arr3)}]");
MoveZerosToEnd(arr3);
Console.WriteLine($"Output: [{string.Join(", ", arr3)}]");
Console.WriteLine("Expected: [0, 0]");
