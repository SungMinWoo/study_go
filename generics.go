package main

import "fmt"
// MapKeys라는 제네릭 함수 생성, comparable은 비교 가능한 변수, V는 어떤 타입이 들어오던 상관없는 변수
// m은 map[K]V 타입의 매개변수이고 키 값이 K타입, 값이 V 타입인 맵을 나타낸다.(dict형태)
// []K는 K타입의 슬라이스
func MapKeys[K comparable, V any](m map[K]V) []K {
    r := make([]K, 0, len(m))
    for k := range m {
        r = append(r, k)
    }
    return r
}
// List는 제네릭 T타입을 받아들이는 리스트
type List[T any] struct {
    head, tail *element[T]
}
// element는 리스트의 각 노드를 나타낸다. 각 노트는 다음 노드를 나타내는 포인터 next와 해당 노드의 값 val을 포함한다.
type element[T any] struct {
    next *element[T]
    val  T
}
// 리스트에 새로운 값을 추가하는 메서드
func (lst *List[T]) Push(v T) {
    if lst.tail == nil { // 리스트가 비어있다면 새로운 노트 생성
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else { // 비어있지 않다면 tail에 새로운 노드 추가
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}
// 리스트에 있는 모든 값을 슬라이스로 반환
func (lst *List[T]) GetAll() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func main() {
    var m = map[int]string{1: "2", 2: "4", 4: "8"}

    fmt.Println("keys:", MapKeys(m)) // [1 2 4]

    _ = MapKeys[int, string](m)
	fmt.Println(_)
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.GetAll()) // [10 13 23]
}