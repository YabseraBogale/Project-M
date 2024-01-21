static bool Finder(string? name){
    if(File.Exists(String.Format($"{name}"))){
        return true;
    } else if(Directory.EnumerateDirectories(name).Any()){
        IEnumerable<string> ListOfDir=Directory.EnumerateDirectories(name);
        foreach(string i in ListOfDir){
            if(Finder(String.Format($"{i}/{name}"))==true){
                return true;
            }
        }
    }
    return false;
}







//Console.Write("Enter File Name: ");
string name = "README";

Console.WriteLine(Finder(name));
