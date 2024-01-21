IEnumerable<string> pp= Directory.EnumerateDirectories(".");

foreach(string i in pp){
    DirectoryInfo list=new DirectoryInfo(i);
    DirectoryInfo []kk=list.GetDirectories("./*",SearchOption.AllDirectories);
    foreach(DirectoryInfo k in kk){
       if(File.Exists(k.FullName+"/Hello World.txt")){
            Console.WriteLine("First Loop");
            break;
       }
    }
    if(File.Exists(i+"/Hello World.txt")){
        Console.WriteLine("Second Loop");
            break;
    }
}