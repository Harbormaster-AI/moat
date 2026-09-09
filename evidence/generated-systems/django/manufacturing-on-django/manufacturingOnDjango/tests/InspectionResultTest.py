import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.InspectionResult import InspectionResult
from manufacturingOnDjango.delegates.InspectionResultDelegate import InspectionResultDelegate

 #======================================================================
# 
# Encapsulates data for model InspectionResult
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionResultTest Declaration
#======================================================================
class InspectionResultTest (TestCase) :
	def test_crud(self) :
		inspectionResult = InspectionResult()
		inspectionResult.resultValue = "default resultValue field value"
		inspectionResult.recordedOn = "default recordedOn field value"
		inspectionResult.notes = "default notes field value"
		inspectionResult.resultStatus = "default resultStatus field value"
		
		delegate = InspectionResultDelegate()
		responseObj = delegate.create(inspectionResult)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


