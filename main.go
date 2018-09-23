//Jungman Berliansyah Nurdin
//Huffman Code

package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "unicode/utf8"
    "sort"
   "strconv"
)

type ValueType int32

// Node in the Huffman tree.
type Node struct {
 Parent *Node     // Optional parent node, for fast code read-out
 Left   *Node     // Optional left node
 Right  *Node     // Optional right node
 Count  int       // Relative frequency
 Value  ValueType // Optional value, set if this is a leaf
}

// Code returns the Huffman code of the node.
// Left children get bit 0, Right children get bit 1.
// Implementation uses Node.Parent to walk "up" in the tree.
func (n *Node) Code() (r uint64, bits byte) {
 for parent := n.Parent; parent != nil; n, parent = parent, parent.Parent {
  if parent.Right == n { // bit 1
   r |= 1 << bits
  } // else bit 0 => nothing to do with r
  bits++
 }
 return
}

// SortNodes implements sort.Interface, order defined by Node.Count.
type SortNodes []*Node

func (sn SortNodes) Len() int           { return len(sn) }
func (sn SortNodes) Less(i, j int) bool { return sn[i].Count < sn[j].Count }
func (sn SortNodes) Swap(i, j int)      { sn[i], sn[j] = sn[j], sn[i] }

// Build builds a Huffman tree from the specified leaves.
// The content of the passed slice is modified, if this is unwanted, pass a copy.
// Guaranteed that the same input slice will result in the same Huffman tree.
func Build(leaves []*Node) *Node {
 // We sort once and use binary insertion later on
 sort.Stable(SortNodes(leaves)) // Note: stable sort for deterministic output!

 return BuildSorted(leaves)
}

// BuildSorted builds a Huffman tree from the specified leaves which must be sorted by Node.Count.
// The content of the passed slice is modified, if this is unwanted, pass a copy.
// Guaranteed that the same input slice will result in the same Huffman tree.
func BuildSorted(leaves []*Node) *Node {
 if len(leaves) == 0 {
  return nil
 }

 for len(leaves) > 1 {
  left, right := leaves[0], leaves[1]
  parentCount := left.Count + right.Count
  parent := &Node{Left: left, Right: right, Count: parentCount}
  left.Parent = parent
  right.Parent = parent

  // Where to insert parent in order to remain sorted?
  ls := leaves[2:]
  idx := sort.Search(len(ls), func(i int) bool { return ls[i].Count >= parentCount })
  idx += 2

  // Insert
  copy(leaves[1:], leaves[2:idx])
  leaves[idx-1] = parent
  leaves = leaves[1:]
 }

 return leaves[0]
}

// Print traverses the Huffman tree and prints the values with their code in binary representation.
// For debugging purposes.
func Print(root *Node) {
 // traverse traverses a subtree from the given node,
 // using the prefix code leading to this node, having the number of bits specified.
 var traverse func(n *Node, code uint64, bits byte)

 traverse = func(n *Node, code uint64, bits byte) {
  if n.Left == nil {
   // Leaf
   fmt.Printf("'%c': %0"+strconv.Itoa(int(bits))+"b\n", n.Value, code)
   return
  }
  bits++
  traverse(n.Left, code<<1, bits)
  traverse(n.Right, code<<1+1, bits)
 }

 traverse(root, 0, 0)
}

// Main functions go here
func main() {
  //reads file
  inp, err :=  ioutil.ReadFile("input.txt")
  if err != nil {
      fmt.Print(err)
  }
  //prints file in Decimal ASCII
  fmt.Println("Decimal ASCII :")
  fmt.Println(inp)
  //prints file in string
  fmt.Println("String : ")
  str := string(inp)
  fmt.Print(str)

  lngth := utf8.RuneCountInString(str)
  var c [29]int
  var p [29]float32

  fmt.Println("Sebaran Peluang :")

  for i := 0; i < 29; i++ {
    if (i == 0){
      c[i] = strings.Count(str, string(32))
      fmt.Print("'",string(32),"'",":")
    } else if (i == 27)  {
      c[i] = strings.Count(str, string(32))
      fmt.Print("'",string(44),"'",":")
    } else if (i == 28)  {
      c[i] = strings.Count(str, string(32))
      fmt.Print("'",string(46),"'",":")
    } else {
      c[i] = strings.Count(str, string(64+i))
      fmt.Print("'",string(64+i),"'",":")
    }
    p[i] = float32(c[i])/float32(lngth)
    fmt.Println(p[i])
  }

  leaves := []*Node{
    {Value: ' ', Count: c[0]},
    {Value: 'A', Count: c[1]},
    {Value: 'B', Count: c[2]},
    {Value: 'C', Count: c[3]},
    {Value: 'D', Count: c[4]},
    {Value: 'E', Count: c[5]},
    {Value: 'F', Count: c[6]},
    {Value: 'G', Count: c[7]},
    {Value: 'H', Count: c[8]},
    {Value: 'I', Count: c[9]},
    {Value: 'J', Count: c[10]},
    {Value: 'K', Count: c[11]},
    {Value: 'L', Count: c[12]},
    {Value: 'M', Count: c[13]},
    {Value: 'N', Count: c[14]},
    {Value: 'O', Count: c[15]},
    {Value: 'P', Count: c[16]},
    {Value: 'Q', Count: c[17]},
    {Value: 'R', Count: c[18]},
    {Value: 'S', Count: c[19]},
    {Value: 'T', Count: c[20]},
    {Value: 'U', Count: c[21]},
    {Value: 'V', Count: c[22]},
    {Value: 'W', Count: c[23]},
    {Value: 'X', Count: c[24]},
    {Value: 'Y', Count: c[25]},
    {Value: 'Z', Count: c[26]},
    {Value: ',', Count: c[27]},
    {Value: '.', Count: c[28]},
  }

  root := Build(leaves)
  fmt.Println("Huffman Encoding :")
  Print(root)
}
