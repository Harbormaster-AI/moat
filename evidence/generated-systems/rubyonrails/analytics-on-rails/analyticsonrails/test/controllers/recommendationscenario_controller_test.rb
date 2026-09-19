require "test_helper"

class RecommendationScenarioControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @recommendationScenario = recommendationScenarios(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create recommendationScenario" do
    assert_difference("RecommendationScenario.count") do
      post recommendationScenarios_url, params: { recommendationScenario: { name:"test string for name", objective:"test string for objective", RecommendationType:RecommendationScenario.RecommendationTypes[0] } }
    end

    assert_redirected_to recommendationScenarios_url
  end

 
  
  test "should destroy recommendationScenario" do
    assert_difference("RecommendationScenario.count", -1) do
      delete recommendationScenario_url(@recommendationScenario)
    end

    assert_redirected_to recommendationScenarios_url
  end
  
end


