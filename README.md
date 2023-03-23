### Sample of Contract Testing message-queue-based service (sampleHandler)

Generate/update the contract: 
go test -run Test_genCon

Validate the contract:
go test -run Test_verifyContract

Observe the error:
[ERROR] API handler start failed: accept tcp [::]:42815: use of closed network connection
