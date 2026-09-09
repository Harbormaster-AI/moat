import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Creditor import Creditor
from fintechOnDjango.delegates.CreditorDelegate import CreditorDelegate

 #======================================================================
# 
# Encapsulates data for model Creditor
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreditorTest Declaration
#======================================================================
class CreditorTest (TestCase) :
	def test_crud(self) :
		creditor = Creditor()
		creditor.name = "default name field value"
		creditor.bic = "default bic field value"
		creditor.address = "default address field value"
		
		delegate = CreditorDelegate()
		responseObj = delegate.create(creditor)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


