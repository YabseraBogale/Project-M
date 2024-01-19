Console.Write("Enter Name Of File: ");

string? name=Console.ReadLine();

var exists=File.Exists("./"+name);
Console.WriteLine(exists);