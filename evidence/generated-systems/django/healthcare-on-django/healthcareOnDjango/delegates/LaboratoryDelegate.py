from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Laboratory import Laboratory
from healthcareOnDjango.models.Facility import Facility
from healthcareOnDjango.models.LaboratoryOrder import LaboratoryOrder
from healthcareOnDjango.models.LabResult import LabResult
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Laboratory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LaboratoryDelegate Declaration
#======================================================================
class LaboratoryDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, laboratoryId ):
		try:	
			laboratory = Laboratory.objects.filter(id=laboratoryId)
			return laboratory.first();
		except Laboratory.DoesNotExist:
			raise ProcessingError("Laboratory with id " + str(laboratoryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, laboratory):
		for model in serializers.deserialize("json", laboratory):
			model.save()
			return model;

	def create(self, laboratory):
		laboratory.save()
		return laboratory;

	def saveFromJson(self, laboratory):
		for model in serializers.deserialize("json", laboratory):
			model.save()
			return laboratory;
	
	def save(self, laboratory):
		laboratory.save()
		return laboratory;
	
	def delete(self, laboratoryId ):
		errMsg = "Failed to delete Laboratory from db using id " + str(laboratoryId)
		
		try:
			laboratory = Laboratory.objects.get(id=laboratoryId)
			laboratory.delete()
			return True
		except Laboratory.DoesNotExist:
			raise ProcessingError("Laboratory with id " + str(laboratoryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Laboratory.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Laboratory from db")
		except Exception:
			return None;
		
	def assignFacility( self, laboratoryId, facilityId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.FacilityDelegate import FacilityDelegate

		errMsg = "Failed to assign element " + str(facilityId) + " for Facility on Laboratory"

		try:
			# get the Laboratory from db
			laboratory = self.get( laboratoryId ).first()	
			
			# get the Facility from db
			facility = FacilityDelegate().get(facilityId).first();
			
			# assign the Facility		
			laboratory.facility = facility
			
			#save it
			laboratory.save()

			# reload and return the appropriate version					
			return self.get( laboratoryId );
		except Laboratory.DoesNotExist:
			raise ProcessingError(errMsg + " : Laboratory with id " + str(laboratoryId) + " does not exist.")
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFacility( self, laboratoryId ):
		errMsg = "Failed to unassign element " + str(facilityId) + " for Facility on Laboratory"

		try:
			# get the Laboratory from db
			laboratory = self.get( laboratoryId ).first()	
			
			# assign to None for unassignment
			laboratory.facility = None			

			#save it
			laboratory.save()

			# reload and return the appropriate version					
			return self.get( laboratoryId );
		except Laboratory.DoesNotExist:
			raise ProcessingError(errMsg + " : Laboratory with id " + str(laboratoryId) + " does not exist.")
		except Exception:
			return None;
		
	def addLaboratoryOrders( self, laboratoryId, laboratoryOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LaboratoryOrderDelegate import LaboratoryOrderDelegate

		errMsg = "Failed to add elements " + str(laboratoryOrdersIds) + " for LaboratoryOrders on Laboratory"

		try:
			# get the Laboratory
			laboratory = self.get( laboratoryId ).first()
				
			# split on a comma with no spaces
			idList = laboratoryOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LaboratoryOrder		
				laboratoryOrder = LaboratoryOrderDelegate().get(id).first();	
				# add the LaboratoryOrder
				laboratory.laboratoryOrders.add(laboratoryOrder)
				
			# save it		
			laboratory.save()
			
			# reload and return the appropriate version
			return self.get( laboratoryId );
		except Laboratory.DoesNotExist:
			raise ProcessingError(errMsg + " : Laboratory with id " + str(laboratoryId) + " does not exist.")
		except LaboratoryOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : LaboratoryOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLaboratoryOrders( self, laboratoryId, laboratoryOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LaboratoryOrderDelegate import LaboratoryOrderDelegate

		errMsg = "Failed to remove elements " + str(laboratoryOrdersIds) + " for LaboratoryOrders on Laboratory"

		try:
			# get the Laboratory
			laboratory = self.get( laboratoryId ).first()
				
			# split on a comma with no spaces
			idList = laboratoryOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LaboratoryOrder		
				laboratoryOrder = LaboratoryOrderDelegate().get(id).first();	
				# add the LaboratoryOrder
				laboratory.laboratoryOrders.remove(laboratoryOrder)
				
			# save it		
			laboratory.save()
			
			# reload and return the appropriate version
			return self.get( laboratoryId );
		except Laboratory.DoesNotExist:
			raise ProcessingError(errMsg + " : Laboratory with id " + str(laboratoryId) + " does not exist.")
		except LaboratoryOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : LaboratoryOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLabResults( self, laboratoryId, labResultsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LabResultDelegate import LabResultDelegate

		errMsg = "Failed to add elements " + str(labResultsIds) + " for LabResults on Laboratory"

		try:
			# get the Laboratory
			laboratory = self.get( laboratoryId ).first()
				
			# split on a comma with no spaces
			idList = labResultsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LabResult		
				labResult = LabResultDelegate().get(id).first();	
				# add the LabResult
				laboratory.labResults.add(labResult)
				
			# save it		
			laboratory.save()
			
			# reload and return the appropriate version
			return self.get( laboratoryId );
		except Laboratory.DoesNotExist:
			raise ProcessingError(errMsg + " : Laboratory with id " + str(laboratoryId) + " does not exist.")
		except LabResult.DoesNotExist:
			raise ProcessingError(errMsg + " : LabResult does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLabResults( self, laboratoryId, labResultsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LabResultDelegate import LabResultDelegate

		errMsg = "Failed to remove elements " + str(labResultsIds) + " for LabResults on Laboratory"

		try:
			# get the Laboratory
			laboratory = self.get( laboratoryId ).first()
				
			# split on a comma with no spaces
			idList = labResultsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LabResult		
				labResult = LabResultDelegate().get(id).first();	
				# add the LabResult
				laboratory.labResults.remove(labResult)
				
			# save it		
			laboratory.save()
			
			# reload and return the appropriate version
			return self.get( laboratoryId );
		except Laboratory.DoesNotExist:
			raise ProcessingError(errMsg + " : Laboratory with id " + str(laboratoryId) + " does not exist.")
		except LabResult.DoesNotExist:
			raise ProcessingError(errMsg + " : LabResult does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
