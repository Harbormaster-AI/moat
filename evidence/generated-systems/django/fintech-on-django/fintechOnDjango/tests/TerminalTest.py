import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Terminal import Terminal
from fintechOnDjango.delegates.TerminalDelegate import TerminalDelegate

 #======================================================================
# 
# Encapsulates data for model Terminal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TerminalTest Declaration
#======================================================================
class TerminalTest (TestCase) :
	def test_crud(self) :
		terminal = Terminal()
		terminal.location = "default location field value"
		terminal.type = "default type field value"
		terminal.status = "default status field value"
		
		delegate = TerminalDelegate()
		responseObj = delegate.create(terminal)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


