public static int linear_search(int[] array, int target) {
	for (int i = 0; i < array.length; i++) {
		if (array[i] == target) {
			return i;
		}
	}
	return -1;
}