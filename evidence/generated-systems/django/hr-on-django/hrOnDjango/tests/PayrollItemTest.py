import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.PayrollItem import PayrollItem
from hrOnDjango.delegates.PayrollItemDelegate import PayrollItemDelegate

 #======================================================================
# 
# Encapsulates data for model PayrollItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayrollItemTest Declaration
#======================================================================
class PayrollItemTest (TestCase) :
	def test_crud(self) :
		payrollItem = PayrollItem()
		payrollItem.amount = "default amount field value"
		payrollItem.taxable = False
		payrollItem.itemType = "default itemType field value"
		
		delegate = PayrollItemDelegate()
		responseObj = delegate.create(payrollItem)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


