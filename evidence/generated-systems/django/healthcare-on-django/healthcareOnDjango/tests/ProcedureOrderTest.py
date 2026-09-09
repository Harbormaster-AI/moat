import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.ProcedureOrder import ProcedureOrder
from healthcareOnDjango.delegates.ProcedureOrderDelegate import ProcedureOrderDelegate

 #======================================================================
# 
# Encapsulates data for model ProcedureOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProcedureOrderTest Declaration
#======================================================================
class ProcedureOrderTest (TestCase) :
	def test_crud(self) :
		procedureOrder = ProcedureOrder()
		procedureOrder.procedureCode = "default procedureCode field value"
		procedureOrder.consentObtained = False
		procedureOrder.anesthesiaType = "default anesthesiaType field value"
		
		delegate = ProcedureOrderDelegate()
		responseObj = delegate.create(procedureOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


