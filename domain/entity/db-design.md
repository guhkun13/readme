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


Photo
 - id
 - name
 - url
 - documentation_id


Attendant
- id PK
- name

table relationship antara 2 table yg independent
Invitation --> orang2 yang hadir dalam suatu event --> Orang - Event
 - id PK
 - attendant_id fk to Attendant
 - event_id fk to Event
 - isConfirm bool

<!-- in most cases, this is the standard -->
List wilayah --> select * from wilayah;

<!-- dengan asumsi kita ada table Branch dan column wilayah -->
List wilayah penyelenggara --> select wilayah from Branch;

<!-- noted -->
PK: Primary Key
FK: Foreign Key

apakah kegiatan itu bisa punya banyak artikel ? atau cuman 1

kalau cuman 1 => Event has one Doc --> langsung jadi field
kalau banyak => Event has many Doc --> FK from Doc to Event

CRUD Documentation --> Entity / 1 Table sendiri
Create, List, Read

1 Documentation itu bisa punya banyak photo (gallery)