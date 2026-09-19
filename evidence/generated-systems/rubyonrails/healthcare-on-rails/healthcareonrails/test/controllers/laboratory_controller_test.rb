require "test_helper"

class LaboratoryControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @laboratory = laboratorys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create laboratory" do
    assert_difference("Laboratory.count") do
      post laboratorys_url, params: { laboratory: { name:"test string for name", cliaNumber:"test string for cliaNumber" } }
    end

    assert_redirected_to laboratorys_url
  end

 
  
  test "should destroy laboratory" do
    assert_difference("Laboratory.count", -1) do
      delete laboratory_url(@laboratory)
    end

    assert_redirected_to laboratorys_url
  end
  
end


