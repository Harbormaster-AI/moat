from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.MedicationDispense import MedicationDispense
from healthcareOnDjango.models.MedicationOrder import MedicationOrder
from healthcareOnDjango.models.Pharmacy import Pharmacy
from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model MedicationDispense
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicationDispenseDelegate Declaration
#======================================================================
class MedicationDispenseDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, medicationDispenseId ):
		try:	
			medicationDispense = MedicationDispense.objects.filter(id=medicationDispenseId)
			return medicationDispense.first();
		except MedicationDispense.DoesNotExist:
			raise ProcessingError("MedicationDispense with id " + str(medicationDispenseId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, medicationDispense):
		for model in serializers.deserialize("json", medicationDispense):
			model.save()
			return model;

	def create(self, medicationDispense):
		medicationDispense.save()
		return medicationDispense;

	def saveFromJson(self, medicationDispense):
		for model in serializers.deserialize("json", medicationDispense):
			model.save()
			return medicationDispense;
	
	def save(self, medicationDispense):
		medicationDispense.save()
		return medicationDispense;
	
	def delete(self, medicationDispenseId ):
		errMsg = "Failed to delete MedicationDispense from db using id " + str(medicationDispenseId)
		
		try:
			medicationDispense = MedicationDispense.objects.get(id=medicationDispenseId)
			medicationDispense.delete()
			return True
		except MedicationDispense.DoesNotExist:
			raise ProcessingError("MedicationDispense with id " + str(medicationDispenseId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = MedicationDispense.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all MedicationDispense from db")
		except Exception:
			return None;
		
	def assignMedicationOrder( self, medicationDispenseId, medicationOrderId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicationOrderDelegate import MedicationOrderDelegate

		errMsg = "Failed to assign element " + str(medicationOrderId) + " for MedicationOrder on MedicationDispense"

		try:
			# get the MedicationDispense from db
			medicationDispense = self.get( medicationDispenseId ).first()	
			
			# get the MedicationOrder from db
			medicationOrder = MedicationOrderDelegate().get(medicationOrderId).first();
			
			# assign the MedicationOrder		
			medicationDispense.medicationOrder = medicationOrder
			
			#save it
			medicationDispense.save()

			# reload and return the appropriate version					
			return self.get( medicationDispenseId );
		except MedicationDispense.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationDispense with id " + str(medicationDispenseId) + " does not exist.")
		except MedicationOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationOrder with id " + str(medicationOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMedicationOrder( self, medicationDispenseId ):
		errMsg = "Failed to unassign element " + str(medicationOrderId) + " for MedicationOrder on MedicationDispense"

		try:
			# get the MedicationDispense from db
			medicationDispense = self.get( medicationDispenseId ).first()	
			
			# assign to None for unassignment
			medicationDispense.medicationOrder = None			

			#save it
			medicationDispense.save()

			# reload and return the appropriate version					
			return self.get( medicationDispenseId );
		except MedicationDispense.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationDispense with id " + str(medicationDispenseId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPharmacy( self, medicationDispenseId, pharmacyId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PharmacyDelegate import PharmacyDelegate

		errMsg = "Failed to assign element " + str(pharmacyId) + " for Pharmacy on MedicationDispense"

		try:
			# get the MedicationDispense from db
			medicationDispense = self.get( medicationDispenseId ).first()	
			
			# get the Pharmacy from db
			pharmacy = PharmacyDelegate().get(pharmacyId).first();
			
			# assign the Pharmacy		
			medicationDispense.pharmacy = pharmacy
			
			#save it
			medicationDispense.save()

			# reload and return the appropriate version					
			return self.get( medicationDispenseId );
		except MedicationDispense.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationDispense with id " + str(medicationDispenseId) + " does not exist.")
		except Pharmacy.DoesNotExist:
			raise ProcessingError(errMsg + " : Pharmacy with id " + str(pharmacyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPharmacy( self, medicationDispenseId ):
		errMsg = "Failed to unassign element " + str(pharmacyId) + " for Pharmacy on MedicationDispense"

		try:
			# get the MedicationDispense from db
			medicationDispense = self.get( medicationDispenseId ).first()	
			
			# assign to None for unassignment
			medicationDispense.pharmacy = None			

			#save it
			medicationDispense.save()

			# reload and return the appropriate version					
			return self.get( medicationDispenseId );
		except MedicationDispense.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationDispense with id " + str(medicationDispenseId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPatient( self, medicationDispenseId, patientId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

		errMsg = "Failed to assign element " + str(patientId) + " for Patient on MedicationDispense"

		try:
			# get the MedicationDispense from db
			medicationDispense = self.get( medicationDispenseId ).first()	
			
			# get the Patient from db
			patient = PatientDelegate().get(patientId).first();
			
			# assign the Patient		
			medicationDispense.patient = patient
			
			#save it
			medicationDispense.save()

			# reload and return the appropriate version					
			return self.get( medicationDispenseId );
		except MedicationDispense.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationDispense with id " + str(medicationDispenseId) + " does not exist.")
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPatient( self, medicationDispenseId ):
		errMsg = "Failed to unassign element " + str(patientId) + " for Patient on MedicationDispense"

		try:
			# get the MedicationDispense from db
			medicationDispense = self.get( medicationDispenseId ).first()	
			
			# assign to None for unassignment
			medicationDispense.patient = None			

			#save it
			medicationDispense.save()

			# reload and return the appropriate version					
			return self.get( medicationDispenseId );
		except MedicationDispense.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationDispense with id " + str(medicationDispenseId) + " does not exist.")
		except Exception:
			return None;
		
