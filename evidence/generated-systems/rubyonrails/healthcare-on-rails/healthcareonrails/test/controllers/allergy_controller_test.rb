require "test_helper"

class AllergyControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @allergy = allergys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create allergy" do
    assert_difference("Allergy.count") do
      post allergys_url, params: { allergy: { substance:"test string for substance", reaction:"test string for reaction", Severity:Allergy.Severitys[0], Status:Allergy.Statuss[0] } }
    end

    assert_redirected_to allergys_url
  end

 
  
  test "should destroy allergy" do
    assert_difference("Allergy.count", -1) do
      delete allergy_url(@allergy)
    end

    assert_redirected_to allergys_url
  end
  
end


