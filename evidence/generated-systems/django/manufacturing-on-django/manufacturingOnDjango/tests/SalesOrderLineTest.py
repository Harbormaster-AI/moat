import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.SalesOrderLine import SalesOrderLine
from manufacturingOnDjango.delegates.SalesOrderLineDelegate import SalesOrderLineDelegate

 #======================================================================
# 
# Encapsulates data for model SalesOrderLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesOrderLineTest Declaration
#======================================================================
class SalesOrderLineTest (TestCase) :
	def test_crud(self) :
		salesOrderLine = SalesOrderLine()
		salesOrderLine.lineNumber = 22
		salesOrderLine.quantity = "default quantity field value"
		salesOrderLine.unitPrice = "default unitPrice field value"
		salesOrderLine.dueDate = datetime.datetime.now()
		
		delegate = SalesOrderLineDelegate()
		responseObj = delegate.create(salesOrderLine)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


