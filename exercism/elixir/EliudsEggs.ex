defmodule EliudsEggs do
  @doc """
  Given the number, count the number of eggs.
  """
  @spec egg_count(number :: integer()) :: non_neg_integer()
  def egg_count(number) do
  num=Integer.to_string(number,2)
  num |> String.split("",trim: true) |> Enum.count(fn x-> String.to_integer(x)==1 end)
  end
end
