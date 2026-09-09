import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Loan import Loan
from fintechOnDjango.delegates.LoanDelegate import LoanDelegate

 #======================================================================
# 
# Encapsulates data for model Loan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LoanTest Declaration
#======================================================================
class LoanTest (TestCase) :
	def test_crud(self) :
		loan = Loan()
		loan.loanNumber = "default loanNumber field value"
		loan.principal = "default principal field value"
		loan.interestRate = "default interestRate field value"
		loan.originationDate = datetime.datetime.now()
		loan.maturityDate = datetime.datetime.now()
		loan.rateType = "default rateType field value"
		loan.status = "default status field value"
		
		delegate = LoanDelegate()
		responseObj = delegate.create(loan)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


