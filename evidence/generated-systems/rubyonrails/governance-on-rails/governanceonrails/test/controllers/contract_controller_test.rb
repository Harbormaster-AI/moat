require "test_helper"

class ContractControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @contract = contracts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create contract" do
    assert_difference("Contract.count") do
      post contracts_url, params: { contract: { title:"test string for title", effectiveDate:1.week.ago, expiryDate:1.week.ago, repositoryUrl:"test value", Status:Contract.Statuss[0] } }
    end

    assert_redirected_to contracts_url
  end

 
  
  test "should destroy contract" do
    assert_difference("Contract.count", -1) do
      delete contract_url(@contract)
    end

    assert_redirected_to contracts_url
  end
  
end


