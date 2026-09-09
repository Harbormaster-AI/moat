import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Collateral import Collateral
from fintechOnDjango.delegates.CollateralDelegate import CollateralDelegate

 #======================================================================
# 
# Encapsulates data for model Collateral
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CollateralTest Declaration
#======================================================================
class CollateralTest (TestCase) :
	def test_crud(self) :
		collateral = Collateral()
		collateral.description = "default description field value"
		collateral.value = "default value field value"
		collateral.collateralType = "default collateralType field value"
		
		delegate = CollateralDelegate()
		responseObj = delegate.create(collateral)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


