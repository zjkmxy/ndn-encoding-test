cd cmd && go build -o ../codegen .
cd .. && ./codegen --input ndn-format-0-3.go --output zz_generated.go --package codegen
rm codegen
