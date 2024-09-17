defmodule Darts do
  @type position :: {number, number}

  @doc """
  Calculate the score of a single dart hitting a target
  """
  @spec score(position) :: integer
  def score({x, y}) do
  cond do
   ( x*x + y*y ) >100 -> 0
   ( x*x + y*y ) >25 -> 1
   ( x*x + y*y ) >1 -> 5
   true -> 10
  end
  end
end
