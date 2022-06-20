package javaworld;

public class model {
  private String word;
  private Integer count;
  public String getWord() {
    return word;
  }
  public void setWord(String word) {
    this.word = word;
  }
  public Integer getCount() {
    return count;
  }
  public void setCount(Integer count) {
    this.count = count;
  }
 
  public model(String word, Integer count) {
    this.word = word;
    this.count = count;
  }
  public Integer getLen(){
    return this.word.length();
  }
  public model(){

  }
@Override
  public String toString(){
    return "word --> "+this.word+", count -->"+this.count;

  }
}
