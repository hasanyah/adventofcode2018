package main

import (
  "fmt"
)

var idList = [250]string{"asgwdcmbrkerohqoutfylvzpnx","asgwjcmbrkejihqoutfylvipne","asgwjcmbrkedihqoutvylizpnz","azgsjcmbrkedihqouifylvzpnx","asgwucmbrktddhqoutfylvzpnx","asgwocmbrkedihqoutfylvzivx","aqgwjcmbrkevihqvutfylvzpnx","tsgljcmbrkedihqourfylvzpnx","asgpjcmbrkedihqoutfnlvzsnx","astwjcmbrktdihqrutfylvzpnx","asgwjcmbrpedhhqoutfylvzynx","xsgwjcmbrkedieqowtfylvzpnx","asgwjcmbvkedihfoutnylvzpnx","asgwjcmtrkedihqouafylvzcnx","asgwjcmbrkedihqoutfylvxpvm","usgwjcmbrkedihqortfyuvzpnx","asgwjcmbrwedihqoutfylizpix","asgrjcvbrkedixqoutfylvzpnx","asgwjcmbrogdihqoutfelvzpnx","aggwjcmbrkesihqoutoylvzpnx","asgtjccbrkedihqoutfrlvzpnx","asgcucmbrbedihqoutfylvzpnx","esgwjcmbrkedihqsutfylvzcnx","asgwjcmbrkedrhqoutfyobzpnx","mngwjcbbrkedihqoutfylvzpnx","asgwjcrbrkeoihqyutfylvzpnx","apgwjcmbrkednhqogtfylvzpnx","asgwjcwbrkedihqoutfylplpnx","asgwjcmbrkfdihqoutfxlvzpyx","aegwjcmbrkedihqoutfylbxpnx","asgljcmbrkedixqoutaylvzpnx","aigwjcmbrkedihqouhfylvzpex","asgwjbmbrkedihqoutfylfzpnp","asgwjcmzroedihqoutcylvzinx","asgwjcwbrieuihqoutfylvzpnx","aagwjcmbrkedjhqdutfylvzpnx","ahgwjcmbrkedihqsutfylvzpfx","asgwjcmbrkedihzosttylvzpnx","aegwjcmbrkedioqnutfylvzpnx","asgwjcmbykidihqoutfysvzpnx","asgwkcxbrkeddhqoutfylvzpnx","ashwjcmbrkeeihqoutfylvzknx","acgwjcmbrqedihqoqtfylvzpnx","asgwjcmtrkedihooutfylszpnx","asgwjcmbrkmdihqfutrylvzpnx","asgwjcmbrkedihqoutjylvapnn","asgwjcmbwkedihqoutkylkzpnx","asgwjrmbrkedihqoutfycnzpnx","asgwtcmbrkedihqoqtfylozpnx","asgajcmbrkedihqoutuylvzpny","asgwjcmbykedihqoutfylfzpwx","asgwjcsbrkedihpoutfylvvpnx","hsdwjcmbrvedihqoutfylvzpnx","asgwjcmbrkedihqoutfdmszpnx","adgwjcmbrtidihqoutfylvzpnx","augwjcmbriedihqoutgylvzpnx","asgwjvmbreedihqoutfllvzpnx","asgwjcnbfkedihqoltfylvzpnx","asgwjcmbykddihqoutqylvzpnx","ajgwjcmbrkedihqoutfylvpvnx","asgwjcmbrkydihqoutfylszpnl","xsgwjcmbrkqdihqoutfylvkpnx","asgwjcmbrkedimqoutfklvzknx","csgwjbmbrkedihqoftfylvzpnx","asgwjcmbjkedihjoutfylvzpnn","asgwjcmprkedihqoulfalvzpnx","asgwjcmbrvediqqoutfyuvzpnx","asgwjambrkedhhqoutkylvzpnx","asgejcmbrkidihqoutfylvzpnk","hsiwjcmbrkedihqoutfylvzpnq","asswjczbrkedihqoutfylczpnx","asgwjnmbrkedyhzoutfylvzpnx","asgwscmbrkedihqoutfklvlpnx","asgwlcmbrktdihqoutfylvzpax","asfwjcmerkedihqoutfylvipnx","asgwjcmbrkeditqoeafylvzpnx","asgwgcmbrkesihqoutfylyzpnx","fsgwjcmbrkedihqouvfyavzpnx","asgwjcmbrpedwhqoutfylmzpnx","asgwjcmbrkzdzhqoucfylvzpnx","asgwjcmnrketmhqoutfylvzpnx","asgwjcmbrkedihxoutsylvzpnh","asgwjcobrkedihqoutfrlvzpox","asgwjcmbrkedihqootfylxzpox","asgjjcmcrkedihqoutfylmzpnx","lsgwjcmbrkedihqoutfyqvzunx","asgwjcmbrwedihqoutoylvzpnu","aszwjcmbtkedihqoutfylczpnx","asgwjcmbykedihqoutfylvgpex","asgijcmbrkedilqoutkylvzpnx","astwxcmzrkedihqoutfylvzpnx","akgwjcmbnkedihqfutfylvzpnx","asgwjcmbrqndivqoutfylvzpnx","asgwjrmbrleqihqoutfylvzpnx","asgwjcmbrkevihqoutfxlvzpvx","asbwjcmbrkedihqoutfelvwpnx","asewjcbbrkmdihqoutfylvzpnx","asgwjcmbrkeaihxoutfylpzpnx","asgwjzmbrkedihqrotfylvzpnx","asgwjcmbrkedihqoutgdxvzpnx","asgwjcwbrkmdihqoutfylvzlnx","asgwjcmbrkegihqoutfylrzpax","ajgwjcmbrkegihqhutfylvzpnx","asgwjcmbrzedihqhutfylvkpnx","asgwjcmwrkedihqouhfylkzpnx","aygwjcmbrkedihqoutfdlvzpnr","asgwjcmbrkednhqoutiylvypnx","aqgwjcmbrkezihqoutfylvzonx","bsgwjcmbrkedihqouhfylvzsnx","asgwjcmcrkedihqokyfylvzpnx","asgsjcmbrkewiyqoutfylvzpnx","asgwpcmbrkejihqoutfylzzpnx","asgwjumbrkedbeqoutfylvzpnx","asgwjcmbrkedihpoutqylqzpnx","awgwjcmbrredihqoutfylvzpna","asgwjsmbraedihqoutfylvzpvx","asgwncmbrkedihqoutfyljzrnx","asgwncmbrkedihqohtfylvzonx","asgwjcmbrkedihqlutfylvypux","asgwjcmbbkedihooutfylkzpnx","asghjcmsryedihqoutfylvzpnx","asgwjcmbrkevihqoulfzlvzpnx","asggjcmbrkedizqoutfylvzknx","asbwjcmbriedihqoutfylvmpnx","asgwjcmbrkedqbqoutfylvzenx","asgwjcmprkedihqoutfylvzknp","asgwjcmbrkerihqoutfwlvzpno","asgwjcmvrkesihqoutrylvzpnx","asgzjcmbrkedihqoutfnlvbpnx","asfwjcmbrkhdihqoutfylpzpnx","asgwjcmbskedihqdutfyyvzpnx","asgwjcmzrkedihqoutcylvzinx","asgwjcmbrkedibqoutfylvjonx","asgwjcmbrbedihqoutfylmzbnx","asgwjcmbrkedhhqoutmylczpnx","asgwjcmbrkbgihqoutzylvzpnx","asgwjcfbrkedihqoupfyxvzpnx","asiwjcmbzkedihqoutfyluzpnx","asvwjcmbrkedihqoitfylvzpns","asgwjcmxikedihqoutfyevzpnx","asgwjcmbrkedioqoutfylvzwox","asgwjcmbrkedivqoutjyuvzpnx","asgwjcmbkkydihqrutfylvzpnx","asgwjcmbrkxdiuqoutfylvopnx","asgwjcmbrkedihqouthylvzpra","asgwjcmbrzedimloutfylvzpnx","asgwjcmbrkedmhqoulfytvzpnx","asgwjcmbrkzdihqrutfysvzpnx","ssgwjcmxrkedihqoutftlvzpnx","asgwjcmbrkedihqoutfajvzynx","asgwjcmbrkqdihqxuufylvzpnx","asmwjcabrkedihqouxfylvzpnx","asgwjcmbrkeeihqoatfycvzpnx","asgwjcjbrgedjhqoutfylvzpnx","asgljcmtrkedihqoutoylvzpnx","asgwjcmbrkedigqouzfylvzpvx","ajgvjcmbkkedihqoutfylvzpnx","asgwjcmbrkedihqtugfygvzpnx","asgbjcmbrkedihboftfylvzpnx","asgwjwmbrkedihqontfylhzpnx","asgwjfmhrkedihqoutfylvqpnx","asgwjxmbrkedihqoutzylvzpnj","asgwjcrlrkedihqoutfylvzpsx","aygwjcmbrkedihqoutsylvzdnx","zsgwjcmbrkedihjogtfylvzpnx","asgwjxmbrkegihqoutfylvopnx","asgwjcmbrkedihqhutfylvzcnr","asgwicmbrkewihvoutfylvzpnx","asqwjcmbvkedihqoutfylvzknx","asgwjcmbrkedihqoktfyevzpnu","asgwjcmbrkudihqoutfylqzznx","asgwjdmbrkedihqoutfylvvdnx","asgwjcmbrkwmihqautfylvzpnx","asgwjcmbrxedihqoctfyldzpnx","asgwjdmbrkedjhqoutfyfvzpnx","asgwjcmtrzedihqoutfylvzpnm","bpgwjcmbrmedihqoutfylvzpnx","asgwjctbrkedihqoqtfynvzpnx","askhjcmbrkedihqoutfylvzrnx","asgkjcmblkehihqoutfylvzpnx","asgwjjmbrkedvhqoutfhlvzpnx","asgwjcmbrkedihqoupzylvzknx","asgwjcmbukedchqoutfylizpnx","askwjcmdrkedihqoutwylvzpnx","asgwjcmbtkcdihloutfylvzpnx","asgwjcmbrkedwgqoutvylvzpnx","asmwjcmbrkedihqoutfylozpnc","asgwjcmbriedibqouofylvzpnx","asgnjcmcrkedihqoupfylvzpnx","asgzjcmbrksdihqoutiylvzpnx","asgwjcrbkkedihqouafylvzpnx","asgwjcmbkkvdihqqutfylvzpnx","astwjcqbrkedihqoutfylvzpvx","asgwjcmhrkehihqoutfylvzvnx","asgwjcmbraeduhqoutmylvzpnx","asgwjcmbrkedihquutnylvzptx","asgwjcmbrkedilqoftfylvzpnz","akgwjmmbrkedihqoutfylxzpnx","asgwjcmbrkhxikqoutfylvzpnx","asgcjcmetkedihqoutfylvzpnx","fsgwjcmsrkedihooutfylvzpnx","gsgwjcmbrkedihdoutfylvzdnx","asgwtccbrkedihqoutfylvwpnx","asuwjcmbrkedihqcutfylvzpox","asgwacmbrkodihqlutfylvzpnx","asgwjcmbrkediuqoutfylqhpnx","asgwjcmbrkwdrhqoutfylvzpno","angwjcsblkedihqoutfylvzpnx","aigwjcmbyoedihqoutfylvzpnx","adgwjcmbrkedihqtutfylyzpnx","asgwjzmbrkeeihqputfylvzpnx","asgwjcmbrkwdihqoutfylvzpwc","asgpjcmbrkgdihqbutfylvzpnx","osgwjmmbrkedijqoutfylvzpnx","asgjjcmbrkkdihqoutfylvzynx","asgwjcnerwedihqoutfylvzpnx","azgwhcmbrkedicqoutfylvzpnx","asnwjcmbrsedihqoutfylvzpnm","hsgwjcmgrkedihqoutfilvzpnx","asgwscmbrkjdihqoutfylvzpnm","asgbjbmbrkedhhqoutfylvzpnx","aswwjcmtrkedihqjutfylvzpnx","asgwicmbrbedihqoutfylvzpnm","asgwjcubrkedihqoutfylvbnnx","asvwjcmbrkehihqoutfylhzpnx","gsgwjcmbrkedihqoutsklvzpnx","asgwjcubikedihqoitfylvzpnx","asgwjpmbskedilqoutfylvzpnx","aigwjcmbrxedihqoutyylvzpnx","asgwjcpbrkedihxoutfynvzpnx","asgwjcmbrkegihqoutfklvzcnx","asgwjvubrkedjhqoutfylvzpnx","asgwjcabrkedihqoutfyivzplx","asgwjcxbrkedihqgutfylvepnx","asgwlcmbrkedihqoutqylvwpnx","asgwjhmbrkydihqhutfylvzpnx","asgwjcmbrkedihqoutfylwzone","asgwycmbrkadihqoutuylvzpnx","asgwjcybrkedihqoftfylvzpne","asgwjcmnrkedihrodtfylvzpnx","asgwicmwrkedihqoutfyovzpnx","aqgwjlmbrkedilqoutfylvzpnx","asgwjcmzskwdihqoutfylvzpnx","asgwjcmdrkebihqoutfylvjpnx","asgwjcmbrkpdihqoutfylxzphx","asgwjcmbrkedixqoutlylfzpnx","asgwjcmbrkadihqoutfylvlpdx","asgejcmqrkedyhqoutfylvzpnx","asgwjcmvroedihpoutfylvzpnx","asgwjcmxrkedihqoutfyivzpmx"}
// var idList = [7]string{"abcde","fghij","klmno","pqrst","fguij","axcye","wvxyz"}
func main() {
  for i := 0; i < len(idList); i++ {
    for j := 0; j < len(idList); j++ {
      if i != j {
        df := Compare(idList[i], idList[j])
        if df == 1 {
          fmt.Printf("Words: %v, %v, Diff: %d \n", idList[i], idList[j], df)
        }
      }
    }
  } 
}

func CountLetters(s string) (int, int) {
  var m map[rune]int
  m = make(map[rune]int)

  for _, c := range s {
    _, ok := m[c]
    if ok {
      m[c]++
    } else {
      m[c] = 1
    }
  }

  twos := 0
  threes := 0

  for i, v := range m{
    if v == 2 {
      twos = 1
      fmt.Printf("twos: " + string(i) + "\n")
    } else if v == 3 {
      threes = 1
      fmt.Printf("threes: " +  string(i) + "\n")
    }
  }

  return twos,threes
}

func Compare(s1, s2 string) int {
  diff := 0
  if len(s1) != len(s2) {
    fmt.Printf("Different sized strings\n")
    return 999
  }

  for i := 0; i < len(s1); i++ {
    if s1[i] != s2[i] {
      diff++
    }
  }

  return diff
}