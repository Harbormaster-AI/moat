from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Competency import Competency
from hrOnDjango.models.JobProfile import JobProfile
from hrOnDjango.models.CompetencyRating import CompetencyRating
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Competency
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CompetencyDelegate Declaration
#======================================================================
class CompetencyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, competencyId ):
		try:	
			competency = Competency.objects.filter(id=competencyId)
			return competency.first();
		except Competency.DoesNotExist:
			raise ProcessingError("Competency with id " + str(competencyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, competency):
		for model in serializers.deserialize("json", competency):
			model.save()
			return model;

	def create(self, competency):
		competency.save()
		return competency;

	def saveFromJson(self, competency):
		for model in serializers.deserialize("json", competency):
			model.save()
			return competency;
	
	def save(self, competency):
		competency.save()
		return competency;
	
	def delete(self, competencyId ):
		errMsg = "Failed to delete Competency from db using id " + str(competencyId)
		
		try:
			competency = Competency.objects.get(id=competencyId)
			competency.delete()
			return True
		except Competency.DoesNotExist:
			raise ProcessingError("Competency with id " + str(competencyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Competency.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Competency from db")
		except Exception:
			return None;
		
	def addJobProfiles( self, competencyId, jobProfilesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobProfileDelegate import JobProfileDelegate

		errMsg = "Failed to add elements " + str(jobProfilesIds) + " for JobProfiles on Competency"

		try:
			# get the Competency
			competency = self.get( competencyId ).first()
				
			# split on a comma with no spaces
			idList = jobProfilesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the JobProfile		
				jobProfile = JobProfileDelegate().get(id).first();	
				# add the JobProfile
				competency.jobProfiles.add(jobProfile)
				
			# save it		
			competency.save()
			
			# reload and return the appropriate version
			return self.get( competencyId );
		except Competency.DoesNotExist:
			raise ProcessingError(errMsg + " : Competency with id " + str(competencyId) + " does not exist.")
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeJobProfiles( self, competencyId, jobProfilesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobProfileDelegate import JobProfileDelegate

		errMsg = "Failed to remove elements " + str(jobProfilesIds) + " for JobProfiles on Competency"

		try:
			# get the Competency
			competency = self.get( competencyId ).first()
				
			# split on a comma with no spaces
			idList = jobProfilesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the JobProfile		
				jobProfile = JobProfileDelegate().get(id).first();	
				# add the JobProfile
				competency.jobProfiles.remove(jobProfile)
				
			# save it		
			competency.save()
			
			# reload and return the appropriate version
			return self.get( competencyId );
		except Competency.DoesNotExist:
			raise ProcessingError(errMsg + " : Competency with id " + str(competencyId) + " does not exist.")
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCompetencyRatings( self, competencyId, competencyRatingsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CompetencyRatingDelegate import CompetencyRatingDelegate

		errMsg = "Failed to add elements " + str(competencyRatingsIds) + " for CompetencyRatings on Competency"

		try:
			# get the Competency
			competency = self.get( competencyId ).first()
				
			# split on a comma with no spaces
			idList = competencyRatingsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CompetencyRating		
				competencyRating = CompetencyRatingDelegate().get(id).first();	
				# add the CompetencyRating
				competency.competencyRatings.add(competencyRating)
				
			# save it		
			competency.save()
			
			# reload and return the appropriate version
			return self.get( competencyId );
		except Competency.DoesNotExist:
			raise ProcessingError(errMsg + " : Competency with id " + str(competencyId) + " does not exist.")
		except CompetencyRating.DoesNotExist:
			raise ProcessingError(errMsg + " : CompetencyRating does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCompetencyRatings( self, competencyId, competencyRatingsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CompetencyRatingDelegate import CompetencyRatingDelegate

		errMsg = "Failed to remove elements " + str(competencyRatingsIds) + " for CompetencyRatings on Competency"

		try:
			# get the Competency
			competency = self.get( competencyId ).first()
				
			# split on a comma with no spaces
			idList = competencyRatingsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CompetencyRating		
				competencyRating = CompetencyRatingDelegate().get(id).first();	
				# add the CompetencyRating
				competency.competencyRatings.remove(competencyRating)
				
			# save it		
			competency.save()
			
			# reload and return the appropriate version
			return self.get( competencyId );
		except Competency.DoesNotExist:
			raise ProcessingError(errMsg + " : Competency with id " + str(competencyId) + " does not exist.")
		except CompetencyRating.DoesNotExist:
			raise ProcessingError(errMsg + " : CompetencyRating does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
