class OnboardingTasksController < ApplicationController
  def index
    @onboardingTasks = OnboardingTask.all
  end
 
  def show
    @onboardingTask = OnboardingTask.find(params[:id])
  end
 
  def new
    @onboardingTask = OnboardingTask.new
  end
 
  def edit
    @onboardingTask = OnboardingTask.find(params[:id])
  end
 
  def create
    @onboardingTask = OnboardingTask.new(onboardingTask_params)
 
    if @onboardingTask.save
      redirect_to onboardingTasks_path
    else
      render 'new'
    end
  end
 
  def update
    @onboardingTask = OnboardingTask.find(params[:id])
 
    if @onboardingTask.update(onboardingTask_params)
      redirect_to onboardingTasks_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @onboardingTask = OnboardingTask.find(params[:id])
    @onboardingTask.destroy
    redirect_to onboardingTasks_path
  end

 
  private
    def onboardingTask_params
      params.require(:onboardingTask).permit(:taskNumber, :name, :dueDate, :Status)
    end
end