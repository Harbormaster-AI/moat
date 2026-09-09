import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.InspectionLot import InspectionLot
from manufacturingOnDjango.delegates.InspectionLotDelegate import InspectionLotDelegate

 #======================================================================
# 
# Encapsulates data for model InspectionLot
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionLotTest Declaration
#======================================================================
class InspectionLotTest (TestCase) :
	def test_crud(self) :
		inspectionLot = InspectionLot()
		inspectionLot.lotNumber = "default lotNumber field value"
		inspectionLot.quantity = "default quantity field value"
		inspectionLot.sampleSize = 22
		inspectionLot.createdOn = "default createdOn field value"
		inspectionLot.inspectionType = "default inspectionType field value"
		inspectionLot.status = "default status field value"
		
		delegate = InspectionLotDelegate()
		responseObj = delegate.create(inspectionLot)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


