import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.ConnectedAircraft import ConnectedAircraft
from aerospaceOnDjango.delegates.ConnectedAircraftDelegate import ConnectedAircraftDelegate

 #======================================================================
# 
# Encapsulates data for model ConnectedAircraft
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConnectedAircraftTest Declaration
#======================================================================
class ConnectedAircraftTest (TestCase) :
	def test_crud(self) :
		connectedAircraft = ConnectedAircraft()
		connectedAircraft.communicationsProvider = "default communicationsProvider field value"
		connectedAircraft.connectivityStatus = "default connectivityStatus field value"
		
		delegate = ConnectedAircraftDelegate()
		responseObj = delegate.create(connectedAircraft)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


