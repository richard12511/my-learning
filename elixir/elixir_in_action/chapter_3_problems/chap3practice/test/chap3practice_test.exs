defmodule Chap3practiceTest do
  use ExUnit.Case
  doctest Chap3practice

  test "greets the world" do
    assert Chap3practice.hello() == :world
  end
end
