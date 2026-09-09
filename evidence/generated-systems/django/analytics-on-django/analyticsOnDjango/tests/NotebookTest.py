import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Notebook import Notebook
from analyticsOnDjango.delegates.NotebookDelegate import NotebookDelegate

 #======================================================================
# 
# Encapsulates data for model Notebook
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class NotebookTest Declaration
#======================================================================
class NotebookTest (TestCase) :
	def test_crud(self) :
		notebook = Notebook()
		notebook.title = "default title field value"
		notebook.repository = "default repository field value"
		notebook.language = "default language field value"
		
		delegate = NotebookDelegate()
		responseObj = delegate.create(notebook)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


