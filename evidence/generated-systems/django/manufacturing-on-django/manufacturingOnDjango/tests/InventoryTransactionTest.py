import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.InventoryTransaction import InventoryTransaction
from manufacturingOnDjango.delegates.InventoryTransactionDelegate import InventoryTransactionDelegate

 #======================================================================
# 
# Encapsulates data for model InventoryTransaction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryTransactionTest Declaration
#======================================================================
class InventoryTransactionTest (TestCase) :
	def test_crud(self) :
		inventoryTransaction = InventoryTransaction()
		inventoryTransaction.transactionNumber = "default transactionNumber field value"
		inventoryTransaction.quantity = "default quantity field value"
		inventoryTransaction.transactionDateTime = "default transactionDateTime field value"
		inventoryTransaction.referenceDocument = "default referenceDocument field value"
		inventoryTransaction.transactionType = "default transactionType field value"
		
		delegate = InventoryTransactionDelegate()
		responseObj = delegate.create(inventoryTransaction)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


