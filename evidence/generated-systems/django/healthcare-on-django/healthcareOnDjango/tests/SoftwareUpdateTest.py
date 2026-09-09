import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.SoftwareUpdate import SoftwareUpdate
from healthcareOnDjango.delegates.SoftwareUpdateDelegate import SoftwareUpdateDelegate

 #======================================================================
# 
# Encapsulates data for model SoftwareUpdate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SoftwareUpdateTest Declaration
#======================================================================
class SoftwareUpdateTest (TestCase) :
	def test_crud(self) :
		softwareUpdate = SoftwareUpdate()
		softwareUpdate.version = "default version field value"
		softwareUpdate.appliedDate = "default appliedDate field value"
		softwareUpdate.updateType = "default updateType field value"
		
		delegate = SoftwareUpdateDelegate()
		responseObj = delegate.create(softwareUpdate)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


