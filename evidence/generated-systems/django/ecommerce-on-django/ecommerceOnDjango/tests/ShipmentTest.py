import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Shipment import Shipment
from ecommerceOnDjango.delegates.ShipmentDelegate import ShipmentDelegate

 #======================================================================
# 
# Encapsulates data for model Shipment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShipmentTest Declaration
#======================================================================
class ShipmentTest (TestCase) :
	def test_crud(self) :
		shipment = Shipment()
		shipment.shipmentNumber = "default shipmentNumber field value"
		shipment.shippedDate = datetime.datetime.now()
		shipment.deliveredDate = datetime.datetime.now()
		shipment.trackingNumber = "default trackingNumber field value"
		shipment.shippingAddress = "default shippingAddress field value"
		shipment.status = "default status field value"
		shipment.carrier = "default carrier field value"
		
		delegate = ShipmentDelegate()
		responseObj = delegate.create(shipment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


