import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Discharge import Discharge
from healthcareOnDjango.delegates.DischargeDelegate import DischargeDelegate

 #======================================================================
# 
# Encapsulates data for model Discharge
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DischargeTest Declaration
#======================================================================
class DischargeTest (TestCase) :
	def test_crud(self) :
		discharge = Discharge()
		discharge.dischargeDateTime = "default dischargeDateTime field value"
		discharge.disposition = "default disposition field value"
		
		delegate = DischargeDelegate()
		responseObj = delegate.create(discharge)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


