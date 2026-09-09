import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Wallet import Wallet
from fintechOnDjango.delegates.WalletDelegate import WalletDelegate

 #======================================================================
# 
# Encapsulates data for model Wallet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WalletTest Declaration
#======================================================================
class WalletTest (TestCase) :
	def test_crud(self) :
		wallet = Wallet()
		wallet.currency = "default currency field value"
		wallet.balance = "default balance field value"
		wallet.status = "default status field value"
		
		delegate = WalletDelegate()
		responseObj = delegate.create(wallet)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


