require "test_helper"

class HealthSystemControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @healthSystem = healthSystems(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create healthSystem" do
    assert_difference("HealthSystem.count") do
      post healthSystems_url, params: { healthSystem: { name:"test string for name", legalName:"test string for legalName", headquartersCountry:"test string for headquartersCountry", website:"test string for website" } }
    end

    assert_redirected_to healthSystems_url
  end

 
  
  test "should destroy healthSystem" do
    assert_difference("HealthSystem.count", -1) do
      delete healthSystem_url(@healthSystem)
    end

    assert_redirected_to healthSystems_url
  end
  
end


