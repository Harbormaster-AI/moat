from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Clinician import Clinician
from healthcareOnDjango.models.CareTeam import CareTeam
from healthcareOnDjango.models.Appointment import Appointment
from healthcareOnDjango.models.Encounter import Encounter
from healthcareOnDjango.models.Procedure import Procedure
from healthcareOnDjango.models.ImagingReport import ImagingReport
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Clinician
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClinicianDelegate Declaration
#======================================================================
class ClinicianDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, clinicianId ):
		try:	
			clinician = Clinician.objects.filter(id=clinicianId)
			return clinician.first();
		except Clinician.DoesNotExist:
			raise ProcessingError("Clinician with id " + str(clinicianId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, clinician):
		for model in serializers.deserialize("json", clinician):
			model.save()
			return model;

	def create(self, clinician):
		clinician.save()
		return clinician;

	def saveFromJson(self, clinician):
		for model in serializers.deserialize("json", clinician):
			model.save()
			return clinician;
	
	def save(self, clinician):
		clinician.save()
		return clinician;
	
	def delete(self, clinicianId ):
		errMsg = "Failed to delete Clinician from db using id " + str(clinicianId)
		
		try:
			clinician = Clinician.objects.get(id=clinicianId)
			clinician.delete()
			return True
		except Clinician.DoesNotExist:
			raise ProcessingError("Clinician with id " + str(clinicianId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Clinician.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Clinician from db")
		except Exception:
			return None;
		
	def addCareTeams( self, clinicianId, careTeamsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CareTeamDelegate import CareTeamDelegate

		errMsg = "Failed to add elements " + str(careTeamsIds) + " for CareTeams on Clinician"

		try:
			# get the Clinician
			clinician = self.get( clinicianId ).first()
				
			# split on a comma with no spaces
			idList = careTeamsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CareTeam		
				careTeam = CareTeamDelegate().get(id).first();	
				# add the CareTeam
				clinician.careTeams.add(careTeam)
				
			# save it		
			clinician.save()
			
			# reload and return the appropriate version
			return self.get( clinicianId );
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(clinicianId) + " does not exist.")
		except CareTeam.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTeam does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCareTeams( self, clinicianId, careTeamsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CareTeamDelegate import CareTeamDelegate

		errMsg = "Failed to remove elements " + str(careTeamsIds) + " for CareTeams on Clinician"

		try:
			# get the Clinician
			clinician = self.get( clinicianId ).first()
				
			# split on a comma with no spaces
			idList = careTeamsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CareTeam		
				careTeam = CareTeamDelegate().get(id).first();	
				# add the CareTeam
				clinician.careTeams.remove(careTeam)
				
			# save it		
			clinician.save()
			
			# reload and return the appropriate version
			return self.get( clinicianId );
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(clinicianId) + " does not exist.")
		except CareTeam.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTeam does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAppointments( self, clinicianId, appointmentsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.AppointmentDelegate import AppointmentDelegate

		errMsg = "Failed to add elements " + str(appointmentsIds) + " for Appointments on Clinician"

		try:
			# get the Clinician
			clinician = self.get( clinicianId ).first()
				
			# split on a comma with no spaces
			idList = appointmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Appointment		
				appointment = AppointmentDelegate().get(id).first();	
				# add the Appointment
				clinician.appointments.add(appointment)
				
			# save it		
			clinician.save()
			
			# reload and return the appropriate version
			return self.get( clinicianId );
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(clinicianId) + " does not exist.")
		except Appointment.DoesNotExist:
			raise ProcessingError(errMsg + " : Appointment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAppointments( self, clinicianId, appointmentsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.AppointmentDelegate import AppointmentDelegate

		errMsg = "Failed to remove elements " + str(appointmentsIds) + " for Appointments on Clinician"

		try:
			# get the Clinician
			clinician = self.get( clinicianId ).first()
				
			# split on a comma with no spaces
			idList = appointmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Appointment		
				appointment = AppointmentDelegate().get(id).first();	
				# add the Appointment
				clinician.appointments.remove(appointment)
				
			# save it		
			clinician.save()
			
			# reload and return the appropriate version
			return self.get( clinicianId );
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(clinicianId) + " does not exist.")
		except Appointment.DoesNotExist:
			raise ProcessingError(errMsg + " : Appointment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEncounters( self, clinicianId, encountersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to add elements " + str(encountersIds) + " for Encounters on Clinician"

		try:
			# get the Clinician
			clinician = self.get( clinicianId ).first()
				
			# split on a comma with no spaces
			idList = encountersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Encounter		
				encounter = EncounterDelegate().get(id).first();	
				# add the Encounter
				clinician.encounters.add(encounter)
				
			# save it		
			clinician.save()
			
			# reload and return the appropriate version
			return self.get( clinicianId );
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(clinicianId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEncounters( self, clinicianId, encountersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to remove elements " + str(encountersIds) + " for Encounters on Clinician"

		try:
			# get the Clinician
			clinician = self.get( clinicianId ).first()
				
			# split on a comma with no spaces
			idList = encountersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Encounter		
				encounter = EncounterDelegate().get(id).first();	
				# add the Encounter
				clinician.encounters.remove(encounter)
				
			# save it		
			clinician.save()
			
			# reload and return the appropriate version
			return self.get( clinicianId );
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(clinicianId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addProcedures( self, clinicianId, proceduresIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ProcedureDelegate import ProcedureDelegate

		errMsg = "Failed to add elements " + str(proceduresIds) + " for Procedures on Clinician"

		try:
			# get the Clinician
			clinician = self.get( clinicianId ).first()
				
			# split on a comma with no spaces
			idList = proceduresIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Procedure		
				procedure = ProcedureDelegate().get(id).first();	
				# add the Procedure
				clinician.procedures.add(procedure)
				
			# save it		
			clinician.save()
			
			# reload and return the appropriate version
			return self.get( clinicianId );
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(clinicianId) + " does not exist.")
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProcedures( self, clinicianId, proceduresIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ProcedureDelegate import ProcedureDelegate

		errMsg = "Failed to remove elements " + str(proceduresIds) + " for Procedures on Clinician"

		try:
			# get the Clinician
			clinician = self.get( clinicianId ).first()
				
			# split on a comma with no spaces
			idList = proceduresIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Procedure		
				procedure = ProcedureDelegate().get(id).first();	
				# add the Procedure
				clinician.procedures.remove(procedure)
				
			# save it		
			clinician.save()
			
			# reload and return the appropriate version
			return self.get( clinicianId );
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(clinicianId) + " does not exist.")
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addImagingReports( self, clinicianId, imagingReportsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingReportDelegate import ImagingReportDelegate

		errMsg = "Failed to add elements " + str(imagingReportsIds) + " for ImagingReports on Clinician"

		try:
			# get the Clinician
			clinician = self.get( clinicianId ).first()
				
			# split on a comma with no spaces
			idList = imagingReportsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ImagingReport		
				imagingReport = ImagingReportDelegate().get(id).first();	
				# add the ImagingReport
				clinician.imagingReports.add(imagingReport)
				
			# save it		
			clinician.save()
			
			# reload and return the appropriate version
			return self.get( clinicianId );
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(clinicianId) + " does not exist.")
		except ImagingReport.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingReport does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeImagingReports( self, clinicianId, imagingReportsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingReportDelegate import ImagingReportDelegate

		errMsg = "Failed to remove elements " + str(imagingReportsIds) + " for ImagingReports on Clinician"

		try:
			# get the Clinician
			clinician = self.get( clinicianId ).first()
				
			# split on a comma with no spaces
			idList = imagingReportsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ImagingReport		
				imagingReport = ImagingReportDelegate().get(id).first();	
				# add the ImagingReport
				clinician.imagingReports.remove(imagingReport)
				
			# save it		
			clinician.save()
			
			# reload and return the appropriate version
			return self.get( clinicianId );
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(clinicianId) + " does not exist.")
		except ImagingReport.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingReport does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
