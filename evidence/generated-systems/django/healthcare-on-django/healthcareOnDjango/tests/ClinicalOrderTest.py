import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.ClinicalOrder import ClinicalOrder
from healthcareOnDjango.delegates.ClinicalOrderDelegate import ClinicalOrderDelegate

 #======================================================================
# 
# Encapsulates data for model ClinicalOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClinicalOrderTest Declaration
#======================================================================
class ClinicalOrderTest (TestCase) :
	def test_crud(self) :
		clinicalOrder = ClinicalOrder()
		clinicalOrder.orderNumber = "default orderNumber field value"
		clinicalOrder.status = "default status field value"
		clinicalOrder.orderType = "default orderType field value"
		clinicalOrder.priority = "default priority field value"
		
		delegate = ClinicalOrderDelegate()
		responseObj = delegate.create(clinicalOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


