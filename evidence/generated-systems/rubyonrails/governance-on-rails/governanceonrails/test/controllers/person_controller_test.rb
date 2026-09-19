require "test_helper"

class PersonControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @person = persons(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create person" do
    assert_difference("Person.count") do
      post persons_url, params: { person: { firstName:"test string for firstName", lastName:"test string for lastName", email:"test value", department:"test string for department" } }
    end

    assert_redirected_to persons_url
  end

 
  
  test "should destroy person" do
    assert_difference("Person.count", -1) do
      delete person_url(@person)
    end

    assert_redirected_to persons_url
  end
  
end


