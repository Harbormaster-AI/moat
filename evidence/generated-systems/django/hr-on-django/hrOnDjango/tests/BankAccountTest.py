import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.BankAccount import BankAccount
from hrOnDjango.delegates.BankAccountDelegate import BankAccountDelegate

 #======================================================================
# 
# Encapsulates data for model BankAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BankAccountTest Declaration
#======================================================================
class BankAccountTest (TestCase) :
	def test_crud(self) :
		bankAccount = BankAccount()
		bankAccount.accountHolder = "default accountHolder field value"
		bankAccount.bankName = "default bankName field value"
		bankAccount.iban = "default iban field value"
		bankAccount.bic = "default bic field value"
		bankAccount.accountNumber = "default accountNumber field value"
		bankAccount.routingNumber = "default routingNumber field value"
		
		delegate = BankAccountDelegate()
		responseObj = delegate.create(bankAccount)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


