# Go arrays and slices

> Go's arrays are values. An array variable denotes the entire array; it is not a pointer to the first array element. This means that when you assign or pass around an array value you will make a copy of its contents. To avoid the copy you could pass a *pointer* to the array, but then that's a pointer to an array, not an array.

> The type specification for a slice is `[]T`, where `T` is the type of the elements of the slice. Unlike an array type, a slice type has no specified length.
