import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.ReplenishmentPolicy import ReplenishmentPolicy
from inventoryOnDjango.delegates.ReplenishmentPolicyDelegate import ReplenishmentPolicyDelegate

 #======================================================================
# 
# Encapsulates data for model ReplenishmentPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReplenishmentPolicyTest Declaration
#======================================================================
class ReplenishmentPolicyTest (TestCase) :
	def test_crud(self) :
		replenishmentPolicy = ReplenishmentPolicy()
		replenishmentPolicy.minLevel = "default minLevel field value"
		replenishmentPolicy.maxLevel = "default maxLevel field value"
		replenishmentPolicy.reorderPoint = "default reorderPoint field value"
		replenishmentPolicy.reorderQuantity = "default reorderQuantity field value"
		replenishmentPolicy.leadTimeDays = 22
		replenishmentPolicy.reviewPeriodDays = 22
		replenishmentPolicy.policyType = "default policyType field value"
		
		delegate = ReplenishmentPolicyDelegate()
		responseObj = delegate.create(replenishmentPolicy)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


