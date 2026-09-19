require "test_helper"

class ContactControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @contact = contacts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create contact" do
    assert_difference("Contact.count") do
      post contacts_url, params: { contact: { firstName:"test string for firstName", lastName:"test string for lastName", title:"test string for title", email:"test value", phone:"test value", mobile:"test value", mailingAddress:"test value", PreferredContactMethod:Contact.PreferredContactMethods[0] } }
    end

    assert_redirected_to contacts_url
  end

 
  
  test "should destroy contact" do
    assert_difference("Contact.count", -1) do
      delete contact_url(@contact)
    end

    assert_redirected_to contacts_url
  end
  
end


