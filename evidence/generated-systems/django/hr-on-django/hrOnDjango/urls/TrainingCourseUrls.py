from django.urls import path
from hrOnDjango.views import TrainingCourseView

urlpatterns = [
    path('', TrainingCourseView.index, name='index'),
	path('create', TrainingCourseView.get, name='create'),
	path('get/<int:trainingCourseId>/', TrainingCourseView.get, name='get'),
	path('save', TrainingCourseView.save, name='save'),
	path('getAll', TrainingCourseView.getAll, name='getAll'),
	path('delete/<int:trainingCourseId>/', TrainingCourseView.delete, name='delete'),
	path('addPrerequisites/<int:trainingCourseId>/<PrerequisitesIds>/', TrainingCourseView.addPrerequisites, name='addPrerequisites'),
	path('removePrerequisites/<int:trainingCourseId>/<PrerequisitesIds>/', TrainingCourseView.removePrerequisites, name='removePrerequisites'),
	path('addEnrollments/<int:trainingCourseId>/<EnrollmentsIds>/', TrainingCourseView.addEnrollments, name='addEnrollments'),
	path('removeEnrollments/<int:trainingCourseId>/<EnrollmentsIds>/', TrainingCourseView.removeEnrollments, name='removeEnrollments'),
	path('addJobProfiles/<int:trainingCourseId>/<JobProfilesIds>/', TrainingCourseView.addJobProfiles, name='addJobProfiles'),
	path('removeJobProfiles/<int:trainingCourseId>/<JobProfilesIds>/', TrainingCourseView.removeJobProfiles, name='removeJobProfiles'),
]
