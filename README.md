<h1>Database Connectors</h1>

<h5>I created this repo to showcase the differences between newer, non-relational DB technology and standard MySQL technolgy, along with the scripts to get up and running.</h3>



| Feature          | MongoDB (NoSQL - Document DB)            | MySQL (SQL - Relational DB)      |
|-----------------|--------------------------------|--------------------------------|
| **Storage Model** | Collections of JSON-like documents (BSON) | Tables with rows and columns |
| **Record Equivalent** | **Document** (stored in a collection) | **Row** (stored in a table) |
| **Field Equivalent** | **Field** (key-value pair in a document) | **Column** (specific data attribute in a row) |
| **Schema** | Dynamic (documents in a collection can have different fields) | Fixed (all rows in a table follow the defined schema) |
| **Joins** | Not supported natively; uses embedded documents or `$lookup` (aggregation) | Supports `JOIN` operations for related tables |
| **Query Language** | MongoDB Query Language (MQL) using JSON-like queries | Structured Query Language (SQL) |
| **Transactions** | Multi-document ACID transactions supported (since v4.0) | Fully ACID-compliant transactions |
| **Scalability** | Horizontal (sharding, distributed clusters) | Vertical (replication, scale-up approach) |
| **Performance** | Faster for unstructured, high-volume data | Efficient for structured, relational data |


<h2> Which DB tech is most likely to disrupt the industry?</h2>  

AND 

<b>Some topics to explore:</b>

<li>Redis Cache and what it uses to make database speed/transaction-efficency that surpass the rest. (IoT, edge devices, Content delivery networks)
<li>MongoDB documents, collections, steaming real-time data. (real-time visualizations and streaming, IoT, edge devices, automation) 
<li></lio>SQLite for handling large amounts of string data. (maybe best for blogs and article writing)

<li>Name a few different businesses across different industries and contemplate which database technology may be used for a specific purpose. </li>


<br>
<hr>


<em>This repo is a work in progress. The end goal is many db connector scripts for many DB technologies in many different languages.</em>
