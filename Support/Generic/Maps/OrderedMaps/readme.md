


    func main() {
    m := NewOrderedMap[int, string]()
    
        m.Set(1, "string1")
        m.Set(2, "string2")
        m.Set(3, "string3")
        m.Set(4, "string4")
        m.Delete(3)
        m.Delete(8)
    
        iterator := m.Iterator()
    
        for {
            i, k, v := iterator()
            if i == nil {
                break
            }
            fmt.Println(*i, *k, v+" is a string")
        }
    }