import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Record_ import Record_
from governanceOnDjango.delegates.Record_Delegate import Record_Delegate

 #======================================================================
# 
# Encapsulates data for model Record_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Record_Test Declaration
#======================================================================
class Record_Test (TestCase) :
	def test_crud(self) :
		record_ = Record_()
		record_.title = "default title field value"
		record_.creationDate = datetime.datetime.now()
		record_.recordType = "default recordType field value"
		record_.classification = "default classification field value"
		record_.status = "default status field value"
		
		delegate = Record_Delegate()
		responseObj = delegate.create(record_)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


