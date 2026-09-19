require "test_helper"

class MedicationDispenseControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @medicationDispense = medicationDispenses(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create medicationDispense" do
    assert_difference("MedicationDispense.count") do
      post medicationDispenses_url, params: { medicationDispense: { dispenseNumber:"test string for dispenseNumber", quantity:"test value", whenPrepared:1.week.ago, Status:MedicationDispense.Statuss[0] } }
    end

    assert_redirected_to medicationDispenses_url
  end

 
  
  test "should destroy medicationDispense" do
    assert_difference("MedicationDispense.count", -1) do
      delete medicationDispense_url(@medicationDispense)
    end

    assert_redirected_to medicationDispenses_url
  end
  
end


