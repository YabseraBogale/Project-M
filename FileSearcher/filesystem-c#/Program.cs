static bool Finder(string? name){
    if(File.Exists(String.Format($"{name}"))){
        return true;
    } else if(Directory.EnumerateDirectories(".").Any()){
        IEnumerable<string> ListOfDir=Directory.EnumerateDirectories(".");
        foreach(string i in ListOfDir){
            if(Finder(i)==true){
                return true;
            }
            
        }
    }
    return false;
}







//Console.Write("Enter File Name: ");
string name = "README";

Console.WriteLine(Finder(name));
