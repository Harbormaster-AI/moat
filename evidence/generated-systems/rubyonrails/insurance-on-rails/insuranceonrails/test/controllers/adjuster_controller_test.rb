require "test_helper"

class AdjusterControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @adjuster = adjusters(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create adjuster" do
    assert_difference("Adjuster.count") do
      post adjusters_url, params: { adjuster: { firstName:"test string for firstName", lastName:"test string for lastName", licenseNumber:"test string for licenseNumber", AdjusterType:Adjuster.AdjusterTypes[0] } }
    end

    assert_redirected_to adjusters_url
  end

 
  
  test "should destroy adjuster" do
    assert_difference("Adjuster.count", -1) do
      delete adjuster_url(@adjuster)
    end

    assert_redirected_to adjusters_url
  end
  
end


