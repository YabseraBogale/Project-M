defmodule BirdCount do
  def today(list) do
    # Please implement the today/1 function
    if length(list)==0 do
      nil
    else
      [head|tail]=list
      head
    end
  end

  def increment_day_count(list) do
    # Please implement the increment_day_count/1 function
    if length(list)==0 do
      [1]
    else
      [head|tail]=list
      head=head+1
      [head]++tail
    end
  end

  def has_day_without_birds?(list) do
    # Please implement the has_day_without_birds?/1 function
    count=length(Enum.filter(list,fn x-> x==0 end))
    if count>=1 do
      true
    else
      false
    end
  end

  def total(list) do
    # Please implement the total/1 function
  Enum.sum(list)
  end

  def busy_days(list) do
    # Please implement the busy_days/1 function
    length(Enum.filter(list,fn x-> x>=5 end))
  end
end
