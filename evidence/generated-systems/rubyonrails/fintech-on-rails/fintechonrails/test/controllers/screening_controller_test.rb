require "test_helper"

class ScreeningControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @screening = screenings(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create screening" do
    assert_difference("Screening.count") do
      post screenings_url, params: { screening: { score:"test value", screenedAt:1.week.ago, ScreeningType:Screening.ScreeningTypes[0], Status:Screening.Statuss[0] } }
    end

    assert_redirected_to screenings_url
  end

 
  
  test "should destroy screening" do
    assert_difference("Screening.count", -1) do
      delete screening_url(@screening)
    end

    assert_redirected_to screenings_url
  end
  
end


