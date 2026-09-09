from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.JobProfile import JobProfile
from hrOnDjango.models.JobFamily import JobFamily
from hrOnDjango.models.Competency import Competency
from hrOnDjango.models.TrainingCourse import TrainingCourse
from hrOnDjango.models.Position import Position
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model JobProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobProfileDelegate Declaration
#======================================================================
class JobProfileDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, jobProfileId ):
		try:	
			jobProfile = JobProfile.objects.filter(id=jobProfileId)
			return jobProfile.first();
		except JobProfile.DoesNotExist:
			raise ProcessingError("JobProfile with id " + str(jobProfileId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, jobProfile):
		for model in serializers.deserialize("json", jobProfile):
			model.save()
			return model;

	def create(self, jobProfile):
		jobProfile.save()
		return jobProfile;

	def saveFromJson(self, jobProfile):
		for model in serializers.deserialize("json", jobProfile):
			model.save()
			return jobProfile;
	
	def save(self, jobProfile):
		jobProfile.save()
		return jobProfile;
	
	def delete(self, jobProfileId ):
		errMsg = "Failed to delete JobProfile from db using id " + str(jobProfileId)
		
		try:
			jobProfile = JobProfile.objects.get(id=jobProfileId)
			jobProfile.delete()
			return True
		except JobProfile.DoesNotExist:
			raise ProcessingError("JobProfile with id " + str(jobProfileId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = JobProfile.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all JobProfile from db")
		except Exception:
			return None;
		
	def assignJobFamily( self, jobProfileId, jobFamilyId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobFamilyDelegate import JobFamilyDelegate

		errMsg = "Failed to assign element " + str(jobFamilyId) + " for JobFamily on JobProfile"

		try:
			# get the JobProfile from db
			jobProfile = self.get( jobProfileId ).first()	
			
			# get the JobFamily from db
			jobFamily = JobFamilyDelegate().get(jobFamilyId).first();
			
			# assign the JobFamily		
			jobProfile.jobFamily = jobFamily
			
			#save it
			jobProfile.save()

			# reload and return the appropriate version					
			return self.get( jobProfileId );
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile with id " + str(jobProfileId) + " does not exist.")
		except JobFamily.DoesNotExist:
			raise ProcessingError(errMsg + " : JobFamily with id " + str(jobFamilyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignJobFamily( self, jobProfileId ):
		errMsg = "Failed to unassign element " + str(jobFamilyId) + " for JobFamily on JobProfile"

		try:
			# get the JobProfile from db
			jobProfile = self.get( jobProfileId ).first()	
			
			# assign to None for unassignment
			jobProfile.jobFamily = None			

			#save it
			jobProfile.save()

			# reload and return the appropriate version					
			return self.get( jobProfileId );
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile with id " + str(jobProfileId) + " does not exist.")
		except Exception:
			return None;
		
	def addCompetencies( self, jobProfileId, competenciesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CompetencyDelegate import CompetencyDelegate

		errMsg = "Failed to add elements " + str(competenciesIds) + " for Competencies on JobProfile"

		try:
			# get the JobProfile
			jobProfile = self.get( jobProfileId ).first()
				
			# split on a comma with no spaces
			idList = competenciesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Competency		
				competency = CompetencyDelegate().get(id).first();	
				# add the Competency
				jobProfile.competencies.add(competency)
				
			# save it		
			jobProfile.save()
			
			# reload and return the appropriate version
			return self.get( jobProfileId );
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile with id " + str(jobProfileId) + " does not exist.")
		except Competency.DoesNotExist:
			raise ProcessingError(errMsg + " : Competency does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCompetencies( self, jobProfileId, competenciesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CompetencyDelegate import CompetencyDelegate

		errMsg = "Failed to remove elements " + str(competenciesIds) + " for Competencies on JobProfile"

		try:
			# get the JobProfile
			jobProfile = self.get( jobProfileId ).first()
				
			# split on a comma with no spaces
			idList = competenciesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Competency		
				competency = CompetencyDelegate().get(id).first();	
				# add the Competency
				jobProfile.competencies.remove(competency)
				
			# save it		
			jobProfile.save()
			
			# reload and return the appropriate version
			return self.get( jobProfileId );
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile with id " + str(jobProfileId) + " does not exist.")
		except Competency.DoesNotExist:
			raise ProcessingError(errMsg + " : Competency does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTrainingRecommendations( self, jobProfileId, trainingRecommendationsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TrainingCourseDelegate import TrainingCourseDelegate

		errMsg = "Failed to add elements " + str(trainingRecommendationsIds) + " for TrainingRecommendations on JobProfile"

		try:
			# get the JobProfile
			jobProfile = self.get( jobProfileId ).first()
				
			# split on a comma with no spaces
			idList = trainingRecommendationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TrainingCourse		
				trainingCourse = TrainingCourseDelegate().get(id).first();	
				# add the TrainingCourse
				jobProfile.trainingRecommendations.add(trainingCourse)
				
			# save it		
			jobProfile.save()
			
			# reload and return the appropriate version
			return self.get( jobProfileId );
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile with id " + str(jobProfileId) + " does not exist.")
		except TrainingCourse.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingCourse does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTrainingRecommendations( self, jobProfileId, trainingRecommendationsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TrainingCourseDelegate import TrainingCourseDelegate

		errMsg = "Failed to remove elements " + str(trainingRecommendationsIds) + " for TrainingRecommendations on JobProfile"

		try:
			# get the JobProfile
			jobProfile = self.get( jobProfileId ).first()
				
			# split on a comma with no spaces
			idList = trainingRecommendationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TrainingCourse		
				trainingCourse = TrainingCourseDelegate().get(id).first();	
				# add the TrainingCourse
				jobProfile.trainingRecommendations.remove(trainingCourse)
				
			# save it		
			jobProfile.save()
			
			# reload and return the appropriate version
			return self.get( jobProfileId );
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile with id " + str(jobProfileId) + " does not exist.")
		except TrainingCourse.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingCourse does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPositions( self, jobProfileId, positionsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to add elements " + str(positionsIds) + " for Positions on JobProfile"

		try:
			# get the JobProfile
			jobProfile = self.get( jobProfileId ).first()
				
			# split on a comma with no spaces
			idList = positionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Position		
				position = PositionDelegate().get(id).first();	
				# add the Position
				jobProfile.positions.add(position)
				
			# save it		
			jobProfile.save()
			
			# reload and return the appropriate version
			return self.get( jobProfileId );
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile with id " + str(jobProfileId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePositions( self, jobProfileId, positionsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to remove elements " + str(positionsIds) + " for Positions on JobProfile"

		try:
			# get the JobProfile
			jobProfile = self.get( jobProfileId ).first()
				
			# split on a comma with no spaces
			idList = positionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Position		
				position = PositionDelegate().get(id).first();	
				# add the Position
				jobProfile.positions.remove(position)
				
			# save it		
			jobProfile.save()
			
			# reload and return the appropriate version
			return self.get( jobProfileId );
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile with id " + str(jobProfileId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
