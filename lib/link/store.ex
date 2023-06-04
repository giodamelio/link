defmodule Link.Store do
  use GenServer

  ## Client API

  def start_link(_) do
    GenServer.start_link(__MODULE__, :ok, name: Link.Store)
  end

  def set(new_link) do
    GenServer.call(__MODULE__, {:set, new_link})
  end

  def get() do
    GenServer.call(__MODULE__, {:get})
  end

  ## Callbacks

  @impl true
  def init(:ok) do
    {:ok, nil}
  end

  @impl true
  def handle_call({:get}, _from, link) do
    {:reply, link, link}
  end

  @impl true
  def handle_call({:set, new_link}, _from, _link) do
    {:reply, new_link, new_link}
  end
end
