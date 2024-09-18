
# todo: add a variety of optional functional languages (and handlers)

defmodule App do
  def render(page, args, layout) do
    # todo: import elixir page templates in production, and render directly
    IO.puts '#{page} #{layout} #{args[:key]}'
  end
end

App.render :index, %{key: "value"}, :layout
