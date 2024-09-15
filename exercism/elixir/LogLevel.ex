defmodule LogLevel do

  def to_label(0, true) do
    # Please implement the to_label/2 function
    :unknown
  end
  def to_label(0, false) do
    # Please implement the to_label/2 function
    :trace
  end
  def to_label(1, _) do
    # Please implement the to_label/2 function
    :debug
  end
  def to_label(2, _) do
    # Please implement the to_label/2 function
    :info
  end
  def to_label(3, _) do
    # Please implement the to_label/2 function
    :warning
  end
  def to_label(4, _) do
  :error
  end
  def to_label(5, false) do
  :fatal
  end
  def to_label(5, true) do
  :unknown
  end
  def to_label(_, _) do
  :unknown
  end
  
  def alert_recipient(level, legacy?) do
    # Please implement the alert_recipient/2 function
    case {to_label(level, legacy?),legacy?} do
      {:error, _} -> :ops
      {:fatal, _} -> :ops
      {:unknown, true} -> :dev1
      {:unknown, false} -> :dev2
      _ -> false
    end
  end
end

