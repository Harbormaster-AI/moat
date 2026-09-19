class DataTasksController < ApplicationController
  def index
    @dataTasks = DataTask.all
  end
 
  def show
    @dataTask = DataTask.find(params[:id])
  end
 
  def new
    @dataTask = DataTask.new
  end
 
  def edit
    @dataTask = DataTask.find(params[:id])
  end
 
  def create
    @dataTask = DataTask.new(dataTask_params)
 
    if @dataTask.save
      redirect_to dataTasks_path
    else
      render 'new'
    end
  end
 
  def update
    @dataTask = DataTask.find(params[:id])
 
    if @dataTask.update(dataTask_params)
      redirect_to dataTasks_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @dataTask = DataTask.find(params[:id])
    @dataTask.destroy
    redirect_to dataTasks_path
  end

 
  private
    def dataTask_params
      params.require(:dataTask).permit(:name, :command, :retries, :TaskType)
    end
end