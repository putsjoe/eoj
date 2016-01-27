#eoj

A LILO stack, used for an attempt at randomness, or unpredictability.
 
## Plan

- Create the main script (main.go), which takes in strings on a certain port, after the stack is full enough, start returning the first string in the array/slice. 

- Create a producer (prod.go or producer.go), which is creating a random string a sending it to the chosen port.

- Improve the scripts by allowing command line arguments for port and for seed or maybe algorithm for randomness.

- Improve further by adding to main the ability to listen on multiple ports. This could then 'randomly' accept on the different ports in different orders without telling the producers what has or hasnt been accepted.



## Old Plan

- Start with reading and writing to a file, as a basic way of getting the values etc and using the input to control this.

- Look into alternatives to using a file - such as writing to memory.

- Try to maybe also write to a file as a backup, in parallel after adding and removing from the main storage/memory.

- Look at actually serving it to be used, possibly as a http server using /get/ and /set/? -- More than likely this wont be suitable and will need to look more into API.


