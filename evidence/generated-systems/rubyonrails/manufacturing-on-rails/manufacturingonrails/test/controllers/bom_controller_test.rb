require "test_helper"

class BOMControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @bOM = bOMs(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create bOM" do
    assert_difference("BOM.count") do
      post bOMs_url, params: { bOM: { bomNumber:"test string for bomNumber", revision:"test string for revision", effectivityStart:1.week.ago, effectivityEnd:1.week.ago, Status:BOM.Statuss[0] } }
    end

    assert_redirected_to bOMs_url
  end

 
  
  test "should destroy bOM" do
    assert_difference("BOM.count", -1) do
      delete bOM_url(@bOM)
    end

    assert_redirected_to bOMs_url
  end
  
end


