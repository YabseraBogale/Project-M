static bool Finder(string? name){
    if(File.Exists(String.Format($"{name}"))){
        return true;
    } else if(Directory.EnumerateDirectories(".").Count()>2){
        IEnumerable<string> ListOfDir=Directory.EnumerateDirectories(".");
        foreach(string i in ListOfDir){
            Console.WriteLine(String.Format($"{i}"));
            if(Finder(String.Format($"i/{name}"))==true){
                return true;
            }
            
        }
    }
    return false;
}







//Console.Write("Enter File Name: ");
string name = "README";

Console.WriteLine(Finder(name));
