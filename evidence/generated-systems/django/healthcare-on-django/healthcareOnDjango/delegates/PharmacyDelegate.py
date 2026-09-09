from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Pharmacy import Pharmacy
from healthcareOnDjango.models.Facility import Facility
from healthcareOnDjango.models.MedicationDispense import MedicationDispense
from healthcareOnDjango.models.MedicationOrder import MedicationOrder
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Pharmacy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PharmacyDelegate Declaration
#======================================================================
class PharmacyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, pharmacyId ):
		try:	
			pharmacy = Pharmacy.objects.filter(id=pharmacyId)
			return pharmacy.first();
		except Pharmacy.DoesNotExist:
			raise ProcessingError("Pharmacy with id " + str(pharmacyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, pharmacy):
		for model in serializers.deserialize("json", pharmacy):
			model.save()
			return model;

	def create(self, pharmacy):
		pharmacy.save()
		return pharmacy;

	def saveFromJson(self, pharmacy):
		for model in serializers.deserialize("json", pharmacy):
			model.save()
			return pharmacy;
	
	def save(self, pharmacy):
		pharmacy.save()
		return pharmacy;
	
	def delete(self, pharmacyId ):
		errMsg = "Failed to delete Pharmacy from db using id " + str(pharmacyId)
		
		try:
			pharmacy = Pharmacy.objects.get(id=pharmacyId)
			pharmacy.delete()
			return True
		except Pharmacy.DoesNotExist:
			raise ProcessingError("Pharmacy with id " + str(pharmacyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Pharmacy.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Pharmacy from db")
		except Exception:
			return None;
		
	def assignFacility( self, pharmacyId, facilityId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.FacilityDelegate import FacilityDelegate

		errMsg = "Failed to assign element " + str(facilityId) + " for Facility on Pharmacy"

		try:
			# get the Pharmacy from db
			pharmacy = self.get( pharmacyId ).first()	
			
			# get the Facility from db
			facility = FacilityDelegate().get(facilityId).first();
			
			# assign the Facility		
			pharmacy.facility = facility
			
			#save it
			pharmacy.save()

			# reload and return the appropriate version					
			return self.get( pharmacyId );
		except Pharmacy.DoesNotExist:
			raise ProcessingError(errMsg + " : Pharmacy with id " + str(pharmacyId) + " does not exist.")
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFacility( self, pharmacyId ):
		errMsg = "Failed to unassign element " + str(facilityId) + " for Facility on Pharmacy"

		try:
			# get the Pharmacy from db
			pharmacy = self.get( pharmacyId ).first()	
			
			# assign to None for unassignment
			pharmacy.facility = None			

			#save it
			pharmacy.save()

			# reload and return the appropriate version					
			return self.get( pharmacyId );
		except Pharmacy.DoesNotExist:
			raise ProcessingError(errMsg + " : Pharmacy with id " + str(pharmacyId) + " does not exist.")
		except Exception:
			return None;
		
	def addMedicationDispenses( self, pharmacyId, medicationDispensesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicationDispenseDelegate import MedicationDispenseDelegate

		errMsg = "Failed to add elements " + str(medicationDispensesIds) + " for MedicationDispenses on Pharmacy"

		try:
			# get the Pharmacy
			pharmacy = self.get( pharmacyId ).first()
				
			# split on a comma with no spaces
			idList = medicationDispensesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MedicationDispense		
				medicationDispense = MedicationDispenseDelegate().get(id).first();	
				# add the MedicationDispense
				pharmacy.medicationDispenses.add(medicationDispense)
				
			# save it		
			pharmacy.save()
			
			# reload and return the appropriate version
			return self.get( pharmacyId );
		except Pharmacy.DoesNotExist:
			raise ProcessingError(errMsg + " : Pharmacy with id " + str(pharmacyId) + " does not exist.")
		except MedicationDispense.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationDispense does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMedicationDispenses( self, pharmacyId, medicationDispensesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicationDispenseDelegate import MedicationDispenseDelegate

		errMsg = "Failed to remove elements " + str(medicationDispensesIds) + " for MedicationDispenses on Pharmacy"

		try:
			# get the Pharmacy
			pharmacy = self.get( pharmacyId ).first()
				
			# split on a comma with no spaces
			idList = medicationDispensesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MedicationDispense		
				medicationDispense = MedicationDispenseDelegate().get(id).first();	
				# add the MedicationDispense
				pharmacy.medicationDispenses.remove(medicationDispense)
				
			# save it		
			pharmacy.save()
			
			# reload and return the appropriate version
			return self.get( pharmacyId );
		except Pharmacy.DoesNotExist:
			raise ProcessingError(errMsg + " : Pharmacy with id " + str(pharmacyId) + " does not exist.")
		except MedicationDispense.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationDispense does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addMedicationOrders( self, pharmacyId, medicationOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicationOrderDelegate import MedicationOrderDelegate

		errMsg = "Failed to add elements " + str(medicationOrdersIds) + " for MedicationOrders on Pharmacy"

		try:
			# get the Pharmacy
			pharmacy = self.get( pharmacyId ).first()
				
			# split on a comma with no spaces
			idList = medicationOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MedicationOrder		
				medicationOrder = MedicationOrderDelegate().get(id).first();	
				# add the MedicationOrder
				pharmacy.medicationOrders.add(medicationOrder)
				
			# save it		
			pharmacy.save()
			
			# reload and return the appropriate version
			return self.get( pharmacyId );
		except Pharmacy.DoesNotExist:
			raise ProcessingError(errMsg + " : Pharmacy with id " + str(pharmacyId) + " does not exist.")
		except MedicationOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMedicationOrders( self, pharmacyId, medicationOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicationOrderDelegate import MedicationOrderDelegate

		errMsg = "Failed to remove elements " + str(medicationOrdersIds) + " for MedicationOrders on Pharmacy"

		try:
			# get the Pharmacy
			pharmacy = self.get( pharmacyId ).first()
				
			# split on a comma with no spaces
			idList = medicationOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MedicationOrder		
				medicationOrder = MedicationOrderDelegate().get(id).first();	
				# add the MedicationOrder
				pharmacy.medicationOrders.remove(medicationOrder)
				
			# save it		
			pharmacy.save()
			
			# reload and return the appropriate version
			return self.get( pharmacyId );
		except Pharmacy.DoesNotExist:
			raise ProcessingError(errMsg + " : Pharmacy with id " + str(pharmacyId) + " does not exist.")
		except MedicationOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
