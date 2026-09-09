import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.CardTokenization import CardTokenization
from fintechOnDjango.delegates.CardTokenizationDelegate import CardTokenizationDelegate

 #======================================================================
# 
# Encapsulates data for model CardTokenization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CardTokenizationTest Declaration
#======================================================================
class CardTokenizationTest (TestCase) :
	def test_crud(self) :
		cardTokenization = CardTokenization()
		cardTokenization.tokenReference = "default tokenReference field value"
		cardTokenization.createdAt = "default createdAt field value"
		cardTokenization.walletProvider = "default walletProvider field value"
		cardTokenization.status = "default status field value"
		
		delegate = CardTokenizationDelegate()
		responseObj = delegate.create(cardTokenization)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


