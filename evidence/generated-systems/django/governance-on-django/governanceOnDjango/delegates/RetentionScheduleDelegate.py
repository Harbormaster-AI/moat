from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.RetentionSchedule import RetentionSchedule
from governanceOnDjango.models.RecordsRepository import RecordsRepository
from governanceOnDjango.models.Record_ import Record_
from governanceOnDjango.models.Exception_ import Exception_
from governanceOnDjango.models.DispositionReview import DispositionReview
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model RetentionSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RetentionScheduleDelegate Declaration
#======================================================================
class RetentionScheduleDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, retentionScheduleId ):
		try:	
			retentionSchedule = RetentionSchedule.objects.filter(id=retentionScheduleId)
			return retentionSchedule.first();
		except RetentionSchedule.DoesNotExist:
			raise ProcessingError("RetentionSchedule with id " + str(retentionScheduleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, retentionSchedule):
		for model in serializers.deserialize("json", retentionSchedule):
			model.save()
			return model;

	def create(self, retentionSchedule):
		retentionSchedule.save()
		return retentionSchedule;

	def saveFromJson(self, retentionSchedule):
		for model in serializers.deserialize("json", retentionSchedule):
			model.save()
			return retentionSchedule;
	
	def save(self, retentionSchedule):
		retentionSchedule.save()
		return retentionSchedule;
	
	def delete(self, retentionScheduleId ):
		errMsg = "Failed to delete RetentionSchedule from db using id " + str(retentionScheduleId)
		
		try:
			retentionSchedule = RetentionSchedule.objects.get(id=retentionScheduleId)
			retentionSchedule.delete()
			return True
		except RetentionSchedule.DoesNotExist:
			raise ProcessingError("RetentionSchedule with id " + str(retentionScheduleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = RetentionSchedule.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all RetentionSchedule from db")
		except Exception:
			return None;
		
	def addRepositories( self, retentionScheduleId, repositoriesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RecordsRepositoryDelegate import RecordsRepositoryDelegate

		errMsg = "Failed to add elements " + str(repositoriesIds) + " for Repositories on RetentionSchedule"

		try:
			# get the RetentionSchedule
			retentionSchedule = self.get( retentionScheduleId ).first()
				
			# split on a comma with no spaces
			idList = repositoriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the RecordsRepository		
				recordsRepository = RecordsRepositoryDelegate().get(id).first();	
				# add the RecordsRepository
				retentionSchedule.repositories.add(recordsRepository)
				
			# save it		
			retentionSchedule.save()
			
			# reload and return the appropriate version
			return self.get( retentionScheduleId );
		except RetentionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RetentionSchedule with id " + str(retentionScheduleId) + " does not exist.")
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRepositories( self, retentionScheduleId, repositoriesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RecordsRepositoryDelegate import RecordsRepositoryDelegate

		errMsg = "Failed to remove elements " + str(repositoriesIds) + " for Repositories on RetentionSchedule"

		try:
			# get the RetentionSchedule
			retentionSchedule = self.get( retentionScheduleId ).first()
				
			# split on a comma with no spaces
			idList = repositoriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the RecordsRepository		
				recordsRepository = RecordsRepositoryDelegate().get(id).first();	
				# add the RecordsRepository
				retentionSchedule.repositories.remove(recordsRepository)
				
			# save it		
			retentionSchedule.save()
			
			# reload and return the appropriate version
			return self.get( retentionScheduleId );
		except RetentionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RetentionSchedule with id " + str(retentionScheduleId) + " does not exist.")
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRecords( self, retentionScheduleId, recordsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Record_Delegate import Record_Delegate

		errMsg = "Failed to add elements " + str(recordsIds) + " for Records on RetentionSchedule"

		try:
			# get the RetentionSchedule
			retentionSchedule = self.get( retentionScheduleId ).first()
				
			# split on a comma with no spaces
			idList = recordsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Record_		
				record_ = Record_Delegate().get(id).first();	
				# add the Record_
				retentionSchedule.records.add(record_)
				
			# save it		
			retentionSchedule.save()
			
			# reload and return the appropriate version
			return self.get( retentionScheduleId );
		except RetentionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RetentionSchedule with id " + str(retentionScheduleId) + " does not exist.")
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRecords( self, retentionScheduleId, recordsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Record_Delegate import Record_Delegate

		errMsg = "Failed to remove elements " + str(recordsIds) + " for Records on RetentionSchedule"

		try:
			# get the RetentionSchedule
			retentionSchedule = self.get( retentionScheduleId ).first()
				
			# split on a comma with no spaces
			idList = recordsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Record_		
				record_ = Record_Delegate().get(id).first();	
				# add the Record_
				retentionSchedule.records.remove(record_)
				
			# save it		
			retentionSchedule.save()
			
			# reload and return the appropriate version
			return self.get( retentionScheduleId );
		except RetentionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RetentionSchedule with id " + str(retentionScheduleId) + " does not exist.")
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addExceptions( self, retentionScheduleId, exceptionsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Exception_Delegate import Exception_Delegate

		errMsg = "Failed to add elements " + str(exceptionsIds) + " for Exceptions on RetentionSchedule"

		try:
			# get the RetentionSchedule
			retentionSchedule = self.get( retentionScheduleId ).first()
				
			# split on a comma with no spaces
			idList = exceptionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Exception_		
				exception_ = Exception_Delegate().get(id).first();	
				# add the Exception_
				retentionSchedule.exceptions.add(exception_)
				
			# save it		
			retentionSchedule.save()
			
			# reload and return the appropriate version
			return self.get( retentionScheduleId );
		except RetentionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RetentionSchedule with id " + str(retentionScheduleId) + " does not exist.")
		except Exception_.DoesNotExist:
			raise ProcessingError(errMsg + " : Exception_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeExceptions( self, retentionScheduleId, exceptionsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Exception_Delegate import Exception_Delegate

		errMsg = "Failed to remove elements " + str(exceptionsIds) + " for Exceptions on RetentionSchedule"

		try:
			# get the RetentionSchedule
			retentionSchedule = self.get( retentionScheduleId ).first()
				
			# split on a comma with no spaces
			idList = exceptionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Exception_		
				exception_ = Exception_Delegate().get(id).first();	
				# add the Exception_
				retentionSchedule.exceptions.remove(exception_)
				
			# save it		
			retentionSchedule.save()
			
			# reload and return the appropriate version
			return self.get( retentionScheduleId );
		except RetentionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RetentionSchedule with id " + str(retentionScheduleId) + " does not exist.")
		except Exception_.DoesNotExist:
			raise ProcessingError(errMsg + " : Exception_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDispositionReviews( self, retentionScheduleId, dispositionReviewsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DispositionReviewDelegate import DispositionReviewDelegate

		errMsg = "Failed to add elements " + str(dispositionReviewsIds) + " for DispositionReviews on RetentionSchedule"

		try:
			# get the RetentionSchedule
			retentionSchedule = self.get( retentionScheduleId ).first()
				
			# split on a comma with no spaces
			idList = dispositionReviewsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DispositionReview		
				dispositionReview = DispositionReviewDelegate().get(id).first();	
				# add the DispositionReview
				retentionSchedule.dispositionReviews.add(dispositionReview)
				
			# save it		
			retentionSchedule.save()
			
			# reload and return the appropriate version
			return self.get( retentionScheduleId );
		except RetentionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RetentionSchedule with id " + str(retentionScheduleId) + " does not exist.")
		except DispositionReview.DoesNotExist:
			raise ProcessingError(errMsg + " : DispositionReview does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDispositionReviews( self, retentionScheduleId, dispositionReviewsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DispositionReviewDelegate import DispositionReviewDelegate

		errMsg = "Failed to remove elements " + str(dispositionReviewsIds) + " for DispositionReviews on RetentionSchedule"

		try:
			# get the RetentionSchedule
			retentionSchedule = self.get( retentionScheduleId ).first()
				
			# split on a comma with no spaces
			idList = dispositionReviewsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DispositionReview		
				dispositionReview = DispositionReviewDelegate().get(id).first();	
				# add the DispositionReview
				retentionSchedule.dispositionReviews.remove(dispositionReview)
				
			# save it		
			retentionSchedule.save()
			
			# reload and return the appropriate version
			return self.get( retentionScheduleId );
		except RetentionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RetentionSchedule with id " + str(retentionScheduleId) + " does not exist.")
		except DispositionReview.DoesNotExist:
			raise ProcessingError(errMsg + " : DispositionReview does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
