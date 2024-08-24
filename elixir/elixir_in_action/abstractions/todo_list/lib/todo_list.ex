defmodule TodoList do
  defstruct next_id: 1, entries: %{}

  def new, do: %TodoList{}

  def add_entry(todo_list, entry) do
    entry = Map.put(entry, :id, todo_list.next_id)

    updated_entries = Map.put(todo_list.entries, todo_list.next_id, entry)
    %TodoList{todo_list | entries: updated_entries, next_id: todo_list.next_id + 1 } 
  end

  def entries(todo_list, date) do
    todo_list.entries
    |> Map.values()
    |> Enum.filter(&(&1.date == date))
  end

end
