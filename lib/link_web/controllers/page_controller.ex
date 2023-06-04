defmodule LinkWeb.PageController do
  use LinkWeb, :controller

  def index(conn, _params) do
    redirect(conn, external: "https://giodamelio.com")
  end

  def see(conn, _params) do
    text(conn, "https://giodamelio.com")
  end
end
