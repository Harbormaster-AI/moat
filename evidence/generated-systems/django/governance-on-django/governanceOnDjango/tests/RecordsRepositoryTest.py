import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.RecordsRepository import RecordsRepository
from governanceOnDjango.delegates.RecordsRepositoryDelegate import RecordsRepositoryDelegate

 #======================================================================
# 
# Encapsulates data for model RecordsRepository
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RecordsRepositoryTest Declaration
#======================================================================
class RecordsRepositoryTest (TestCase) :
	def test_crud(self) :
		recordsRepository = RecordsRepository()
		recordsRepository.name = "default name field value"
		recordsRepository.location = "default location field value"
		recordsRepository.ownerDepartment = "default ownerDepartment field value"
		recordsRepository.repositoryType = "default repositoryType field value"
		
		delegate = RecordsRepositoryDelegate()
		responseObj = delegate.create(recordsRepository)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


