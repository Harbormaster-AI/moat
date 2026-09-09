from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Screening import Screening
from hrOnDjango.models.JobApplication import JobApplication
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Screening
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ScreeningDelegate Declaration
#======================================================================
class ScreeningDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, screeningId ):
		try:	
			screening = Screening.objects.filter(id=screeningId)
			return screening.first();
		except Screening.DoesNotExist:
			raise ProcessingError("Screening with id " + str(screeningId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, screening):
		for model in serializers.deserialize("json", screening):
			model.save()
			return model;

	def create(self, screening):
		screening.save()
		return screening;

	def saveFromJson(self, screening):
		for model in serializers.deserialize("json", screening):
			model.save()
			return screening;
	
	def save(self, screening):
		screening.save()
		return screening;
	
	def delete(self, screeningId ):
		errMsg = "Failed to delete Screening from db using id " + str(screeningId)
		
		try:
			screening = Screening.objects.get(id=screeningId)
			screening.delete()
			return True
		except Screening.DoesNotExist:
			raise ProcessingError("Screening with id " + str(screeningId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Screening.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Screening from db")
		except Exception:
			return None;
		
	def assignApplication( self, screeningId, applicationId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobApplicationDelegate import JobApplicationDelegate

		errMsg = "Failed to assign element " + str(applicationId) + " for Application on Screening"

		try:
			# get the Screening from db
			screening = self.get( screeningId ).first()	
			
			# get the JobApplication from db
			jobApplication = JobApplicationDelegate().get(applicationId).first();
			
			# assign the Application		
			screening.application = jobApplication
			
			#save it
			screening.save()

			# reload and return the appropriate version					
			return self.get( screeningId );
		except Screening.DoesNotExist:
			raise ProcessingError(errMsg + " : Screening with id " + str(screeningId) + " does not exist.")
		except JobApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : JobApplication with id " + str(applicationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignApplication( self, screeningId ):
		errMsg = "Failed to unassign element " + str(applicationId) + " for Application on Screening"

		try:
			# get the Screening from db
			screening = self.get( screeningId ).first()	
			
			# assign to None for unassignment
			screening.jobApplication = None			

			#save it
			screening.save()

			# reload and return the appropriate version					
			return self.get( screeningId );
		except Screening.DoesNotExist:
			raise ProcessingError(errMsg + " : Screening with id " + str(screeningId) + " does not exist.")
		except Exception:
			return None;
		
