function swap(items, left, right){
    let temp = items[left]
    items[left] = items[right]
    items[right] = temp
}

function partition(items, left, right) {
    let pivot = items[Math.floor((right + left) / 2)]
    let i = left
    let j = right
    while (i <= j) {
        while (items[i] < pivot) {
            i++
        }
        while (items[j] > pivot) {
            j--
        }
        if (i <= j) {
            swap(items, i, j)
            i++
            j--
        }
    }
    return i
}

function quickSort(items, left = 0, right = items.length - 1) {
    let index
    if (items.length > 1) {
        index = partition(items, left, right)
        if (left < index - 1) {
            quickSort(items, left, index - 1)
        }
        if (index < right) {
            quickSort(items, index, right)
        }
    }
    return items
}