require "test_helper"

class ReservationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @reservation = reservations(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create reservation" do
    assert_difference("Reservation.count") do
      post reservations_url, params: { reservation: { referenceNumber:"test string for referenceNumber", reservedQuantity:"test value", promisedDate:1.week.ago, ReservationStatus:Reservation.ReservationStatuss[0], ReservationType:Reservation.ReservationTypes[0] } }
    end

    assert_redirected_to reservations_url
  end

 
  
  test "should destroy reservation" do
    assert_difference("Reservation.count", -1) do
      delete reservation_url(@reservation)
    end

    assert_redirected_to reservations_url
  end
  
end


