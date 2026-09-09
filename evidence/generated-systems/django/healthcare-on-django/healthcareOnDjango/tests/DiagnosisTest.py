import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Diagnosis import Diagnosis
from healthcareOnDjango.delegates.DiagnosisDelegate import DiagnosisDelegate

 #======================================================================
# 
# Encapsulates data for model Diagnosis
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DiagnosisTest Declaration
#======================================================================
class DiagnosisTest (TestCase) :
	def test_crud(self) :
		diagnosis = Diagnosis()
		diagnosis.code = "default code field value"
		diagnosis.description = "default description field value"
		diagnosis.onsetDate = datetime.datetime.now()
		diagnosis.certainty = "default certainty field value"
		
		delegate = DiagnosisDelegate()
		responseObj = delegate.create(diagnosis)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


