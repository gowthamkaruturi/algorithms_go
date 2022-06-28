package javaworld;

import java.util.Comparator;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;
import java.util.stream.Collector;
import java.util.stream.Collectors;

class word{
  public static void main(String []args){
    String s = "In 1992, Tim Berners-Lee circulated a document titled “HTML Tags,” which outlined just 20 tags, many of which are now obsolete or have taken other forms. The first surviving tag to be defined in the document, after the crucial anchor tag, is the paragraph tag. It wasn’t until 1993 that a discussion emerged on the proposed image tag.";

    TreeMap<String,Integer> m = new TreeMap<>();

    String words[] = s.split("\\W+");
    
    for (String w :words){
      m.put(w, m.getOrDefault(w,0)+1);
    }
   List<model> l = m.entrySet().stream().map(e -> new model(e.getKey(),e.getValue())).collect(Collectors.toList());
   System.out.println("before sorting");
   l.forEach(word -> System.out.println(word));
   Comparator<model> compareBy = Comparator.comparing(model::getLen).thenComparing(model::getWord);
   System.out.println("after sorting");
   l.stream().sorted(compareBy).collect(Collectors.toList()).forEach(System.out::println);
  }

}

