require "test_helper"

class LeadControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @lead = leads(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create lead" do
    assert_difference("Lead.count") do
      post leads_url, params: { lead: { firstName:"test string for firstName", lastName:"test string for lastName", company:"test string for company", email:"test value", phone:"test value", converted:true, Status:Lead.Statuss[0], Source:Lead.Sources[0], Rating:Lead.Ratings[0] } }
    end

    assert_redirected_to leads_url
  end

 
  
  test "should destroy lead" do
    assert_difference("Lead.count", -1) do
      delete lead_url(@lead)
    end

    assert_redirected_to leads_url
  end
  
end


