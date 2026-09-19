require "test_helper"

class CompetencyControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @competency = competencys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create competency" do
    assert_difference("Competency.count") do
      post competencys_url, params: { competency: { name:"test string for name", category:"test string for category" } }
    end

    assert_redirected_to competencys_url
  end

 
  
  test "should destroy competency" do
    assert_difference("Competency.count", -1) do
      delete competency_url(@competency)
    end

    assert_redirected_to competencys_url
  end
  
end


