package list

import "fmt"

type MemBlock struct{
   From int
   To   int
   NextBlock *MemBlock
}
 // Slice 

type Memeory struct{
   b byte
   Next ByteNode
}

type MemoryManager struct{
   Store [1024*1024]bytes // 4
   MemeBlocks * MemBlock
}

func (mm *MemoryManager) Alloc(noOfBytes int) *Block {
      
    var  f,t int
    if mm.MemeBlocks ==nil {
      f=0
      t=f+noOfBytes
      return  &MemBlock{From: f, To:t }
    }     
    if mm.MemeBlocks.From>0 && mm.MemeBlocks.From > noOfBytes {
       newBlock :=&MemBlock{From: mm.MemeBlocks.From, To:t }
       newBlock.Next=mm.MemeBlocks;
	   mm.MemeBlocks=newBlock;
       return  newBlock
     }

   temp := mm.MemeBlocks
   tempNext :=temp.Next
   for temp !=nil && tempNext!=nil  { //  //0 99, 100-199. 
        
    if tempNext.From - temp.From>= noOfBytes{
		f=temp.From+1
		t=f+noOfBytes
		//break
      }
	temp=temp.Next

         }

   }
    //block := & MemBlock{From: f, To:t }

   return nill
}


funct (mm *MemoryManager) free(block *Block) bool {
      
     temp :=mm.MemeBlocks
     for mm.MemeBlocks !=block {
	       //deletion 
     }
 
   return block
}



func main() {
	fmt.Println("Hello, 世界")
}