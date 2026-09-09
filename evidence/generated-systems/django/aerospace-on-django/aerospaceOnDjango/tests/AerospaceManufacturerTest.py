import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.AerospaceManufacturer import AerospaceManufacturer
from aerospaceOnDjango.delegates.AerospaceManufacturerDelegate import AerospaceManufacturerDelegate

 #======================================================================
# 
# Encapsulates data for model AerospaceManufacturer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AerospaceManufacturerTest Declaration
#======================================================================
class AerospaceManufacturerTest (TestCase) :
	def test_crud(self) :
		aerospaceManufacturer = AerospaceManufacturer()
		aerospaceManufacturer.name = "default name field value"
		aerospaceManufacturer.legalName = "default legalName field value"
		aerospaceManufacturer.headquartersCountry = "default headquartersCountry field value"
		aerospaceManufacturer.website = "default website field value"
		
		delegate = AerospaceManufacturerDelegate()
		responseObj = delegate.create(aerospaceManufacturer)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


