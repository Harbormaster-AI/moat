from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.RecordsRepository import RecordsRepository
from governanceOnDjango.models.Organization import Organization
from governanceOnDjango.models.Record_ import Record_
from governanceOnDjango.models.System_ import System_
from governanceOnDjango.models.RetentionSchedule import RetentionSchedule
from governanceOnDjango.models.LegalHold import LegalHold
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model RecordsRepository
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RecordsRepositoryDelegate Declaration
#======================================================================
class RecordsRepositoryDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, recordsRepositoryId ):
		try:	
			recordsRepository = RecordsRepository.objects.filter(id=recordsRepositoryId)
			return recordsRepository.first();
		except RecordsRepository.DoesNotExist:
			raise ProcessingError("RecordsRepository with id " + str(recordsRepositoryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, recordsRepository):
		for model in serializers.deserialize("json", recordsRepository):
			model.save()
			return model;

	def create(self, recordsRepository):
		recordsRepository.save()
		return recordsRepository;

	def saveFromJson(self, recordsRepository):
		for model in serializers.deserialize("json", recordsRepository):
			model.save()
			return recordsRepository;
	
	def save(self, recordsRepository):
		recordsRepository.save()
		return recordsRepository;
	
	def delete(self, recordsRepositoryId ):
		errMsg = "Failed to delete RecordsRepository from db using id " + str(recordsRepositoryId)
		
		try:
			recordsRepository = RecordsRepository.objects.get(id=recordsRepositoryId)
			recordsRepository.delete()
			return True
		except RecordsRepository.DoesNotExist:
			raise ProcessingError("RecordsRepository with id " + str(recordsRepositoryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = RecordsRepository.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all RecordsRepository from db")
		except Exception:
			return None;
		
	def assignOrganization( self, recordsRepositoryId, organizationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on RecordsRepository"

		try:
			# get the RecordsRepository from db
			recordsRepository = self.get( recordsRepositoryId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			recordsRepository.organization = organization
			
			#save it
			recordsRepository.save()

			# reload and return the appropriate version					
			return self.get( recordsRepositoryId );
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository with id " + str(recordsRepositoryId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, recordsRepositoryId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on RecordsRepository"

		try:
			# get the RecordsRepository from db
			recordsRepository = self.get( recordsRepositoryId ).first()	
			
			# assign to None for unassignment
			recordsRepository.organization = None			

			#save it
			recordsRepository.save()

			# reload and return the appropriate version					
			return self.get( recordsRepositoryId );
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository with id " + str(recordsRepositoryId) + " does not exist.")
		except Exception:
			return None;
		
	def addRecords( self, recordsRepositoryId, recordsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Record_Delegate import Record_Delegate

		errMsg = "Failed to add elements " + str(recordsIds) + " for Records on RecordsRepository"

		try:
			# get the RecordsRepository
			recordsRepository = self.get( recordsRepositoryId ).first()
				
			# split on a comma with no spaces
			idList = recordsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Record_		
				record_ = Record_Delegate().get(id).first();	
				# add the Record_
				recordsRepository.records.add(record_)
				
			# save it		
			recordsRepository.save()
			
			# reload and return the appropriate version
			return self.get( recordsRepositoryId );
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository with id " + str(recordsRepositoryId) + " does not exist.")
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRecords( self, recordsRepositoryId, recordsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Record_Delegate import Record_Delegate

		errMsg = "Failed to remove elements " + str(recordsIds) + " for Records on RecordsRepository"

		try:
			# get the RecordsRepository
			recordsRepository = self.get( recordsRepositoryId ).first()
				
			# split on a comma with no spaces
			idList = recordsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Record_		
				record_ = Record_Delegate().get(id).first();	
				# add the Record_
				recordsRepository.records.remove(record_)
				
			# save it		
			recordsRepository.save()
			
			# reload and return the appropriate version
			return self.get( recordsRepositoryId );
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository with id " + str(recordsRepositoryId) + " does not exist.")
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSystems( self, recordsRepositoryId, systemsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.System_Delegate import System_Delegate

		errMsg = "Failed to add elements " + str(systemsIds) + " for Systems on RecordsRepository"

		try:
			# get the RecordsRepository
			recordsRepository = self.get( recordsRepositoryId ).first()
				
			# split on a comma with no spaces
			idList = systemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the System_		
				system_ = System_Delegate().get(id).first();	
				# add the System_
				recordsRepository.systems.add(system_)
				
			# save it		
			recordsRepository.save()
			
			# reload and return the appropriate version
			return self.get( recordsRepositoryId );
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository with id " + str(recordsRepositoryId) + " does not exist.")
		except System_.DoesNotExist:
			raise ProcessingError(errMsg + " : System_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSystems( self, recordsRepositoryId, systemsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.System_Delegate import System_Delegate

		errMsg = "Failed to remove elements " + str(systemsIds) + " for Systems on RecordsRepository"

		try:
			# get the RecordsRepository
			recordsRepository = self.get( recordsRepositoryId ).first()
				
			# split on a comma with no spaces
			idList = systemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the System_		
				system_ = System_Delegate().get(id).first();	
				# add the System_
				recordsRepository.systems.remove(system_)
				
			# save it		
			recordsRepository.save()
			
			# reload and return the appropriate version
			return self.get( recordsRepositoryId );
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository with id " + str(recordsRepositoryId) + " does not exist.")
		except System_.DoesNotExist:
			raise ProcessingError(errMsg + " : System_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRetentionSchedules( self, recordsRepositoryId, retentionSchedulesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RetentionScheduleDelegate import RetentionScheduleDelegate

		errMsg = "Failed to add elements " + str(retentionSchedulesIds) + " for RetentionSchedules on RecordsRepository"

		try:
			# get the RecordsRepository
			recordsRepository = self.get( recordsRepositoryId ).first()
				
			# split on a comma with no spaces
			idList = retentionSchedulesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the RetentionSchedule		
				retentionSchedule = RetentionScheduleDelegate().get(id).first();	
				# add the RetentionSchedule
				recordsRepository.retentionSchedules.add(retentionSchedule)
				
			# save it		
			recordsRepository.save()
			
			# reload and return the appropriate version
			return self.get( recordsRepositoryId );
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository with id " + str(recordsRepositoryId) + " does not exist.")
		except RetentionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RetentionSchedule does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRetentionSchedules( self, recordsRepositoryId, retentionSchedulesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RetentionScheduleDelegate import RetentionScheduleDelegate

		errMsg = "Failed to remove elements " + str(retentionSchedulesIds) + " for RetentionSchedules on RecordsRepository"

		try:
			# get the RecordsRepository
			recordsRepository = self.get( recordsRepositoryId ).first()
				
			# split on a comma with no spaces
			idList = retentionSchedulesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the RetentionSchedule		
				retentionSchedule = RetentionScheduleDelegate().get(id).first();	
				# add the RetentionSchedule
				recordsRepository.retentionSchedules.remove(retentionSchedule)
				
			# save it		
			recordsRepository.save()
			
			# reload and return the appropriate version
			return self.get( recordsRepositoryId );
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository with id " + str(recordsRepositoryId) + " does not exist.")
		except RetentionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RetentionSchedule does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLegalHolds( self, recordsRepositoryId, legalHoldsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.LegalHoldDelegate import LegalHoldDelegate

		errMsg = "Failed to add elements " + str(legalHoldsIds) + " for LegalHolds on RecordsRepository"

		try:
			# get the RecordsRepository
			recordsRepository = self.get( recordsRepositoryId ).first()
				
			# split on a comma with no spaces
			idList = legalHoldsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LegalHold		
				legalHold = LegalHoldDelegate().get(id).first();	
				# add the LegalHold
				recordsRepository.legalHolds.add(legalHold)
				
			# save it		
			recordsRepository.save()
			
			# reload and return the appropriate version
			return self.get( recordsRepositoryId );
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository with id " + str(recordsRepositoryId) + " does not exist.")
		except LegalHold.DoesNotExist:
			raise ProcessingError(errMsg + " : LegalHold does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLegalHolds( self, recordsRepositoryId, legalHoldsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.LegalHoldDelegate import LegalHoldDelegate

		errMsg = "Failed to remove elements " + str(legalHoldsIds) + " for LegalHolds on RecordsRepository"

		try:
			# get the RecordsRepository
			recordsRepository = self.get( recordsRepositoryId ).first()
				
			# split on a comma with no spaces
			idList = legalHoldsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LegalHold		
				legalHold = LegalHoldDelegate().get(id).first();	
				# add the LegalHold
				recordsRepository.legalHolds.remove(legalHold)
				
			# save it		
			recordsRepository.save()
			
			# reload and return the appropriate version
			return self.get( recordsRepositoryId );
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository with id " + str(recordsRepositoryId) + " does not exist.")
		except LegalHold.DoesNotExist:
			raise ProcessingError(errMsg + " : LegalHold does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
