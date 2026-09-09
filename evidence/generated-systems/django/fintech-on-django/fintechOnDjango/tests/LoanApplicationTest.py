import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.LoanApplication import LoanApplication
from fintechOnDjango.delegates.LoanApplicationDelegate import LoanApplicationDelegate

 #======================================================================
# 
# Encapsulates data for model LoanApplication
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LoanApplicationTest Declaration
#======================================================================
class LoanApplicationTest (TestCase) :
	def test_crud(self) :
		loanApplication = LoanApplication()
		loanApplication.applicationNumber = "default applicationNumber field value"
		loanApplication.amountRequested = "default amountRequested field value"
		loanApplication.termMonths = 22
		loanApplication.submittedAt = "default submittedAt field value"
		loanApplication.product = "default product field value"
		loanApplication.purpose = "default purpose field value"
		loanApplication.status = "default status field value"
		
		delegate = LoanApplicationDelegate()
		responseObj = delegate.create(loanApplication)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


