import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.ShipmentItem import ShipmentItem
from ecommerceOnDjango.delegates.ShipmentItemDelegate import ShipmentItemDelegate

 #======================================================================
# 
# Encapsulates data for model ShipmentItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShipmentItemTest Declaration
#======================================================================
class ShipmentItemTest (TestCase) :
	def test_crud(self) :
		shipmentItem = ShipmentItem()
		shipmentItem.quantity = 22
		
		delegate = ShipmentItemDelegate()
		responseObj = delegate.create(shipmentItem)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


