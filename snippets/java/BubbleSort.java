public int[] bubble_sort(int[] arr) {
	for (int i = 0; i < arr.length-1; i++) {
		boolean sorted = true;
		for (int j = 0; j < arr.length-1; j++) {
			if (arr[j] > arr[j+1]) {
				int temp = arr[j+1];
				arr[j+1] = arr[j];
				arr[j] = temp;
				sorted = false;
			}
		}
		if (sorted) {
			return arr
		}
	}
	return arr;
}