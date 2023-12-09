import { Name, Data } from "@ndn/packet"
import { Encoder } from "@ndn/tlv"

const oneRun = (i: number) => {
  const nameStr = "/go-ndn/test/" + i.toLocaleString('en-US', { minimumIntegerDigits: 10, useGrouping: false })
  const content = new Uint8Array(4000)
  const data = new Data(
    new Name(nameStr),
    Data.ContentType(0),
    Data.FreshnessPeriod(4000),
    Data.FinalBlock,
    "10000",
    content
  )
  Encoder.encode(data)
}

if (import.meta.main) {
  const tic = performance.now()
  for (let i = 0; i < 1000000; i++) {
    oneRun(i)
  }
  const toc = performance.now()
  console.log(`Total time: ${(toc - tic) / 1000.0} seconds`)
}
