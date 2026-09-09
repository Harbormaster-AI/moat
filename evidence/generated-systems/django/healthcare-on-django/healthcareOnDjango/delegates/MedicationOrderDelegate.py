from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.MedicationOrder import MedicationOrder
from healthcareOnDjango.models.ClinicalOrder import ClinicalOrder
from healthcareOnDjango.models.Pharmacy import Pharmacy
from healthcareOnDjango.models.MedicationDispense import MedicationDispense
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model MedicationOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicationOrderDelegate Declaration
#======================================================================
class MedicationOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, medicationOrderId ):
		try:	
			medicationOrder = MedicationOrder.objects.filter(id=medicationOrderId)
			return medicationOrder.first();
		except MedicationOrder.DoesNotExist:
			raise ProcessingError("MedicationOrder with id " + str(medicationOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, medicationOrder):
		for model in serializers.deserialize("json", medicationOrder):
			model.save()
			return model;

	def create(self, medicationOrder):
		medicationOrder.save()
		return medicationOrder;

	def saveFromJson(self, medicationOrder):
		for model in serializers.deserialize("json", medicationOrder):
			model.save()
			return medicationOrder;
	
	def save(self, medicationOrder):
		medicationOrder.save()
		return medicationOrder;
	
	def delete(self, medicationOrderId ):
		errMsg = "Failed to delete MedicationOrder from db using id " + str(medicationOrderId)
		
		try:
			medicationOrder = MedicationOrder.objects.get(id=medicationOrderId)
			medicationOrder.delete()
			return True
		except MedicationOrder.DoesNotExist:
			raise ProcessingError("MedicationOrder with id " + str(medicationOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = MedicationOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all MedicationOrder from db")
		except Exception:
			return None;
		
	def assignOrder( self, medicationOrderId, orderId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClinicalOrderDelegate import ClinicalOrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on MedicationOrder"

		try:
			# get the MedicationOrder from db
			medicationOrder = self.get( medicationOrderId ).first()	
			
			# get the ClinicalOrder from db
			clinicalOrder = ClinicalOrderDelegate().get(orderId).first();
			
			# assign the Order		
			medicationOrder.order = clinicalOrder
			
			#save it
			medicationOrder.save()

			# reload and return the appropriate version					
			return self.get( medicationOrderId );
		except MedicationOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationOrder with id " + str(medicationOrderId) + " does not exist.")
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, medicationOrderId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on MedicationOrder"

		try:
			# get the MedicationOrder from db
			medicationOrder = self.get( medicationOrderId ).first()	
			
			# assign to None for unassignment
			medicationOrder.clinicalOrder = None			

			#save it
			medicationOrder.save()

			# reload and return the appropriate version					
			return self.get( medicationOrderId );
		except MedicationOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationOrder with id " + str(medicationOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPharmacy( self, medicationOrderId, pharmacyId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PharmacyDelegate import PharmacyDelegate

		errMsg = "Failed to assign element " + str(pharmacyId) + " for Pharmacy on MedicationOrder"

		try:
			# get the MedicationOrder from db
			medicationOrder = self.get( medicationOrderId ).first()	
			
			# get the Pharmacy from db
			pharmacy = PharmacyDelegate().get(pharmacyId).first();
			
			# assign the Pharmacy		
			medicationOrder.pharmacy = pharmacy
			
			#save it
			medicationOrder.save()

			# reload and return the appropriate version					
			return self.get( medicationOrderId );
		except MedicationOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationOrder with id " + str(medicationOrderId) + " does not exist.")
		except Pharmacy.DoesNotExist:
			raise ProcessingError(errMsg + " : Pharmacy with id " + str(pharmacyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPharmacy( self, medicationOrderId ):
		errMsg = "Failed to unassign element " + str(pharmacyId) + " for Pharmacy on MedicationOrder"

		try:
			# get the MedicationOrder from db
			medicationOrder = self.get( medicationOrderId ).first()	
			
			# assign to None for unassignment
			medicationOrder.pharmacy = None			

			#save it
			medicationOrder.save()

			# reload and return the appropriate version					
			return self.get( medicationOrderId );
		except MedicationOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationOrder with id " + str(medicationOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def addDispenses( self, medicationOrderId, dispensesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicationDispenseDelegate import MedicationDispenseDelegate

		errMsg = "Failed to add elements " + str(dispensesIds) + " for Dispenses on MedicationOrder"

		try:
			# get the MedicationOrder
			medicationOrder = self.get( medicationOrderId ).first()
				
			# split on a comma with no spaces
			idList = dispensesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MedicationDispense		
				medicationDispense = MedicationDispenseDelegate().get(id).first();	
				# add the MedicationDispense
				medicationOrder.dispenses.add(medicationDispense)
				
			# save it		
			medicationOrder.save()
			
			# reload and return the appropriate version
			return self.get( medicationOrderId );
		except MedicationOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationOrder with id " + str(medicationOrderId) + " does not exist.")
		except MedicationDispense.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationDispense does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDispenses( self, medicationOrderId, dispensesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicationDispenseDelegate import MedicationDispenseDelegate

		errMsg = "Failed to remove elements " + str(dispensesIds) + " for Dispenses on MedicationOrder"

		try:
			# get the MedicationOrder
			medicationOrder = self.get( medicationOrderId ).first()
				
			# split on a comma with no spaces
			idList = dispensesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MedicationDispense		
				medicationDispense = MedicationDispenseDelegate().get(id).first();	
				# add the MedicationDispense
				medicationOrder.dispenses.remove(medicationDispense)
				
			# save it		
			medicationOrder.save()
			
			# reload and return the appropriate version
			return self.get( medicationOrderId );
		except MedicationOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationOrder with id " + str(medicationOrderId) + " does not exist.")
		except MedicationDispense.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationDispense does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
