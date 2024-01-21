string Finder(string name){
    string path=String.Format($"./{name}");
    if(File.Exists(path)){
        return path;
    } else{
        IEnumerable<string> ListOfDir=Directory.EnumerateDirectories(path);
        int OuterBound=ListOfDir.Count();
        while(OuterBound>0){
            OuterBound-=1;
        }
    }
}







Console.Write("Enter File Name: ");
string? name = Console.ReadLine();


