## NSQ Windower

Is a L2 (considering L1 to be payload agnostic, and L2 to be payload aware) utility focused
on event time windowing.  Windowing is a primitive operation in stream processing frameworks,
dataflow API provides a clean API to trivially window messages.  

The goal is to get windowing (or some variant) close to the NSQ community, so that 
people can stop creating their own home grown application level solutions to core
stream processing problems.