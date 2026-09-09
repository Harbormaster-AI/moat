import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.LineItem import LineItem
from advertisingOnDjango.delegates.LineItemDelegate import LineItemDelegate

 #======================================================================
# 
# Encapsulates data for model LineItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LineItemTest Declaration
#======================================================================
class LineItemTest (TestCase) :
	def test_crud(self) :
		lineItem = LineItem()
		lineItem.name = "default name field value"
		lineItem.bidAmount = "default bidAmount field value"
		lineItem.dailyBudget = "default dailyBudget field value"
		lineItem.frequencyCap = "default frequencyCap field value"
		lineItem.status = "default status field value"
		lineItem.pricingModel = "default pricingModel field value"
		lineItem.bidStrategy = "default bidStrategy field value"
		lineItem.pacing = "default pacing field value"
		
		delegate = LineItemDelegate()
		responseObj = delegate.create(lineItem)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


