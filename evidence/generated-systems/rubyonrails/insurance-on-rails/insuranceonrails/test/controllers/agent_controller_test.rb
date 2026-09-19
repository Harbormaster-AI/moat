require "test_helper"

class AgentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @agent = agents(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create agent" do
    assert_difference("Agent.count") do
      post agents_url, params: { agent: { firstName:"test string for firstName", lastName:"test string for lastName", licenseId:"test string for licenseId", Status:Agent.Statuss[0] } }
    end

    assert_redirected_to agents_url
  end

 
  
  test "should destroy agent" do
    assert_difference("Agent.count", -1) do
      delete agent_url(@agent)
    end

    assert_redirected_to agents_url
  end
  
end


