from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.DataCategory import DataCategory
from governanceOnDjango.models.DataProcessingActivity import DataProcessingActivity
from governanceOnDjango.models.Record_ import Record_
from governanceOnDjango.models.DataBreach import DataBreach
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model DataCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataCategoryDelegate Declaration
#======================================================================
class DataCategoryDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dataCategoryId ):
		try:	
			dataCategory = DataCategory.objects.filter(id=dataCategoryId)
			return dataCategory.first();
		except DataCategory.DoesNotExist:
			raise ProcessingError("DataCategory with id " + str(dataCategoryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, dataCategory):
		for model in serializers.deserialize("json", dataCategory):
			model.save()
			return model;

	def create(self, dataCategory):
		dataCategory.save()
		return dataCategory;

	def saveFromJson(self, dataCategory):
		for model in serializers.deserialize("json", dataCategory):
			model.save()
			return dataCategory;
	
	def save(self, dataCategory):
		dataCategory.save()
		return dataCategory;
	
	def delete(self, dataCategoryId ):
		errMsg = "Failed to delete DataCategory from db using id " + str(dataCategoryId)
		
		try:
			dataCategory = DataCategory.objects.get(id=dataCategoryId)
			dataCategory.delete()
			return True
		except DataCategory.DoesNotExist:
			raise ProcessingError("DataCategory with id " + str(dataCategoryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = DataCategory.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all DataCategory from db")
		except Exception:
			return None;
		
	def addProcessingActivities( self, dataCategoryId, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to add elements " + str(processingActivitiesIds) + " for ProcessingActivities on DataCategory"

		try:
			# get the DataCategory
			dataCategory = self.get( dataCategoryId ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				dataCategory.processingActivities.add(dataProcessingActivity)
				
			# save it		
			dataCategory.save()
			
			# reload and return the appropriate version
			return self.get( dataCategoryId );
		except DataCategory.DoesNotExist:
			raise ProcessingError(errMsg + " : DataCategory with id " + str(dataCategoryId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProcessingActivities( self, dataCategoryId, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to remove elements " + str(processingActivitiesIds) + " for ProcessingActivities on DataCategory"

		try:
			# get the DataCategory
			dataCategory = self.get( dataCategoryId ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				dataCategory.processingActivities.remove(dataProcessingActivity)
				
			# save it		
			dataCategory.save()
			
			# reload and return the appropriate version
			return self.get( dataCategoryId );
		except DataCategory.DoesNotExist:
			raise ProcessingError(errMsg + " : DataCategory with id " + str(dataCategoryId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRecords( self, dataCategoryId, recordsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Record_Delegate import Record_Delegate

		errMsg = "Failed to add elements " + str(recordsIds) + " for Records on DataCategory"

		try:
			# get the DataCategory
			dataCategory = self.get( dataCategoryId ).first()
				
			# split on a comma with no spaces
			idList = recordsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Record_		
				record_ = Record_Delegate().get(id).first();	
				# add the Record_
				dataCategory.records.add(record_)
				
			# save it		
			dataCategory.save()
			
			# reload and return the appropriate version
			return self.get( dataCategoryId );
		except DataCategory.DoesNotExist:
			raise ProcessingError(errMsg + " : DataCategory with id " + str(dataCategoryId) + " does not exist.")
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRecords( self, dataCategoryId, recordsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Record_Delegate import Record_Delegate

		errMsg = "Failed to remove elements " + str(recordsIds) + " for Records on DataCategory"

		try:
			# get the DataCategory
			dataCategory = self.get( dataCategoryId ).first()
				
			# split on a comma with no spaces
			idList = recordsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Record_		
				record_ = Record_Delegate().get(id).first();	
				# add the Record_
				dataCategory.records.remove(record_)
				
			# save it		
			dataCategory.save()
			
			# reload and return the appropriate version
			return self.get( dataCategoryId );
		except DataCategory.DoesNotExist:
			raise ProcessingError(errMsg + " : DataCategory with id " + str(dataCategoryId) + " does not exist.")
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDataBreaches( self, dataCategoryId, dataBreachesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataBreachDelegate import DataBreachDelegate

		errMsg = "Failed to add elements " + str(dataBreachesIds) + " for DataBreaches on DataCategory"

		try:
			# get the DataCategory
			dataCategory = self.get( dataCategoryId ).first()
				
			# split on a comma with no spaces
			idList = dataBreachesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataBreach		
				dataBreach = DataBreachDelegate().get(id).first();	
				# add the DataBreach
				dataCategory.dataBreaches.add(dataBreach)
				
			# save it		
			dataCategory.save()
			
			# reload and return the appropriate version
			return self.get( dataCategoryId );
		except DataCategory.DoesNotExist:
			raise ProcessingError(errMsg + " : DataCategory with id " + str(dataCategoryId) + " does not exist.")
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDataBreaches( self, dataCategoryId, dataBreachesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataBreachDelegate import DataBreachDelegate

		errMsg = "Failed to remove elements " + str(dataBreachesIds) + " for DataBreaches on DataCategory"

		try:
			# get the DataCategory
			dataCategory = self.get( dataCategoryId ).first()
				
			# split on a comma with no spaces
			idList = dataBreachesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataBreach		
				dataBreach = DataBreachDelegate().get(id).first();	
				# add the DataBreach
				dataCategory.dataBreaches.remove(dataBreach)
				
			# save it		
			dataCategory.save()
			
			# reload and return the appropriate version
			return self.get( dataCategoryId );
		except DataCategory.DoesNotExist:
			raise ProcessingError(errMsg + " : DataCategory with id " + str(dataCategoryId) + " does not exist.")
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
