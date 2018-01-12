package qsort

func QSort(values []int) {
    qsortMain(values, 0, len(values))
}

func qsortMain(v [] int, vl int, vr int) {
    if vl + 1 < vr {
        vp := qsortPart(v, vl, vr)
        qsortMain(v, vl, vp)
        qsortMain(v, vp + 1, vr)
    }
}

func qsortPart(v [] int, vl int, vr int) int {
    vr --
    val := v[vl]
    for vl < vr {
        for vl < vr && v[vr] >= val {
            vr --
        }
        v[vl] = v[vr]
        for vl < vr && v[vl] <= val {
            vl ++
        } 
        v[vr] = v[vl]
    }
    v[vl] = val
    return vl
}