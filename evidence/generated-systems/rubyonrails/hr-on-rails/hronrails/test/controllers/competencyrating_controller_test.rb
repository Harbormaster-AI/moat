require "test_helper"

class CompetencyRatingControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @competencyRating = competencyRatings(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create competencyRating" do
    assert_difference("CompetencyRating.count") do
      post competencyRatings_url, params: { competencyRating: { comment:"test string for comment", Rating:CompetencyRating.Ratings[0] } }
    end

    assert_redirected_to competencyRatings_url
  end

 
  
  test "should destroy competencyRating" do
    assert_difference("CompetencyRating.count", -1) do
      delete competencyRating_url(@competencyRating)
    end

    assert_redirected_to competencyRatings_url
  end
  
end


