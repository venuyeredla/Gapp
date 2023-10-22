# GoProject
Implementing all data storage systems like RDBMS,NOSQL,Information Retrieval as POC.

# Data managment
    Data integrity (CRC, checksums), Security(Encryption/decryption), encoding/decoding(XML,JSON,AVRO,Thrift,Protocol) and compresssion are imiportant aspects to be consider either storing to disk or transferring over network.
# Storage Systems for scalbilty
    Scalability, Fault tolerance, Latency.
# Scalability
    1. Partioning/Sharding 
    2. Replication
    3. Transactions
    4. Consistancy and consensus. 

# Data structures for creating indexes.

    Any storage system cosists of main data record and meta data called index for fast retrival.
    Storage system = Main data record + Index.

    1. Key value storage - In memory and on Disk storage. Map<K,Offset value main file> , SortedMap
    2. RDBMS - B+ Trees for indexing.
    3. Append only writes.
        a. SSTables - Sorted string tables.
        b. LSMTree - Log structured Merge Tree 

# RDBMS
    Using B+Tree based index.

# NoSQL 
  Append only for high write throghputs
 