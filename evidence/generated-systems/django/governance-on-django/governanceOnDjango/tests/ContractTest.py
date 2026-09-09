import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Contract import Contract
from governanceOnDjango.delegates.ContractDelegate import ContractDelegate

 #======================================================================
# 
# Encapsulates data for model Contract
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ContractTest Declaration
#======================================================================
class ContractTest (TestCase) :
	def test_crud(self) :
		contract = Contract()
		contract.title = "default title field value"
		contract.effectiveDate = datetime.datetime.now()
		contract.expiryDate = datetime.datetime.now()
		contract.repositoryUrl = "default repositoryUrl field value"
		contract.status = "default status field value"
		
		delegate = ContractDelegate()
		responseObj = delegate.create(contract)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


