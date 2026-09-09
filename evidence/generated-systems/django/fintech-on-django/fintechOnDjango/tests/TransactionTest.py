import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Transaction import Transaction
from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

 #======================================================================
# 
# Encapsulates data for model Transaction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TransactionTest Declaration
#======================================================================
class TransactionTest (TestCase) :
	def test_crud(self) :
		transaction = Transaction()
		transaction.amount = "default amount field value"
		transaction.fee = "default fee field value"
		transaction.exchangeRate = "default exchangeRate field value"
		transaction.createdAt = "default createdAt field value"
		transaction.completedAt = "default completedAt field value"
		transaction.narrative = "default narrative field value"
		transaction.transactionType = "default transactionType field value"
		transaction.status = "default status field value"
		
		delegate = TransactionDelegate()
		responseObj = delegate.create(transaction)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


