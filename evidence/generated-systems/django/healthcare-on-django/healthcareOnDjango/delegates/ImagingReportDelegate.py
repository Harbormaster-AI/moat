from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.ImagingReport import ImagingReport
from healthcareOnDjango.models.ImagingOrder import ImagingOrder
from healthcareOnDjango.models.Clinician import Clinician
from healthcareOnDjango.models.Encounter import Encounter
from healthcareOnDjango.models.ImagingCenter import ImagingCenter
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ImagingReport
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ImagingReportDelegate Declaration
#======================================================================
class ImagingReportDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, imagingReportId ):
		try:	
			imagingReport = ImagingReport.objects.filter(id=imagingReportId)
			return imagingReport.first();
		except ImagingReport.DoesNotExist:
			raise ProcessingError("ImagingReport with id " + str(imagingReportId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, imagingReport):
		for model in serializers.deserialize("json", imagingReport):
			model.save()
			return model;

	def create(self, imagingReport):
		imagingReport.save()
		return imagingReport;

	def saveFromJson(self, imagingReport):
		for model in serializers.deserialize("json", imagingReport):
			model.save()
			return imagingReport;
	
	def save(self, imagingReport):
		imagingReport.save()
		return imagingReport;
	
	def delete(self, imagingReportId ):
		errMsg = "Failed to delete ImagingReport from db using id " + str(imagingReportId)
		
		try:
			imagingReport = ImagingReport.objects.get(id=imagingReportId)
			imagingReport.delete()
			return True
		except ImagingReport.DoesNotExist:
			raise ProcessingError("ImagingReport with id " + str(imagingReportId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ImagingReport.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ImagingReport from db")
		except Exception:
			return None;
		
	def assignImagingOrder( self, imagingReportId, imagingOrderId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingOrderDelegate import ImagingOrderDelegate

		errMsg = "Failed to assign element " + str(imagingOrderId) + " for ImagingOrder on ImagingReport"

		try:
			# get the ImagingReport from db
			imagingReport = self.get( imagingReportId ).first()	
			
			# get the ImagingOrder from db
			imagingOrder = ImagingOrderDelegate().get(imagingOrderId).first();
			
			# assign the ImagingOrder		
			imagingReport.imagingOrder = imagingOrder
			
			#save it
			imagingReport.save()

			# reload and return the appropriate version					
			return self.get( imagingReportId );
		except ImagingReport.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingReport with id " + str(imagingReportId) + " does not exist.")
		except ImagingOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingOrder with id " + str(imagingOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignImagingOrder( self, imagingReportId ):
		errMsg = "Failed to unassign element " + str(imagingOrderId) + " for ImagingOrder on ImagingReport"

		try:
			# get the ImagingReport from db
			imagingReport = self.get( imagingReportId ).first()	
			
			# assign to None for unassignment
			imagingReport.imagingOrder = None			

			#save it
			imagingReport.save()

			# reload and return the appropriate version					
			return self.get( imagingReportId );
		except ImagingReport.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingReport with id " + str(imagingReportId) + " does not exist.")
		except Exception:
			return None;
		
	def assignClinician( self, imagingReportId, clinicianId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClinicianDelegate import ClinicianDelegate

		errMsg = "Failed to assign element " + str(clinicianId) + " for Clinician on ImagingReport"

		try:
			# get the ImagingReport from db
			imagingReport = self.get( imagingReportId ).first()	
			
			# get the Clinician from db
			clinician = ClinicianDelegate().get(clinicianId).first();
			
			# assign the Clinician		
			imagingReport.clinician = clinician
			
			#save it
			imagingReport.save()

			# reload and return the appropriate version					
			return self.get( imagingReportId );
		except ImagingReport.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingReport with id " + str(imagingReportId) + " does not exist.")
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(clinicianId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignClinician( self, imagingReportId ):
		errMsg = "Failed to unassign element " + str(clinicianId) + " for Clinician on ImagingReport"

		try:
			# get the ImagingReport from db
			imagingReport = self.get( imagingReportId ).first()	
			
			# assign to None for unassignment
			imagingReport.clinician = None			

			#save it
			imagingReport.save()

			# reload and return the appropriate version					
			return self.get( imagingReportId );
		except ImagingReport.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingReport with id " + str(imagingReportId) + " does not exist.")
		except Exception:
			return None;
		
	def assignEncounter( self, imagingReportId, encounterId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to assign element " + str(encounterId) + " for Encounter on ImagingReport"

		try:
			# get the ImagingReport from db
			imagingReport = self.get( imagingReportId ).first()	
			
			# get the Encounter from db
			encounter = EncounterDelegate().get(encounterId).first();
			
			# assign the Encounter		
			imagingReport.encounter = encounter
			
			#save it
			imagingReport.save()

			# reload and return the appropriate version					
			return self.get( imagingReportId );
		except ImagingReport.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingReport with id " + str(imagingReportId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEncounter( self, imagingReportId ):
		errMsg = "Failed to unassign element " + str(encounterId) + " for Encounter on ImagingReport"

		try:
			# get the ImagingReport from db
			imagingReport = self.get( imagingReportId ).first()	
			
			# assign to None for unassignment
			imagingReport.encounter = None			

			#save it
			imagingReport.save()

			# reload and return the appropriate version					
			return self.get( imagingReportId );
		except ImagingReport.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingReport with id " + str(imagingReportId) + " does not exist.")
		except Exception:
			return None;
		
	def assignImagingCenter( self, imagingReportId, imagingCenterId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingCenterDelegate import ImagingCenterDelegate

		errMsg = "Failed to assign element " + str(imagingCenterId) + " for ImagingCenter on ImagingReport"

		try:
			# get the ImagingReport from db
			imagingReport = self.get( imagingReportId ).first()	
			
			# get the ImagingCenter from db
			imagingCenter = ImagingCenterDelegate().get(imagingCenterId).first();
			
			# assign the ImagingCenter		
			imagingReport.imagingCenter = imagingCenter
			
			#save it
			imagingReport.save()

			# reload and return the appropriate version					
			return self.get( imagingReportId );
		except ImagingReport.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingReport with id " + str(imagingReportId) + " does not exist.")
		except ImagingCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingCenter with id " + str(imagingCenterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignImagingCenter( self, imagingReportId ):
		errMsg = "Failed to unassign element " + str(imagingCenterId) + " for ImagingCenter on ImagingReport"

		try:
			# get the ImagingReport from db
			imagingReport = self.get( imagingReportId ).first()	
			
			# assign to None for unassignment
			imagingReport.imagingCenter = None			

			#save it
			imagingReport.save()

			# reload and return the appropriate version					
			return self.get( imagingReportId );
		except ImagingReport.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingReport with id " + str(imagingReportId) + " does not exist.")
		except Exception:
			return None;
		
