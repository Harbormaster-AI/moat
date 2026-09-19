class TrainingCoursesController < ApplicationController
  def index
    @trainingCourses = TrainingCourse.all
  end
 
  def show
    @trainingCourse = TrainingCourse.find(params[:id])
  end
 
  def new
    @trainingCourse = TrainingCourse.new
  end
 
  def edit
    @trainingCourse = TrainingCourse.find(params[:id])
  end
 
  def create
    @trainingCourse = TrainingCourse.new(trainingCourse_params)
 
    if @trainingCourse.save
      redirect_to trainingCourses_path
    else
      render 'new'
    end
  end
 
  def update
    @trainingCourse = TrainingCourse.find(params[:id])
 
    if @trainingCourse.update(trainingCourse_params)
      redirect_to trainingCourses_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @trainingCourse = TrainingCourse.find(params[:id])
    @trainingCourse.destroy
    redirect_to trainingCourses_path
  end

 
  private
    def trainingCourse_params
      params.require(:trainingCourse).permit(:code, :title, :durationHours, :DeliveryMethod)
    end
end