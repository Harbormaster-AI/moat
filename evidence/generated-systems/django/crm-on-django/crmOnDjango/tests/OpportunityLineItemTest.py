import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.OpportunityLineItem import OpportunityLineItem
from crmOnDjango.delegates.OpportunityLineItemDelegate import OpportunityLineItemDelegate

 #======================================================================
# 
# Encapsulates data for model OpportunityLineItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OpportunityLineItemTest Declaration
#======================================================================
class OpportunityLineItemTest (TestCase) :
	def test_crud(self) :
		opportunityLineItem = OpportunityLineItem()
		opportunityLineItem.quantity = "default quantity field value"
		opportunityLineItem.unitPrice = "default unitPrice field value"
		opportunityLineItem.discountPercent = "default discountPercent field value"
		opportunityLineItem.totalPrice = "default totalPrice field value"
		
		delegate = OpportunityLineItemDelegate()
		responseObj = delegate.create(opportunityLineItem)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


