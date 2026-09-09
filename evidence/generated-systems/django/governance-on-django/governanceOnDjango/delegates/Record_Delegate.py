from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Record_ import Record_
from governanceOnDjango.models.RecordsRepository import RecordsRepository
from governanceOnDjango.models.RetentionSchedule import RetentionSchedule
from governanceOnDjango.models.DataProcessingActivity import DataProcessingActivity
from governanceOnDjango.models.DataCategory import DataCategory
from governanceOnDjango.models.LegalHold import LegalHold
from governanceOnDjango.models.DataSubjectRequest import DataSubjectRequest
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Record_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Record_Delegate Declaration
#======================================================================
class Record_Delegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, record_Id ):
		try:	
			record_ = Record_.objects.filter(id=record_Id)
			return record_.first();
		except Record_.DoesNotExist:
			raise ProcessingError("Record_ with id " + str(record_Id) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, record_):
		for model in serializers.deserialize("json", record_):
			model.save()
			return model;

	def create(self, record_):
		record_.save()
		return record_;

	def saveFromJson(self, record_):
		for model in serializers.deserialize("json", record_):
			model.save()
			return record_;
	
	def save(self, record_):
		record_.save()
		return record_;
	
	def delete(self, record_Id ):
		errMsg = "Failed to delete Record_ from db using id " + str(record_Id)
		
		try:
			record_ = Record_.objects.get(id=record_Id)
			record_.delete()
			return True
		except Record_.DoesNotExist:
			raise ProcessingError("Record_ with id " + str(record_Id) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Record_.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Record_ from db")
		except Exception:
			return None;
		
	def assignRepository( self, record_Id, repositoryId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RecordsRepositoryDelegate import RecordsRepositoryDelegate

		errMsg = "Failed to assign element " + str(repositoryId) + " for Repository on Record_"

		try:
			# get the Record_ from db
			record_ = self.get( record_Id ).first()	
			
			# get the RecordsRepository from db
			recordsRepository = RecordsRepositoryDelegate().get(repositoryId).first();
			
			# assign the Repository		
			record_.repository = recordsRepository
			
			#save it
			record_.save()

			# reload and return the appropriate version					
			return self.get( record_Id );
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ with id " + str(record_Id) + " does not exist.")
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository with id " + str(repositoryId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRepository( self, record_Id ):
		errMsg = "Failed to unassign element " + str(repositoryId) + " for Repository on Record_"

		try:
			# get the Record_ from db
			record_ = self.get( record_Id ).first()	
			
			# assign to None for unassignment
			record_.recordsRepository = None			

			#save it
			record_.save()

			# reload and return the appropriate version					
			return self.get( record_Id );
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ with id " + str(record_Id) + " does not exist.")
		except Exception:
			return None;
		
	def assignRetentionSchedule( self, record_Id, retentionScheduleId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RetentionScheduleDelegate import RetentionScheduleDelegate

		errMsg = "Failed to assign element " + str(retentionScheduleId) + " for RetentionSchedule on Record_"

		try:
			# get the Record_ from db
			record_ = self.get( record_Id ).first()	
			
			# get the RetentionSchedule from db
			retentionSchedule = RetentionScheduleDelegate().get(retentionScheduleId).first();
			
			# assign the RetentionSchedule		
			record_.retentionSchedule = retentionSchedule
			
			#save it
			record_.save()

			# reload and return the appropriate version					
			return self.get( record_Id );
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ with id " + str(record_Id) + " does not exist.")
		except RetentionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RetentionSchedule with id " + str(retentionScheduleId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRetentionSchedule( self, record_Id ):
		errMsg = "Failed to unassign element " + str(retentionScheduleId) + " for RetentionSchedule on Record_"

		try:
			# get the Record_ from db
			record_ = self.get( record_Id ).first()	
			
			# assign to None for unassignment
			record_.retentionSchedule = None			

			#save it
			record_.save()

			# reload and return the appropriate version					
			return self.get( record_Id );
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ with id " + str(record_Id) + " does not exist.")
		except Exception:
			return None;
		
	def addProcessingActivities( self, record_Id, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to add elements " + str(processingActivitiesIds) + " for ProcessingActivities on Record_"

		try:
			# get the Record_
			record_ = self.get( record_Id ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				record_.processingActivities.add(dataProcessingActivity)
				
			# save it		
			record_.save()
			
			# reload and return the appropriate version
			return self.get( record_Id );
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ with id " + str(record_Id) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProcessingActivities( self, record_Id, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to remove elements " + str(processingActivitiesIds) + " for ProcessingActivities on Record_"

		try:
			# get the Record_
			record_ = self.get( record_Id ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				record_.processingActivities.remove(dataProcessingActivity)
				
			# save it		
			record_.save()
			
			# reload and return the appropriate version
			return self.get( record_Id );
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ with id " + str(record_Id) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDataCategories( self, record_Id, dataCategoriesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataCategoryDelegate import DataCategoryDelegate

		errMsg = "Failed to add elements " + str(dataCategoriesIds) + " for DataCategories on Record_"

		try:
			# get the Record_
			record_ = self.get( record_Id ).first()
				
			# split on a comma with no spaces
			idList = dataCategoriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataCategory		
				dataCategory = DataCategoryDelegate().get(id).first();	
				# add the DataCategory
				record_.dataCategories.add(dataCategory)
				
			# save it		
			record_.save()
			
			# reload and return the appropriate version
			return self.get( record_Id );
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ with id " + str(record_Id) + " does not exist.")
		except DataCategory.DoesNotExist:
			raise ProcessingError(errMsg + " : DataCategory does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDataCategories( self, record_Id, dataCategoriesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataCategoryDelegate import DataCategoryDelegate

		errMsg = "Failed to remove elements " + str(dataCategoriesIds) + " for DataCategories on Record_"

		try:
			# get the Record_
			record_ = self.get( record_Id ).first()
				
			# split on a comma with no spaces
			idList = dataCategoriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataCategory		
				dataCategory = DataCategoryDelegate().get(id).first();	
				# add the DataCategory
				record_.dataCategories.remove(dataCategory)
				
			# save it		
			record_.save()
			
			# reload and return the appropriate version
			return self.get( record_Id );
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ with id " + str(record_Id) + " does not exist.")
		except DataCategory.DoesNotExist:
			raise ProcessingError(errMsg + " : DataCategory does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLegalHolds( self, record_Id, legalHoldsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.LegalHoldDelegate import LegalHoldDelegate

		errMsg = "Failed to add elements " + str(legalHoldsIds) + " for LegalHolds on Record_"

		try:
			# get the Record_
			record_ = self.get( record_Id ).first()
				
			# split on a comma with no spaces
			idList = legalHoldsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LegalHold		
				legalHold = LegalHoldDelegate().get(id).first();	
				# add the LegalHold
				record_.legalHolds.add(legalHold)
				
			# save it		
			record_.save()
			
			# reload and return the appropriate version
			return self.get( record_Id );
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ with id " + str(record_Id) + " does not exist.")
		except LegalHold.DoesNotExist:
			raise ProcessingError(errMsg + " : LegalHold does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLegalHolds( self, record_Id, legalHoldsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.LegalHoldDelegate import LegalHoldDelegate

		errMsg = "Failed to remove elements " + str(legalHoldsIds) + " for LegalHolds on Record_"

		try:
			# get the Record_
			record_ = self.get( record_Id ).first()
				
			# split on a comma with no spaces
			idList = legalHoldsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LegalHold		
				legalHold = LegalHoldDelegate().get(id).first();	
				# add the LegalHold
				record_.legalHolds.remove(legalHold)
				
			# save it		
			record_.save()
			
			# reload and return the appropriate version
			return self.get( record_Id );
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ with id " + str(record_Id) + " does not exist.")
		except LegalHold.DoesNotExist:
			raise ProcessingError(errMsg + " : LegalHold does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDataSubjectRequests( self, record_Id, dataSubjectRequestsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataSubjectRequestDelegate import DataSubjectRequestDelegate

		errMsg = "Failed to add elements " + str(dataSubjectRequestsIds) + " for DataSubjectRequests on Record_"

		try:
			# get the Record_
			record_ = self.get( record_Id ).first()
				
			# split on a comma with no spaces
			idList = dataSubjectRequestsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSubjectRequest		
				dataSubjectRequest = DataSubjectRequestDelegate().get(id).first();	
				# add the DataSubjectRequest
				record_.dataSubjectRequests.add(dataSubjectRequest)
				
			# save it		
			record_.save()
			
			# reload and return the appropriate version
			return self.get( record_Id );
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ with id " + str(record_Id) + " does not exist.")
		except DataSubjectRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSubjectRequest does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDataSubjectRequests( self, record_Id, dataSubjectRequestsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataSubjectRequestDelegate import DataSubjectRequestDelegate

		errMsg = "Failed to remove elements " + str(dataSubjectRequestsIds) + " for DataSubjectRequests on Record_"

		try:
			# get the Record_
			record_ = self.get( record_Id ).first()
				
			# split on a comma with no spaces
			idList = dataSubjectRequestsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSubjectRequest		
				dataSubjectRequest = DataSubjectRequestDelegate().get(id).first();	
				# add the DataSubjectRequest
				record_.dataSubjectRequests.remove(dataSubjectRequest)
				
			# save it		
			record_.save()
			
			# reload and return the appropriate version
			return self.get( record_Id );
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ with id " + str(record_Id) + " does not exist.")
		except DataSubjectRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSubjectRequest does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
