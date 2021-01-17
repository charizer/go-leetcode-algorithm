package lengthOfLongestSubstring


// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func LengthOfLongestSubstring(s string) int {
	longest := 0
	for i:=0; i< len(s); i++ {
		sum := 1
		m := make(map[byte]int)
		m[s[i]] = 1
		for j:=i+1;j<len(s);j++{
			if m[s[j]] > 0 {
				break
			}else{
				m[s[j]] = 1
				sum++
			}
		}
		if sum > longest {
			longest = sum
		}
	}
	return longest
}

func LengthOfLongestSubstringn(s string) int {
	longest := 0
	n := len(s)
	rk, k := -1, 0
	m := make(map[byte]int)
	for ; k < n; k++ {
		if k != 0 {
			delete(m, s[k-1])
		}
		for rk+1 < n && m[s[rk+1]] <= 0 {
			m[s[rk+1]]++
			rk++
		}
		longest = max(longest, rk+1 - k)
	}
	return longest
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}


/**
  1、首先，判断当前字符是否包含在map中，如果不包含，将该字符添加到map（字符，字符在数组下标）,
   此时没有出现重复的字符，左指针不需要变化。此时不重复子串的长度为：right-left+1，与原来的ans比较，取最大值；

  2、如果当前字符 ch 包含在 map中，此时有2类情况：
   1）当前字符包含在当前有效的子段中，如：abca，当我们遍历到第二个a，当前有效最长子段是 abc，我们又遍历到a，
   那么此时更新 left 为 map.get(a)+1=1，当前有效子段更新为 bca；
   2）当前字符不包含在当前最长有效子段中，如：abba，我们先添加a,b进map，此时left=0，我们再添加b，发现map中包含b，
   而且b包含在最长有效子段中，就是1）的情况，我们更新 left=map.get(b)+1=2，此时子段更新为 b，而且map中仍然包含a，map.get(a)=0；
   随后，我们遍历到a，发现a包含在map中，且map.get(a)=0，如果我们像1）一样处理，就会发现 left=map.get(a)+1=1，实际上，left此时
   应该不变，left始终为2，子段变成 ba才对。

   为了处理以上2类情况，我们每次更新left，left=Math.max(left , map.get(ch)+1).
   另外，更新left后，不管原来的 s.charAt(i) 是否在最长子段中，我们都要将 s.charAt(i) 的位置更新为当前的i，
   因此此时新的 s.charAt(i) 已经进入到 当前最长的子段中！
*/

func LengthOfLongestSubstring2(s string) int {
	ans := 0
	m := make(map[byte]int)
	for left,right:=0,0; right<len(s); right++  {
		// 如果有重复，移动left到此重复元素原来所记录位置的下一个位置
		if _, ok := m[s[right]]; ok {
			// 原来left及其左边的子串丢弃， 更新left指针到原来left的的下一位重新计算
			// 有可能这个元素的下一个位置比left小， 因为map中历史元素没有delete
			left = max(left, m[s[right]] + 1)
		}
		//更新当前元素的下标
		m[s[right]] = right
		ans = max(ans, right - left + 1)
	}
	return ans
}