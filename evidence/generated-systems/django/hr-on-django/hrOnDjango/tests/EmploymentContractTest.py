import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.EmploymentContract import EmploymentContract
from hrOnDjango.delegates.EmploymentContractDelegate import EmploymentContractDelegate

 #======================================================================
# 
# Encapsulates data for model EmploymentContract
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmploymentContractTest Declaration
#======================================================================
class EmploymentContractTest (TestCase) :
	def test_crud(self) :
		employmentContract = EmploymentContract()
		employmentContract.contractNumber = "default contractNumber field value"
		employmentContract.startDate = datetime.datetime.now()
		employmentContract.endDate = datetime.datetime.now()
		employmentContract.workHoursPerWeek = "default workHoursPerWeek field value"
		employmentContract.employmentType = "default employmentType field value"
		employmentContract.status = "default status field value"
		employmentContract.payFrequency = "default payFrequency field value"
		
		delegate = EmploymentContractDelegate()
		responseObj = delegate.create(employmentContract)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


