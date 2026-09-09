from django.urls import path
from hrOnDjango.views import TrainingEnrollmentView

urlpatterns = [
    path('', TrainingEnrollmentView.index, name='index'),
	path('create', TrainingEnrollmentView.get, name='create'),
	path('get/<int:trainingEnrollmentId>/', TrainingEnrollmentView.get, name='get'),
	path('save', TrainingEnrollmentView.save, name='save'),
	path('getAll', TrainingEnrollmentView.getAll, name='getAll'),
	path('delete/<int:trainingEnrollmentId>/', TrainingEnrollmentView.delete, name='delete'),
	path('assignCourse/<int:trainingEnrollmentId>/<int:CourseId>/', TrainingEnrollmentView.assignCourse, name='assignCourse'),
	path('unassignCourse/<int:trainingEnrollmentId>/', TrainingEnrollmentView.unassignCourse, name='unassignCourse'),
	path('assignEmployee/<int:trainingEnrollmentId>/<int:EmployeeId>/', TrainingEnrollmentView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:trainingEnrollmentId>/', TrainingEnrollmentView.unassignEmployee, name='unassignEmployee'),
	path('assignInstructor/<int:trainingEnrollmentId>/<int:InstructorId>/', TrainingEnrollmentView.assignInstructor, name='assignInstructor'),
	path('unassignInstructor/<int:trainingEnrollmentId>/', TrainingEnrollmentView.unassignInstructor, name='unassignInstructor'),
]
