import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.CarrierService import CarrierService
from ecommerceOnDjango.delegates.CarrierServiceDelegate import CarrierServiceDelegate

 #======================================================================
# 
# Encapsulates data for model CarrierService
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CarrierServiceTest Declaration
#======================================================================
class CarrierServiceTest (TestCase) :
	def test_crud(self) :
		carrierService = CarrierService()
		carrierService.name = "default name field value"
		carrierService.code = "default code field value"
		carrierService.carrier = "default carrier field value"
		carrierService.serviceLevel = "default serviceLevel field value"
		
		delegate = CarrierServiceDelegate()
		responseObj = delegate.create(carrierService)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


