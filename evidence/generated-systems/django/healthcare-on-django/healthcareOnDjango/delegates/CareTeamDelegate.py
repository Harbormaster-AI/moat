from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.CareTeam import CareTeam
from healthcareOnDjango.models.Department import Department
from healthcareOnDjango.models.Clinician import Clinician
from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CareTeam
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CareTeamDelegate Declaration
#======================================================================
class CareTeamDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, careTeamId ):
		try:	
			careTeam = CareTeam.objects.filter(id=careTeamId)
			return careTeam.first();
		except CareTeam.DoesNotExist:
			raise ProcessingError("CareTeam with id " + str(careTeamId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, careTeam):
		for model in serializers.deserialize("json", careTeam):
			model.save()
			return model;

	def create(self, careTeam):
		careTeam.save()
		return careTeam;

	def saveFromJson(self, careTeam):
		for model in serializers.deserialize("json", careTeam):
			model.save()
			return careTeam;
	
	def save(self, careTeam):
		careTeam.save()
		return careTeam;
	
	def delete(self, careTeamId ):
		errMsg = "Failed to delete CareTeam from db using id " + str(careTeamId)
		
		try:
			careTeam = CareTeam.objects.get(id=careTeamId)
			careTeam.delete()
			return True
		except CareTeam.DoesNotExist:
			raise ProcessingError("CareTeam with id " + str(careTeamId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CareTeam.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CareTeam from db")
		except Exception:
			return None;
		
	def assignDepartment( self, careTeamId, departmentId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.DepartmentDelegate import DepartmentDelegate

		errMsg = "Failed to assign element " + str(departmentId) + " for Department on CareTeam"

		try:
			# get the CareTeam from db
			careTeam = self.get( careTeamId ).first()	
			
			# get the Department from db
			department = DepartmentDelegate().get(departmentId).first();
			
			# assign the Department		
			careTeam.department = department
			
			#save it
			careTeam.save()

			# reload and return the appropriate version					
			return self.get( careTeamId );
		except CareTeam.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTeam with id " + str(careTeamId) + " does not exist.")
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDepartment( self, careTeamId ):
		errMsg = "Failed to unassign element " + str(departmentId) + " for Department on CareTeam"

		try:
			# get the CareTeam from db
			careTeam = self.get( careTeamId ).first()	
			
			# assign to None for unassignment
			careTeam.department = None			

			#save it
			careTeam.save()

			# reload and return the appropriate version					
			return self.get( careTeamId );
		except CareTeam.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTeam with id " + str(careTeamId) + " does not exist.")
		except Exception:
			return None;
		
	def addClinicians( self, careTeamId, cliniciansIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClinicianDelegate import ClinicianDelegate

		errMsg = "Failed to add elements " + str(cliniciansIds) + " for Clinicians on CareTeam"

		try:
			# get the CareTeam
			careTeam = self.get( careTeamId ).first()
				
			# split on a comma with no spaces
			idList = cliniciansIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Clinician		
				clinician = ClinicianDelegate().get(id).first();	
				# add the Clinician
				careTeam.clinicians.add(clinician)
				
			# save it		
			careTeam.save()
			
			# reload and return the appropriate version
			return self.get( careTeamId );
		except CareTeam.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTeam with id " + str(careTeamId) + " does not exist.")
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeClinicians( self, careTeamId, cliniciansIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClinicianDelegate import ClinicianDelegate

		errMsg = "Failed to remove elements " + str(cliniciansIds) + " for Clinicians on CareTeam"

		try:
			# get the CareTeam
			careTeam = self.get( careTeamId ).first()
				
			# split on a comma with no spaces
			idList = cliniciansIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Clinician		
				clinician = ClinicianDelegate().get(id).first();	
				# add the Clinician
				careTeam.clinicians.remove(clinician)
				
			# save it		
			careTeam.save()
			
			# reload and return the appropriate version
			return self.get( careTeamId );
		except CareTeam.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTeam with id " + str(careTeamId) + " does not exist.")
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPatients( self, careTeamId, patientsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

		errMsg = "Failed to add elements " + str(patientsIds) + " for Patients on CareTeam"

		try:
			# get the CareTeam
			careTeam = self.get( careTeamId ).first()
				
			# split on a comma with no spaces
			idList = patientsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Patient		
				patient = PatientDelegate().get(id).first();	
				# add the Patient
				careTeam.patients.add(patient)
				
			# save it		
			careTeam.save()
			
			# reload and return the appropriate version
			return self.get( careTeamId );
		except CareTeam.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTeam with id " + str(careTeamId) + " does not exist.")
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePatients( self, careTeamId, patientsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

		errMsg = "Failed to remove elements " + str(patientsIds) + " for Patients on CareTeam"

		try:
			# get the CareTeam
			careTeam = self.get( careTeamId ).first()
				
			# split on a comma with no spaces
			idList = patientsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Patient		
				patient = PatientDelegate().get(id).first();	
				# add the Patient
				careTeam.patients.remove(patient)
				
			# save it		
			careTeam.save()
			
			# reload and return the appropriate version
			return self.get( careTeamId );
		except CareTeam.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTeam with id " + str(careTeamId) + " does not exist.")
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
