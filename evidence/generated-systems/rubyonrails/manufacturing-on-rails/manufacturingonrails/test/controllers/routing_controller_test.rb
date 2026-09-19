require "test_helper"

class RoutingControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @routing = routings(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create routing" do
    assert_difference("Routing.count") do
      post routings_url, params: { routing: { routingNumber:"test string for routingNumber", revision:"test string for revision", effectivityStart:1.week.ago, effectivityEnd:1.week.ago, RoutingType:Routing.RoutingTypes[0], Status:Routing.Statuss[0] } }
    end

    assert_redirected_to routings_url
  end

 
  
  test "should destroy routing" do
    assert_difference("Routing.count", -1) do
      delete routing_url(@routing)
    end

    assert_redirected_to routings_url
  end
  
end


