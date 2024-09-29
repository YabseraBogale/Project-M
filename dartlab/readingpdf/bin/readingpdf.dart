import 'dart:io';

void main(List<String> arguments) {
  String data =
      String.fromCharCodes(File("sample.pdf").readAsBytesSync().toList());
  print(data);
}
