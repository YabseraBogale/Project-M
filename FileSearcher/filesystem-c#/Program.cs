Console.Write("Enter File Name to Look For: ");

string? name=Console.ReadLine();

bool exits=File.Exists("./"+name);

if(exits==false){
    string path=".././";
    int MaxSearch=Path.GetFullPath(".").ToString().Split("/").Length;
    for(int i=0;i<MaxSearch;i++){
        exits=File.Exists(path+name);
        if(exits==true){
            Console.WriteLine($"Found {name} at {Path.GetFullPath(path).ToString()}");
            break;
        }
       else{
            IEnumerable<string>? ListDir=Directory.EnumerateDirectories(path);
            foreach(string k in ListDir){
                    string []files=Directory.GetDirectories(k);
                    foreach(string j in files){
                        Console.WriteLine(j);
                    }
            }
    
            path+=".././";
       }
        
    }
}
else{

    Console.WriteLine($"Found {name} in {Directory.GetCurrentDirectory()}");
}