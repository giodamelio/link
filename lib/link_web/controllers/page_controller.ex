defmodule LinkWeb.PageController do
  use LinkWeb, :controller
  alias Link.Store

  def index(conn, _params) do
    redirect(conn, external: Store.get())
  end

  def see(conn, _params) do
    text(conn, Store.get())
  end
end
