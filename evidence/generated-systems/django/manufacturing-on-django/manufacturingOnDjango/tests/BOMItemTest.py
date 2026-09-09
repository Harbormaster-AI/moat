import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.BOMItem import BOMItem
from manufacturingOnDjango.delegates.BOMItemDelegate import BOMItemDelegate

 #======================================================================
# 
# Encapsulates data for model BOMItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BOMItemTest Declaration
#======================================================================
class BOMItemTest (TestCase) :
	def test_crud(self) :
		bOMItem = BOMItem()
		bOMItem.lineNumber = 22
		bOMItem.quantity = "default quantity field value"
		bOMItem.scrapPercent = "default scrapPercent field value"
		
		delegate = BOMItemDelegate()
		responseObj = delegate.create(bOMItem)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


