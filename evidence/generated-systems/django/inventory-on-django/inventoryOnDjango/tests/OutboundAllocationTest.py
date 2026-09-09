import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.OutboundAllocation import OutboundAllocation
from inventoryOnDjango.delegates.OutboundAllocationDelegate import OutboundAllocationDelegate

 #======================================================================
# 
# Encapsulates data for model OutboundAllocation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OutboundAllocationTest Declaration
#======================================================================
class OutboundAllocationTest (TestCase) :
	def test_crud(self) :
		outboundAllocation = OutboundAllocation()
		outboundAllocation.allocationNumber = "default allocationNumber field value"
		outboundAllocation.allocatedQuantity = "default allocatedQuantity field value"
		outboundAllocation.allocationDate = datetime.datetime.now()
		outboundAllocation.status = "default status field value"
		
		delegate = OutboundAllocationDelegate()
		responseObj = delegate.create(outboundAllocation)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


