require "test_helper"

class MedicalSupplierControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @medicalSupplier = medicalSuppliers(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create medicalSupplier" do
    assert_difference("MedicalSupplier.count") do
      post medicalSuppliers_url, params: { medicalSupplier: { name:"test string for name", website:"test string for website", SupplierTier:MedicalSupplier.SupplierTiers[0] } }
    end

    assert_redirected_to medicalSuppliers_url
  end

 
  
  test "should destroy medicalSupplier" do
    assert_difference("MedicalSupplier.count", -1) do
      delete medicalSupplier_url(@medicalSupplier)
    end

    assert_redirected_to medicalSuppliers_url
  end
  
end


