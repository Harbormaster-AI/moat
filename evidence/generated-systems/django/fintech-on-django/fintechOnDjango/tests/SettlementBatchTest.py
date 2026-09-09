import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.SettlementBatch import SettlementBatch
from fintechOnDjango.delegates.SettlementBatchDelegate import SettlementBatchDelegate

 #======================================================================
# 
# Encapsulates data for model SettlementBatch
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SettlementBatchTest Declaration
#======================================================================
class SettlementBatchTest (TestCase) :
	def test_crud(self) :
		settlementBatch = SettlementBatch()
		settlementBatch.batchId = "default batchId field value"
		settlementBatch.periodStart = "default periodStart field value"
		settlementBatch.periodEnd = "default periodEnd field value"
		settlementBatch.totalVolume = "default totalVolume field value"
		settlementBatch.totalCount = 22
		settlementBatch.status = "default status field value"
		
		delegate = SettlementBatchDelegate()
		responseObj = delegate.create(settlementBatch)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


