import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Document import Document
from hrOnDjango.delegates.DocumentDelegate import DocumentDelegate

 #======================================================================
# 
# Encapsulates data for model Document
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DocumentTest Declaration
#======================================================================
class DocumentTest (TestCase) :
	def test_crud(self) :
		document = Document()
		document.name = "default name field value"
		document.fileUrl = "default fileUrl field value"
		document.uploadedDate = datetime.datetime.now()
		document.documentType = "default documentType field value"
		
		delegate = DocumentDelegate()
		responseObj = delegate.create(document)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


