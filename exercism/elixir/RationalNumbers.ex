defmodule RationalNumbers do
  @type rational :: {integer, integer}

  @doc """
  Add two rational numbers
  """
  @spec add(a :: rational, b :: rational) :: rational
  def add(a, b) do
    {x,y}=a
    {w,z}=b
    {e,f}={(x*z)+(w*y),z*y}
    reduce({e,f})
  end

  @doc """
  Subtract two rational numbers
  """
  @spec subtract(a :: rational, b :: rational) :: rational
  def subtract(a, b) do
    {x,y}=a
    {w,z}=b
    {e,f}={(x*z)-(w*y),(z*y)}
    reduce({e,f})
  end

  @doc """
  Multiply two rational numbers
  """
  @spec multiply(a :: rational, b :: rational) :: rational
  def multiply(a, b) do
    {x,y}=a
    {w,z}=b
    {e,f}={(x*w),(z*y)}
    reduce({e,f})
  end

  @doc """
  Divide two rational numbers
  """
  @spec divide_by(num :: rational, den :: rational) :: rational
  def divide_by(num, den) do
    {x,y}=num
    {w,z}=den
    {e,f}={(x*z),(w*y)}
    reduce({e,f})
  end

  @doc """
  Absolute value of a rational number
  """
  @spec abs(a :: rational) :: rational
  def abs(a) do
    {x,y}=a
    reduce({Kernel.abs(x),Kernel.abs(y)})
  end

  @doc """
  Exponentiation of a rational number by an integer
  """
  @spec pow_rational(a :: rational, n :: integer) :: rational
  def pow_rational({a,b}, n) when n<0 do
      {b ** (-n), a ** (-n)}
    |> reduce()
  end
  @spec pow_rational(a :: rational, n :: integer) :: rational
  def pow_rational({a,b}, n) do
    {a ** n, b ** n}
    |> reduce()
  end
  @doc """
  Exponentiation of a real number by a rational number
  """
  @spec pow_real(x :: integer, n :: rational) :: float
  def pow_real(x, {a,b}) do
   x ** (a/b)
  end
  
  @doc """
  Reduce a rational number to its lowest terms
  """
  @spec reduce(a :: rational) :: rational
  def reduce({a,b}) when b< 0 do
      reduce({-a, -b})    
  end
  
  @doc """
  Reduce a rational number to its lowest terms
  """
  @spec reduce(a :: rational) :: rational
  def reduce(a) do
    {x,y}=a
    number=Integer.gcd(x,y)
    {(x/number),(y/number)}
  end
end
