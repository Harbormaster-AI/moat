from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Encounter import Encounter
from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.models.Clinician import Clinician
from healthcareOnDjango.models.Facility import Facility
from healthcareOnDjango.models.Appointment import Appointment
from healthcareOnDjango.models.Diagnosis import Diagnosis
from healthcareOnDjango.models.Procedure import Procedure
from healthcareOnDjango.models.Observation import Observation
from healthcareOnDjango.models.ClinicalOrder import ClinicalOrder
from healthcareOnDjango.models.Admission import Admission
from healthcareOnDjango.models.Discharge import Discharge
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Encounter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EncounterDelegate Declaration
#======================================================================
class EncounterDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, encounterId ):
		try:	
			encounter = Encounter.objects.filter(id=encounterId)
			return encounter.first();
		except Encounter.DoesNotExist:
			raise ProcessingError("Encounter with id " + str(encounterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, encounter):
		for model in serializers.deserialize("json", encounter):
			model.save()
			return model;

	def create(self, encounter):
		encounter.save()
		return encounter;

	def saveFromJson(self, encounter):
		for model in serializers.deserialize("json", encounter):
			model.save()
			return encounter;
	
	def save(self, encounter):
		encounter.save()
		return encounter;
	
	def delete(self, encounterId ):
		errMsg = "Failed to delete Encounter from db using id " + str(encounterId)
		
		try:
			encounter = Encounter.objects.get(id=encounterId)
			encounter.delete()
			return True
		except Encounter.DoesNotExist:
			raise ProcessingError("Encounter with id " + str(encounterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Encounter.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Encounter from db")
		except Exception:
			return None;
		
	def assignPatient( self, encounterId, patientId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

		errMsg = "Failed to assign element " + str(patientId) + " for Patient on Encounter"

		try:
			# get the Encounter from db
			encounter = self.get( encounterId ).first()	
			
			# get the Patient from db
			patient = PatientDelegate().get(patientId).first();
			
			# assign the Patient		
			encounter.patient = patient
			
			#save it
			encounter.save()

			# reload and return the appropriate version					
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPatient( self, encounterId ):
		errMsg = "Failed to unassign element " + str(patientId) + " for Patient on Encounter"

		try:
			# get the Encounter from db
			encounter = self.get( encounterId ).first()	
			
			# assign to None for unassignment
			encounter.patient = None			

			#save it
			encounter.save()

			# reload and return the appropriate version					
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
		
	def assignClinician( self, encounterId, clinicianId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClinicianDelegate import ClinicianDelegate

		errMsg = "Failed to assign element " + str(clinicianId) + " for Clinician on Encounter"

		try:
			# get the Encounter from db
			encounter = self.get( encounterId ).first()	
			
			# get the Clinician from db
			clinician = ClinicianDelegate().get(clinicianId).first();
			
			# assign the Clinician		
			encounter.clinician = clinician
			
			#save it
			encounter.save()

			# reload and return the appropriate version					
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(clinicianId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignClinician( self, encounterId ):
		errMsg = "Failed to unassign element " + str(clinicianId) + " for Clinician on Encounter"

		try:
			# get the Encounter from db
			encounter = self.get( encounterId ).first()	
			
			# assign to None for unassignment
			encounter.clinician = None			

			#save it
			encounter.save()

			# reload and return the appropriate version					
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
		
	def assignFacility( self, encounterId, facilityId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.FacilityDelegate import FacilityDelegate

		errMsg = "Failed to assign element " + str(facilityId) + " for Facility on Encounter"

		try:
			# get the Encounter from db
			encounter = self.get( encounterId ).first()	
			
			# get the Facility from db
			facility = FacilityDelegate().get(facilityId).first();
			
			# assign the Facility		
			encounter.facility = facility
			
			#save it
			encounter.save()

			# reload and return the appropriate version					
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFacility( self, encounterId ):
		errMsg = "Failed to unassign element " + str(facilityId) + " for Facility on Encounter"

		try:
			# get the Encounter from db
			encounter = self.get( encounterId ).first()	
			
			# assign to None for unassignment
			encounter.facility = None			

			#save it
			encounter.save()

			# reload and return the appropriate version					
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAppointment( self, encounterId, appointmentId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.AppointmentDelegate import AppointmentDelegate

		errMsg = "Failed to assign element " + str(appointmentId) + " for Appointment on Encounter"

		try:
			# get the Encounter from db
			encounter = self.get( encounterId ).first()	
			
			# get the Appointment from db
			appointment = AppointmentDelegate().get(appointmentId).first();
			
			# assign the Appointment		
			encounter.appointment = appointment
			
			#save it
			encounter.save()

			# reload and return the appropriate version					
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Appointment.DoesNotExist:
			raise ProcessingError(errMsg + " : Appointment with id " + str(appointmentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAppointment( self, encounterId ):
		errMsg = "Failed to unassign element " + str(appointmentId) + " for Appointment on Encounter"

		try:
			# get the Encounter from db
			encounter = self.get( encounterId ).first()	
			
			# assign to None for unassignment
			encounter.appointment = None			

			#save it
			encounter.save()

			# reload and return the appropriate version					
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAdmission( self, encounterId, admissionId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.AdmissionDelegate import AdmissionDelegate

		errMsg = "Failed to assign element " + str(admissionId) + " for Admission on Encounter"

		try:
			# get the Encounter from db
			encounter = self.get( encounterId ).first()	
			
			# get the Admission from db
			admission = AdmissionDelegate().get(admissionId).first();
			
			# assign the Admission		
			encounter.admission = admission
			
			#save it
			encounter.save()

			# reload and return the appropriate version					
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Admission.DoesNotExist:
			raise ProcessingError(errMsg + " : Admission with id " + str(admissionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAdmission( self, encounterId ):
		errMsg = "Failed to unassign element " + str(admissionId) + " for Admission on Encounter"

		try:
			# get the Encounter from db
			encounter = self.get( encounterId ).first()	
			
			# assign to None for unassignment
			encounter.admission = None			

			#save it
			encounter.save()

			# reload and return the appropriate version					
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDischarge( self, encounterId, dischargeId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.DischargeDelegate import DischargeDelegate

		errMsg = "Failed to assign element " + str(dischargeId) + " for Discharge on Encounter"

		try:
			# get the Encounter from db
			encounter = self.get( encounterId ).first()	
			
			# get the Discharge from db
			discharge = DischargeDelegate().get(dischargeId).first();
			
			# assign the Discharge		
			encounter.discharge = discharge
			
			#save it
			encounter.save()

			# reload and return the appropriate version					
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Discharge.DoesNotExist:
			raise ProcessingError(errMsg + " : Discharge with id " + str(dischargeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDischarge( self, encounterId ):
		errMsg = "Failed to unassign element " + str(dischargeId) + " for Discharge on Encounter"

		try:
			# get the Encounter from db
			encounter = self.get( encounterId ).first()	
			
			# assign to None for unassignment
			encounter.discharge = None			

			#save it
			encounter.save()

			# reload and return the appropriate version					
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
		
	def addDiagnoses( self, encounterId, diagnosesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.DiagnosisDelegate import DiagnosisDelegate

		errMsg = "Failed to add elements " + str(diagnosesIds) + " for Diagnoses on Encounter"

		try:
			# get the Encounter
			encounter = self.get( encounterId ).first()
				
			# split on a comma with no spaces
			idList = diagnosesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Diagnosis		
				diagnosis = DiagnosisDelegate().get(id).first();	
				# add the Diagnosis
				encounter.diagnoses.add(diagnosis)
				
			# save it		
			encounter.save()
			
			# reload and return the appropriate version
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Diagnosis.DoesNotExist:
			raise ProcessingError(errMsg + " : Diagnosis does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDiagnoses( self, encounterId, diagnosesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.DiagnosisDelegate import DiagnosisDelegate

		errMsg = "Failed to remove elements " + str(diagnosesIds) + " for Diagnoses on Encounter"

		try:
			# get the Encounter
			encounter = self.get( encounterId ).first()
				
			# split on a comma with no spaces
			idList = diagnosesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Diagnosis		
				diagnosis = DiagnosisDelegate().get(id).first();	
				# add the Diagnosis
				encounter.diagnoses.remove(diagnosis)
				
			# save it		
			encounter.save()
			
			# reload and return the appropriate version
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Diagnosis.DoesNotExist:
			raise ProcessingError(errMsg + " : Diagnosis does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addProcedures( self, encounterId, proceduresIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ProcedureDelegate import ProcedureDelegate

		errMsg = "Failed to add elements " + str(proceduresIds) + " for Procedures on Encounter"

		try:
			# get the Encounter
			encounter = self.get( encounterId ).first()
				
			# split on a comma with no spaces
			idList = proceduresIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Procedure		
				procedure = ProcedureDelegate().get(id).first();	
				# add the Procedure
				encounter.procedures.add(procedure)
				
			# save it		
			encounter.save()
			
			# reload and return the appropriate version
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProcedures( self, encounterId, proceduresIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ProcedureDelegate import ProcedureDelegate

		errMsg = "Failed to remove elements " + str(proceduresIds) + " for Procedures on Encounter"

		try:
			# get the Encounter
			encounter = self.get( encounterId ).first()
				
			# split on a comma with no spaces
			idList = proceduresIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Procedure		
				procedure = ProcedureDelegate().get(id).first();	
				# add the Procedure
				encounter.procedures.remove(procedure)
				
			# save it		
			encounter.save()
			
			# reload and return the appropriate version
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addObservations( self, encounterId, observationsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ObservationDelegate import ObservationDelegate

		errMsg = "Failed to add elements " + str(observationsIds) + " for Observations on Encounter"

		try:
			# get the Encounter
			encounter = self.get( encounterId ).first()
				
			# split on a comma with no spaces
			idList = observationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Observation		
				observation = ObservationDelegate().get(id).first();	
				# add the Observation
				encounter.observations.add(observation)
				
			# save it		
			encounter.save()
			
			# reload and return the appropriate version
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeObservations( self, encounterId, observationsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ObservationDelegate import ObservationDelegate

		errMsg = "Failed to remove elements " + str(observationsIds) + " for Observations on Encounter"

		try:
			# get the Encounter
			encounter = self.get( encounterId ).first()
				
			# split on a comma with no spaces
			idList = observationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Observation		
				observation = ObservationDelegate().get(id).first();	
				# add the Observation
				encounter.observations.remove(observation)
				
			# save it		
			encounter.save()
			
			# reload and return the appropriate version
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOrders( self, encounterId, ordersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClinicalOrderDelegate import ClinicalOrderDelegate

		errMsg = "Failed to add elements " + str(ordersIds) + " for Orders on Encounter"

		try:
			# get the Encounter
			encounter = self.get( encounterId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ClinicalOrder		
				clinicalOrder = ClinicalOrderDelegate().get(id).first();	
				# add the ClinicalOrder
				encounter.orders.add(clinicalOrder)
				
			# save it		
			encounter.save()
			
			# reload and return the appropriate version
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOrders( self, encounterId, ordersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClinicalOrderDelegate import ClinicalOrderDelegate

		errMsg = "Failed to remove elements " + str(ordersIds) + " for Orders on Encounter"

		try:
			# get the Encounter
			encounter = self.get( encounterId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ClinicalOrder		
				clinicalOrder = ClinicalOrderDelegate().get(id).first();	
				# add the ClinicalOrder
				encounter.orders.remove(clinicalOrder)
				
			# save it		
			encounter.save()
			
			# reload and return the appropriate version
			return self.get( encounterId );
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
