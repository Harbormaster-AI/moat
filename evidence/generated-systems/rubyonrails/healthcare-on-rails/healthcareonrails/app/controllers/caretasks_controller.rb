class CareTasksController < ApplicationController
  def index
    @careTasks = CareTask.all
  end
 
  def show
    @careTask = CareTask.find(params[:id])
  end
 
  def new
    @careTask = CareTask.new
  end
 
  def edit
    @careTask = CareTask.find(params[:id])
  end
 
  def create
    @careTask = CareTask.new(careTask_params)
 
    if @careTask.save
      redirect_to careTasks_path
    else
      render 'new'
    end
  end
 
  def update
    @careTask = CareTask.find(params[:id])
 
    if @careTask.update(careTask_params)
      redirect_to careTasks_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @careTask = CareTask.find(params[:id])
    @careTask.destroy
    redirect_to careTasks_path
  end

 
  private
    def careTask_params
      params.require(:careTask).permit(:description, :dueDate, :Status, :Priority)
    end
end