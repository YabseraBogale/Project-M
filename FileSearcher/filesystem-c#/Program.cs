IEnumerable<string> p=Directory.EnumerateDirectories(".");

foreach(string i in p){
    Console.WriteLine(i);
}

