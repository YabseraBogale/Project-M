var exits=File.Exists("./filesystem-c#.sln");
var parentDir=Path.GetFullPath("..");
Console.WriteLine(exits.ToString());
Console.WriteLine(parentDir.ToString());