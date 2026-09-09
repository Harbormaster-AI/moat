import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.Reservation import Reservation
from inventoryOnDjango.delegates.ReservationDelegate import ReservationDelegate

 #======================================================================
# 
# Encapsulates data for model Reservation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReservationTest Declaration
#======================================================================
class ReservationTest (TestCase) :
	def test_crud(self) :
		reservation = Reservation()
		reservation.referenceNumber = "default referenceNumber field value"
		reservation.reservedQuantity = "default reservedQuantity field value"
		reservation.promisedDate = datetime.datetime.now()
		reservation.reservationStatus = "default reservationStatus field value"
		reservation.reservationType = "default reservationType field value"
		
		delegate = ReservationDelegate()
		responseObj = delegate.create(reservation)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


