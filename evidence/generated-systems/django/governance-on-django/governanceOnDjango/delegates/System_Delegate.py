from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.System_ import System_
from governanceOnDjango.models.DataProcessingActivity import DataProcessingActivity
from governanceOnDjango.models.RecordsRepository import RecordsRepository
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model System_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class System_Delegate Declaration
#======================================================================
class System_Delegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, system_Id ):
		try:	
			system_ = System_.objects.filter(id=system_Id)
			return system_.first();
		except System_.DoesNotExist:
			raise ProcessingError("System_ with id " + str(system_Id) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, system_):
		for model in serializers.deserialize("json", system_):
			model.save()
			return model;

	def create(self, system_):
		system_.save()
		return system_;

	def saveFromJson(self, system_):
		for model in serializers.deserialize("json", system_):
			model.save()
			return system_;
	
	def save(self, system_):
		system_.save()
		return system_;
	
	def delete(self, system_Id ):
		errMsg = "Failed to delete System_ from db using id " + str(system_Id)
		
		try:
			system_ = System_.objects.get(id=system_Id)
			system_.delete()
			return True
		except System_.DoesNotExist:
			raise ProcessingError("System_ with id " + str(system_Id) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = System_.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all System_ from db")
		except Exception:
			return None;
		
	def addProcessingActivities( self, system_Id, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to add elements " + str(processingActivitiesIds) + " for ProcessingActivities on System_"

		try:
			# get the System_
			system_ = self.get( system_Id ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				system_.processingActivities.add(dataProcessingActivity)
				
			# save it		
			system_.save()
			
			# reload and return the appropriate version
			return self.get( system_Id );
		except System_.DoesNotExist:
			raise ProcessingError(errMsg + " : System_ with id " + str(system_Id) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProcessingActivities( self, system_Id, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to remove elements " + str(processingActivitiesIds) + " for ProcessingActivities on System_"

		try:
			# get the System_
			system_ = self.get( system_Id ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				system_.processingActivities.remove(dataProcessingActivity)
				
			# save it		
			system_.save()
			
			# reload and return the appropriate version
			return self.get( system_Id );
		except System_.DoesNotExist:
			raise ProcessingError(errMsg + " : System_ with id " + str(system_Id) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRecordsRepositories( self, system_Id, recordsRepositoriesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RecordsRepositoryDelegate import RecordsRepositoryDelegate

		errMsg = "Failed to add elements " + str(recordsRepositoriesIds) + " for RecordsRepositories on System_"

		try:
			# get the System_
			system_ = self.get( system_Id ).first()
				
			# split on a comma with no spaces
			idList = recordsRepositoriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the RecordsRepository		
				recordsRepository = RecordsRepositoryDelegate().get(id).first();	
				# add the RecordsRepository
				system_.recordsRepositories.add(recordsRepository)
				
			# save it		
			system_.save()
			
			# reload and return the appropriate version
			return self.get( system_Id );
		except System_.DoesNotExist:
			raise ProcessingError(errMsg + " : System_ with id " + str(system_Id) + " does not exist.")
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRecordsRepositories( self, system_Id, recordsRepositoriesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RecordsRepositoryDelegate import RecordsRepositoryDelegate

		errMsg = "Failed to remove elements " + str(recordsRepositoriesIds) + " for RecordsRepositories on System_"

		try:
			# get the System_
			system_ = self.get( system_Id ).first()
				
			# split on a comma with no spaces
			idList = recordsRepositoriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the RecordsRepository		
				recordsRepository = RecordsRepositoryDelegate().get(id).first();	
				# add the RecordsRepository
				system_.recordsRepositories.remove(recordsRepository)
				
			# save it		
			system_.save()
			
			# reload and return the appropriate version
			return self.get( system_Id );
		except System_.DoesNotExist:
			raise ProcessingError(errMsg + " : System_ with id " + str(system_Id) + " does not exist.")
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
