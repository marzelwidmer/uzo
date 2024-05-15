# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Uzo < Formula
  desc ""
  homepage "https://github.com/marzelwidmer/uzo"
  version "1.0.7"

  on_macos do
    url "https://github.com/marzelwidmer/uzo/releases/download/1.0.7/uzo_1.0.7_darwin_all.tar.gz"
    sha256 "18acbee9e90c9c8c6f6348a4ce07fd6eb9fe095deb69872d900d880be7dc5927"

    def install
      bin.install "uzo"
    end
  end

  on_linux do
    on_intel do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/marzelwidmer/uzo/releases/download/1.0.7/uzo_1.0.7_linux_amd64.tar.gz"
        sha256 "4c3999efb7cdc04dac7168e9398ef30b70c3c343869652cbc5ca195ef3a5dd25"

        def install
          bin.install "uzo"
        end
      end
    end
    on_arm do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/marzelwidmer/uzo/releases/download/1.0.7/uzo_1.0.7_linux_arm64.tar.gz"
        sha256 "4a1999eb953bebad8542349b2eec09453c1d13c8a1142a96384aa2e407a85f6c"

        def install
          bin.install "uzo"
        end
      end
    end
  end
end
