func minTimeToVisitAllPoints(points [][]int) int {
    num := 0.0 
    for i:=0;i<len(points)-1;i++ {
        dx, dy := math.Abs(float64(points[i+1][0]-points[i][0])), math.Abs(float64(points[i+1][1]-points[i][1]))
        num += math.Max(dx,dy)
    }
    return int(num)
}