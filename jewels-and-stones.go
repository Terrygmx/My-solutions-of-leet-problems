func numJewelsInStones(J string, S string) int {
    var n = 0
    for _,i := range S {
        for _,j := range J {
            if i == j{
                n++
                break
            }
        }
    }
    return n
}
