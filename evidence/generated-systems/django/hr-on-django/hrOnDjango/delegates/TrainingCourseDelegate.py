from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.TrainingCourse import TrainingCourse
from hrOnDjango.models.TrainingEnrollment import TrainingEnrollment
from hrOnDjango.models.JobProfile import JobProfile
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model TrainingCourse
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrainingCourseDelegate Declaration
#======================================================================
class TrainingCourseDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, trainingCourseId ):
		try:	
			trainingCourse = TrainingCourse.objects.filter(id=trainingCourseId)
			return trainingCourse.first();
		except TrainingCourse.DoesNotExist:
			raise ProcessingError("TrainingCourse with id " + str(trainingCourseId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, trainingCourse):
		for model in serializers.deserialize("json", trainingCourse):
			model.save()
			return model;

	def create(self, trainingCourse):
		trainingCourse.save()
		return trainingCourse;

	def saveFromJson(self, trainingCourse):
		for model in serializers.deserialize("json", trainingCourse):
			model.save()
			return trainingCourse;
	
	def save(self, trainingCourse):
		trainingCourse.save()
		return trainingCourse;
	
	def delete(self, trainingCourseId ):
		errMsg = "Failed to delete TrainingCourse from db using id " + str(trainingCourseId)
		
		try:
			trainingCourse = TrainingCourse.objects.get(id=trainingCourseId)
			trainingCourse.delete()
			return True
		except TrainingCourse.DoesNotExist:
			raise ProcessingError("TrainingCourse with id " + str(trainingCourseId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = TrainingCourse.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all TrainingCourse from db")
		except Exception:
			return None;
		
	def addPrerequisites( self, trainingCourseId, prerequisitesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TrainingCourseDelegate import TrainingCourseDelegate

		errMsg = "Failed to add elements " + str(prerequisitesIds) + " for Prerequisites on TrainingCourse"

		try:
			# get the TrainingCourse
			trainingCourse = self.get( trainingCourseId ).first()
				
			# split on a comma with no spaces
			idList = prerequisitesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TrainingCourse		
				trainingCourse = TrainingCourseDelegate().get(id).first();	
				# add the TrainingCourse
				trainingCourse.prerequisites.add(trainingCourse)
				
			# save it		
			trainingCourse.save()
			
			# reload and return the appropriate version
			return self.get( trainingCourseId );
		except TrainingCourse.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingCourse with id " + str(trainingCourseId) + " does not exist.")
		except TrainingCourse.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingCourse does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePrerequisites( self, trainingCourseId, prerequisitesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TrainingCourseDelegate import TrainingCourseDelegate

		errMsg = "Failed to remove elements " + str(prerequisitesIds) + " for Prerequisites on TrainingCourse"

		try:
			# get the TrainingCourse
			trainingCourse = self.get( trainingCourseId ).first()
				
			# split on a comma with no spaces
			idList = prerequisitesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TrainingCourse		
				trainingCourse = TrainingCourseDelegate().get(id).first();	
				# add the TrainingCourse
				trainingCourse.prerequisites.remove(trainingCourse)
				
			# save it		
			trainingCourse.save()
			
			# reload and return the appropriate version
			return self.get( trainingCourseId );
		except TrainingCourse.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingCourse with id " + str(trainingCourseId) + " does not exist.")
		except TrainingCourse.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingCourse does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEnrollments( self, trainingCourseId, enrollmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TrainingEnrollmentDelegate import TrainingEnrollmentDelegate

		errMsg = "Failed to add elements " + str(enrollmentsIds) + " for Enrollments on TrainingCourse"

		try:
			# get the TrainingCourse
			trainingCourse = self.get( trainingCourseId ).first()
				
			# split on a comma with no spaces
			idList = enrollmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TrainingEnrollment		
				trainingEnrollment = TrainingEnrollmentDelegate().get(id).first();	
				# add the TrainingEnrollment
				trainingCourse.enrollments.add(trainingEnrollment)
				
			# save it		
			trainingCourse.save()
			
			# reload and return the appropriate version
			return self.get( trainingCourseId );
		except TrainingCourse.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingCourse with id " + str(trainingCourseId) + " does not exist.")
		except TrainingEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingEnrollment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEnrollments( self, trainingCourseId, enrollmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TrainingEnrollmentDelegate import TrainingEnrollmentDelegate

		errMsg = "Failed to remove elements " + str(enrollmentsIds) + " for Enrollments on TrainingCourse"

		try:
			# get the TrainingCourse
			trainingCourse = self.get( trainingCourseId ).first()
				
			# split on a comma with no spaces
			idList = enrollmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TrainingEnrollment		
				trainingEnrollment = TrainingEnrollmentDelegate().get(id).first();	
				# add the TrainingEnrollment
				trainingCourse.enrollments.remove(trainingEnrollment)
				
			# save it		
			trainingCourse.save()
			
			# reload and return the appropriate version
			return self.get( trainingCourseId );
		except TrainingCourse.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingCourse with id " + str(trainingCourseId) + " does not exist.")
		except TrainingEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingEnrollment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addJobProfiles( self, trainingCourseId, jobProfilesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobProfileDelegate import JobProfileDelegate

		errMsg = "Failed to add elements " + str(jobProfilesIds) + " for JobProfiles on TrainingCourse"

		try:
			# get the TrainingCourse
			trainingCourse = self.get( trainingCourseId ).first()
				
			# split on a comma with no spaces
			idList = jobProfilesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the JobProfile		
				jobProfile = JobProfileDelegate().get(id).first();	
				# add the JobProfile
				trainingCourse.jobProfiles.add(jobProfile)
				
			# save it		
			trainingCourse.save()
			
			# reload and return the appropriate version
			return self.get( trainingCourseId );
		except TrainingCourse.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingCourse with id " + str(trainingCourseId) + " does not exist.")
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeJobProfiles( self, trainingCourseId, jobProfilesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobProfileDelegate import JobProfileDelegate

		errMsg = "Failed to remove elements " + str(jobProfilesIds) + " for JobProfiles on TrainingCourse"

		try:
			# get the TrainingCourse
			trainingCourse = self.get( trainingCourseId ).first()
				
			# split on a comma with no spaces
			idList = jobProfilesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the JobProfile		
				jobProfile = JobProfileDelegate().get(id).first();	
				# add the JobProfile
				trainingCourse.jobProfiles.remove(jobProfile)
				
			# save it		
			trainingCourse.save()
			
			# reload and return the appropriate version
			return self.get( trainingCourseId );
		except TrainingCourse.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingCourse with id " + str(trainingCourseId) + " does not exist.")
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
