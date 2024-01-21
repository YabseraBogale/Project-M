IEnumerable<string> pp= Directory.EnumerateDirectories(".");

foreach(string i in pp){
    DirectoryInfo list=new DirectoryInfo(i);
    DirectoryInfo []kk=list.GetDirectories("./*",SearchOption.AllDirectories);
    foreach(DirectoryInfo k in kk){
        Console.WriteLine(k.Name);
    }
}