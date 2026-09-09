import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

 #======================================================================
# 
# Encapsulates data for model Patient
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PatientTest Declaration
#======================================================================
class PatientTest (TestCase) :
	def test_crud(self) :
		patient = Patient()
		patient.firstName = "default firstName field value"
		patient.lastName = "default lastName field value"
		patient.mrn = "default mrn field value"
		patient.dateOfBirth = datetime.datetime.now()
		patient.address = "default address field value"
		patient.primaryLanguage = "default primaryLanguage field value"
		patient.sexAtBirth = "default sexAtBirth field value"
		patient.bloodType = "default bloodType field value"
		
		delegate = PatientDelegate()
		responseObj = delegate.create(patient)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


