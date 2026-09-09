import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.LoanTransaction import LoanTransaction
from fintechOnDjango.delegates.LoanTransactionDelegate import LoanTransactionDelegate

 #======================================================================
# 
# Encapsulates data for model LoanTransaction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LoanTransactionTest Declaration
#======================================================================
class LoanTransactionTest (TestCase) :
	def test_crud(self) :
		loanTransaction = LoanTransaction()
		loanTransaction.transactionId = "default transactionId field value"
		loanTransaction.amount = "default amount field value"
		loanTransaction.postingDate = datetime.datetime.now()
		loanTransaction.type = "default type field value"
		loanTransaction.status = "default status field value"
		
		delegate = LoanTransactionDelegate()
		responseObj = delegate.create(loanTransaction)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


