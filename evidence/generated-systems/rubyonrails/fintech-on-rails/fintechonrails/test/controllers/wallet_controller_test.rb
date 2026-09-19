require "test_helper"

class WalletControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @wallet = wallets(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create wallet" do
    assert_difference("Wallet.count") do
      post wallets_url, params: { wallet: { currency:"test string for currency", balance:"test value", Status:Wallet.Statuss[0] } }
    end

    assert_redirected_to wallets_url
  end

 
  
  test "should destroy wallet" do
    assert_difference("Wallet.count", -1) do
      delete wallet_url(@wallet)
    end

    assert_redirected_to wallets_url
  end
  
end


