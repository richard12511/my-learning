defmodule Chap3practice do
  def large_lines!(path) do
    File.stream!(path)
    |> Stream.map(&String.trim_trailing(&1, "\n"))
    |> Enum.filter(&(String.length(&1) > 80))
  end

  def line_lengths!(path) do
    File.stream!(path)
    |> Enum.map(fn ln -> String.length(ln) end)
  end

  def longest_line_length!(path) do
    File.stream!(path)
    |> Stream.map(&String.length/1)
    |> Enum.reduce(fn x, acc -> max(x,acc) end)
  end

  def longest_line!(path) do
    File.stream!(path)
    |> Enum.max_by(&String.length/1)
  end

  def words_per_line!(path) do
    File.stream!(path)
    |> Stream.map(&String.trim_trailing(&1, "\n"))
    |> Enum.map(fn line -> words_in_line(line) end)
  end

  def words_in_line(line) do
    String.split(line) |> length
  end
  
end
