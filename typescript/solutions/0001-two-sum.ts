function twoSum(nums: number[], target: number): number[] {
  let map: { [key: number]: number } = {}

  for (let i = 0; i < nums.length; i++){
    const diff = target - nums[i]
    if(diff in map){
      return [map[diff], i]
    }

    map[nums[i]] = i
  }

  return []
}
