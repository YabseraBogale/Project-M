defmodule KitchenCalculator do
  def get_volume(volume_pair) do
    # Please implement the get_volume/1 function
    {_,number}=volume_pair
    number
  end
  

  def to_milliliter(volume_pair) do
    # Please implement the to_milliliter/1 functions
    {unit,number}=volume_pair
    cond do
      unit==:cup -> {:milliliter,number*240}
      unit==:milliliter -> {:milliliter,number}
      unit==:fluid_ounce -> {:milliliter,number*30}
      unit==:teaspoon -> {:milliliter,number*5}
      unit==:tablespoon -> {:milliliter,number*15}
    end
  
  end
  
  def from_milliliter(volume_pair, unit) do
    # Please implement the from_milliliter/2 functions
    {_,volume}=volume_pair
    cond do
    unit==:cup ->{:cup,volume/240} 
    unit==:teaspoon -> {:teaspoon,volume/5}
    unit==:tablespoon -> {:tablespoon,volume/15}
    unit==:fluid_ounce -> {:fluid_ounce,volume/30}
    unit==:milliliter -> {:milliliter,volume}
    end
  end

  def convert(volume_pair, unit) do
    # Please implement the convert/2 function
    milliliter=KitchenCalculator.to_milliliter(volume_pair)
    cond do
    unit==:cup -> KitchenCalculator.from_milliliter(milliliter,unit)
    unit==:teaspoon -> KitchenCalculator.from_milliliter(milliliter,unit)
    unit==:tablespoon -> KitchenCalculator.from_milliliter(milliliter,unit)
    unit==:fluid_ounce -> KitchenCalculator.from_milliliter(milliliter,unit)
    unit==:milliliter -> KitchenCalculator.from_milliliter(milliliter,unit)
    end
    
  end
end
