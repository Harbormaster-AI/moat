import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.Warranty import Warranty
from aerospaceOnDjango.delegates.WarrantyDelegate import WarrantyDelegate

 #======================================================================
# 
# Encapsulates data for model Warranty
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WarrantyTest Declaration
#======================================================================
class WarrantyTest (TestCase) :
	def test_crud(self) :
		warranty = Warranty()
		warranty.coverageMonths = 22
		warranty.warrantyType = "default warrantyType field value"
		
		delegate = WarrantyDelegate()
		responseObj = delegate.create(warranty)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


