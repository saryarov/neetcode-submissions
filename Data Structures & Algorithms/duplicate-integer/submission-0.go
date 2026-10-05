func hasDuplicate(nums []int) bool {
    m := make(map[int]bool)
    var flag bool = false
    for _, a := range nums { 
        if m[a] == true {
            flag = true
        }
        m[a] = true
    }

    return flag
}
