# Lamport Clocks

### Context
    - Spatially seprarated processes.
    - Transmission delay is NOT negligible


### Terminology
"Process" is a set of events.
Ci is a clock for each process Pi is a fnction that assigns a number C(i) to 
any event 'a' in that process.

C is a function of entire system of clocks that assigns ant event v the number 
C(b), where C(b) = Cj(b) if b is an event in process Pj.

Timestamp Tm time at which a message was sent. Each message must contain Tm.

Relation =>: event 'a' of process Pi and event 'b'
 of proess Pj, then  a => b iff(Ci(a) < cJ(b) || Ci(a) = Cj(b)) && 
arbitrary tie breaking.

Total Ordering

Clock Condition 

𝒢: the set of all system events.   

-> denotes partial ordering.

Ci(t) physical time t. 

### Description 

Happened before relationship is deonted as ->.

"happened before" relation, ->, om a set of events of a system is the smallest relation satisfying the following three conditions:

    i) if a and b are events in the same process, and a comes before 
       b, then a -> b.

    ii) If a is is the sender and b is the recipient, then a -> b.

    iii) if a -> b and b -> c, then a -> c.

NOTE: Order is not defined by physical time, but rather by one event 
 occured before the other.

Clock Condition: 
    For any events a, b: if a -> b then C(a) < C(b).

Clock Conditionals:
    i) If 'a' and 'b' are events in process Pi, and 'a' comes before
       'b', then Ci(a) < Ci(b).

    ii) If 'a' if process Pi is the sender and 'b' of process Pj is the 
        receipent, then Ci(a) < Cj(b).


Rule 1: Each process Pi increments Ci between any two successive events.



Rule 2: If event 'a' of process Pi is the sender, then the message 'm' must contain a timestamp Tm = Ci(a); once receipent event 'b' of process Pj receives 
time-stamped message Tm, process Pj must update Cj to be greather than Tm 
of the message.

Context: N processes which require exclusive access to a resource.
Five Items:
    1) In order to request a resource, process Pi sends a request-resource-message Tm:Pi to every other process, and puts it on its request queue.

    2) When process Pj receives message Tm:pi of process Pi, it places it on its request queue and sends an ack with timestamp to process Pi. No ACK is needed 
        when process Pj has sent a resoruce request to process Pi whose timestamp is greater than Pi's resource-request timestamp.

    3) To release the resource, process Pi removes any Tm:Pi request resource message from its request queue and sneds a release- message with timesstam to
every other process. 

    4) When process Pj receivves process Pi's relese message, it removes any 
Tm:Pi request-resource message.

    5) Process Pi is granted the resource when the following two conditions are met: (i) There is a Tm:Pi request-resource message in its request queue which is ordered before any other in its queue by the reltion =>, .i.e., the sending event. (ii) Pi has received a message or an ACK from every other process time-stamped later than process Pi's resoruce-request Tm timestamp. 


Con: Active participation is required by all processes. 

Strong Clock Condition: For any events a, b in 𝒢, if a -> b then C(a) < C(b).

 
