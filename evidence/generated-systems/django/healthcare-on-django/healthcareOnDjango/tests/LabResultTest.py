import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.LabResult import LabResult
from healthcareOnDjango.delegates.LabResultDelegate import LabResultDelegate

 #======================================================================
# 
# Encapsulates data for model LabResult
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LabResultTest Declaration
#======================================================================
class LabResultTest (TestCase) :
	def test_crud(self) :
		labResult = LabResult()
		labResult.resultCode = "default resultCode field value"
		labResult.issuedDate = "default issuedDate field value"
		labResult.status = "default status field value"
		
		delegate = LabResultDelegate()
		responseObj = delegate.create(labResult)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


