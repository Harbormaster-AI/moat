import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Procedure import Procedure
from governanceOnDjango.delegates.ProcedureDelegate import ProcedureDelegate

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
		procedure.title = "default title field value"
		procedure.versionLabel = "default versionLabel field value"
		procedure.status = "default status field value"
		
		delegate = ProcedureDelegate()
		responseObj = delegate.create(procedure)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


