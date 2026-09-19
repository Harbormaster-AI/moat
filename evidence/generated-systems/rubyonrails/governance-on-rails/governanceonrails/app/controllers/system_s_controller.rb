class System_sController < ApplicationController
  def index
    @system_s = System_.all
  end
 
  def show
    @system_ = System_.find(params[:id])
  end
 
  def new
    @system_ = System_.new
  end
 
  def edit
    @system_ = System_.find(params[:id])
  end
 
  def create
    @system_ = System_.new(system__params)
 
    if @system_.save
      redirect_to system_s_path
    else
      render 'new'
    end
  end
 
  def update
    @system_ = System_.find(params[:id])
 
    if @system_.update(system__params)
      redirect_to system_s_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @system_ = System_.find(params[:id])
    @system_.destroy
    redirect_to system_s_path
  end

 
  private
    def system__params
      params.require(:system_).permit(:name, :ownerDepartment, :SystemType)
    end
end