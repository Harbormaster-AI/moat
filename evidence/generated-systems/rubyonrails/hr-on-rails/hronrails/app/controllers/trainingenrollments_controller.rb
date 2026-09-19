class TrainingEnrollmentsController < ApplicationController
  def index
    @trainingEnrollments = TrainingEnrollment.all
  end
 
  def show
    @trainingEnrollment = TrainingEnrollment.find(params[:id])
  end
 
  def new
    @trainingEnrollment = TrainingEnrollment.new
  end
 
  def edit
    @trainingEnrollment = TrainingEnrollment.find(params[:id])
  end
 
  def create
    @trainingEnrollment = TrainingEnrollment.new(trainingEnrollment_params)
 
    if @trainingEnrollment.save
      redirect_to trainingEnrollments_path
    else
      render 'new'
    end
  end
 
  def update
    @trainingEnrollment = TrainingEnrollment.find(params[:id])
 
    if @trainingEnrollment.update(trainingEnrollment_params)
      redirect_to trainingEnrollments_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @trainingEnrollment = TrainingEnrollment.find(params[:id])
    @trainingEnrollment.destroy
    redirect_to trainingEnrollments_path
  end

 
  private
    def trainingEnrollment_params
      params.require(:trainingEnrollment).permit(:enrollmentNumber, :completionDate, :score, :Status)
    end
end