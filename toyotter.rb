class Toyotter < Formula
  desc "CUI based Twitter client for Golang"
  homepage "https://github.com/KeisukeToyota/toyotter"
  url "https://github.com/KeisukeToyota/toyotter/releases/download/v0.3.2/toyotter_0.3.2_Darwin_x86_64.tar.gz"
  version "0.3.2"
  sha256 "b44947f21081f064113cd7457b1699b02fbf7f58f7d3e00cad87be8585738afa"
  
  depends_on "fzf"

  def install
    bin.install 'toyotter'
  end
end
