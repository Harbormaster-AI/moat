import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.Aircraft import Aircraft
from aerospaceOnDjango.delegates.AircraftDelegate import AircraftDelegate

 #======================================================================
# 
# Encapsulates data for model Aircraft
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftTest Declaration
#======================================================================
class AircraftTest (TestCase) :
	def test_crud(self) :
		aircraft = Aircraft()
		aircraft.msn = "default msn field value"
		aircraft.deliveryDate = datetime.datetime.now()
		
		delegate = AircraftDelegate()
		responseObj = delegate.create(aircraft)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


