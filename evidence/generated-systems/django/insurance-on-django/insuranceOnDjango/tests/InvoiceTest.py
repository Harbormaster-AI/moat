import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.Invoice import Invoice
from insuranceOnDjango.delegates.InvoiceDelegate import InvoiceDelegate

 #======================================================================
# 
# Encapsulates data for model Invoice
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InvoiceTest Declaration
#======================================================================
class InvoiceTest (TestCase) :
	def test_crud(self) :
		invoice = Invoice()
		invoice.invoiceNumber = "default invoiceNumber field value"
		invoice.dueDate = datetime.datetime.now()
		invoice.totalDue = "default totalDue field value"
		invoice.status = "default status field value"
		
		delegate = InvoiceDelegate()
		responseObj = delegate.create(invoice)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


