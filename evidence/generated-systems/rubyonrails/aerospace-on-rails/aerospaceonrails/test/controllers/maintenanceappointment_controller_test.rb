require "test_helper"

class MaintenanceAppointmentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @maintenanceAppointment = maintenanceAppointments(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create maintenanceAppointment" do
    assert_difference("MaintenanceAppointment.count") do
      post maintenanceAppointments_url, params: { maintenanceAppointment: { appointmentDate:1.week.ago, Status:MaintenanceAppointment.Statuss[0] } }
    end

    assert_redirected_to maintenanceAppointments_url
  end

 
  
  test "should destroy maintenanceAppointment" do
    assert_difference("MaintenanceAppointment.count", -1) do
      delete maintenanceAppointment_url(@maintenanceAppointment)
    end

    assert_redirected_to maintenanceAppointments_url
  end
  
end


