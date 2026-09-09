import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.InspectionCharacteristic import InspectionCharacteristic
from manufacturingOnDjango.delegates.InspectionCharacteristicDelegate import InspectionCharacteristicDelegate

 #======================================================================
# 
# Encapsulates data for model InspectionCharacteristic
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionCharacteristicTest Declaration
#======================================================================
class InspectionCharacteristicTest (TestCase) :
	def test_crud(self) :
		inspectionCharacteristic = InspectionCharacteristic()
		inspectionCharacteristic.characteristicCode = "default characteristicCode field value"
		inspectionCharacteristic.name = "default name field value"
		inspectionCharacteristic.lowerSpecLimit = "default lowerSpecLimit field value"
		inspectionCharacteristic.upperSpecLimit = "default upperSpecLimit field value"
		inspectionCharacteristic.target = "default target field value"
		inspectionCharacteristic.measurementType = "default measurementType field value"
		
		delegate = InspectionCharacteristicDelegate()
		responseObj = delegate.create(inspectionCharacteristic)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


