static bool Finder(string? name){
    if(File.Exists(String.Format($"{name}"))){
        return true;
    } else if(Directory.EnumerateDirectories(String.Format($"./{name}")).Any()){
        IEnumerable<string> ListOfDir=Directory.EnumerateDirectories(String.Format($"./{name}"));
        foreach(string i in ListOfDir){
            
            Console.WriteLine(String.Format($"{i}/{name}"));
        }
    }
    return false;
}







//Console.Write("Enter File Name: ");
string name = "README";

Console.WriteLine(Finder(name));
