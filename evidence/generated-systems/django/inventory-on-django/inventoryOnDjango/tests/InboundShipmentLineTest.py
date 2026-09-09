import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.InboundShipmentLine import InboundShipmentLine
from inventoryOnDjango.delegates.InboundShipmentLineDelegate import InboundShipmentLineDelegate

 #======================================================================
# 
# Encapsulates data for model InboundShipmentLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InboundShipmentLineTest Declaration
#======================================================================
class InboundShipmentLineTest (TestCase) :
	def test_crud(self) :
		inboundShipmentLine = InboundShipmentLine()
		inboundShipmentLine.lineNumber = 22
		inboundShipmentLine.quantity = "default quantity field value"
		inboundShipmentLine.unitOfMeasure = "default unitOfMeasure field value"
		inboundShipmentLine.stockStatus = "default stockStatus field value"
		
		delegate = InboundShipmentLineDelegate()
		responseObj = delegate.create(inboundShipmentLine)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


