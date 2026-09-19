require "test_helper"

class CreditorControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @creditor = creditors(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create creditor" do
    assert_difference("Creditor.count") do
      post creditors_url, params: { creditor: { name:"test string for name", bic:"test value", address:"test value" } }
    end

    assert_redirected_to creditors_url
  end

 
  
  test "should destroy creditor" do
    assert_difference("Creditor.count", -1) do
      delete creditor_url(@creditor)
    end

    assert_redirected_to creditors_url
  end
  
end


