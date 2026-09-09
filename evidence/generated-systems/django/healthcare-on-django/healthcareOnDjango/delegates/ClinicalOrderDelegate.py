from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.ClinicalOrder import ClinicalOrder
from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.models.Encounter import Encounter
from healthcareOnDjango.models.Clinician import Clinician
from healthcareOnDjango.models.MedicationOrder import MedicationOrder
from healthcareOnDjango.models.LaboratoryOrder import LaboratoryOrder
from healthcareOnDjango.models.ImagingOrder import ImagingOrder
from healthcareOnDjango.models.ProcedureOrder import ProcedureOrder
from healthcareOnDjango.models.Authorization import Authorization
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ClinicalOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClinicalOrderDelegate Declaration
#======================================================================
class ClinicalOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, clinicalOrderId ):
		try:	
			clinicalOrder = ClinicalOrder.objects.filter(id=clinicalOrderId)
			return clinicalOrder.first();
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError("ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, clinicalOrder):
		for model in serializers.deserialize("json", clinicalOrder):
			model.save()
			return model;

	def create(self, clinicalOrder):
		clinicalOrder.save()
		return clinicalOrder;

	def saveFromJson(self, clinicalOrder):
		for model in serializers.deserialize("json", clinicalOrder):
			model.save()
			return clinicalOrder;
	
	def save(self, clinicalOrder):
		clinicalOrder.save()
		return clinicalOrder;
	
	def delete(self, clinicalOrderId ):
		errMsg = "Failed to delete ClinicalOrder from db using id " + str(clinicalOrderId)
		
		try:
			clinicalOrder = ClinicalOrder.objects.get(id=clinicalOrderId)
			clinicalOrder.delete()
			return True
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError("ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ClinicalOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ClinicalOrder from db")
		except Exception:
			return None;
		
	def assignPatient( self, clinicalOrderId, patientId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

		errMsg = "Failed to assign element " + str(patientId) + " for Patient on ClinicalOrder"

		try:
			# get the ClinicalOrder from db
			clinicalOrder = self.get( clinicalOrderId ).first()	
			
			# get the Patient from db
			patient = PatientDelegate().get(patientId).first();
			
			# assign the Patient		
			clinicalOrder.patient = patient
			
			#save it
			clinicalOrder.save()

			# reload and return the appropriate version					
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPatient( self, clinicalOrderId ):
		errMsg = "Failed to unassign element " + str(patientId) + " for Patient on ClinicalOrder"

		try:
			# get the ClinicalOrder from db
			clinicalOrder = self.get( clinicalOrderId ).first()	
			
			# assign to None for unassignment
			clinicalOrder.patient = None			

			#save it
			clinicalOrder.save()

			# reload and return the appropriate version					
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignEncounter( self, clinicalOrderId, encounterId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to assign element " + str(encounterId) + " for Encounter on ClinicalOrder"

		try:
			# get the ClinicalOrder from db
			clinicalOrder = self.get( clinicalOrderId ).first()	
			
			# get the Encounter from db
			encounter = EncounterDelegate().get(encounterId).first();
			
			# assign the Encounter		
			clinicalOrder.encounter = encounter
			
			#save it
			clinicalOrder.save()

			# reload and return the appropriate version					
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEncounter( self, clinicalOrderId ):
		errMsg = "Failed to unassign element " + str(encounterId) + " for Encounter on ClinicalOrder"

		try:
			# get the ClinicalOrder from db
			clinicalOrder = self.get( clinicalOrderId ).first()	
			
			# assign to None for unassignment
			clinicalOrder.encounter = None			

			#save it
			clinicalOrder.save()

			# reload and return the appropriate version					
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOrderingClinician( self, clinicalOrderId, orderingClinicianId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClinicianDelegate import ClinicianDelegate

		errMsg = "Failed to assign element " + str(orderingClinicianId) + " for OrderingClinician on ClinicalOrder"

		try:
			# get the ClinicalOrder from db
			clinicalOrder = self.get( clinicalOrderId ).first()	
			
			# get the Clinician from db
			clinician = ClinicianDelegate().get(orderingClinicianId).first();
			
			# assign the OrderingClinician		
			clinicalOrder.orderingClinician = clinician
			
			#save it
			clinicalOrder.save()

			# reload and return the appropriate version					
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(orderingClinicianId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrderingClinician( self, clinicalOrderId ):
		errMsg = "Failed to unassign element " + str(orderingClinicianId) + " for OrderingClinician on ClinicalOrder"

		try:
			# get the ClinicalOrder from db
			clinicalOrder = self.get( clinicalOrderId ).first()	
			
			# assign to None for unassignment
			clinicalOrder.clinician = None			

			#save it
			clinicalOrder.save()

			# reload and return the appropriate version					
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def addMedicationOrders( self, clinicalOrderId, medicationOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicationOrderDelegate import MedicationOrderDelegate

		errMsg = "Failed to add elements " + str(medicationOrdersIds) + " for MedicationOrders on ClinicalOrder"

		try:
			# get the ClinicalOrder
			clinicalOrder = self.get( clinicalOrderId ).first()
				
			# split on a comma with no spaces
			idList = medicationOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MedicationOrder		
				medicationOrder = MedicationOrderDelegate().get(id).first();	
				# add the MedicationOrder
				clinicalOrder.medicationOrders.add(medicationOrder)
				
			# save it		
			clinicalOrder.save()
			
			# reload and return the appropriate version
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except MedicationOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMedicationOrders( self, clinicalOrderId, medicationOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicationOrderDelegate import MedicationOrderDelegate

		errMsg = "Failed to remove elements " + str(medicationOrdersIds) + " for MedicationOrders on ClinicalOrder"

		try:
			# get the ClinicalOrder
			clinicalOrder = self.get( clinicalOrderId ).first()
				
			# split on a comma with no spaces
			idList = medicationOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MedicationOrder		
				medicationOrder = MedicationOrderDelegate().get(id).first();	
				# add the MedicationOrder
				clinicalOrder.medicationOrders.remove(medicationOrder)
				
			# save it		
			clinicalOrder.save()
			
			# reload and return the appropriate version
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except MedicationOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicationOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLaboratoryOrders( self, clinicalOrderId, laboratoryOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LaboratoryOrderDelegate import LaboratoryOrderDelegate

		errMsg = "Failed to add elements " + str(laboratoryOrdersIds) + " for LaboratoryOrders on ClinicalOrder"

		try:
			# get the ClinicalOrder
			clinicalOrder = self.get( clinicalOrderId ).first()
				
			# split on a comma with no spaces
			idList = laboratoryOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LaboratoryOrder		
				laboratoryOrder = LaboratoryOrderDelegate().get(id).first();	
				# add the LaboratoryOrder
				clinicalOrder.laboratoryOrders.add(laboratoryOrder)
				
			# save it		
			clinicalOrder.save()
			
			# reload and return the appropriate version
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except LaboratoryOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : LaboratoryOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLaboratoryOrders( self, clinicalOrderId, laboratoryOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LaboratoryOrderDelegate import LaboratoryOrderDelegate

		errMsg = "Failed to remove elements " + str(laboratoryOrdersIds) + " for LaboratoryOrders on ClinicalOrder"

		try:
			# get the ClinicalOrder
			clinicalOrder = self.get( clinicalOrderId ).first()
				
			# split on a comma with no spaces
			idList = laboratoryOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LaboratoryOrder		
				laboratoryOrder = LaboratoryOrderDelegate().get(id).first();	
				# add the LaboratoryOrder
				clinicalOrder.laboratoryOrders.remove(laboratoryOrder)
				
			# save it		
			clinicalOrder.save()
			
			# reload and return the appropriate version
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except LaboratoryOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : LaboratoryOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addImagingOrders( self, clinicalOrderId, imagingOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingOrderDelegate import ImagingOrderDelegate

		errMsg = "Failed to add elements " + str(imagingOrdersIds) + " for ImagingOrders on ClinicalOrder"

		try:
			# get the ClinicalOrder
			clinicalOrder = self.get( clinicalOrderId ).first()
				
			# split on a comma with no spaces
			idList = imagingOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ImagingOrder		
				imagingOrder = ImagingOrderDelegate().get(id).first();	
				# add the ImagingOrder
				clinicalOrder.imagingOrders.add(imagingOrder)
				
			# save it		
			clinicalOrder.save()
			
			# reload and return the appropriate version
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except ImagingOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeImagingOrders( self, clinicalOrderId, imagingOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingOrderDelegate import ImagingOrderDelegate

		errMsg = "Failed to remove elements " + str(imagingOrdersIds) + " for ImagingOrders on ClinicalOrder"

		try:
			# get the ClinicalOrder
			clinicalOrder = self.get( clinicalOrderId ).first()
				
			# split on a comma with no spaces
			idList = imagingOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ImagingOrder		
				imagingOrder = ImagingOrderDelegate().get(id).first();	
				# add the ImagingOrder
				clinicalOrder.imagingOrders.remove(imagingOrder)
				
			# save it		
			clinicalOrder.save()
			
			# reload and return the appropriate version
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except ImagingOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addProcedureOrders( self, clinicalOrderId, procedureOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ProcedureOrderDelegate import ProcedureOrderDelegate

		errMsg = "Failed to add elements " + str(procedureOrdersIds) + " for ProcedureOrders on ClinicalOrder"

		try:
			# get the ClinicalOrder
			clinicalOrder = self.get( clinicalOrderId ).first()
				
			# split on a comma with no spaces
			idList = procedureOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ProcedureOrder		
				procedureOrder = ProcedureOrderDelegate().get(id).first();	
				# add the ProcedureOrder
				clinicalOrder.procedureOrders.add(procedureOrder)
				
			# save it		
			clinicalOrder.save()
			
			# reload and return the appropriate version
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except ProcedureOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProcedureOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProcedureOrders( self, clinicalOrderId, procedureOrdersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ProcedureOrderDelegate import ProcedureOrderDelegate

		errMsg = "Failed to remove elements " + str(procedureOrdersIds) + " for ProcedureOrders on ClinicalOrder"

		try:
			# get the ClinicalOrder
			clinicalOrder = self.get( clinicalOrderId ).first()
				
			# split on a comma with no spaces
			idList = procedureOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ProcedureOrder		
				procedureOrder = ProcedureOrderDelegate().get(id).first();	
				# add the ProcedureOrder
				clinicalOrder.procedureOrders.remove(procedureOrder)
				
			# save it		
			clinicalOrder.save()
			
			# reload and return the appropriate version
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except ProcedureOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProcedureOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAuthorizations( self, clinicalOrderId, authorizationsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.AuthorizationDelegate import AuthorizationDelegate

		errMsg = "Failed to add elements " + str(authorizationsIds) + " for Authorizations on ClinicalOrder"

		try:
			# get the ClinicalOrder
			clinicalOrder = self.get( clinicalOrderId ).first()
				
			# split on a comma with no spaces
			idList = authorizationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Authorization		
				authorization = AuthorizationDelegate().get(id).first();	
				# add the Authorization
				clinicalOrder.authorizations.add(authorization)
				
			# save it		
			clinicalOrder.save()
			
			# reload and return the appropriate version
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except Authorization.DoesNotExist:
			raise ProcessingError(errMsg + " : Authorization does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAuthorizations( self, clinicalOrderId, authorizationsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.AuthorizationDelegate import AuthorizationDelegate

		errMsg = "Failed to remove elements " + str(authorizationsIds) + " for Authorizations on ClinicalOrder"

		try:
			# get the ClinicalOrder
			clinicalOrder = self.get( clinicalOrderId ).first()
				
			# split on a comma with no spaces
			idList = authorizationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Authorization		
				authorization = AuthorizationDelegate().get(id).first();	
				# add the Authorization
				clinicalOrder.authorizations.remove(authorization)
				
			# save it		
			clinicalOrder.save()
			
			# reload and return the appropriate version
			return self.get( clinicalOrderId );
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(clinicalOrderId) + " does not exist.")
		except Authorization.DoesNotExist:
			raise ProcessingError(errMsg + " : Authorization does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
