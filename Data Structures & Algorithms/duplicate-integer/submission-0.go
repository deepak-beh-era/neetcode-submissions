func hasDuplicate(nums []int) bool {
    frequencyMap := make(map[int]struct{})
    for _, value := range nums {
        if _, ok := frequencyMap[value]; ok {
            return true
        }
        frequencyMap[value] = struct{}{} 
    }
    return false
}
