string Finder(string? name){
    string path=String.Format($"./{name}");
    if(File.Exists(path)){
        return path;
    } else{
        IEnumerable<string> ListOfDir=Directory.EnumerateDirectories(path);
        foreach(string i in ListOfDir){
            Console.WriteLine(i);
        }
    }
    return "Done";
}







Console.Write("Enter File Name: ");
string? name = Console.ReadLine();

Console.WriteLine(Finder(name));
