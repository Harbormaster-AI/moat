import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.PayrollRun import PayrollRun
from hrOnDjango.delegates.PayrollRunDelegate import PayrollRunDelegate

 #======================================================================
# 
# Encapsulates data for model PayrollRun
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayrollRunTest Declaration
#======================================================================
class PayrollRunTest (TestCase) :
	def test_crud(self) :
		payrollRun = PayrollRun()
		payrollRun.runNumber = "default runNumber field value"
		payrollRun.periodStart = datetime.datetime.now()
		payrollRun.periodEnd = datetime.datetime.now()
		payrollRun.paymentDate = datetime.datetime.now()
		payrollRun.status = "default status field value"
		
		delegate = PayrollRunDelegate()
		responseObj = delegate.create(payrollRun)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


