import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.CostCenter import CostCenter
from hrOnDjango.delegates.CostCenterDelegate import CostCenterDelegate

 #======================================================================
# 
# Encapsulates data for model CostCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CostCenterTest Declaration
#======================================================================
class CostCenterTest (TestCase) :
	def test_crud(self) :
		costCenter = CostCenter()
		costCenter.code = "default code field value"
		costCenter.name = "default name field value"
		
		delegate = CostCenterDelegate()
		responseObj = delegate.create(costCenter)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


