
Class Diagram: 
 Event 
 - id PK 
 - organizer_id FK to Organizer.id 
 - name 
 - startTime
 - endTime
 - attendant
 - location
 - article (null) 


Organizer (Branch)
 - id
 - name

Documentation
 - id
 - event_id FK Event
 - article
 - photo_name
 - photo_caption
 - photo_desc
 - photo_url

PK: Primary Key
FK: Foreign Key

apakah kegiatan itu bisa punya banyak artikel ? atau cuman 1

kalau cuman 1 => Event has one Article (Doc) --> langsung jadi field
kalau banyakj => Event has many Articles (Doc) --> FK from Doc to Event

CRUD Documentation --> Entity / 1 Table sendiri
Create, List, Read

misalkan photo cukup 1, maka 

