defmodule HighSchoolSweetheart do
  def first_letter(name) do
    # Please implement the first_letter/1 function
    name|> String.trim |> String.first
  end

  def initial(name) do
    # Please implement the initial/1 function
    name |> String.first|> String.upcase() |> Kernel.<> "."
  end

  def initials(full_name) do
    # Please implement the initials/1 function
    [first,last]=full_name |> String.split(" ", trim: true)
    initial(first)<> " " <> initial(last) 
  end

  def pair(full_name1, full_name2) do
       "     ******       ******\n   **      **   **      **\n **         ** **         **\n**            *            **\n**                         **\n**     #{initials(full_name1)}  +  #{initials(full_name2)}     **\n **                       **\n   **                   **\n     **               **\n       **           **\n         **       **\n           **   **\n             ***\n              *\n"
    # Please implement the pair/2 function
  end
end
