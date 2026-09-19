require "test_helper"

class DependentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @dependent = dependents(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create dependent" do
    assert_difference("Dependent.count") do
      post dependents_url, params: { dependent: { firstName:"test string for firstName", lastName:"test string for lastName", birthDate:1.week.ago, Relationship:Dependent.Relationships[0] } }
    end

    assert_redirected_to dependents_url
  end

 
  
  test "should destroy dependent" do
    assert_difference("Dependent.count", -1) do
      delete dependent_url(@dependent)
    end

    assert_redirected_to dependents_url
  end
  
end


