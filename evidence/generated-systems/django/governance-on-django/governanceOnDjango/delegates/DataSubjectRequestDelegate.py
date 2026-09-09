from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.DataSubjectRequest import DataSubjectRequest
from governanceOnDjango.models.Organization import Organization
from governanceOnDjango.models.DataProcessingActivity import DataProcessingActivity
from governanceOnDjango.models.Record_ import Record_
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model DataSubjectRequest
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataSubjectRequestDelegate Declaration
#======================================================================
class DataSubjectRequestDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dataSubjectRequestId ):
		try:	
			dataSubjectRequest = DataSubjectRequest.objects.filter(id=dataSubjectRequestId)
			return dataSubjectRequest.first();
		except DataSubjectRequest.DoesNotExist:
			raise ProcessingError("DataSubjectRequest with id " + str(dataSubjectRequestId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, dataSubjectRequest):
		for model in serializers.deserialize("json", dataSubjectRequest):
			model.save()
			return model;

	def create(self, dataSubjectRequest):
		dataSubjectRequest.save()
		return dataSubjectRequest;

	def saveFromJson(self, dataSubjectRequest):
		for model in serializers.deserialize("json", dataSubjectRequest):
			model.save()
			return dataSubjectRequest;
	
	def save(self, dataSubjectRequest):
		dataSubjectRequest.save()
		return dataSubjectRequest;
	
	def delete(self, dataSubjectRequestId ):
		errMsg = "Failed to delete DataSubjectRequest from db using id " + str(dataSubjectRequestId)
		
		try:
			dataSubjectRequest = DataSubjectRequest.objects.get(id=dataSubjectRequestId)
			dataSubjectRequest.delete()
			return True
		except DataSubjectRequest.DoesNotExist:
			raise ProcessingError("DataSubjectRequest with id " + str(dataSubjectRequestId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = DataSubjectRequest.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all DataSubjectRequest from db")
		except Exception:
			return None;
		
	def assignOrganization( self, dataSubjectRequestId, organizationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on DataSubjectRequest"

		try:
			# get the DataSubjectRequest from db
			dataSubjectRequest = self.get( dataSubjectRequestId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			dataSubjectRequest.organization = organization
			
			#save it
			dataSubjectRequest.save()

			# reload and return the appropriate version					
			return self.get( dataSubjectRequestId );
		except DataSubjectRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSubjectRequest with id " + str(dataSubjectRequestId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, dataSubjectRequestId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on DataSubjectRequest"

		try:
			# get the DataSubjectRequest from db
			dataSubjectRequest = self.get( dataSubjectRequestId ).first()	
			
			# assign to None for unassignment
			dataSubjectRequest.organization = None			

			#save it
			dataSubjectRequest.save()

			# reload and return the appropriate version					
			return self.get( dataSubjectRequestId );
		except DataSubjectRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSubjectRequest with id " + str(dataSubjectRequestId) + " does not exist.")
		except Exception:
			return None;
		
	def addProcessingActivities( self, dataSubjectRequestId, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to add elements " + str(processingActivitiesIds) + " for ProcessingActivities on DataSubjectRequest"

		try:
			# get the DataSubjectRequest
			dataSubjectRequest = self.get( dataSubjectRequestId ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				dataSubjectRequest.processingActivities.add(dataProcessingActivity)
				
			# save it		
			dataSubjectRequest.save()
			
			# reload and return the appropriate version
			return self.get( dataSubjectRequestId );
		except DataSubjectRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSubjectRequest with id " + str(dataSubjectRequestId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProcessingActivities( self, dataSubjectRequestId, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to remove elements " + str(processingActivitiesIds) + " for ProcessingActivities on DataSubjectRequest"

		try:
			# get the DataSubjectRequest
			dataSubjectRequest = self.get( dataSubjectRequestId ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				dataSubjectRequest.processingActivities.remove(dataProcessingActivity)
				
			# save it		
			dataSubjectRequest.save()
			
			# reload and return the appropriate version
			return self.get( dataSubjectRequestId );
		except DataSubjectRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSubjectRequest with id " + str(dataSubjectRequestId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRecords( self, dataSubjectRequestId, recordsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Record_Delegate import Record_Delegate

		errMsg = "Failed to add elements " + str(recordsIds) + " for Records on DataSubjectRequest"

		try:
			# get the DataSubjectRequest
			dataSubjectRequest = self.get( dataSubjectRequestId ).first()
				
			# split on a comma with no spaces
			idList = recordsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Record_		
				record_ = Record_Delegate().get(id).first();	
				# add the Record_
				dataSubjectRequest.records.add(record_)
				
			# save it		
			dataSubjectRequest.save()
			
			# reload and return the appropriate version
			return self.get( dataSubjectRequestId );
		except DataSubjectRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSubjectRequest with id " + str(dataSubjectRequestId) + " does not exist.")
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRecords( self, dataSubjectRequestId, recordsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Record_Delegate import Record_Delegate

		errMsg = "Failed to remove elements " + str(recordsIds) + " for Records on DataSubjectRequest"

		try:
			# get the DataSubjectRequest
			dataSubjectRequest = self.get( dataSubjectRequestId ).first()
				
			# split on a comma with no spaces
			idList = recordsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Record_		
				record_ = Record_Delegate().get(id).first();	
				# add the Record_
				dataSubjectRequest.records.remove(record_)
				
			# save it		
			dataSubjectRequest.save()
			
			# reload and return the appropriate version
			return self.get( dataSubjectRequestId );
		except DataSubjectRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSubjectRequest with id " + str(dataSubjectRequestId) + " does not exist.")
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
