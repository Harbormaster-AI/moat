class RoleAssignmentsController < ApplicationController
  def index
    @roleAssignments = RoleAssignment.all
  end
 
  def show
    @roleAssignment = RoleAssignment.find(params[:id])
  end
 
  def new
    @roleAssignment = RoleAssignment.new
  end
 
  def edit
    @roleAssignment = RoleAssignment.find(params[:id])
  end
 
  def create
    @roleAssignment = RoleAssignment.new(roleAssignment_params)
 
    if @roleAssignment.save
      redirect_to roleAssignments_path
    else
      render 'new'
    end
  end
 
  def update
    @roleAssignment = RoleAssignment.find(params[:id])
 
    if @roleAssignment.update(roleAssignment_params)
      redirect_to roleAssignments_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @roleAssignment = RoleAssignment.find(params[:id])
    @roleAssignment.destroy
    redirect_to roleAssignments_path
  end

 
  private
    def roleAssignment_params
      params.require(:roleAssignment).permit(:effectiveFrom, :effectiveTo)
    end
end