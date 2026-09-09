import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.InboundShipment import InboundShipment
from inventoryOnDjango.delegates.InboundShipmentDelegate import InboundShipmentDelegate

 #======================================================================
# 
# Encapsulates data for model InboundShipment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InboundShipmentTest Declaration
#======================================================================
class InboundShipmentTest (TestCase) :
	def test_crud(self) :
		inboundShipment = InboundShipment()
		inboundShipment.shipmentNumber = "default shipmentNumber field value"
		inboundShipment.expectedArrivalDate = datetime.datetime.now()
		inboundShipment.arrivalDate = datetime.datetime.now()
		inboundShipment.carrierName = "default carrierName field value"
		inboundShipment.status = "default status field value"
		
		delegate = InboundShipmentDelegate()
		responseObj = delegate.create(inboundShipment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


