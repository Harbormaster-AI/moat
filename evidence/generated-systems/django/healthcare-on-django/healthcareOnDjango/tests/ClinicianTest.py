import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Clinician import Clinician
from healthcareOnDjango.delegates.ClinicianDelegate import ClinicianDelegate

 #======================================================================
# 
# Encapsulates data for model Clinician
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClinicianTest Declaration
#======================================================================
class ClinicianTest (TestCase) :
	def test_crud(self) :
		clinician = Clinician()
		clinician.firstName = "default firstName field value"
		clinician.lastName = "default lastName field value"
		clinician.licenseNumber = "default licenseNumber field value"
		clinician.clinicianType = "default clinicianType field value"
		clinician.specialty = "default specialty field value"
		
		delegate = ClinicianDelegate()
		responseObj = delegate.create(clinician)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


