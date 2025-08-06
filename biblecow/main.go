package main

import(
	"fmt"
	"os"
	"bufio"
	"math/rand"
	"time"
	"strings"
)

func main(){
	// args:= os.Args
	// content := args[1]
	content:="bible"
	// figure := args[2]
	fmt.Println("今天你想和谁对话：cow / dragon / jesus")
	fmt.Print("> ")
	reader:=bufio.NewReader(os.Stdin)
	figure, _ := reader.ReadString('\n')
	figure = strings.TrimSpace(figure)
	display:=chatbox(content)
	cow:=drawCow(figure)
	fmt.Println(display)
	fmt.Println(cow)
}
func RangeRandom(min, max int) (number int) {
	//创建随机种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	number = r.Intn(max-min) + min
	return number
}
func wrapText(text string, width int) []string {
	words := strings.Fields(text)
	var lines []string
	var current string

	for _, word := range words {
		if len(current)+len(word)+1 > width {
			lines = append(lines, current)
			current = word
		} else {
			if current == "" {
				current = word
			} else {
				current += " " + word
			}
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}

func formatSpeechBubble(text string) string {
	const maxWidth = 40
	lines := wrapText(text, maxWidth)

	border := "+" + strings.Repeat("-", maxWidth+2) + "+"
	var result []string
	result = append(result, border)

	for _, line := range lines {
		padded := fmt.Sprintf("| %-*s |", maxWidth, line)
		result = append(result, padded)
	}

	result = append(result, border)
	return strings.Join(result, "\n")
}
func chatbox(content string)string{
	if content == "bible"{
		file,err := os.ReadFile("xxxxx.txt")
		if err != nil{
			fmt.Println("打开文件失败",err)
			return ""
		}else{
			text:=string(file)
			sentences := strings.Split(text, ".")
			random:=RangeRandom(0,len(sentences))
			res := strings.TrimSpace(sentences[random]+".")
			// res:=sentences[random]
			return formatSpeechBubble(res)
		}
		// defer text.Close()
	}else{
		border:=strings.Repeat("-",len(content)+1)
		return fmt.Sprintf("%v\n%v\n%v", border,content,border)
	}
	return ""
}

func drawCow(figure string) string{
	if figure == "cow"{
		return `
       \
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||`
	}
	if figure == "dragon"{
		return `                               .       .
          \                    / ` + "`" + `.   .' "
           \           .---.  <    > <    >  .---.
            \          |    \  \ - ~ ~ - /  /    |
          _____           ..-~             ~-..-~
         |     |   \~~~\\.'                    ` + "`" + `./~~~/
        ---------   \__/                         \__/
       .'  O    \     /               /       \  "
      (_____,    ` + "`" + `._.'               |         }  \/~~~/
       ` + "`" + `----.          /       }     |        /    \__/
             ` + "`" + `-.      |       /      |       /      ` + "`" + `. ,~~|
                 ~-.__|      /_ - ~ ^|      /- _      ` + "`" + `..-'
                      |     /        |     /     ~-.     ` + "`" + `-. _  _  _
                      |_____|        |_____|         ~ - . _ _ _ _ _>

			`
	}
	if figure == "jesus"{
		return `
		\
		 \
		  \  _______
		    |       |
		    |  O O  |
		    |   ^   |
		    |  \_/  |
		    |_______|
		   /         \
		  /  -------  \
		 /  |     |    \
		/___|_____|_____\
		   /  || ||  \
		  /   || ||   \
		 /    [] []    \

		     JESUS
		`
	}
	return "none"
}
