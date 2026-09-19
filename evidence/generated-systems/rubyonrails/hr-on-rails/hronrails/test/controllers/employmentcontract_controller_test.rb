require "test_helper"

class EmploymentContractControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @employmentContract = employmentContracts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create employmentContract" do
    assert_difference("EmploymentContract.count") do
      post employmentContracts_url, params: { employmentContract: { contractNumber:"test string for contractNumber", startDate:1.week.ago, endDate:1.week.ago, workHoursPerWeek:"test value", EmploymentType:EmploymentContract.EmploymentTypes[0], Status:EmploymentContract.Statuss[0], PayFrequency:EmploymentContract.PayFrequencys[0] } }
    end

    assert_redirected_to employmentContracts_url
  end

 
  
  test "should destroy employmentContract" do
    assert_difference("EmploymentContract.count", -1) do
      delete employmentContract_url(@employmentContract)
    end

    assert_redirected_to employmentContracts_url
  end
  
end


