import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Admission import Admission
from healthcareOnDjango.delegates.AdmissionDelegate import AdmissionDelegate

 #======================================================================
# 
# Encapsulates data for model Admission
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdmissionTest Declaration
#======================================================================
class AdmissionTest (TestCase) :
	def test_crud(self) :
		admission = Admission()
		admission.admitDateTime = "default admitDateTime field value"
		admission.bed = "default bed field value"
		admission.admissionType = "default admissionType field value"
		
		delegate = AdmissionDelegate()
		responseObj = delegate.create(admission)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


