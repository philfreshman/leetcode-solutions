function isAnagram(s: string, t: string): boolean {
  const mapS = mapLetters(s)
  const mapT = mapLetters(t)

  if (Object.keys(mapT).length !== Object.keys(mapS).length){
    return false
  }

  for (const mapSKey in mapS) {
    if(mapS[mapSKey] !== mapT[mapSKey]){
      return false
    }
  }

  return true
}

function mapLetters(s: string): {[key: string]: number}{
  const map: {[key: string]: number} = {}
  for (let i = 0; i < s.length; i++) {
    if (s[i] in map){
      map[s[i]]++
    }else{
      map[s[i]] = 1
    }
  }
  return map
}
