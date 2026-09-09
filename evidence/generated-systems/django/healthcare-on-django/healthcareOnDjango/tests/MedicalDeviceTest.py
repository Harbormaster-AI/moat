import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.MedicalDevice import MedicalDevice
from healthcareOnDjango.delegates.MedicalDeviceDelegate import MedicalDeviceDelegate

 #======================================================================
# 
# Encapsulates data for model MedicalDevice
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicalDeviceTest Declaration
#======================================================================
class MedicalDeviceTest (TestCase) :
	def test_crud(self) :
		medicalDevice = MedicalDevice()
		medicalDevice.udi = "default udi field value"
		medicalDevice.manufacturer = "default manufacturer field value"
		medicalDevice.deviceType = "default deviceType field value"
		medicalDevice.connectivityStatus = "default connectivityStatus field value"
		
		delegate = MedicalDeviceDelegate()
		responseObj = delegate.create(medicalDevice)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


