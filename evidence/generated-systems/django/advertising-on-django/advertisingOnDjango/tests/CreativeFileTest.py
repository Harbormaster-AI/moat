import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.CreativeFile import CreativeFile
from advertisingOnDjango.delegates.CreativeFileDelegate import CreativeFileDelegate

 #======================================================================
# 
# Encapsulates data for model CreativeFile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeFileTest Declaration
#======================================================================
class CreativeFileTest (TestCase) :
	def test_crud(self) :
		creativeFile = CreativeFile()
		creativeFile.uri = "default uri field value"
		creativeFile.fileSizeKB = 22
		creativeFile.mimeType = "default mimeType field value"
		creativeFile.checksum = "default checksum field value"
		
		delegate = CreativeFileDelegate()
		responseObj = delegate.create(creativeFile)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


