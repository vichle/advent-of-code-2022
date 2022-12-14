const fs = require('fs')
const path = require('path')

function isPairOrderCorrect(left, right) {
  let i = 0
  for (; i < left.length; i++) {
    const l = left[i]
    const r = right[i]
    if (r === undefined) return false
    if (typeof l === "number" && typeof r === "number") {
      if (l < r) return true
      if (l > r) return false
      continue
    } else if (Array.isArray(l) && Array.isArray(r)) {
      const ipoc = isPairOrderCorrect(l, r)
      if (ipoc === null) {
        continue
      }
      return ipoc
    } else {
      if (typeof l === "number") {
        const ipoc = isPairOrderCorrect([l], r)
        if (ipoc === null) {
          continue
        }
        return ipoc
      }
      if (typeof r === "number") {
        const ipoc = isPairOrderCorrect(l, [r])
        if (ipoc === null) {
          continue
        }
        return ipoc
      }
    }
  }
  if (right.length > i) {
    return true
  }
  return null
}

const contents = fs.readFileSync(path.join(__dirname, "input.txt")).toString()

const pairs = contents
  .split("\n\n")
  .map(v => v.split("\n"))
  .map(([a, b]) => [JSON.parse(a), JSON.parse(b)])


const indexSum = pairs
  .map(([a, b]) => isPairOrderCorrect(a, b))
  .map((v, i) => [v, i + 1])
  .filter(([v, i]) => v)
  .map(([v, i]) => i)
  .reduce((sum, v) => sum + v, 0)

console.log("Solution1: ", indexSum)

const all = contents
  .split("\n")
  .filter(v => v != "")
  .map(v => JSON.parse(v))

all.push([[2]], [[6]]) // Push "divider" packets

const indexProduct = all
  .sort((a, b) => {
    ipoc = isPairOrderCorrect(a,b)
    if (ipoc === true) return -1
    if (ipoc === false) return 1
    return 0
  })
  .map((v, i) => [v, i+1])
  .filter(([v, i]) => v.length === 1 && v[0].length === 1 && (v[0][0] === 2 || v[0][0] === 6))
  .map(([v, i]) => i)
  .reduce((acc, i) => acc * i, 1)

console.log("Solution2: ", indexProduct)