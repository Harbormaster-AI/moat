import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Procedure import Procedure
from healthcareOnDjango.delegates.ProcedureDelegate import ProcedureDelegate

 #======================================================================
# 
# Encapsulates data for model Procedure
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProcedureTest Declaration
#======================================================================
class ProcedureTest (TestCase) :
	def test_crud(self) :
		procedure = Procedure()
		procedure.procedureCode = "default procedureCode field value"
		procedure.startDateTime = "default startDateTime field value"
		procedure.endDateTime = "default endDateTime field value"
		procedure.status = "default status field value"
		
		delegate = ProcedureDelegate()
		responseObj = delegate.create(procedure)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


