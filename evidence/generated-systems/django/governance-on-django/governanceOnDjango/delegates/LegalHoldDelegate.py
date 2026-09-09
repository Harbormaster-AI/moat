from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.LegalHold import LegalHold
from governanceOnDjango.models.RecordsRepository import RecordsRepository
from governanceOnDjango.models.Record_ import Record_
from governanceOnDjango.models.Matter import Matter
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model LegalHold
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LegalHoldDelegate Declaration
#======================================================================
class LegalHoldDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, legalHoldId ):
		try:	
			legalHold = LegalHold.objects.filter(id=legalHoldId)
			return legalHold.first();
		except LegalHold.DoesNotExist:
			raise ProcessingError("LegalHold with id " + str(legalHoldId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, legalHold):
		for model in serializers.deserialize("json", legalHold):
			model.save()
			return model;

	def create(self, legalHold):
		legalHold.save()
		return legalHold;

	def saveFromJson(self, legalHold):
		for model in serializers.deserialize("json", legalHold):
			model.save()
			return legalHold;
	
	def save(self, legalHold):
		legalHold.save()
		return legalHold;
	
	def delete(self, legalHoldId ):
		errMsg = "Failed to delete LegalHold from db using id " + str(legalHoldId)
		
		try:
			legalHold = LegalHold.objects.get(id=legalHoldId)
			legalHold.delete()
			return True
		except LegalHold.DoesNotExist:
			raise ProcessingError("LegalHold with id " + str(legalHoldId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = LegalHold.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all LegalHold from db")
		except Exception:
			return None;
		
	def assignMatter( self, legalHoldId, matterId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.MatterDelegate import MatterDelegate

		errMsg = "Failed to assign element " + str(matterId) + " for Matter on LegalHold"

		try:
			# get the LegalHold from db
			legalHold = self.get( legalHoldId ).first()	
			
			# get the Matter from db
			matter = MatterDelegate().get(matterId).first();
			
			# assign the Matter		
			legalHold.matter = matter
			
			#save it
			legalHold.save()

			# reload and return the appropriate version					
			return self.get( legalHoldId );
		except LegalHold.DoesNotExist:
			raise ProcessingError(errMsg + " : LegalHold with id " + str(legalHoldId) + " does not exist.")
		except Matter.DoesNotExist:
			raise ProcessingError(errMsg + " : Matter with id " + str(matterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMatter( self, legalHoldId ):
		errMsg = "Failed to unassign element " + str(matterId) + " for Matter on LegalHold"

		try:
			# get the LegalHold from db
			legalHold = self.get( legalHoldId ).first()	
			
			# assign to None for unassignment
			legalHold.matter = None			

			#save it
			legalHold.save()

			# reload and return the appropriate version					
			return self.get( legalHoldId );
		except LegalHold.DoesNotExist:
			raise ProcessingError(errMsg + " : LegalHold with id " + str(legalHoldId) + " does not exist.")
		except Exception:
			return None;
		
	def addRepositories( self, legalHoldId, repositoriesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RecordsRepositoryDelegate import RecordsRepositoryDelegate

		errMsg = "Failed to add elements " + str(repositoriesIds) + " for Repositories on LegalHold"

		try:
			# get the LegalHold
			legalHold = self.get( legalHoldId ).first()
				
			# split on a comma with no spaces
			idList = repositoriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the RecordsRepository		
				recordsRepository = RecordsRepositoryDelegate().get(id).first();	
				# add the RecordsRepository
				legalHold.repositories.add(recordsRepository)
				
			# save it		
			legalHold.save()
			
			# reload and return the appropriate version
			return self.get( legalHoldId );
		except LegalHold.DoesNotExist:
			raise ProcessingError(errMsg + " : LegalHold with id " + str(legalHoldId) + " does not exist.")
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRepositories( self, legalHoldId, repositoriesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RecordsRepositoryDelegate import RecordsRepositoryDelegate

		errMsg = "Failed to remove elements " + str(repositoriesIds) + " for Repositories on LegalHold"

		try:
			# get the LegalHold
			legalHold = self.get( legalHoldId ).first()
				
			# split on a comma with no spaces
			idList = repositoriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the RecordsRepository		
				recordsRepository = RecordsRepositoryDelegate().get(id).first();	
				# add the RecordsRepository
				legalHold.repositories.remove(recordsRepository)
				
			# save it		
			legalHold.save()
			
			# reload and return the appropriate version
			return self.get( legalHoldId );
		except LegalHold.DoesNotExist:
			raise ProcessingError(errMsg + " : LegalHold with id " + str(legalHoldId) + " does not exist.")
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRecords( self, legalHoldId, recordsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Record_Delegate import Record_Delegate

		errMsg = "Failed to add elements " + str(recordsIds) + " for Records on LegalHold"

		try:
			# get the LegalHold
			legalHold = self.get( legalHoldId ).first()
				
			# split on a comma with no spaces
			idList = recordsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Record_		
				record_ = Record_Delegate().get(id).first();	
				# add the Record_
				legalHold.records.add(record_)
				
			# save it		
			legalHold.save()
			
			# reload and return the appropriate version
			return self.get( legalHoldId );
		except LegalHold.DoesNotExist:
			raise ProcessingError(errMsg + " : LegalHold with id " + str(legalHoldId) + " does not exist.")
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRecords( self, legalHoldId, recordsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Record_Delegate import Record_Delegate

		errMsg = "Failed to remove elements " + str(recordsIds) + " for Records on LegalHold"

		try:
			# get the LegalHold
			legalHold = self.get( legalHoldId ).first()
				
			# split on a comma with no spaces
			idList = recordsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Record_		
				record_ = Record_Delegate().get(id).first();	
				# add the Record_
				legalHold.records.remove(record_)
				
			# save it		
			legalHold.save()
			
			# reload and return the appropriate version
			return self.get( legalHoldId );
		except LegalHold.DoesNotExist:
			raise ProcessingError(errMsg + " : LegalHold with id " + str(legalHoldId) + " does not exist.")
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
