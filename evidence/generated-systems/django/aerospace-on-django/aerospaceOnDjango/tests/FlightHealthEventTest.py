import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.FlightHealthEvent import FlightHealthEvent
from aerospaceOnDjango.delegates.FlightHealthEventDelegate import FlightHealthEventDelegate

 #======================================================================
# 
# Encapsulates data for model FlightHealthEvent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FlightHealthEventTest Declaration
#======================================================================
class FlightHealthEventTest (TestCase) :
	def test_crud(self) :
		flightHealthEvent = FlightHealthEvent()
		flightHealthEvent.eventCode = "default eventCode field value"
		flightHealthEvent.severity = "default severity field value"
		
		delegate = FlightHealthEventDelegate()
		responseObj = delegate.create(flightHealthEvent)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


