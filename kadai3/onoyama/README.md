# 解説 (Gopher道場 課題 3-1) 
 
## コマンド 
./Questions     
(上記コマンドを実行後、画面の指示に従ってください) 
   
## ライブラリ 
 
### PACKAGE DOCUMENTATION
 
package utils 
    import "./utils" 
 
### FUNCTIONS 
 
func CheckAnswer(question Question, answer string) bool    
　回答があっているか確認して結果を表示    
 
func InputChannel(r io.Reader) <-chan string     
　入力用のチャネル    
 
func LoadQuestions(path string) []Question    
　JSONファイルから問題をロード    
 
func ShowResult(result Result)    
　正解の合計を表示    
 
func TimeoutChannel(timeout int) <-chan bool    
　タイムアウト用のチャネル    
 
### TYPES 
 
type Question struct {     
　　Question string ``json:"question"``      
　　Answer   string ``json:"answer"``      
}   
　質問格納用struct 
 
func GetQuestion(questions []Question) Question     
　問題の配列からランダムに一問抽出     
 
func PopQuestion(questions []Question, q_count int) Question     
　問題の表示   
 
type Result struct {    
　　Count   int    
　　Correct int   
}   
　回答結果の格納用struct 


