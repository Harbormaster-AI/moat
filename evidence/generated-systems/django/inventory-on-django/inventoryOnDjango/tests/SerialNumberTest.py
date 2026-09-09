import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.SerialNumber import SerialNumber
from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

 #======================================================================
# 
# Encapsulates data for model SerialNumber
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SerialNumberTest Declaration
#======================================================================
class SerialNumberTest (TestCase) :
	def test_crud(self) :
		serialNumber = SerialNumber()
		serialNumber.serial = "default serial field value"
		serialNumber.activationDate = datetime.datetime.now()
		serialNumber.status = "default status field value"
		
		delegate = SerialNumberDelegate()
		responseObj = delegate.create(serialNumber)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


