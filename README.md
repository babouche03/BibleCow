# 🐮 BibleCow

一个有趣的 CLI 命令行工具，让圣经的金句通过牛（或龙、耶稣）形象以气泡对话框的方式呈现！

---

## ✨ 项目特点

- 📖 随机输出圣经金句（从文本中提取）
- 🐄 支持多个角色形象（cow、dragon、jesus）
- 💬 气泡样式自动适配句长，文字自动换行
- 🖥️ 支持在终端中运行并互动选择角色

---

## 🚀 安装与使用

### 克隆项目

```bash
git clone https://github.com/你的用户名/BibleCow.git
cd BibleCow

2. 构建可执行文件

go build -o biblecow

3. 全局使用（可选）

sudo mv biblecow /usr/local/bin/

现在你可以直接在终端运行：

biblecow

⸻

🐉 运行示例

$ biblecow
请选择一个形象：cow / dragon / jesus
> cow
 ________________________________________
/ For God so loved the world that He     \
\ gave His only begotten Son...          /
 ----------------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||


```
### 📂 圣经语录文件

请在项目中准备一份圣经名言文件（例如 bible.txt），每条语句以句号（.）分隔：

For God so loved the world that He gave His only begotten Son.
The Lord is my shepherd; I shall not want.
I can do all things through Christ who strengthens me.
...

你可以放在项目目录，或者修改代码中对应路径。


## 🛠️ 技术栈
	• Go 语言
	• 标准库：os, fmt, math/rand, time, bufio, strings


## 📜 License

MIT License


## 🙌 致谢

本项目灵感来自于经典 Unix 命令 cowsay，结合圣经金句作为输出内容。

欢迎 PR、Star 或 Fork！愿神祝福你 ✝️



