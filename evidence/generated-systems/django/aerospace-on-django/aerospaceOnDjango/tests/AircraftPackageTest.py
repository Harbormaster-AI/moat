import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.AircraftPackage import AircraftPackage
from aerospaceOnDjango.delegates.AircraftPackageDelegate import AircraftPackageDelegate

 #======================================================================
# 
# Encapsulates data for model AircraftPackage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftPackageTest Declaration
#======================================================================
class AircraftPackageTest (TestCase) :
	def test_crud(self) :
		aircraftPackage = AircraftPackage()
		aircraftPackage.name = "default name field value"
		aircraftPackage.packageType = "default packageType field value"
		
		delegate = AircraftPackageDelegate()
		responseObj = delegate.create(aircraftPackage)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


