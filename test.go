package simhash

import "fmt"
const t_v1 = 12
const t_v2 = 12
type Item struct {
    id  string
    price   float64
    quantity    int

}

func (item Item) Cost() float64{
    return item.price * float64(item.quantity)
}

func (item *Item) inc(){
    item.price++
}


type SpItem struct {
    Item
    cata    int
}

func (s SpItem) String() string{
    return fmt.Sprintf("%f , %d", s.price, s.quantity)
}
