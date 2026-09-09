import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.Contract import Contract
from crmOnDjango.delegates.ContractDelegate import ContractDelegate

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
		contract.contractNumber = "default contractNumber field value"
		contract.startDate = datetime.datetime.now()
		contract.endDate = datetime.datetime.now()
		contract.renewalTermMonths = 22
		contract.autoRenew = False
		contract.status = "default status field value"
		
		delegate = ContractDelegate()
		responseObj = delegate.create(contract)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


